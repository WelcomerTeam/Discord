package discord

import (
	"golang.org/x/xerrors"
	"net/http"
)

func CreateWebhook(s *Session, channelID Snowflake, webhookParam WebhookParam, reason *string) (webhook *Webhook, err error) {
	endpoint := EndpointChannelWebhooks(channelID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err = s.Interface.FetchJJ(s, http.MethodGet, endpoint, webhookParam, headers, &webhook)
	if err != nil {
		return nil, xerrors.Errorf("Failed to create webhook: %v", err)
	}

	return
}

func GetChannelWebhooks(s *Session, channelID Snowflake) (webhooks []*Webhook, err error) {
	endpoint := EndpointChannelWebhooks(channelID.String())

	err = s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &webhooks)
	if err != nil {
		return nil, xerrors.Errorf("Failed to get channel webhooks: %v", err)
	}

	return
}

func GetGuildWebhooks(s *Session, guildID Snowflake) (webhooks []*Webhook, err error) {
	endpoint := EndpointGuildWebhooks(guildID.String())

	err = s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &webhooks)
	if err != nil {
		return nil, xerrors.Errorf("Failed to get guild webhooks: %v", err)
	}

	return
}

func GetWebhook(s *Session, webhookID Snowflake) (webhook *Webhook, err error) {
	endpoint := EndpointWebhook(webhookID.String())

	err = s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &webhook)
	if err != nil {
		return nil, xerrors.Errorf("Failed to get webhook: %v", err)
	}

	return
}

func GetWebhookWithToken(s *Session, webhookID Snowflake, webhookToken string) (webhook *Webhook, err error) {
	endpoint := EndpointWebhookToken(webhookID.String(), webhookToken)

	err = s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &webhook)
	if err != nil {
		return nil, xerrors.Errorf("Failed to get webhook with token: %v", err)
	}

	return
}

func ModifyWebhook(s *Session, webhookID Snowflake, webhookParam WebhookParam, reason *string) (webhook *Webhook, err error) {
	endpoint := EndpointWebhook(webhookID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err = s.Interface.FetchJJ(s, http.MethodPatch, endpoint, webhookParam, headers, &webhook)
	if err != nil {
		return nil, xerrors.Errorf("Failed to modify webhook: %v", err)
	}

	return
}

func ModifyWebhookWithToken(s *Session, webhookID Snowflake, webhookToken string, webhookParam WebhookParam) (webhook *Webhook, err error) {
	endpoint := EndpointWebhookToken(webhookID.String(), webhookToken)

	err = s.Interface.FetchJJ(s, http.MethodPatch, endpoint, webhookParam, nil, &webhook)
	if err != nil {
		return nil, xerrors.Errorf("Failed to modify webhook with token: %v", err)
	}

	return
}

func DeleteWebhook(s *Session, webhookID Snowflake, reason *string) (err error) {
	endpoint := EndpointWebhook(webhookID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err = s.Interface.FetchJJ(s, http.MethodDelete, endpoint, nil, headers, nil)
	if err != nil {
		return xerrors.Errorf("Failed to delete webhook: %v", err)
	}

	return
}

func DeleteWebhookWithToken(s *Session, webhookID Snowflake, webhookToken string) (err error) {
	endpoint := EndpointWebhookToken(webhookID.String(), webhookToken)

	err = s.Interface.FetchJJ(s, http.MethodDelete, endpoint, nil, nil, nil)
	if err != nil {
		return xerrors.Errorf("Failed to delete webhook with token: %v", err)
	}

	return
}

func ExecuteWebhook(s *Session, webhookID Snowflake, webhookToken string, messageParams WebhookMessageParams) (message *WebhookMessage, err error) {
	endpoint := EndpointWebhookToken(webhookID.String(), webhookToken)

	if len(messageParams.Files) > 0 {
		contentType, body, err := multipartBodyWithJSON(messageParams, messageParams.Files)
		if err != nil {
			return nil, err
		}

		err = s.Interface.FetchBJ(s, http.MethodPost, endpoint, contentType, body, nil, &message)
		if err != nil {
			return nil, xerrors.Errorf("Failed to execute webhook: %v", err)
		}
	} else {
		err = s.Interface.FetchJJ(s, http.MethodPost, endpoint, messageParams, nil, &message)
		if err != nil {
			return nil, xerrors.Errorf("Failed to execute webhook: %v", err)
		}
	}

	return
}

func GetWebhookMessage(s *Session, webhookID Snowflake, webhookToken string, messageID Snowflake) (message *WebhookMessage, err error) {
	endpoint := EndpointWebhookMessage(webhookID.String(), webhookToken, messageID.String())

	err = s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &message)
	if err != nil {
		return nil, xerrors.Errorf("Failed to get webhook message: %v", err)
	}

	return
}

func EditWebhookMessage(s *Session, webhookID Snowflake, webhookToken string, messageID Snowflake, messageParam WebhookMessageParams) (message *WebhookMessage, err error) {
	endpoint := EndpointWebhookMessage(webhookID.String(), webhookToken, messageID.String())

	if len(messageParam.Files) > 0 {
		contentType, body, err := multipartBodyWithJSON(messageParam, messageParam.Files)
		if err != nil {
			return nil, err
		}

		err = s.Interface.FetchBJ(s, http.MethodPatch, endpoint, contentType, body, nil, &message)
		if err != nil {
			return nil, xerrors.Errorf("Failed to edit webhook message: %v", err)
		}

	} else {
		err = s.Interface.FetchJJ(s, http.MethodPatch, endpoint, messageParam, nil, &message)
		if err != nil {
			return nil, xerrors.Errorf("Failed to edit webhook message: %v", err)
		}
	}

	return
}

func DeleteWebhookMessage(s *Session, webhookID Snowflake, webhookToken string, messageID Snowflake) (err error) {
	endpoint := EndpointWebhookMessage(webhookID.String(), webhookToken, messageID.String())

	err = s.Interface.FetchJJ(s, http.MethodDelete, endpoint, nil, nil, nil)
	if err != nil {
		return xerrors.Errorf("Failed to delete webhook message: %v", err)
	}

	return
}
