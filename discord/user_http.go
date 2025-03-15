package discord

import (
	"context"
	"fmt"
	"net/http"
)

func GetCurrentUser(ctx context.Context, s *Session) (*User, error) {
	endpoint := EndpointUser("@me")

	var user *User

	err := s.Interface.FetchJJ(ctx, s, http.MethodGet, endpoint, nil, nil, &user)
	if err != nil {
		return nil, fmt.Errorf("failed to get current user: %w", err)
	}

	return user, nil
}

func GetUser(ctx context.Context, s *Session, userID Snowflake) (*User, error) {
	endpoint := EndpointUser(userID.String())

	var user *User

	err := s.Interface.FetchJJ(ctx, s, http.MethodGet, endpoint, nil, nil, &user)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return user, nil
}

func ModifyCurrentUser(ctx context.Context, s *Session, userParam UserParam) (*User, error) {
	endpoint := EndpointUser("@me")

	var user *User

	err := s.Interface.FetchJJ(ctx, s, http.MethodPatch, endpoint, userParam, nil, &user)
	if err != nil {
		return nil, fmt.Errorf("failed to modify current user: %w", err)
	}

	return user, nil
}

func GetCurrentUserGuilds(ctx context.Context, s *Session) ([]Guild, error) {
	endpoint := EndpointUserGuilds("@me")

	var guilds []Guild

	err := s.Interface.FetchJJ(ctx, s, http.MethodGet, endpoint, nil, nil, &guilds)
	if err != nil {
		return nil, fmt.Errorf("failed to get current user guilds: %w", err)
	}

	return guilds, nil
}

func GetCurrentUserGuildMember(ctx context.Context, s *Session, guildID Snowflake) (*GuildMember, error) {
	endpoint := EndpointUserGuildMember("@me", guildID.String())

	var guildMember *GuildMember

	err := s.Interface.FetchJJ(ctx, s, http.MethodGet, endpoint, nil, nil, &guildMember)
	if err != nil {
		return nil, fmt.Errorf("failed to get current user guild member: %w", err)
	}

	return guildMember, nil
}

func LeaveGuild(ctx context.Context, s *Session, guildID Snowflake) error {
	endpoint := EndpointUserGuild("@me", guildID.String())

	err := s.Interface.FetchJJ(ctx, s, http.MethodDelete, endpoint, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to leave guild: %w", err)
	}

	return nil
}

func CreateDM(ctx context.Context, s *Session, recipientID Snowflake) (*Channel, error) {
	endpoint := EndpointUserChannels("@me")

	createDMStruct := struct {
		RecipientID Snowflake `json:"recipient_id"`
	}{recipientID}

	var channel *Channel

	err := s.Interface.FetchJJ(ctx, s, http.MethodPost, endpoint, createDMStruct, nil, &channel)
	if err != nil {
		return nil, fmt.Errorf("failed to create dm: %w", err)
	}

	return channel, nil
}
