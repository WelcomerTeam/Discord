package http

import (
	"net/http"

	"github.com/WelcomerTeam/Discord/discord"
	"github.com/WelcomerTeam/Discord/structs"
	"golang.org/x/xerrors"
)

func (s *Session) CreateInteractionResponse(interactionID discord.Snowflake, interactionToken string, interactionResponse structs.InteractionResponse) (interaction *structs.Interaction, err error) {
	endpoint := EndpointInteractionResponse(interactionID.String(), interactionToken)

	err = s.Interface.FetchJJ(s, http.MethodPost, endpoint, interactionResponse, nil, &interaction)
	if err != nil {
		return nil, xerrors.Errorf("Failed to create interaction response: %v", err)
	}

	return
}

func (s *Session) GetOrigionalInteractionResponse(applicationID discord.Snowflake, interactionToken string) (interaction *structs.Interaction, err error) {
	endpoint := EndpointInteractionResponseActions(applicationID.String(), interactionToken)

	err = s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &interaction)
	if err != nil {
		return nil, xerrors.Errorf("Failed to get original interaction response: %v", err)
	}

	return
}

func (s *Session) EditOriginalInteractionResponse(applicationID discord.Snowflake, interactionToken string, messageParams structs.WebhookMessageParams) (interaction *structs.Interaction, err error) {
	endpoint := EndpointInteractionResponseActions(applicationID.String(), interactionToken)

	err = s.Interface.FetchJJ(s, http.MethodPatch, endpoint, messageParams, nil, &interaction)
	if err != nil {
		return nil, xerrors.Errorf("Failed to edit original interaction response: %v", err)
	}

	return
}

func (s *Session) DeleteOriginalInteractionResponse(applicationID discord.Snowflake, interactionToken string) (err error) {
	endpoint := EndpointInteractionResponseActions(applicationID.String(), interactionToken)

	err = s.Interface.FetchJJ(s, http.MethodDelete, endpoint, nil, nil, nil)
	if err != nil {
		return xerrors.Errorf("Failed to create interaction response: %v", err)
	}

	return
}

func (s *Session) CreateFollowupMessage(applicationID discord.Snowflake, interactionToken string, messageParams structs.WebhookMessageParams) (message *structs.Message, err error) {
	endpoint := EndpointFollowupMessage(applicationID.String(), interactionToken)

	err = s.Interface.FetchJJ(s, http.MethodPost, endpoint, messageParams, nil, &message)
	if err != nil {
		return nil, xerrors.Errorf("Failed to create followup message: %v", err)
	}

	return
}

func (s *Session) GetFollowupMessage(applicationID discord.Snowflake, interactionToken string, messageID discord.Snowflake) (message *structs.Message, err error) {
	endpoint := EndpointFollowupMessageActions(applicationID.String(), interactionToken, messageID.String())

	err = s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &message)
	if err != nil {
		return nil, xerrors.Errorf("Failed to get followup message: %v", err)
	}

	return
}

func (s *Session) EditFollowupMessage(applicationID discord.Snowflake, interactionToken string, messageID discord.Snowflake, messageParams structs.WebhookMessageParams) (message *structs.Message, err error) {
	endpoint := EndpointFollowupMessageActions(applicationID.String(), interactionToken, messageID.String())

	err = s.Interface.FetchJJ(s, http.MethodPatch, endpoint, messageParams, nil, &message)
	if err != nil {
		return nil, xerrors.Errorf("Failed to edit followup message: %v", err)
	}

	return
}

func (s *Session) DeleteFollowupMessage(applicationID discord.Snowflake, interactionToken string, messageID discord.Snowflake) (err error) {
	endpoint := EndpointFollowupMessageActions(applicationID.String(), interactionToken, messageID.String())

	err = s.Interface.FetchJJ(s, http.MethodDelete, endpoint, nil, nil, nil)
	if err != nil {
		return xerrors.Errorf("Failed to delete followup message: %v", err)
	}

	return
}
