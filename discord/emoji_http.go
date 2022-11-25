package discord

import (
	"fmt"
	"net/http"
)

func ListGuildEmojis(s *Session, guildID Snowflake) ([]*Emoji, error) {
	endpoint := EndpointGuildEmojis(guildID.String())

	var emojis []*Emoji

	err := s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &emojis)
	if err != nil {
		return nil, fmt.Errorf("failed to list guild emojis: %v", err)
	}

	return emojis, nil
}

func GetGuildEmoji(s *Session, guildID Snowflake, emojiID Snowflake) (*Emoji, error) {
	endpoint := EndpointGuildEmoji(guildID.String(), emojiID.String())

	var emoji *Emoji

	err := s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &emoji)
	if err != nil {
		return nil, fmt.Errorf("failed to get get guild emoji: %v", err)
	}

	return emoji, nil
}

func CreateGuildEmoji(s *Session, guildID Snowflake, emojiParams EmojiParams, reason *string) (*Emoji, error) {
	endpoint := EndpointGuildEmojis(guildID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	var emoji *Emoji

	err := s.Interface.FetchJJ(s, http.MethodPost, endpoint, emojiParams, headers, &emoji)
	if err != nil {
		return nil, fmt.Errorf("failed to create guild emoji: %v", err)
	}

	return emoji, nil
}

func ModifyGuildEmoji(s *Session, guildID Snowflake, emojiID Snowflake, emojiParams EmojiParams, reason *string) (*Emoji, error) {
	endpoint := EndpointGuildEmoji(guildID.String(), emojiID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	var emoji *Emoji

	err := s.Interface.FetchJJ(s, http.MethodPatch, endpoint, emojiParams, headers, &emoji)
	if err != nil {
		return nil, fmt.Errorf("failed to modify guild emoji: %v", err)
	}

	return emoji, nil
}

func DeleteGuildEmoji(s *Session, guildID Snowflake, emojiID Snowflake, reason *string) error {
	endpoint := EndpointGuildEmoji(guildID.String(), emojiID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err := s.Interface.FetchJJ(s, http.MethodDelete, endpoint, nil, headers, nil)
	if err != nil {
		return fmt.Errorf("failed to delete guild emoji: %v", err)
	}

	return nil
}
