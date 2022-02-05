package discord

import (
	"net/http"

	"golang.org/x/xerrors"
)

func CreateInteractionResponse(s *Session, interactionID Snowflake, interactionToken string, interactionResponse InteractionResponse) (err error) {
	endpoint := EndpointInteractionResponse(interactionID.String(), interactionToken)

	// TODO: Handle file uploads

	err = s.Interface.FetchJJ(s, http.MethodPost, endpoint, interactionResponse, nil, nil)
	if err != nil {
		return xerrors.Errorf("Failed to create interaction response: %v", err)
	}

	return
}

func GetoriginalInteractionResponse(s *Session, applicationID Snowflake, interactionToken string) (message *Message, err error) {
	endpoint := EndpointInteractionResponseActions(applicationID.String(), interactionToken)

	err = s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &message)
	if err != nil {
		return nil, xerrors.Errorf("Failed to get original interaction response: %v", err)
	}

	return
}

func EditOriginalInteractionResponse(s *Session, applicationID Snowflake, interactionToken string, messageParams WebhookMessageParams) (message *Message, err error) {
	endpoint := EndpointInteractionResponseActions(applicationID.String(), interactionToken)

	// TODO: Handle file uploads

	err = s.Interface.FetchJJ(s, http.MethodPatch, endpoint, messageParams, nil, &message)
	if err != nil {
		return nil, xerrors.Errorf("Failed to edit original interaction response: %v", err)
	}

	return
}

func DeleteOriginalInteractionResponse(s *Session, applicationID Snowflake, interactionToken string) (err error) {
	endpoint := EndpointInteractionResponseActions(applicationID.String(), interactionToken)

	err = s.Interface.FetchJJ(s, http.MethodDelete, endpoint, nil, nil, nil)
	if err != nil {
		return xerrors.Errorf("Failed to create interaction response: %v", err)
	}

	return
}

func CreateFollowupMessage(s *Session, applicationID Snowflake, interactionToken string, messageParams WebhookMessageParams) (message *Message, err error) {
	endpoint := EndpointFollowupMessage(applicationID.String(), interactionToken)

	// TODO: Handle file uploads

	err = s.Interface.FetchJJ(s, http.MethodPost, endpoint, messageParams, nil, &message)
	if err != nil {
		return nil, xerrors.Errorf("Failed to create followup message: %v", err)
	}

	return
}

func GetFollowupMessage(s *Session, applicationID Snowflake, interactionToken string, messageID Snowflake) (message *Message, err error) {
	endpoint := EndpointFollowupMessageActions(applicationID.String(), interactionToken, messageID.String())

	err = s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &message)
	if err != nil {
		return nil, xerrors.Errorf("Failed to get followup message: %v", err)
	}

	return
}

func EditFollowupMessage(s *Session, applicationID Snowflake, interactionToken string, messageID Snowflake, messageParams WebhookMessageParams) (message *Message, err error) {
	endpoint := EndpointFollowupMessageActions(applicationID.String(), interactionToken, messageID.String())

	// TODO: Handle file uploads

	err = s.Interface.FetchJJ(s, http.MethodPatch, endpoint, messageParams, nil, &message)
	if err != nil {
		return nil, xerrors.Errorf("Failed to edit followup message: %v", err)
	}

	return
}

func DeleteFollowupMessage(s *Session, applicationID Snowflake, interactionToken string, messageID Snowflake) (err error) {
	endpoint := EndpointFollowupMessageActions(applicationID.String(), interactionToken, messageID.String())

	err = s.Interface.FetchJJ(s, http.MethodDelete, endpoint, nil, nil, nil)
	if err != nil {
		return xerrors.Errorf("Failed to delete followup message: %v", err)
	}

	return
}
