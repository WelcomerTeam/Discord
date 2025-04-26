package discord

import (
	"context"
	"fmt"
	"net/http"
)

func CreateInteractionResponse(ctx context.Context, session *Session, interactionID Snowflake, interactionToken string, interactionResponse InteractionResponse) error {
	endpoint := EndpointInteractionResponse(interactionID.String(), interactionToken)

	if len(interactionResponse.Data.Files) > 0 {
		contentType, body, err := multipartBodyWithJSON(interactionResponse, interactionResponse.Data.Files)
		if err != nil {
			return err
		}

		err = session.Interface.FetchBJ(ctx, session, http.MethodPost, endpoint, contentType, body, nil, nil)
		if err != nil {
			return fmt.Errorf("failed to create interaction response: %w", err)
		}
	} else {
		err := session.Interface.FetchJJ(ctx, session, http.MethodPost, endpoint, interactionResponse, nil, nil)
		if err != nil {
			return fmt.Errorf("failed to create interaction response: %w", err)
		}
	}

	return nil
}

func GetOriginalInteractionResponse(ctx context.Context, session *Session, applicationID Snowflake, interactionToken string) (*Message, error) {
	endpoint := EndpointInteractionResponseActions(applicationID.String(), interactionToken)

	var message *Message

	err := session.Interface.FetchJJ(ctx, session, http.MethodGet, endpoint, nil, nil, &message)
	if err != nil {
		return nil, fmt.Errorf("failed to get original interaction response: %w", err)
	}

	return message, nil
}

func EditOriginalInteractionResponse(ctx context.Context, session *Session, applicationID Snowflake, interactionToken string, messageParam WebhookMessageParams) (*Message, error) {
	endpoint := EndpointInteractionResponseActions(applicationID.String(), interactionToken)

	var message *Message

	if len(messageParam.Files) > 0 {
		contentType, body, err := multipartBodyWithJSON(messageParam, messageParam.Files)
		if err != nil {
			return nil, err
		}

		err = session.Interface.FetchBJ(ctx, session, http.MethodPatch, endpoint, contentType, body, nil, &message)
		if err != nil {
			return nil, fmt.Errorf("failed to edit original interaction response: %w", err)
		}
	} else {
		err := session.Interface.FetchJJ(ctx, session, http.MethodPatch, endpoint, messageParam, nil, &message)
		if err != nil {
			return nil, fmt.Errorf("failed to edit original interaction response: %w", err)
		}
	}

	return message, nil
}

func DeleteOriginalInteractionResponse(ctx context.Context, session *Session, applicationID Snowflake, interactionToken string) error {
	endpoint := EndpointInteractionResponseActions(applicationID.String(), interactionToken)

	err := session.Interface.FetchJJ(ctx, session, http.MethodDelete, endpoint, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to create interaction response: %w", err)
	}

	return nil
}

func CreateFollowupMessage(ctx context.Context, session *Session, applicationID Snowflake, interactionToken string, messageParams WebhookMessageParams) (*Message, error) {
	endpoint := EndpointFollowupMessage(applicationID.String(), interactionToken)

	var message *Message

	if len(messageParams.Files) > 0 {
		contentType, body, err := multipartBodyWithJSON(messageParams, messageParams.Files)
		if err != nil {
			return nil, err
		}

		err = session.Interface.FetchBJ(ctx, session, http.MethodPost, endpoint, contentType, body, nil, &message)
		if err != nil {
			return nil, fmt.Errorf("failed to create followup message: %w", err)
		}
	} else {
		err := session.Interface.FetchJJ(ctx, session, http.MethodPost, endpoint, messageParams, nil, &message)
		if err != nil {
			return nil, fmt.Errorf("failed to create followup message: %w", err)
		}
	}

	return message, nil
}

func GetFollowupMessage(ctx context.Context, session *Session, applicationID Snowflake, interactionToken string, messageID Snowflake) (*Message, error) {
	endpoint := EndpointFollowupMessageActions(applicationID.String(), interactionToken, messageID.String())

	var message *Message

	err := session.Interface.FetchJJ(ctx, session, http.MethodGet, endpoint, nil, nil, &message)
	if err != nil {
		return nil, fmt.Errorf("failed to get followup message: %w", err)
	}

	return message, nil
}

func EditFollowupMessage(ctx context.Context, session *Session, applicationID Snowflake, interactionToken string, messageID Snowflake, messageParams WebhookMessageParams) (*Message, error) {
	endpoint := EndpointFollowupMessageActions(applicationID.String(), interactionToken, messageID.String())

	var message *Message

	if len(messageParams.Files) > 0 {
		contentType, body, err := multipartBodyWithJSON(messageParams, messageParams.Files)
		if err != nil {
			return nil, err
		}

		err = session.Interface.FetchBJ(ctx, session, http.MethodPatch, endpoint, contentType, body, nil, &message)
		if err != nil {
			return nil, fmt.Errorf("failed to edit followup message: %w", err)
		}
	} else {
		err := session.Interface.FetchJJ(ctx, session, http.MethodPatch, endpoint, messageParams, nil, &message)
		if err != nil {
			return nil, fmt.Errorf("failed to edit followup message: %w", err)
		}
	}

	return message, nil
}

func DeleteFollowupMessage(ctx context.Context, session *Session, applicationID Snowflake, interactionToken string, messageID Snowflake) error {
	endpoint := EndpointFollowupMessageActions(applicationID.String(), interactionToken, messageID.String())

	err := session.Interface.FetchJJ(ctx, session, http.MethodDelete, endpoint, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete followup message: %w", err)
	}

	return nil
}
