package http

import (
	"context"
	"net/http"
	"net/url"

	"github.com/WelcomerTeam/Discord/discord"
	"github.com/WelcomerTeam/Discord/structs"
	"golang.org/x/xerrors"
)

func (s *Session) GetCurrentUser(ctx context.Context) (user *structs.User, err error) {
	endpoint := EndpointUser("@me")

	err = s.Interface.FetchJJ(ctx, http.MethodGet, endpoint, nil, nil, &user)
	if err != nil {
		return nil, xerrors.Errorf("Failed to get current user: %v", err)
	}

	return
}

func (s *Session) GetUser(ctx context.Context, userID discord.Snowflake) (user *structs.User, err error) {
	endpoint := EndpointUser(userID.String())

	err = s.Interface.FetchJJ(ctx, http.MethodGet, endpoint, nil, nil, &user)
	if err != nil {
		return nil, xerrors.Errorf("Failed to get user: %v", err)
	}

	return
}

func (s *Session) ModifyCurrentUser(ctx context.Context, userParam structs.UserParam) (user *structs.User, err error) {
	endpoint := EndpointUser("@me")

	err = s.Interface.FetchJJ(ctx, http.MethodPatch, endpoint, userParam, nil, &user)
	if err != nil {
		return nil, xerrors.Errorf("Failed to modify current user: %v", err)
	}

	return
}

func (s *Session) GetCurrentUserGuilds(ctx context.Context) (guilds []*structs.Guild, err error) {
	endpoint := EndpointUserGuilds("@me")

	err = s.Interface.FetchJJ(ctx, http.MethodGet, endpoint, nil, nil, &guilds)
	if err != nil {
		return nil, xerrors.Errorf("Failed to get current user guilds: %v", err)
	}

	return
}

func (s *Session) GetCurrentUserGuildMember(ctx context.Context, guildID discord.Snowflake) (guildMember *structs.GuildMember, err error) {
	endpoint := EndpointUserGuildMember("@me", guildID.String())

	err = s.Interface.FetchJJ(ctx, http.MethodGet, endpoint, nil, nil, &guildMember)
	if err != nil {
		return nil, xerrors.Errorf("Failed to get current user guild member: %v", err)
	}

	return
}

func (s *Session) LeaveGuild(ctx context.Context, guildID discord.Snowflake) (err error) {
	endpoint := EndpointUserGuild("@me", guildID.String())

	err = s.Interface.FetchJJ(ctx, http.MethodDelete, endpoint, nil, nil, nil)
	if err != nil {
		return xerrors.Errorf("Failed to leave guild: %v", err)
	}

	return
}

func (s *Session) CreateDM(ctx context.Context, recipientID discord.Snowflake) (channel *structs.Channel, err error) {
	endpoint := EndpointUserChannels("@me")

	var values url.Values

	values.Add("recipient_id", recipientID.String())

	endpoint += "?" + values.Encode()

	err = s.Interface.FetchJJ(ctx, http.MethodPost, endpoint, nil, nil, &channel)
	if err != nil {
		return nil, xerrors.Errorf("Failed to create dm: %v", err)
	}

	return
}

// TODO: CreateGroupDM
// TODO: GetUserConnections
