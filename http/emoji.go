package http

import (
	"net/http"

	"github.com/WelcomerTeam/Discord/discord"
	"github.com/WelcomerTeam/Discord/structs"
	"golang.org/x/xerrors"
)

func (s *Session) ListGuildEmojis(guildID discord.Snowflake) (emojis []*structs.Emoji, err error) {
	endpoint := EndpointGuildEmojis(guildID.String())

	err = s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &emojis)
	if err != nil {
		return nil, xerrors.Errorf("Failed to list guild emojis: %v", err)
	}

	return
}

func (s *Session) GetGuildEmoji(guildID discord.Snowflake, emojiID discord.Snowflake) (emoji *structs.Emoji, err error) {
	endpoint := EndpointGuildEmoji(guildID.String(), emojiID.String())

	err = s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &emoji)
	if err != nil {
		return nil, xerrors.Errorf("Failed to get get guild emoji: %v", err)
	}

	return
}

func (s *Session) CreateGuildEmoji(guildID discord.Snowflake, emojiParams structs.EmojiParams, reason *string) (emoji *structs.Emoji, err error) {
	endpoint := EndpointGuildEmojis(guildID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(structs.AuditLogReasonHeader, *reason)
	}

	err = s.Interface.FetchJJ(s, http.MethodPost, endpoint, emojiParams, headers, &emoji)
	if err != nil {
		return nil, xerrors.Errorf("Failed to create guild emoji: %v", err)
	}

	return
}

func (s *Session) ModifyGuildEmoji(guildID discord.Snowflake, emojiID discord.Snowflake, emojiParams structs.EmojiParams, reason *string) (emoji *structs.Emoji, err error) {
	endpoint := EndpointGuildEmoji(guildID.String(), emojiID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(structs.AuditLogReasonHeader, *reason)
	}

	err = s.Interface.FetchJJ(s, http.MethodPatch, endpoint, emojiParams, headers, &emoji)
	if err != nil {
		return nil, xerrors.Errorf("Failed to modify guild emoji: %v", err)
	}

	return
}

func (s *Session) DeleteGuildEmoji(guildID discord.Snowflake, emojiID discord.Snowflake, reason *string) (err error) {
	endpoint := EndpointGuildEmoji(guildID.String(), emojiID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(structs.AuditLogReasonHeader, *reason)
	}

	err = s.Interface.FetchJJ(s, http.MethodDelete, endpoint, nil, headers, nil)
	if err != nil {
		return xerrors.Errorf("Failed to delete guild emoji: %v", err)
	}

	return
}
