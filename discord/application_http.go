package discord

import (
	"fmt"
	"net/http"
	"net/url"
)

func GetGlobalApplicationCommands(s *Session, applicationID Snowflake, withLocalizations bool) ([]*ApplicationCommand, error) {
	endpoint := EndpointApplicationGlobalCommands(applicationID.String())

	values := url.Values{}

	if withLocalizations {
		values.Set("with_localizations", "true")
	}

	if len(values) > 0 {
		endpoint += "?" + values.Encode()
	}

	var commands []*ApplicationCommand

	err := s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &commands)
	if err != nil {
		return nil, fmt.Errorf("failed to get global application commands: %v", err)
	}

	return commands, nil
}

func CreateGlobalApplicationCommand(s *Session, applicationID Snowflake, commandArg ApplicationCommand) (*ApplicationCommand, error) {
	endpoint := EndpointApplicationGlobalCommands(applicationID.String())

	var command *ApplicationCommand

	err := s.Interface.FetchJJ(s, http.MethodPost, endpoint, commandArg, nil, &command)
	if err != nil {
		return nil, fmt.Errorf("failed to create global application command: %v", err)
	}

	return command, nil
}

func GetGlobalApplicationCommand(s *Session, applicationID Snowflake, commandID Snowflake, withLocalizations bool) (*ApplicationCommand, error) {
	endpoint := EndpointApplicationGlobalCommand(applicationID.String(), commandID.String())

	values := url.Values{}

	if withLocalizations {
		values.Set("with_localizations", "true")
	}

	if len(values) > 0 {
		endpoint += "?" + values.Encode()
	}

	var command *ApplicationCommand

	err := s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &command)
	if err != nil {
		return nil, fmt.Errorf("failed to get global application command: %v", err)
	}

	return command, nil
}

func EditGlobalApplicationCommand(s *Session, applicationID Snowflake, commandID Snowflake, commandArg ApplicationCommand) (*ApplicationCommand, error) {
	endpoint := EndpointApplicationGlobalCommand(applicationID.String(), commandID.String())

	var command *ApplicationCommand

	err := s.Interface.FetchJJ(s, http.MethodPatch, endpoint, commandArg, nil, &command)
	if err != nil {
		return nil, fmt.Errorf("failed to edit global application command: %v", err)
	}

	return command, nil
}

func DeleteGlobalApplicationCommand(s *Session, applicationID Snowflake, commandID Snowflake) error {
	endpoint := EndpointApplicationGlobalCommand(applicationID.String(), commandID.String())

	err := s.Interface.FetchJJ(s, http.MethodDelete, endpoint, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete global application command: %v", err)
	}

	return nil
}

func BulkOverwriteGlobalApplicationCommands(s *Session, applicationID Snowflake, commandArgs []*ApplicationCommand) ([]*ApplicationCommand, error) {
	endpoint := EndpointApplicationGlobalCommands(applicationID.String())

	var commands []*ApplicationCommand

	err := s.Interface.FetchJJ(s, http.MethodPut, endpoint, commandArgs, nil, &commands)
	if err != nil {
		return nil, fmt.Errorf("failed to bulk overwrite global application commands: %v", err)
	}

	return commands, nil
}

func GetGuildApplicationCommands(s *Session, applicationID Snowflake, guildID Snowflake) ([]*ApplicationCommand, error) {
	endpoint := EndpointApplicationGuildCommands(applicationID.String(), guildID.String())

	var commands []*ApplicationCommand

	err := s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &commands)
	if err != nil {
		return nil, fmt.Errorf("failed to get guild application commands: %v", err)
	}

	return commands, nil
}

func CreateGuildApplicationCommand(s *Session, applicationID Snowflake, guildID Snowflake, commandArg ApplicationCommand) (*ApplicationCommand, error) {
	endpoint := EndpointApplicationGuildCommands(applicationID.String(), guildID.String())

	var command *ApplicationCommand

	err := s.Interface.FetchJJ(s, http.MethodPost, endpoint, commandArg, nil, &command)
	if err != nil {
		return nil, fmt.Errorf("failed to create guild application command: %v", err)
	}

	return command, nil
}

