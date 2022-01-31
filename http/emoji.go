package http

import (
	"context"
	"net/http"

	"github.com/WelcomerTeam/Discord/discord"
	"github.com/WelcomerTeam/Discord/structs"
	"golang.org/x/xerrors"
)

func (s *Session) ListGuildEmojis(ctx context.Context, guildID discord.Snowflake) (emojis []*structs.Emoji, err error) {
	endpoint := EndpointGuildEmojis(guildID.String())

	err = s.Interface.FetchJJ(ctx, http.MethodGet, endpoint, nil, nil, &emojis)
	if err != nil {
		return nil, xerrors.Errorf("Failed to list guild emojis: %v", err)
	}

	return
}

func (s *Session) GetGuildEmoji(ctx context.Context, guildID discord.Snowflake, emojiID discord.Snowflake) (emoji *structs.Emoji, err error) {
	endpoint := EndpointGuildEmoji(guildID.String(), emojiID.String())

	err = s.Interface.FetchJJ(ctx, http.MethodGet, endpoint, nil, nil, &emoji)
	if err != nil {
		return nil, xerrors.Errorf("Failed to get get guild emoji: %v", err)
	}

	return
}

func (s *Session) CreateGuildEmoji(ctx context.Context, guildID discord.Snowflake, emojiParams structs.EmojiParams, reason *string) (emoji *structs.Emoji, err error) {
	endpoint := EndpointGuildEmojis(guildID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(structs.AuditLogReasonHeader, *reason)
	}

	err = s.Interface.FetchJJ(ctx, http.MethodPost, endpoint, emojiParams, headers, &emoji)
	if err != nil {
		return nil, xerrors.Errorf("Failed to create guild emoji: %v", err)
	}

	return
}

func (s *Session) ModifyGuildEmoji(ctx context.Context, guildID discord.Snowflake, emojiID discord.Snowflake, emojiParams structs.EmojiParams, reason *string) (emoji *structs.Emoji, err error) {
	endpoint := EndpointGuildEmoji(guildID.String(), emojiID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(structs.AuditLogReasonHeader, *reason)
	}

	err = s.Interface.FetchJJ(ctx, http.MethodPatch, endpoint, emojiParams, headers, &emoji)
	if err != nil {
		return nil, xerrors.Errorf("Failed to modify guild emoji: %v", err)
	}

	return
}

func (s *Session) DeleteGuildEmoji(ctx context.Context, guildID discord.Snowflake, emojiID discord.Snowflake, reason *string) (err error) {
	endpoint := EndpointGuildEmoji(guildID.String(), emojiID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(structs.AuditLogReasonHeader, *reason)
	}

	err = s.Interface.FetchJJ(ctx, http.MethodDelete, endpoint, nil, headers, nil)
	if err != nil {
		return xerrors.Errorf("Failed to delete guild emoji: %v", err)
	}

	return
}
