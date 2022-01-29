package http

import (
	"context"
	"net/http"

	"github.com/WelcomerTeam/Discord/discord"
	"github.com/WelcomerTeam/Discord/structs"
	"golang.org/x/xerrors"
)

func (s *Session) GetGlobalApplicationCommands(ctx context.Context, applicationID discord.Snowflake) (commands []*structs.ApplicationCommand, err error) {
	endpoint := EndpointApplicationGlobalCommands(applicationID.String())

	err = s.Interface.FetchJJ(ctx, http.MethodGet, endpoint, nil, nil, &commands)
	if err != nil {
		return nil, xerrors.Errorf("Failed to get global application commands: %v", err)
	}

	return
}

func (s *Session) CreateGlobalApplicationCommand(ctx context.Context, applicationID discord.Snowflake, commandArgs structs.ApplicationCommand) (command *structs.ApplicationCommand, err error) {
	endpoint := EndpointApplicationGlobalCommands(applicationID.String())

	err = s.Interface.FetchJJ(ctx, http.MethodPost, endpoint, commandArgs, nil, &command)
	if err != nil {
		return nil, xerrors.Errorf("Failed to create global application command: %v", err)
	}

	return
}

func (s *Session) GetGlobalApplicationCommand(ctx context.Context, applicationID discord.Snowflake, commandID discord.Snowflake) (command *structs.ApplicationCommand, err error) {
	endpoint := EndpointApplicationGlobalCommand(applicationID.String(), commandID.String())

	err = s.Interface.FetchJJ(ctx, http.MethodGet, endpoint, nil, nil, &command)
	if err != nil {
		return nil, xerrors.Errorf("Failed to get global application command: %v", err)
	}

	return
}

func (s *Session) EditGlobalApplicationCommand(ctx context.Context, applicationID discord.Snowflake, commandID discord.Snowflake, commandArg structs.ApplicationCommand) (command *structs.ApplicationCommand, err error) {
	endpoint := EndpointApplicationGlobalCommand(applicationID.String(), commandID.String())

	err = s.Interface.FetchJJ(ctx, http.MethodPatch, endpoint, commandArg, nil, &command)
	if err != nil {
		return nil, xerrors.Errorf("Failed to edit global application command: %v", err)
	}

	return
}

func (s *Session) DeleteGlobalApplicationCommand(ctx context.Context, applicationID discord.Snowflake, commandID discord.Snowflake) (err error) {
	endpoint := EndpointApplicationGlobalCommand(applicationID.String(), commandID.String())

	err = s.Interface.FetchJJ(ctx, http.MethodDelete, endpoint, nil, nil, nil)
	if err != nil {
		return xerrors.Errorf("Failed to delete global application command: %v", err)
	}

	return
}

func (s *Session) BulkOverwriteGloblApplicationCommands(ctx context.Context, applicationID discord.Snowflake, commandArgs []structs.ApplicationCommand) (commands []*structs.ApplicationCommand, err error) {
	endpoint := EndpointApplicationGlobalCommands(applicationID.String())

	err = s.Interface.FetchJJ(ctx, http.MethodPut, endpoint, commandArgs, nil, &commands)
	if err != nil {
		return nil, xerrors.Errorf("Failed to bulk overwrite global application commands: %v", err)
	}

	return
}

func (s *Session) GetGuildApplicationCommands(ctx context.Context, applicationID discord.Snowflake, guildID discord.Snowflake) (commands []*structs.ApplicationCommand, err error) {
	endpoint := EndpointApplicationGuildCommands(applicationID.String(), guildID.String())

	err = s.Interface.FetchJJ(ctx, http.MethodGet, endpoint, nil, nil, &commands)
	if err != nil {
		return nil, xerrors.Errorf("Failed to get guild application commands: %v", err)
	}

	return
}

func (s *Session) CreateGuildApplicationCommand(ctx context.Context, applicationID discord.Snowflake, guildID discord.Snowflake, commandArg structs.ApplicationCommand) (command *structs.ApplicationCommand, err error) {
	endpoint := EndpointApplicationGuildCommands(applicationID.String(), guildID.String())

	err = s.Interface.FetchJJ(ctx, http.MethodPost, endpoint, commandArg, nil, &command)
	if err != nil {
		return nil, xerrors.Errorf("Failed to create guild application command: %v", err)
	}

	return
}

