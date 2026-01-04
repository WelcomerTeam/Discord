package discord

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

func CreateWebhook(ctx context.Context, session *Session, channelID Snowflake, webhookParam WebhookParam, reason *string) (*Webhook, error) {
	endpoint := EndpointChannelWebhooks(channelID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	var webhook *Webhook

	err := session.Interface.FetchJJ(ctx, session, http.MethodGet, endpoint, webhookParam, headers, &webhook)
	if err != nil {
		return nil, fmt.Errorf("failed to create webhook: %w", err)
	}

	return webhook, nil
}

func GetChannelWebhooks(ctx context.Context, session *Session, channelID Snowflake) ([]Webhook, error) {
	endpoint := EndpointChannelWebhooks(channelID.String())

	var webhooks []Webhook

	err := session.Interface.FetchJJ(ctx, session, http.MethodGet, endpoint, nil, nil, &webhooks)
	if err != nil {
		return nil, fmt.Errorf("failed to get channel webhooks: %w", err)
	}

	return webhooks, nil
}

func GetGuildWebhooks(ctx context.Context, session *Session, guildID Snowflake) ([]Webhook, error) {
	endpoint := EndpointGuildWebhooks(guildID.String())

	var webhooks []Webhook

	err := session.Interface.FetchJJ(ctx, session, http.MethodGet, endpoint, nil, nil, &webhooks)
	if err != nil {
		return nil, fmt.Errorf("failed to get guild webhooks: %w", err)
	}

	return webhooks, nil
}

func GetWebhook(ctx context.Context, session *Session, webhookID Snowflake) (*Webhook, error) {
	endpoint := EndpointWebhook(webhookID.String())

	var webhook *Webhook

	err := session.Interface.FetchJJ(ctx, session, http.MethodGet, endpoint, nil, nil, &webhook)
	if err != nil {
		return nil, fmt.Errorf("failed to get webhook: %w", err)
	}

	return webhook, nil
}

func GetWebhookWithToken(ctx context.Context, session *Session, webhookID Snowflake, webhookToken string) (*Webhook, error) {
	endpoint := EndpointWebhookToken(webhookID.String(), webhookToken)

	var webhook *Webhook

	err := session.Interface.FetchJJ(ctx, session, http.MethodGet, endpoint, nil, nil, &webhook)
	if err != nil {
		return nil, fmt.Errorf("failed to get webhook with token: %w", err)
	}

	return webhook, nil
}

func ModifyWebhook(ctx context.Context, session *Session, webhookID Snowflake, webhookParam WebhookParam, reason *string) (*Webhook, error) {
	endpoint := EndpointWebhook(webhookID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	var webhook *Webhook

	err := session.Interface.FetchJJ(ctx, session, http.MethodPatch, endpoint, webhookParam, headers, &webhook)
	if err != nil {
		return nil, fmt.Errorf("failed to modify webhook: %w", err)
	}

	return webhook, nil
}

func ModifyWebhookWithToken(ctx context.Context, session *Session, webhookID Snowflake, webhookToken string, webhookParam WebhookParam) (*Webhook, error) {
	endpoint := EndpointWebhookToken(webhookID.String(), webhookToken)

	var webhook *Webhook

	err := session.Interface.FetchJJ(ctx, session, http.MethodPatch, endpoint, webhookParam, nil, &webhook)
	if err != nil {
		return nil, fmt.Errorf("failed to modify webhook with token: %w", err)
	}

	return webhook, nil
}

func DeleteWebhook(ctx context.Context, session *Session, webhookID Snowflake, reason *string) error {
	endpoint := EndpointWebhook(webhookID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err := session.Interface.FetchJJ(ctx, session, http.MethodDelete, endpoint, nil, headers, nil)
	if err != nil {
		return fmt.Errorf("failed to delete webhook: %w", err)
	}

	return nil
}

func DeleteWebhookWithToken(ctx context.Context, session *Session, webhookID Snowflake, webhookToken string) error {
	endpoint := EndpointWebhookToken(webhookID.String(), webhookToken)

	err := session.Interface.FetchJJ(ctx, session, http.MethodDelete, endpoint, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete webhook with token: %w", err)
	}

	return nil
}

func ExecuteWebhook(ctx context.Context, session *Session, webhookID Snowflake, webhookToken string, messageParams WebhookMessageParams, wait bool) (*WebhookMessage, error) {
	endpoint := EndpointWebhookToken(webhookID.String(), webhookToken)

	values := url.Values{}

	if wait {
		values.Set("wait", "true")
	}

	if len(values) > 0 {
		endpoint += "?" + values.Encode()
	}

	var message *WebhookMessage

	var err error

	if len(messageParams.Files) > 0 {
		contentType, body, err := multipartBodyWithJSON(messageParams, messageParams.Files)
		if err != nil {
			return nil, err
		}

		if wait {
			err = session.Interface.FetchBJ(ctx, session, http.MethodPost, endpoint, contentType, body, nil, &message)
		} else {
			err = session.Interface.FetchBJ(ctx, session, http.MethodPost, endpoint, contentType, body, nil, nil)
		}

		if err != nil {
			return nil, fmt.Errorf("failed to execute webhook: %w", err)
		}
	} else {
		if wait {
			err = session.Interface.FetchJJ(ctx, session, http.MethodPost, endpoint, messageParams, nil, &message)
		} else {
			err = session.Interface.FetchJJ(ctx, session, http.MethodPost, endpoint, messageParams, nil, nil)
		}

		if err != nil {
			return nil, fmt.Errorf("failed to execute webhook: %w", err)
		}
	}

	return message, nil
}

func GetWebhookMessage(ctx context.Context, session *Session, webhookID Snowflake, webhookToken string, messageID Snowflake) (*WebhookMessage, error) {
	endpoint := EndpointWebhookMessage(webhookID.String(), webhookToken, messageID.String())

	var message *WebhookMessage

	err := session.Interface.FetchJJ(ctx, session, http.MethodGet, endpoint, nil, nil, &message)
	if err != nil {
		return nil, fmt.Errorf("failed to get webhook message: %w", err)
	}

	return message, nil
}

func EditWebhookMessage(ctx context.Context, session *Session, webhookID Snowflake, webhookToken string, messageID Snowflake, messageParam WebhookMessageParams) (*WebhookMessage, error) {
	endpoint := EndpointWebhookMessage(webhookID.String(), webhookToken, messageID.String())

	var message *WebhookMessage

	var err error

	if len(messageParam.Files) > 0 {
		contentType, body, err := multipartBodyWithJSON(messageParam, messageParam.Files)
		if err != nil {
			return nil, err
		}

		err = session.Interface.FetchBJ(ctx, session, http.MethodPatch, endpoint, contentType, body, nil, &message)
		if err != nil {
			return nil, fmt.Errorf("failed to edit webhook message: %w", err)
		}
	} else {
		err = session.Interface.FetchJJ(ctx, session, http.MethodPatch, endpoint, messageParam, nil, &message)
		if err != nil {
			return nil, fmt.Errorf("failed to edit webhook message: %w", err)
		}
	}

	return message, nil
}

func DeleteWebhookMessage(ctx context.Context, session *Session, webhookID Snowflake, webhookToken string, messageID Snowflake) error {
	endpoint := EndpointWebhookMessage(webhookID.String(), webhookToken, messageID.String())

	err := session.Interface.FetchJJ(ctx, session, http.MethodDelete, endpoint, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete webhook message: %w", err)
	}

	return nil
}
