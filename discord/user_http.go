package discord

import (
	"net/http"

	"golang.org/x/xerrors"
)

func GetCurrentUser(s *Session) (user *User, err error) {
	endpoint := EndpointUser("@me")

	err = s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &user)
	if err != nil {
		return nil, xerrors.Errorf("Failed to get current user: %v", err)
	}

	return
}

func GetUser(s *Session, userID Snowflake) (user *User, err error) {
	endpoint := EndpointUser(userID.String())

	err = s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &user)
	if err != nil {
		return nil, xerrors.Errorf("Failed to get user: %v", err)
	}

	return
}

func ModifyCurrentUser(s *Session, userParam UserParam) (user *User, err error) {
	endpoint := EndpointUser("@me")

	err = s.Interface.FetchJJ(s, http.MethodPatch, endpoint, userParam, nil, &user)
	if err != nil {
		return nil, xerrors.Errorf("Failed to modify current user: %v", err)
	}

	return
}

func GetCurrentUserGuilds(s *Session) (guilds []*Guild, err error) {
	endpoint := EndpointUserGuilds("@me")

	err = s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &guilds)
	if err != nil {
		return nil, xerrors.Errorf("Failed to get current user guilds: %v", err)
	}

	return
}

func GetCurrentUserGuildMember(s *Session, guildID Snowflake) (guildMember *GuildMember, err error) {
	endpoint := EndpointUserGuildMember("@me", guildID.String())

	err = s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &guildMember)
	if err != nil {
		return nil, xerrors.Errorf("Failed to get current user guild member: %v", err)
	}

	return
}

func LeaveGuild(s *Session, guildID Snowflake) (err error) {
	endpoint := EndpointUserGuild("@me", guildID.String())

	err = s.Interface.FetchJJ(s, http.MethodDelete, endpoint, nil, nil, nil)
	if err != nil {
		return xerrors.Errorf("Failed to leave guild: %v", err)
	}

	return
}

func CreateDM(s *Session, recipientID Snowflake) (channel *Channel, err error) {
	endpoint := EndpointUserChannels("@me")

	createDMStruct := struct {
		RecipientID Snowflake `json:"recipient_id"`
	}{recipientID}

	err = s.Interface.FetchJJ(s, http.MethodPost, endpoint, createDMStruct, nil, &channel)
	if err != nil {
		return nil, xerrors.Errorf("Failed to create dm: %v", err)
	}

	return
}
