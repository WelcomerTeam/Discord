package discord

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

func GetGlobalApplicationCommands(ctx context.Context, session *Session, applicationID Snowflake, withLocalizations bool) ([]ApplicationCommand, error) {
	endpoint := EndpointApplicationGlobalCommands(applicationID.String())

	values := url.Values{}

	if withLocalizations {
		values.Set("with_localizations", "true")
	}

	if len(values) > 0 {
		endpoint += "?" + values.Encode()
	}

	var commands []ApplicationCommand

	err := session.Interface.FetchJJ(ctx, session, http.MethodGet, endpoint, nil, nil, &commands)
	if err != nil {
		return nil, fmt.Errorf("failed to get global application commands: %w", err)
	}

	return commands, nil
}

func CreateGlobalApplicationCommand(ctx context.Context, session *Session, applicationID Snowflake, commandArg ApplicationCommand) (*ApplicationCommand, error) {
	endpoint := EndpointApplicationGlobalCommands(applicationID.String())

	var command *ApplicationCommand

	err := session.Interface.FetchJJ(ctx, session, http.MethodPost, endpoint, commandArg, nil, &command)
	if err != nil {
		return nil, fmt.Errorf("failed to create global application command: %w", err)
	}

	return command, nil
}

func GetGlobalApplicationCommand(ctx context.Context, session *Session, applicationID, commandID Snowflake, withLocalizations bool) (*ApplicationCommand, error) {
	endpoint := EndpointApplicationGlobalCommand(applicationID.String(), commandID.String())

	values := url.Values{}

	if withLocalizations {
		values.Set("with_localizations", "true")
	}

	if len(values) > 0 {
		endpoint += "?" + values.Encode()
	}

	var command *ApplicationCommand

	err := session.Interface.FetchJJ(ctx, session, http.MethodGet, endpoint, nil, nil, &command)
	if err != nil {
		return nil, fmt.Errorf("failed to get global application command: %w", err)
	}

	return command, nil
}

func EditGlobalApplicationCommand(ctx context.Context, session *Session, applicationID, commandID Snowflake, commandArg ApplicationCommand) (*ApplicationCommand, error) {
	endpoint := EndpointApplicationGlobalCommand(applicationID.String(), commandID.String())

	var command *ApplicationCommand

	err := session.Interface.FetchJJ(ctx, session, http.MethodPatch, endpoint, commandArg, nil, &command)
	if err != nil {
		return nil, fmt.Errorf("failed to edit global application command: %w", err)
	}

	return command, nil
}

func DeleteGlobalApplicationCommand(ctx context.Context, session *Session, applicationID, commandID Snowflake) error {
	endpoint := EndpointApplicationGlobalCommand(applicationID.String(), commandID.String())

	err := session.Interface.FetchJJ(ctx, session, http.MethodDelete, endpoint, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete global application command: %w", err)
	}

	return nil
}

func BulkOverwriteGlobalApplicationCommands(ctx context.Context, session *Session, applicationID Snowflake, commandArgs []ApplicationCommand) ([]ApplicationCommand, error) {
	endpoint := EndpointApplicationGlobalCommands(applicationID.String())

	var commands []ApplicationCommand

	err := session.Interface.FetchJJ(ctx, session, http.MethodPut, endpoint, commandArgs, nil, &commands)
	if err != nil {
		return nil, fmt.Errorf("failed to bulk overwrite global application commands: %w", err)
	}

	return commands, nil
}

func GetGuildApplicationCommands(ctx context.Context, session *Session, applicationID, guildID Snowflake) ([]ApplicationCommand, error) {
	endpoint := EndpointApplicationGuildCommands(applicationID.String(), guildID.String())

	var commands []ApplicationCommand

	err := session.Interface.FetchJJ(ctx, session, http.MethodGet, endpoint, nil, nil, &commands)
	if err != nil {
		return nil, fmt.Errorf("failed to get guild application commands: %w", err)
	}

	return commands, nil
}

func CreateGuildApplicationCommand(ctx context.Context, session *Session, applicationID, guildID Snowflake, commandArg ApplicationCommand) (*ApplicationCommand, error) {
	endpoint := EndpointApplicationGuildCommands(applicationID.String(), guildID.String())

	var command *ApplicationCommand

	err := session.Interface.FetchJJ(ctx, session, http.MethodPost, endpoint, commandArg, nil, &command)
	if err != nil {
		return nil, fmt.Errorf("failed to create guild application command: %w", err)
	}

	return command, nil
}

