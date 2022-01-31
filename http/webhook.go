package http

import (
	"context"
	"net/http"

	"github.com/WelcomerTeam/Discord/discord"
	"github.com/WelcomerTeam/Discord/structs"
	"golang.org/x/xerrors"
)

func (s *Session) CreateWebhook(ctx context.Context, channelID discord.Snowflake, webhookParam structs.WebhookParam, reason *string) (webhook *structs.Webhook, err error) {
	endpoint := EndpointChannelWebhooks(channelID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(structs.AuditLogReasonHeader, *reason)
	}

	err = s.Interface.FetchJJ(ctx, http.MethodGet, endpoint, webhookParam, headers, &webhook)
	if err != nil {
		return nil, xerrors.Errorf("Failed to create webhook: %v", err)
	}

	return
}

func (s *Session) GetChannelWebhooks(ctx context.Context, channelID discord.Snowflake) (webhooks []*structs.Webhook, err error) {
	endpoint := EndpointChannelWebhooks(channelID.String())

	err = s.Interface.FetchJJ(ctx, http.MethodGet, endpoint, nil, nil, &webhooks)
	if err != nil {
		return nil, xerrors.Errorf("Failed to get channel webhooks: %v", err)
	}

	return
}

func (s *Session) GetGuildWebhooks(ctx context.Context, guildID discord.Snowflake) (webhooks []*structs.Webhook, err error) {
	endpoint := EndpointGuildWebhooks(guildID.String())

	err = s.Interface.FetchJJ(ctx, http.MethodGet, endpoint, nil, nil, &webhooks)
	if err != nil {
		return nil, xerrors.Errorf("Failed to get guild webhooks: %v", err)
	}

	return
}

func (s *Session) GetWebhook(ctx context.Context, webhookID discord.Snowflake) (webhook *structs.Webhook, err error) {
	endpoint := EndpointWebhook(webhookID.String())

	err = s.Interface.FetchJJ(ctx, http.MethodGet, endpoint, nil, nil, &webhook)
	if err != nil {
		return nil, xerrors.Errorf("Failed to get webhook: %v", err)
	}

	return
}

func (s *Session) GetWebhookWithToken(ctx context.Context, webhookID discord.Snowflake, webhookToken string) (webhook *structs.Webhook, err error) {
	endpoint := EndpointWebhookToken(webhookID.String(), webhookToken)

	err = s.Interface.FetchJJ(ctx, http.MethodGet, endpoint, nil, nil, &webhook)
	if err != nil {
		return nil, xerrors.Errorf("Failed to get webhook with token: %v", err)
	}

	return
}

func (s *Session) ModifyWebhook(ctx context.Context, webhookID discord.Snowflake, webhookParam structs.WebhookParam, reason *string) (webhook *structs.Webhook, err error) {
	endpoint := EndpointWebhook(webhookID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(structs.AuditLogReasonHeader, *reason)
	}

	err = s.Interface.FetchJJ(ctx, http.MethodPatch, endpoint, webhookParam, headers, &webhook)
	if err != nil {
		return nil, xerrors.Errorf("Failed to modify webhook: %v", err)
	}

	return
}

func (s *Session) ModifyWebhookWithToken(ctx context.Context, webhookID discord.Snowflake, webhookToken string, webhookParam structs.WebhookParam) (webhook *structs.Webhook, err error) {
	endpoint := EndpointWebhookToken(webhookID.String(), webhookToken)

	err = s.Interface.FetchJJ(ctx, http.MethodPatch, endpoint, webhookParam, nil, &webhook)
	if err != nil {
		return nil, xerrors.Errorf("Failed to modify webhook with token: %v", err)
	}

	return
}

func (s *Session) DeleteWebhook(ctx context.Context, webhookID discord.Snowflake, reason *string) (err error) {
	endpoint := EndpointWebhook(webhookID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(structs.AuditLogReasonHeader, *reason)
	}

	err = s.Interface.FetchJJ(ctx, http.MethodDelete, endpoint, nil, headers, nil)
	if err != nil {
		return xerrors.Errorf("Failed to delete webhook: %v", err)
	}

	return
}

func (s *Session) DeleteWebhookWithToken(ctx context.Context, webhookID discord.Snowflake, webhookToken string) (err error) {
	endpoint := EndpointWebhookToken(webhookID.String(), webhookToken)

	err = s.Interface.FetchJJ(ctx, http.MethodDelete, endpoint, nil, nil, nil)
	if err != nil {
		return xerrors.Errorf("Failed to delete webhook with token: %v", err)
	}

	return
}

func (s *Session) ExecuteWebhook(ctx context.Context, webhookID discord.Snowflake, webhookToken string, messageParam structs.WebhookMessageParams) (message *structs.Message, err error) {
	endpoint := EndpointWebhookToken(webhookID.String(), webhookToken)

	err = s.Interface.FetchJJ(ctx, http.MethodPost, endpoint, messageParam, nil, &message)
	if err != nil {
		return nil, xerrors.Errorf("Failed to execute webhook: %v", err)
	}

	return
}

func (s *Session) GetWebhookMessage(ctx context.Context, webhookID discord.Snowflake, webhookToken string, messageID discord.Snowflake) (message *structs.Message, err error) {
	endpoint := EndpointWebhookMessage(webhookID.String(), webhookToken, messageID.String())

	err = s.Interface.FetchJJ(ctx, http.MethodGet, endpoint, nil, nil, &message)
	if err != nil {
		return nil, xerrors.Errorf("Failed to get webhook message: %v", err)
	}

	return
}

func (s *Session) EditWebhookMessage(ctx context.Context, webhookID discord.Snowflake, webhookToken string, messageID discord.Snowflake, messageParam structs.WebhookMessageParams) (message *structs.Message, err error) {
	endpoint := EndpointWebhookMessage(webhookID.String(), webhookToken, messageID.String())

	err = s.Interface.FetchJJ(ctx, http.MethodPatch, endpoint, messageParam, nil, &message)
	if err != nil {
		return nil, xerrors.Errorf("Failed to edit webhook message: %v", err)
	}

	return
}

func (s *Session) DeleteWebhookMessage(ctx context.Context, webhookID discord.Snowflake, webhookToken string, messageID discord.Snowflake) (err error) {
	endpoint := EndpointWebhookMessage(webhookID.String(), webhookToken, messageID.String())

	err = s.Interface.FetchJJ(ctx, http.MethodDelete, endpoint, nil, nil, nil)
	if err != nil {
		return xerrors.Errorf("Failed to delete webhook message: %v", err)
	}

	return
}
