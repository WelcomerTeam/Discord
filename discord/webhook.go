package discord

import (
	"context"
	"encoding/json"
)

// webhook.go represents all structures to create a webhook and interact with it.

// WebhookType is the type of webhook.
type WebhookType uint16

// Webhook type.
const (
	WebhookTypeIncoming WebhookType = iota + 1
	WebhookTypeChannelFollower
)

// Webhook represents a webhook on discord.
type Webhook struct {
	GuildID       *Snowflake  `json:"guild_id,omitempty"`
	ChannelID     *Snowflake  `json:"channel_id,omitempty"`
	User          *User       `json:"user,omitempty"`
	ApplicationID *Snowflake  `json:"application_id,omitempty"`
	Name          string      `json:"name"`
	Avatar        string      `json:"avatar"`
	Token         string      `json:"token"`
	ID            Snowflake   `json:"id"`
	Type          WebhookType `json:"type"`
}

// Delete deletes this webhook.
// reason: The reason for deleting this webhook.
func (w *Webhook) Delete(ctx context.Context, session *Session, reason *string) error {
	return DeleteWebhook(ctx, session, w.ID, reason)
}

// Edit edits this webhook.
// name: The webhooks new default name.
// avatar: bytes representing the webhooks new avatar.
// reason: The reason for editing this webhook.
func (w *Webhook) Edit(ctx context.Context, session *Session, name *string, avatar *[]byte, reason *string) error {
	params := WebhookParam{
		Name: name,
	}

	if avatar != nil {
		avatarBase64, err := bytesToBase64Data(*avatar)
		if err != nil {
			return err
		}

		params.Avatar = &avatarBase64
	}

	var newWebhook *Webhook

	var err error

	if w.Token != "" {
		newWebhook, err = ModifyWebhookWithToken(ctx, session, w.ID, w.Token, params)
	} else {
		newWebhook, err = ModifyWebhook(ctx, session, w.ID, params, reason)
	}

	if err != nil {
		return err
	}

	*w = *newWebhook

	return nil
}

// Send sends a webhook message.
// params: The message parameters to send.
func (w *Webhook) Send(ctx context.Context, session *Session, params WebhookMessageParams, wait bool) (*WebhookMessage, error) {
	return ExecuteWebhook(ctx, session, w.ID, w.Token, params, wait)
}

// EditMessage edits a webhook message.
// messageID: The message id you are editing.
// params: The message parameters used to update the message.
func (w *Webhook) EditMessage(ctx context.Context, session *Session, messageID Snowflake, params WebhookMessageParams) (*WebhookMessage, error) {
	return EditWebhookMessage(ctx, session, w.ID, w.Token, messageID, params)
}

// DeleteMessage deletes a webhook message.
// messageID: The message id you are deleting.
func (w *Webhook) DeleteMessage(ctx context.Context, session *Session, messageID Snowflake) error {
	return DeleteWebhookMessage(ctx, session, w.ID, w.Token, messageID)
}

// WebhookMessage aliases Message to provide webhook specific methods.
type WebhookMessage Message

// Edit edits a webhook message.
// token: The token of the parent webhook.
// params: The message parameters used to update the message.
func (wm *WebhookMessage) Edit(ctx context.Context, session *Session, token string, params WebhookMessageParams) (*WebhookMessage, error) {
	return EditWebhookMessage(ctx, session, *wm.WebhookID, token, wm.ID, params)
}

// Delete deletes a webhook message.
// token: The token of the parent webhook.
func (wm *WebhookMessage) Delete(ctx context.Context, session *Session, token string) error {
	return DeleteWebhookMessage(ctx, session, *wm.WebhookID, token, wm.ID)
}

// WebhookMessage represents the structure for sending a webhook message.
type WebhookMessageParams struct {
	PayloadJSON     *json.RawMessage         `json:"payload_json,omitempty"`
	Content         string                   `json:"content,omitempty"`
	Username        string                   `json:"username,omitempty"`
	AvatarURL       string                   `json:"avatar_url,omitempty"`
	Embeds          []Embed                  `json:"embeds"`
	AllowedMentions []MessageAllowedMentions `json:"allowed_mentions,omitempty"`
	Components      []InteractionComponent   `json:"components"`
	Files           []File                   `json:"-"`
	Attachments     []MessageAttachment      `json:"attachments,omitempty"`
	TTS             bool                     `json:"tts,omitempty"`
}

// WebhookParam represents the data sent to discord to create a webhook.
type WebhookParam struct {
	Name   *string `json:"name,omitempty"`
	Avatar *string `json:"avatar,omitempty"`
}
