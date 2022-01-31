package http

import (
	"net/http"

	"github.com/WelcomerTeam/Discord/discord"
	"github.com/WelcomerTeam/Discord/structs"
	"golang.org/x/xerrors"
)

func (s *Session) CreateWebhook(channelID discord.Snowflake, webhookParam structs.WebhookParam, reason *string) (webhook *structs.Webhook, err error) {
	endpoint := EndpointChannelWebhooks(channelID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(structs.AuditLogReasonHeader, *reason)
	}

	err = s.Interface.FetchJJ(s, http.MethodGet, endpoint, webhookParam, headers, &webhook)
	if err != nil {
		return nil, xerrors.Errorf("Failed to create webhook: %v", err)
	}

	return
}

func (s *Session) GetChannelWebhooks(channelID discord.Snowflake) (webhooks []*structs.Webhook, err error) {
	endpoint := EndpointChannelWebhooks(channelID.String())

	err = s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &webhooks)
	if err != nil {
		return nil, xerrors.Errorf("Failed to get channel webhooks: %v", err)
	}

	return
}

func (s *Session) GetGuildWebhooks(guildID discord.Snowflake) (webhooks []*structs.Webhook, err error) {
	endpoint := EndpointGuildWebhooks(guildID.String())

	err = s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &webhooks)
	if err != nil {
		return nil, xerrors.Errorf("Failed to get guild webhooks: %v", err)
	}

	return
}

func (s *Session) GetWebhook(webhookID discord.Snowflake) (webhook *structs.Webhook, err error) {
	endpoint := EndpointWebhook(webhookID.String())

	err = s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &webhook)
	if err != nil {
		return nil, xerrors.Errorf("Failed to get webhook: %v", err)
	}

	return
}

func (s *Session) GetWebhookWithToken(webhookID discord.Snowflake, webhookToken string) (webhook *structs.Webhook, err error) {
	endpoint := EndpointWebhookToken(webhookID.String(), webhookToken)

	err = s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &webhook)
	if err != nil {
		return nil, xerrors.Errorf("Failed to get webhook with token: %v", err)
	}

	return
}

func (s *Session) ModifyWebhook(webhookID discord.Snowflake, webhookParam structs.WebhookParam, reason *string) (webhook *structs.Webhook, err error) {
	endpoint := EndpointWebhook(webhookID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(structs.AuditLogReasonHeader, *reason)
	}

	err = s.Interface.FetchJJ(s, http.MethodPatch, endpoint, webhookParam, headers, &webhook)
	if err != nil {
		return nil, xerrors.Errorf("Failed to modify webhook: %v", err)
	}

	return
}

func (s *Session) ModifyWebhookWithToken(webhookID discord.Snowflake, webhookToken string, webhookParam structs.WebhookParam) (webhook *structs.Webhook, err error) {
	endpoint := EndpointWebhookToken(webhookID.String(), webhookToken)

	err = s.Interface.FetchJJ(s, http.MethodPatch, endpoint, webhookParam, nil, &webhook)
	if err != nil {
		return nil, xerrors.Errorf("Failed to modify webhook with token: %v", err)
	}

	return
}

func (s *Session) DeleteWebhook(webhookID discord.Snowflake, reason *string) (err error) {
	endpoint := EndpointWebhook(webhookID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(structs.AuditLogReasonHeader, *reason)
	}

	err = s.Interface.FetchJJ(s, http.MethodDelete, endpoint, nil, headers, nil)
	if err != nil {
		return xerrors.Errorf("Failed to delete webhook: %v", err)
	}

	return
}

func (s *Session) DeleteWebhookWithToken(webhookID discord.Snowflake, webhookToken string) (err error) {
	endpoint := EndpointWebhookToken(webhookID.String(), webhookToken)

	err = s.Interface.FetchJJ(s, http.MethodDelete, endpoint, nil, nil, nil)
	if err != nil {
		return xerrors.Errorf("Failed to delete webhook with token: %v", err)
	}

	return
}

func (s *Session) ExecuteWebhook(webhookID discord.Snowflake, webhookToken string, messageParam structs.WebhookMessageParams) (message *structs.Message, err error) {
	endpoint := EndpointWebhookToken(webhookID.String(), webhookToken)

	err = s.Interface.FetchJJ(s, http.MethodPost, endpoint, messageParam, nil, &message)
	if err != nil {
		return nil, xerrors.Errorf("Failed to execute webhook: %v", err)
	}

	return
}

func (s *Session) GetWebhookMessage(webhookID discord.Snowflake, webhookToken string, messageID discord.Snowflake) (message *structs.Message, err error) {
	endpoint := EndpointWebhookMessage(webhookID.String(), webhookToken, messageID.String())

	err = s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &message)
	if err != nil {
		return nil, xerrors.Errorf("Failed to get webhook message: %v", err)
	}

	return
}

func (s *Session) EditWebhookMessage(webhookID discord.Snowflake, webhookToken string, messageID discord.Snowflake, messageParam structs.WebhookMessageParams) (message *structs.Message, err error) {
	endpoint := EndpointWebhookMessage(webhookID.String(), webhookToken, messageID.String())

	err = s.Interface.FetchJJ(s, http.MethodPatch, endpoint, messageParam, nil, &message)
	if err != nil {
		return nil, xerrors.Errorf("Failed to edit webhook message: %v", err)
	}

	return
}

func (s *Session) DeleteWebhookMessage(webhookID discord.Snowflake, webhookToken string, messageID discord.Snowflake) (err error) {
	endpoint := EndpointWebhookMessage(webhookID.String(), webhookToken, messageID.String())

	err = s.Interface.FetchJJ(s, http.MethodDelete, endpoint, nil, nil, nil)
	if err != nil {
		return xerrors.Errorf("Failed to delete webhook message: %v", err)
	}

	return
}
