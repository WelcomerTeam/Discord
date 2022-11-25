package discord

import (
	"fmt"
	"net/http"
)

func CreateInteractionResponse(s *Session, interactionID Snowflake, interactionToken string, interactionResponse InteractionResponse) error {
	endpoint := EndpointInteractionResponse(interactionID.String(), interactionToken)

	if len(interactionResponse.Data.Files) > 0 {
		contentType, body, err := multipartBodyWithJSON(interactionResponse, interactionResponse.Data.Files)
		if err != nil {
			return err
		}

		err = s.Interface.FetchBJ(s, http.MethodPost, endpoint, contentType, body, nil, nil)
		if err != nil {
			return fmt.Errorf("failed to create interaction response: %v", err)
		}
	} else {
		err := s.Interface.FetchJJ(s, http.MethodPost, endpoint, interactionResponse, nil, nil)
		if err != nil {
			return fmt.Errorf("failed to create interaction response: %v", err)
		}
	}

	return nil
}

func GetoriginalInteractionResponse(s *Session, applicationID Snowflake, interactionToken string) (*Message, error) {
	endpoint := EndpointInteractionResponseActions(applicationID.String(), interactionToken)

	var message *Message

	err := s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &message)
	if err != nil {
		return nil, fmt.Errorf("failed to get original interaction response: %v", err)
	}

	return message, nil
}

func EditOriginalInteractionResponse(s *Session, applicationID Snowflake, interactionToken string, messageParam WebhookMessageParams) (*Message, error) {
	endpoint := EndpointInteractionResponseActions(applicationID.String(), interactionToken)

	var message *Message

	if len(messageParam.Files) > 0 {
		contentType, body, err := multipartBodyWithJSON(messageParam, messageParam.Files)
		if err != nil {
			return nil, err
		}

		err = s.Interface.FetchBJ(s, http.MethodPatch, endpoint, contentType, body, nil, &message)
		if err != nil {
			return nil, fmt.Errorf("failed to edit original interaction response: %v", err)
		}
	} else {
		err := s.Interface.FetchJJ(s, http.MethodPatch, endpoint, messageParam, nil, &message)
		if err != nil {
			return nil, fmt.Errorf("failed to edit original interaction response: %v", err)
		}
	}

	return message, nil
}

func DeleteOriginalInteractionResponse(s *Session, applicationID Snowflake, interactionToken string) error {
	endpoint := EndpointInteractionResponseActions(applicationID.String(), interactionToken)

	err := s.Interface.FetchJJ(s, http.MethodDelete, endpoint, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to create interaction response: %v", err)
	}

	return nil
}

func CreateFollowupMessage(s *Session, applicationID Snowflake, interactionToken string, messageParams WebhookMessageParams) (*Message, error) {
	endpoint := EndpointFollowupMessage(applicationID.String(), interactionToken)

	var message *Message

	if len(messageParams.Files) > 0 {
		contentType, body, err := multipartBodyWithJSON(messageParams, messageParams.Files)
		if err != nil {
			return nil, err
		}

		err = s.Interface.FetchBJ(s, http.MethodPost, endpoint, contentType, body, nil, &message)
		if err != nil {
			return nil, fmt.Errorf("failed to create followup message: %v", err)
		}
	} else {
		err := s.Interface.FetchJJ(s, http.MethodPost, endpoint, messageParams, nil, &message)
		if err != nil {
			return nil, fmt.Errorf("failed to create followup message: %v", err)
		}
	}

	return message, nil
}

func GetFollowupMessage(s *Session, applicationID Snowflake, interactionToken string, messageID Snowflake) (*Message, error) {
	endpoint := EndpointFollowupMessageActions(applicationID.String(), interactionToken, messageID.String())

	var message *Message

	err := s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &message)
	if err != nil {
		return nil, fmt.Errorf("failed to get followup message: %v", err)
	}

	return message, nil
}

func EditFollowupMessage(s *Session, applicationID Snowflake, interactionToken string, messageID Snowflake, messageParams WebhookMessageParams) (*Message, error) {
	endpoint := EndpointFollowupMessageActions(applicationID.String(), interactionToken, messageID.String())

	var message *Message

	if len(messageParams.Files) > 0 {
		contentType, body, err := multipartBodyWithJSON(messageParams, messageParams.Files)
		if err != nil {
			return nil, err
		}

		err = s.Interface.FetchBJ(s, http.MethodPatch, endpoint, contentType, body, nil, &message)
		if err != nil {
			return nil, fmt.Errorf("failed to edit followup message: %v", err)
		}
	} else {
		err := s.Interface.FetchJJ(s, http.MethodPatch, endpoint, messageParams, nil, &message)
		if err != nil {
			return nil, fmt.Errorf("failed to edit followup message: %v", err)
		}
	}

	return message, nil
}

func DeleteFollowupMessage(s *Session, applicationID Snowflake, interactionToken string, messageID Snowflake) error {
	endpoint := EndpointFollowupMessageActions(applicationID.String(), interactionToken, messageID.String())

	err := s.Interface.FetchJJ(s, http.MethodDelete, endpoint, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete followup message: %v", err)
	}

	return nil
}