func GetGuildApplicationCommand(s *Session, applicationID Snowflake, guildID Snowflake, commandID Snowflake) (*ApplicationCommand, error) {
	endpoint := EndpointApplicationGuildCommand(applicationID.String(), guildID.String(), commandID.String())

	var command *ApplicationCommand

	err := s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &command)
	if err != nil {
		return nil, fmt.Errorf("failed to get guild application command: %v", err)
	}

	return command, nil
}

func EditGuildApplicationCommand(s *Session, applicationID Snowflake, guildID Snowflake, commandID Snowflake, commandArg ApplicationCommand) (*ApplicationCommand, error) {
	endpoint := EndpointApplicationGuildCommand(applicationID.String(), guildID.String(), commandID.String())

	var command *ApplicationCommand

	err := s.Interface.FetchJJ(s, http.MethodPatch, endpoint, commandArg, nil, &command)
	if err != nil {
		return nil, fmt.Errorf("failed to edit guild application command: %v", err)
	}

	return command, nil
}

func DeleteGuildApplicationCommand(s *Session, applicationID Snowflake, guildID Snowflake, commandID Snowflake) error {
	endpoint := EndpointApplicationGuildCommand(applicationID.String(), guildID.String(), commandID.String())

	err := s.Interface.FetchJJ(s, http.MethodDelete, endpoint, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete guild application command: %v", err)
	}

	return nil
}

func BulkOverwriteGuildApplicationCommands(s *Session, applicationID Snowflake, guildID Snowflake, commandArgs []*ApplicationCommand) ([]*ApplicationCommand, error) {
	endpoint := EndpointApplicationGuildCommands(applicationID.String(), guildID.String())

	var commands []*ApplicationCommand

	err := s.Interface.FetchJJ(s, http.MethodPut, endpoint, commandArgs, nil, &commands)
	if err != nil {
		return nil, fmt.Errorf("failed to bulk overwrite guild application commands: %v", err)
	}

	return commands, nil
}

func GetGuildApplicationCommandPermissions(s *Session, applicationID Snowflake, guildID Snowflake) ([]GuildApplicationCommandPermissions, error) {
	endpoint := EndpointApplicationGuildCommandsPermissions(applicationID.String(), guildID.String())

	var permissions []GuildApplicationCommandPermissions

	err := s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &permissions)
	if err != nil {
		return nil, fmt.Errorf("failed to get guild application command permissions: %v", err)
	}

	return permissions, nil
}

func GetApplicationCommandPermissions(s *Session, applicationID Snowflake, guildID Snowflake, commandID Snowflake) ([]*GuildApplicationCommandPermissions, error) {
	endpoint := EndpointApplicationGuildCommandPermissions(applicationID.String(), guildID.String(), commandID.String())

	var permissions []*GuildApplicationCommandPermissions

	err := s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &permissions)
	if err != nil {
		return nil, fmt.Errorf("failed to get application command permissions: %v", err)
	}

	return permissions, nil
}

func EditApplicationCommandPermissions(s *Session, applicationID Snowflake, guildID Snowflake, commandID Snowflake, permissionsArg []*GuildApplicationCommandPermissions) ([]*GuildApplicationCommandPermissions, error) {
	endpoint := EndpointApplicationGuildCommandPermissions(applicationID.String(), guildID.String(), commandID.String())

	var permissions []*GuildApplicationCommandPermissions

	err := s.Interface.FetchJJ(s, http.MethodPut, endpoint, permissionsArg, nil, &permissions)
	if err != nil {
		return nil, fmt.Errorf("failed to edit application command permissions: %v", err)
	}

	return permissions, nil
}

func BatchEditApplicationCommandPermissions(s *Session, applicationID Snowflake, guildID Snowflake, permissionsArg []GuildApplicationCommandPermissions) ([]*GuildApplicationCommandPermissions, error) {
	endpoint := EndpointApplicationGuildCommandsPermissions(applicationID.String(), guildID.String())

	var permissions []*GuildApplicationCommandPermissions

	err := s.Interface.FetchJJ(s, http.MethodPut, endpoint, permissionsArg, nil, &permissions)
	if err != nil {
		return nil, fmt.Errorf("failed to batch edit application command permissions: %v", err)
	}

	return permissions, nil
}
