package discord

import (
	"context"
	"fmt"
	"net/http"
)

func GetCurrentBotApplicationInformation(ctx context.Context, s *Session) (*Application, error) {
	endpoint := EndpointOAuth2Application("@me")

	var application *Application

	err := s.Interface.FetchBJ(ctx, s, http.MethodGet, endpoint, "", nil, nil, &application)
	if err != nil {
		return nil, fmt.Errorf("failed to get current bot application information: %w", err)
	}

	return application, nil
}

func GetCurrentAuthorizationInformation(ctx context.Context, s *Session) (*AuthorizationInformation, error) {
	endpoint := EndpointOAuth2Me

	var authorizationInformation *AuthorizationInformation

	err := s.Interface.FetchBJ(ctx, s, http.MethodGet, endpoint, "", nil, nil, &authorizationInformation)
	if err != nil {
		return nil, fmt.Errorf("failed to get current authorization information: %w", err)
	}

	return authorizationInformation, nil
}
