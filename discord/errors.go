package discord

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"golang.org/x/xerrors"
	"net/http"
)

var (
	ErrUnauthorized         = xerrors.New("Inproper token was passed")
	ErrUnsupportedImageType = xerrors.New("Unsupported image type given")
)

// RestError contains the error structure that is returned by discord.
type RestError struct {
	Request      *http.Request
	Response     *http.Response
	ResponseBody []byte

	Message *ErrorMessage
}

// ErrorMessage represents a basic error message.
type ErrorMessage struct {
	Code    int32               `json:"code"`
	Message string              `json:"message"`
	Errors  jsoniter.RawMessage `json:"errors"`
}

func NewRestError(req *http.Request, resp *http.Response, body []byte) *RestError {
	var errorMessage ErrorMessage

	_ = jsoniter.Unmarshal(body, errorMessage)

	return &RestError{
		Request:      req,
		Response:     resp,
		ResponseBody: body,
		Message:      &errorMessage,
	}
}

func (r *RestError) Error() string {
	return fmt.Sprintf("%s: %s", r.Response.Status, r.Message.Message)
}
