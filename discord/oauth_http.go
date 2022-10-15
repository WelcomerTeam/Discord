package discord

import (
	"fmt"
	"net/http"
)

func GetCurrentBotApplicationInformation(s *Session) (application *Application, err error) {
	endpoint := EndpointOAuth2Application("@me")

	err = s.Interface.FetchBJ(s, http.MethodGet, endpoint, "", nil, nil, &application)
	if err != nil {
		return nil, fmt.Errorf("Failed to get current bot application information: %v", err)
	}

	return
}

func GetCurrentAuthorizationInformation(s *Session) (authorizationInformation *AuthorizationInformation, err error) {
	endpoint := EndpointOAuth2Me

	err = s.Interface.FetchBJ(s, http.MethodGet, endpoint, "", nil, nil, &authorizationInformation)
	if err != nil {
		return nil, fmt.Errorf("Failed to get current authorization information: %v", err)
	}

	return
}