func (s *Session) GetGuildApplicationCommand(ctx context.Context, applicationID discord.Snowflake, guildID discord.Snowflake, commandID discord.Snowflake) (command *structs.ApplicationCommand, err error) {
	endpoint := EndpointApplicationGuildCommand(applicationID.String(), guildID.String(), commandID.String())

	err = s.Interface.FetchJJ(ctx, http.MethodGet, endpoint, nil, nil, &command)
	if err != nil {
		return nil, xerrors.Errorf("Failed to get guild application command: %v", err)
	}

	return
}

func (s *Session) EditGuildApplicationCommand(ctx context.Context, applicationID discord.Snowflake, guildID discord.Snowflake, commandID discord.Snowflake, commandArg structs.ApplicationCommand) (command *structs.ApplicationCommand, err error) {
	endpoint := EndpointApplicationGuildCommand(applicationID.String(), guildID.String(), commandID.String())

	err = s.Interface.FetchJJ(ctx, http.MethodPatch, endpoint, commandArg, nil, &command)
	if err != nil {
		return nil, xerrors.Errorf("Failed to edit guild application command: %v", err)
	}

	return
}

func (s *Session) DeleteGuildApplicationCommand(ctx context.Context, applicationID discord.Snowflake, guildID discord.Snowflake, commandID discord.Snowflake) (err error) {
	endpoint := EndpointApplicationGuildCommand(applicationID.String(), guildID.String(), commandID.String())

	err = s.Interface.FetchJJ(ctx, http.MethodDelete, endpoint, nil, nil, nil)
	if err != nil {
		return xerrors.Errorf("Failed to delete guild application command: %v", err)
	}

	return
}

func (s *Session) BulkOverwriteGuildApplicationCommands(ctx context.Context, applicationID discord.Snowflake, guildID discord.Snowflake, commandArgs []structs.ApplicationCommand) (commands []*structs.ApplicationCommand, err error) {
	endpoint := EndpointApplicationGuildCommands(applicationID.String(), guildID.String())

	err = s.Interface.FetchJJ(ctx, http.MethodPut, endpoint, commandArgs, nil, &commands)
	if err != nil {
		return nil, xerrors.Errorf("Failed to bulk overwrite guild application commands: %v", err)
	}

	return
}

func (s *Session) GetGuildApplicationCommandPermissions(ctx context.Context, applicationID discord.Snowflake, guildID discord.Snowflake) (permissions []structs.GuildApplicationCommandPermissions, err error) {
	endpoint := EndpointApplicationGuildCommandsPermissions(applicationID.String(), guildID.String())

	err = s.Interface.FetchJJ(ctx, http.MethodGet, endpoint, nil, nil, &permissions)
	if err != nil {
		return nil, xerrors.Errorf("Failed to get guild application command permissions: %v", err)
	}

	return
}

func (s *Session) GetApplicationCommandPermissions(ctx context.Context, applicationID discord.Snowflake, guildID discord.Snowflake, commandID discord.Snowflake) (permissions []structs.GuildApplicationCommandPermissions, err error) {
	endpoint := EndpointApplicationGuildCommandPermissions(applicationID.String(), guildID.String(), commandID.String())

	err = s.Interface.FetchJJ(ctx, http.MethodGet, endpoint, nil, nil, &permissions)
	if err != nil {
		return nil, xerrors.Errorf("Failed to get application command permissions: %v", err)
	}

	return
}

func (s *Session) EditApplicationCommandPermissions(ctx context.Context, applicationID discord.Snowflake, guildID discord.Snowflake, commandID discord.Snowflake, permissionsArg []structs.GuildApplicationCommandPermissions) (permissions []*structs.GuildApplicationCommandPermissions, err error) {
	endpoint := EndpointApplicationGuildCommandPermissions(applicationID.String(), guildID.String(), commandID.String())

	err = s.Interface.FetchJJ(ctx, http.MethodPut, endpoint, permissionsArg, nil, &permissions)
	if err != nil {
		return nil, xerrors.Errorf("Failed to edit application command permissions: %v", err)
	}

	return
}

func (s *Session) BatchEditApplicationCommandPermissions(ctx context.Context, applicationID discord.Snowflake, guildID discord.Snowflake, permissionsArg []structs.GuildApplicationCommandPermissions) (permissions []*structs.GuildApplicationCommandPermissions, err error) {
	endpoint := EndpointApplicationGuildCommandsPermissions(applicationID.String(), guildID.String())

	err = s.Interface.FetchJJ(ctx, http.MethodPut, endpoint, permissionsArg, nil, &permissions)
	if err != nil {
		return nil, xerrors.Errorf("Failed to batch edit application command permissions: %v", err)
	}

	return
}
