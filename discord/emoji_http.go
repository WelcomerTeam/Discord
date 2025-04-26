package discord

import (
	"context"
	"fmt"
	"net/http"
)

func ListGuildEmojis(ctx context.Context, session *Session, guildID Snowflake) ([]Emoji, error) {
	endpoint := EndpointGuildEmojis(guildID.String())

	var emojis []Emoji

	err := session.Interface.FetchJJ(ctx, session, http.MethodGet, endpoint, nil, nil, &emojis)
	if err != nil {
		return nil, fmt.Errorf("failed to list guild emojis: %w", err)
	}

	return emojis, nil
}

func GetGuildEmoji(ctx context.Context, session *Session, guildID, emojiID Snowflake) (*Emoji, error) {
	endpoint := EndpointGuildEmoji(guildID.String(), emojiID.String())

	var emoji *Emoji

	err := session.Interface.FetchJJ(ctx, session, http.MethodGet, endpoint, nil, nil, &emoji)
	if err != nil {
		return nil, fmt.Errorf("failed to get guild emoji: %w", err)
	}

	return emoji, nil
}

func CreateGuildEmoji(ctx context.Context, session *Session, guildID Snowflake, emojiParams EmojiParams, reason *string) (*Emoji, error) {
	endpoint := EndpointGuildEmojis(guildID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	var emoji *Emoji

	err := session.Interface.FetchJJ(ctx, session, http.MethodPost, endpoint, emojiParams, headers, &emoji)
	if err != nil {
		return nil, fmt.Errorf("failed to create guild emoji: %w", err)
	}

	return emoji, nil
}

func ModifyGuildEmoji(ctx context.Context, session *Session, guildID, emojiID Snowflake, emojiParams EmojiParams, reason *string) (*Emoji, error) {
	endpoint := EndpointGuildEmoji(guildID.String(), emojiID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	var emoji *Emoji

	err := session.Interface.FetchJJ(ctx, session, http.MethodPatch, endpoint, emojiParams, headers, &emoji)
	if err != nil {
		return nil, fmt.Errorf("failed to modify guild emoji: %w", err)
	}

	return emoji, nil
}

func DeleteGuildEmoji(ctx context.Context, session *Session, guildID, emojiID Snowflake, reason *string) error {
	endpoint := EndpointGuildEmoji(guildID.String(), emojiID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err := session.Interface.FetchJJ(ctx, session, http.MethodDelete, endpoint, nil, headers, nil)
	if err != nil {
		return fmt.Errorf("failed to delete guild emoji: %w", err)
	}

	return nil
}
