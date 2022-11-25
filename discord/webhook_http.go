package discord

import (
	"fmt"
	"net/http"
	"net/url"
)

func CreateWebhook(s *Session, channelID Snowflake, webhookParam WebhookParam, reason *string) (*Webhook, error) {
	endpoint := EndpointChannelWebhooks(channelID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	var webhook *Webhook

	err := s.Interface.FetchJJ(s, http.MethodGet, endpoint, webhookParam, headers, &webhook)
	if err != nil {
		return nil, fmt.Errorf("failed to create webhook: %v", err)
	}

	return webhook, nil
}

func GetChannelWebhooks(s *Session, channelID Snowflake) ([]*Webhook, error) {
	endpoint := EndpointChannelWebhooks(channelID.String())

	var webhooks []*Webhook

	err := s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &webhooks)
	if err != nil {
		return nil, fmt.Errorf("failed to get channel webhooks: %v", err)
	}

	return webhooks, nil
}

func GetGuildWebhooks(s *Session, guildID Snowflake) ([]*Webhook, error) {
	endpoint := EndpointGuildWebhooks(guildID.String())

	var webhooks []*Webhook

	err := s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &webhooks)
	if err != nil {
		return nil, fmt.Errorf("failed to get guild webhooks: %v", err)
	}

	return webhooks, nil
}

func GetWebhook(s *Session, webhookID Snowflake) (*Webhook, error) {
	endpoint := EndpointWebhook(webhookID.String())

	var webhook *Webhook

	err := s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &webhook)
	if err != nil {
		return nil, fmt.Errorf("failed to get webhook: %v", err)
	}

	return webhook, nil
}

func GetWebhookWithToken(s *Session, webhookID Snowflake, webhookToken string) (*Webhook, error) {
	endpoint := EndpointWebhookToken(webhookID.String(), webhookToken)

	var webhook *Webhook

	err := s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &webhook)
	if err != nil {
		return nil, fmt.Errorf("failed to get webhook with token: %v", err)
	}

	return webhook, nil
}

func ModifyWebhook(s *Session, webhookID Snowflake, webhookParam WebhookParam, reason *string) (*Webhook, error) {
	endpoint := EndpointWebhook(webhookID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	var webhook *Webhook

	err := s.Interface.FetchJJ(s, http.MethodPatch, endpoint, webhookParam, headers, &webhook)
	if err != nil {
		return nil, fmt.Errorf("failed to modify webhook: %v", err)
	}

	return webhook, nil
}

func ModifyWebhookWithToken(s *Session, webhookID Snowflake, webhookToken string, webhookParam WebhookParam) (*Webhook, error) {
	endpoint := EndpointWebhookToken(webhookID.String(), webhookToken)

	var webhook *Webhook

	err := s.Interface.FetchJJ(s, http.MethodPatch, endpoint, webhookParam, nil, &webhook)
	if err != nil {
		return nil, fmt.Errorf("failed to modify webhook with token: %v", err)
	}

	return webhook, nil
}

func DeleteWebhook(s *Session, webhookID Snowflake, reason *string) error {
	endpoint := EndpointWebhook(webhookID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err := s.Interface.FetchJJ(s, http.MethodDelete, endpoint, nil, headers, nil)
	if err != nil {
		return fmt.Errorf("failed to delete webhook: %v", err)
	}

	return nil
}

func DeleteWebhookWithToken(s *Session, webhookID Snowflake, webhookToken string) error {
	endpoint := EndpointWebhookToken(webhookID.String(), webhookToken)

	err := s.Interface.FetchJJ(s, http.MethodDelete, endpoint, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete webhook with token: %v", err)
	}

	return nil
}

func ExecuteWebhook(s *Session, webhookID Snowflake, webhookToken string, messageParams WebhookMessageParams, wait bool) (*WebhookMessage, error) {
	endpoint := EndpointWebhookToken(webhookID.String(), webhookToken)

	var values url.Values

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
			err = s.Interface.FetchBJ(s, http.MethodPost, endpoint, contentType, body, nil, &message)
		} else {
			err = s.Interface.FetchBJ(s, http.MethodPost, endpoint, contentType, body, nil, nil)
		}

		if err != nil {
			return nil, fmt.Errorf("failed to execute webhook: %v", err)
		}
	} else {
		if wait {
			err = s.Interface.FetchJJ(s, http.MethodPost, endpoint, messageParams, nil, &message)
		} else {
			err = s.Interface.FetchJJ(s, http.MethodPost, endpoint, messageParams, nil, nil)
		}

		if err != nil {
			return nil, fmt.Errorf("failed to execute webhook: %v", err)
		}
	}

	return message, nil
}

func GetWebhookMessage(s *Session, webhookID Snowflake, webhookToken string, messageID Snowflake) (*WebhookMessage, error) {
	endpoint := EndpointWebhookMessage(webhookID.String(), webhookToken, messageID.String())

	var message *WebhookMessage

	err := s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &message)
	if err != nil {
		return nil, fmt.Errorf("failed to get webhook message: %v", err)
	}

	return message, nil
}

func EditWebhookMessage(s *Session, webhookID Snowflake, webhookToken string, messageID Snowflake, messageParam WebhookMessageParams) (*WebhookMessage, error) {
	endpoint := EndpointWebhookMessage(webhookID.String(), webhookToken, messageID.String())

	var message *WebhookMessage
	var err error

	if len(messageParam.Files) > 0 {
		contentType, body, err := multipartBodyWithJSON(messageParam, messageParam.Files)
		if err != nil {
			return nil, err
		}

		err = s.Interface.FetchBJ(s, http.MethodPatch, endpoint, contentType, body, nil, &message)
		if err != nil {
			return nil, fmt.Errorf("failed to edit webhook message: %v", err)
		}

	} else {
		err = s.Interface.FetchJJ(s, http.MethodPatch, endpoint, messageParam, nil, &message)
		if err != nil {
			return nil, fmt.Errorf("failed to edit webhook message: %v", err)
		}
	}

	return message, nil
}

func DeleteWebhookMessage(s *Session, webhookID Snowflake, webhookToken string, messageID Snowflake) error {
	endpoint := EndpointWebhookMessage(webhookID.String(), webhookToken, messageID.String())

	err := s.Interface.FetchJJ(s, http.MethodDelete, endpoint, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete webhook message: %v", err)
	}

	return nil
}
