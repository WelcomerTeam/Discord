package discord

import (
	"context"
	"fmt"
	"net/http"
)

func GetCurrentUser(ctx context.Context, session *Session) (*User, error) {
	endpoint := EndpointUser("@me")

	var user *User

	err := session.Interface.FetchJJ(ctx, session, http.MethodGet, endpoint, nil, nil, &user)
	if err != nil {
		return nil, fmt.Errorf("failed to get current user: %w", err)
	}

	return user, nil
}

func GetUser(ctx context.Context, session *Session, userID Snowflake) (*User, error) {
	endpoint := EndpointUser(userID.String())

	var user *User

	err := session.Interface.FetchJJ(ctx, session, http.MethodGet, endpoint, nil, nil, &user)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return user, nil
}

func ModifyCurrentUser(ctx context.Context, session *Session, userParam UserParam) (*User, error) {
	endpoint := EndpointUser("@me")

	var user *User

	err := session.Interface.FetchJJ(ctx, session, http.MethodPatch, endpoint, userParam, nil, &user)
	if err != nil {
		return nil, fmt.Errorf("failed to modify current user: %w", err)
	}

	return user, nil
}

func GetCurrentUserGuilds(ctx context.Context, session *Session) ([]Guild, error) {
	endpoint := EndpointUserGuilds("@me")

	var guilds []Guild

	err := session.Interface.FetchJJ(ctx, session, http.MethodGet, endpoint, nil, nil, &guilds)
	if err != nil {
		return nil, fmt.Errorf("failed to get current user guilds: %w", err)
	}

	return guilds, nil
}

func GetCurrentUserGuildMember(ctx context.Context, session *Session, guildID Snowflake) (*GuildMember, error) {
	endpoint := EndpointUserGuildMember("@me", guildID.String())

	var guildMember *GuildMember

	err := session.Interface.FetchJJ(ctx, session, http.MethodGet, endpoint, nil, nil, &guildMember)
	if err != nil {
		return nil, fmt.Errorf("failed to get current user guild member: %w", err)
	}

	return guildMember, nil
}

func LeaveGuild(ctx context.Context, session *Session, guildID Snowflake) error {
	endpoint := EndpointUserGuild("@me", guildID.String())

	err := session.Interface.FetchJJ(ctx, session, http.MethodDelete, endpoint, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to leave guild: %w", err)
	}

	return nil
}

func CreateDM(ctx context.Context, session *Session, recipientID Snowflake) (*Channel, error) {
	endpoint := EndpointUserChannels("@me")

	createDMStruct := struct {
		RecipientID Snowflake `json:"recipient_id"`
	}{recipientID}

	var channel *Channel

	err := session.Interface.FetchJJ(ctx, session, http.MethodPost, endpoint, createDMStruct, nil, &channel)
	if err != nil {
		return nil, fmt.Errorf("failed to create dm: %w", err)
	}

	return channel, nil
}

// GetCurrentUserConnections returns a list of the current user's connections.
func GetCurrentUserConnections(ctx context.Context, session *Session) ([]UserConnection, error) {
	endpoint := EndpointUserConnections

	var connections []UserConnection

	err := session.Interface.FetchJJ(ctx, session, http.MethodGet, endpoint, nil, nil, &connections)
	if err != nil {
		return nil, fmt.Errorf("failed to get current user connections: %w", err)
	}

	return connections, nil
}

// GetCurrentUserApplicationRoleConnection returns the role connection for the current user to an application.
func GetCurrentUserApplicationRoleConnection(ctx context.Context, session *Session, applicationID Snowflake) (*ApplicationRoleConnection, error) {
	endpoint := EndpointUserApplicationRoleConnection(applicationID.String())

	var roleConnection *ApplicationRoleConnection

	err := session.Interface.FetchJJ(ctx, session, http.MethodGet, endpoint, nil, nil, &roleConnection)
	if err != nil {
		return nil, fmt.Errorf("failed to get current user application role connection: %w", err)
	}

	return roleConnection, nil
}

// UpdateCurrentUserApplicationRoleConnection updates the role connection for the current user to an application.
func UpdateCurrentUserApplicationRoleConnection(ctx context.Context, session *Session, applicationID Snowflake, params ApplicationRoleConnection) (*ApplicationRoleConnection, error) {
	endpoint := EndpointUserApplicationRoleConnection(applicationID.String())

	var roleConnection *ApplicationRoleConnection

	err := session.Interface.FetchJJ(ctx, session, http.MethodPut, endpoint, params, nil, &roleConnection)
	if err != nil {
		return nil, fmt.Errorf("failed to update current user application role connection: %w", err)
	}

	return roleConnection, nil
}

// DeleteCurrentUserApplicationRoleConnection deletes the role connection for the current user to an application.
func DeleteCurrentUserApplicationRoleConnection(ctx context.Context, session *Session, applicationID Snowflake) error {
	endpoint := EndpointUserApplicationRoleConnection(applicationID.String())

	err := session.Interface.FetchJJ(ctx, session, http.MethodDelete, endpoint, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete current user application role connection: %w", err)
	}

	return nil
}

// GetCurrentUserEntitlements returns entitlements for the current user for a given application.
func GetCurrentUserEntitlements(ctx context.Context, session *Session, applicationID Snowflake) ([]Entitlement, error) {
	endpoint := EndpointUserCurrentEntitlements(applicationID.String())

	var entitlements []Entitlement

	err := session.Interface.FetchJJ(ctx, session, http.MethodGet, endpoint, nil, nil, &entitlements)
	if err != nil {
		return nil, fmt.Errorf("failed to get current user entitlements: %w", err)
	}

	return entitlements, nil
}
