package discord

import (
	"fmt"
	"net/http"
)

func GetCurrentUser(s *Session) (*User, error) {
	endpoint := EndpointUser("@me")

	var user *User

	err := s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &user)
	if err != nil {
		return nil, fmt.Errorf("failed to get current user: %v", err)
	}

	return user, nil
}

func GetUser(s *Session, userID Snowflake) (*User, error) {
	endpoint := EndpointUser(userID.String())

	var user *User

	err := s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &user)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %v", err)
	}

	return user, nil
}

func ModifyCurrentUser(s *Session, userParam UserParam) (*User, error) {
	endpoint := EndpointUser("@me")

	var user *User

	err := s.Interface.FetchJJ(s, http.MethodPatch, endpoint, userParam, nil, &user)
	if err != nil {
		return nil, fmt.Errorf("failed to modify current user: %v", err)
	}

	return user, nil
}

func GetCurrentUserGuilds(s *Session) ([]*Guild, error) {
	endpoint := EndpointUserGuilds("@me")

	var guilds []*Guild

	err := s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &guilds)
	if err != nil {
		return nil, fmt.Errorf("failed to get current user guilds: %v", err)
	}

	return guilds, nil
}

func GetCurrentUserGuildMember(s *Session, guildID Snowflake) (*GuildMember, error) {
	endpoint := EndpointUserGuildMember("@me", guildID.String())

	var guildMember *GuildMember

	err := s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &guildMember)
	if err != nil {
		return nil, fmt.Errorf("failed to get current user guild member: %v", err)
	}

	return guildMember, nil
}

func LeaveGuild(s *Session, guildID Snowflake) error {
	endpoint := EndpointUserGuild("@me", guildID.String())

	err := s.Interface.FetchJJ(s, http.MethodDelete, endpoint, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to leave guild: %v", err)
	}

	return nil
}

func CreateDM(s *Session, recipientID Snowflake) (*Channel, error) {
	endpoint := EndpointUserChannels("@me")

	createDMStruct := struct {
		RecipientID Snowflake `json:"recipient_id"`
	}{recipientID}

	var channel *Channel

	err := s.Interface.FetchJJ(s, http.MethodPost, endpoint, createDMStruct, nil, &channel)
	if err != nil {
		return nil, fmt.Errorf("failed to create dm: %v", err)
	}

	return channel, nil
}