func GetGuildApplicationCommand(ctx context.Context, session *Session, applicationID, guildID, commandID Snowflake) (*ApplicationCommand, error) {
	endpoint := EndpointApplicationGuildCommand(applicationID.String(), guildID.String(), commandID.String())

	var command *ApplicationCommand

	err := session.Interface.FetchJJ(ctx, session, http.MethodGet, endpoint, nil, nil, &command)
	if err != nil {
		return nil, fmt.Errorf("failed to get guild application command: %w", err)
	}

	return command, nil
}

func EditGuildApplicationCommand(ctx context.Context, session *Session, applicationID, guildID, commandID Snowflake, commandArg ApplicationCommand) (*ApplicationCommand, error) {
	endpoint := EndpointApplicationGuildCommand(applicationID.String(), guildID.String(), commandID.String())

	var command *ApplicationCommand

	err := session.Interface.FetchJJ(ctx, session, http.MethodPatch, endpoint, commandArg, nil, &command)
	if err != nil {
		return nil, fmt.Errorf("failed to edit guild application command: %w", err)
	}

	return command, nil
}

func DeleteGuildApplicationCommand(ctx context.Context, session *Session, applicationID, guildID, commandID Snowflake) error {
	endpoint := EndpointApplicationGuildCommand(applicationID.String(), guildID.String(), commandID.String())

	err := session.Interface.FetchJJ(ctx, session, http.MethodDelete, endpoint, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete guild application command: %w", err)
	}

	return nil
}

func BulkOverwriteGuildApplicationCommands(ctx context.Context, session *Session, applicationID, guildID Snowflake, commandArgs []ApplicationCommand) ([]ApplicationCommand, error) {
	endpoint := EndpointApplicationGuildCommands(applicationID.String(), guildID.String())

	var commands []ApplicationCommand

	err := session.Interface.FetchJJ(ctx, session, http.MethodPut, endpoint, commandArgs, nil, &commands)
	if err != nil {
		return nil, fmt.Errorf("failed to bulk overwrite guild application commands: %w", err)
	}

	return commands, nil
}

func GetGuildApplicationCommandPermissions(ctx context.Context, session *Session, applicationID, guildID Snowflake) ([]GuildApplicationCommandPermissions, error) {
	endpoint := EndpointApplicationGuildCommandsPermissions(applicationID.String(), guildID.String())

	var permissions []GuildApplicationCommandPermissions

	err := session.Interface.FetchJJ(ctx, session, http.MethodGet, endpoint, nil, nil, &permissions)
	if err != nil {
		return nil, fmt.Errorf("failed to get guild application command permissions: %w", err)
	}

	return permissions, nil
}

func GetApplicationCommandPermissions(ctx context.Context, session *Session, applicationID, guildID, commandID Snowflake) ([]GuildApplicationCommandPermissions, error) {
	endpoint := EndpointApplicationGuildCommandPermissions(applicationID.String(), guildID.String(), commandID.String())

	var permissions []GuildApplicationCommandPermissions

	err := session.Interface.FetchJJ(ctx, session, http.MethodGet, endpoint, nil, nil, &permissions)
	if err != nil {
		return nil, fmt.Errorf("failed to get application command permissions: %w", err)
	}

	return permissions, nil
}

func EditApplicationCommandPermissions(ctx context.Context, session *Session, applicationID, guildID, commandID Snowflake, permissionsArg []GuildApplicationCommandPermissions) ([]GuildApplicationCommandPermissions, error) {
	endpoint := EndpointApplicationGuildCommandPermissions(applicationID.String(), guildID.String(), commandID.String())

	var permissions []GuildApplicationCommandPermissions

	err := session.Interface.FetchJJ(ctx, session, http.MethodPut, endpoint, permissionsArg, nil, &permissions)
	if err != nil {
		return nil, fmt.Errorf("failed to edit application command permissions: %w", err)
	}

	return permissions, nil
}

func BatchEditApplicationCommandPermissions(ctx context.Context, session *Session, applicationID, guildID Snowflake, permissionsArg []GuildApplicationCommandPermissions) ([]GuildApplicationCommandPermissions, error) {
	endpoint := EndpointApplicationGuildCommandsPermissions(applicationID.String(), guildID.String())

	var permissions []GuildApplicationCommandPermissions

	err := session.Interface.FetchJJ(ctx, session, http.MethodPut, endpoint, permissionsArg, nil, &permissions)
	if err != nil {
		return nil, fmt.Errorf("failed to batch edit application command permissions: %w", err)
	}

	return permissions, nil
}
