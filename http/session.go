package http

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/WelcomerTeam/Discord/discord"
	jsoniter "github.com/json-iterator/go"
	"github.com/rs/zerolog"
	"golang.org/x/xerrors"
)

const (
	APIVersion = "9"
)

type RESTInterface interface {
	// Fetch constructs a request. It will return a response body along with any errors.
	// Errors can include ErrInvalidToken, ErrRateLimited,
	Fetch(ctx context.Context, method, endpoint, contentType string, body []byte, headers http.Header) (response []byte, err error)
	FetchBJ(ctx context.Context, method, endpoint, contentType string, body []byte, headers http.Header, response interface{}) (err error)
	FetchJJ(ctx context.Context, method, endpoint string, payload interface{}, headers http.Header, response interface{}) (err error)
}

// Session contains the context for the discord rest interface.
type Session struct {
	Token string

	Interface RESTInterface
	Logger    zerolog.Logger
}

func NewSession(token string, httpInterface RESTInterface, logger zerolog.Logger) *Session {
	return &Session{
		Token:     token,
		Interface: httpInterface,
		Logger:    logger,
	}
}

// TwilightProxy is a proxy that requests are sent through, instead of directly to discord that will handle
// distributed requests and ratelimits automatically. See more at: https://github.com/twilight-rs/http-proxy
type TwilightProxy struct {
	HTTP       *http.Client
	APIVersion string
	URLHost    string
	URLScheme  string
	UserAgent  string
}

func NewTwilightProxy(url url.URL) RESTInterface {
	return &TwilightProxy{
		HTTP: &http.Client{
			Timeout: 20 * time.Second,
		},
		APIVersion: APIVersion,
		URLHost:    url.Host,
		URLScheme:  url.Scheme,
		UserAgent:  "Sandwich (github.com/WelcomerTeam/Discord)",
	}
}

func (tl *TwilightProxy) Fetch(ctx context.Context, method, endpoint, contentType string, body []byte, headers http.Header) (response []byte, err error) {
	req, err := http.NewRequestWithContext(ctx, method, endpoint, bytes.NewBuffer(body))
	if err != nil {
		return nil, xerrors.Errorf("Failed to create new request: %v", err)
	}

	req.URL.Host = tl.URLHost
	req.URL.Scheme = tl.URLScheme

	for name, values := range headers {
		for _, value := range values {
			req.Header.Add(name, value)
		}
	}

	if body != nil && len(req.Header.Get("content-type")) == 0 {
		req.Header.Set("content-type", contentType)
	}

	resp, err := tl.HTTP.Do(req)
	if err != nil {
		return nil, xerrors.Errorf("Failed to do request: %v", err)
	}

	defer resp.Body.Close()

	response, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, xerrors.Errorf("Failed to read body: %v", err)
	}

	switch resp.StatusCode {
	case http.StatusOK:
	case http.StatusCreated:
	case http.StatusNoContent:
	case http.StatusUnauthorized:
		return response, discord.ErrUnauthorized
	default:
		return response, NewRestError(req, resp, body)
	}

	return response, nil
}

func (tl *TwilightProxy) FetchBJ(ctx context.Context, method, endpoint, contentType string, body []byte, headers http.Header, response interface{}) (err error) {
	resp, err := tl.Fetch(ctx, method, endpoint, contentType, body, headers)
	if err != nil {
		return err
	}

	err = jsoniter.Unmarshal(resp, response)
	if err != nil {
		return xerrors.Errorf("Failed to unmarshal response: %v", err)
	}

	return nil
}

func (tl *TwilightProxy) FetchJJ(ctx context.Context, method, endpoint string, payload interface{}, headers http.Header, response interface{}) (err error) {
	body, err := jsoniter.Marshal(payload)
	if err != nil {
		return xerrors.Errorf("Failed to marshal payload: %v", err)
	}

	return tl.FetchBJ(ctx, method, endpoint, "application/json", body, headers, response)
}
