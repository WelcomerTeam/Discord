package discord

import (
	"net/http"

	"golang.org/x/xerrors"
)

func ListGuildEmojis(s *Session, guildID Snowflake) (emojis []*Emoji, err error) {
	endpoint := EndpointGuildEmojis(guildID.String())

	err = s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &emojis)
	if err != nil {
		return nil, xerrors.Errorf("Failed to list guild emojis: %v", err)
	}

	return
}

func GetGuildEmoji(s *Session, guildID Snowflake, emojiID Snowflake) (emoji *Emoji, err error) {
	endpoint := EndpointGuildEmoji(guildID.String(), emojiID.String())

	err = s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &emoji)
	if err != nil {
		return nil, xerrors.Errorf("Failed to get get guild emoji: %v", err)
	}

	return
}

func CreateGuildEmoji(s *Session, guildID Snowflake, emojiParams EmojiParams, reason *string) (emoji *Emoji, err error) {
	endpoint := EndpointGuildEmojis(guildID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err = s.Interface.FetchJJ(s, http.MethodPost, endpoint, emojiParams, headers, &emoji)
	if err != nil {
		return nil, xerrors.Errorf("Failed to create guild emoji: %v", err)
	}

	return
}

func ModifyGuildEmoji(s *Session, guildID Snowflake, emojiID Snowflake, emojiParams EmojiParams, reason *string) (emoji *Emoji, err error) {
	endpoint := EndpointGuildEmoji(guildID.String(), emojiID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err = s.Interface.FetchJJ(s, http.MethodPatch, endpoint, emojiParams, headers, &emoji)
	if err != nil {
		return nil, xerrors.Errorf("Failed to modify guild emoji: %v", err)
	}

	return
}

func DeleteGuildEmoji(s *Session, guildID Snowflake, emojiID Snowflake, reason *string) (err error) {
	endpoint := EndpointGuildEmoji(guildID.String(), emojiID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err = s.Interface.FetchJJ(s, http.MethodDelete, endpoint, nil, headers, nil)
	if err != nil {
		return xerrors.Errorf("Failed to delete guild emoji: %v", err)
	}

	return
}
