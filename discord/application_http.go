package discord

import (
	"fmt"
	"net/http"
	"net/url"
)

func GetGlobalApplicationCommands(s *Session, applicationID Snowflake, withLocalizations bool) (commands []*ApplicationCommand, err error) {
	endpoint := EndpointApplicationGlobalCommands(applicationID.String())

	values := url.Values{}

	if withLocalizations {
		values.Set("with_localizations", "true")
	}

	if len(values) > 0 {
		endpoint += "?" + values.Encode()
	}

	err = s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &commands)
	if err != nil {
		return nil, fmt.Errorf("Failed to get global application commands: %v", err)
	}

	return
}

func CreateGlobalApplicationCommand(s *Session, applicationID Snowflake, commandArg ApplicationCommand) (command *ApplicationCommand, err error) {
	endpoint := EndpointApplicationGlobalCommands(applicationID.String())

	err = s.Interface.FetchJJ(s, http.MethodPost, endpoint, commandArg, nil, &command)
	if err != nil {
		return nil, fmt.Errorf("Failed to create global application command: %v", err)
	}

	return
}

func GetGlobalApplicationCommand(s *Session, applicationID Snowflake, commandID Snowflake, withLocalizations bool) (command *ApplicationCommand, err error) {
	endpoint := EndpointApplicationGlobalCommand(applicationID.String(), commandID.String())

	values := url.Values{}

	if withLocalizations {
		values.Set("with_localizations", "true")
	}

	if len(values) > 0 {
		endpoint += "?" + values.Encode()
	}

	err = s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &command)
	if err != nil {
		return nil, fmt.Errorf("Failed to get global application command: %v", err)
	}

	return
}

func EditGlobalApplicationCommand(s *Session, applicationID Snowflake, commandID Snowflake, commandArg ApplicationCommand) (command *ApplicationCommand, err error) {
	endpoint := EndpointApplicationGlobalCommand(applicationID.String(), commandID.String())

	err = s.Interface.FetchJJ(s, http.MethodPatch, endpoint, commandArg, nil, &command)
	if err != nil {
		return nil, fmt.Errorf("Failed to edit global application command: %v", err)
	}

	return
}

func DeleteGlobalApplicationCommand(s *Session, applicationID Snowflake, commandID Snowflake) (err error) {
	endpoint := EndpointApplicationGlobalCommand(applicationID.String(), commandID.String())

	err = s.Interface.FetchJJ(s, http.MethodDelete, endpoint, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("Failed to delete global application command: %v", err)
	}

	return
}

func BulkOverwriteGlobalApplicationCommands(s *Session, applicationID Snowflake, commandArgs []ApplicationCommand) (commands []*ApplicationCommand, err error) {
	endpoint := EndpointApplicationGlobalCommands(applicationID.String())

	err = s.Interface.FetchJJ(s, http.MethodPut, endpoint, commandArgs, nil, &commands)
	if err != nil {
		return nil, fmt.Errorf("Failed to bulk overwrite global application commands: %v", err)
	}

	return
}

func GetGuildApplicationCommands(s *Session, applicationID Snowflake, guildID Snowflake) (commands []*ApplicationCommand, err error) {
	endpoint := EndpointApplicationGuildCommands(applicationID.String(), guildID.String())

	err = s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &commands)
	if err != nil {
		return nil, fmt.Errorf("Failed to get guild application commands: %v", err)
	}

	return
}

func CreateGuildApplicationCommand(s *Session, applicationID Snowflake, guildID Snowflake, commandArg ApplicationCommand) (command *ApplicationCommand, err error) {
	endpoint := EndpointApplicationGuildCommands(applicationID.String(), guildID.String())

	err = s.Interface.FetchJJ(s, http.MethodPost, endpoint, commandArg, nil, &command)
	if err != nil {
		return nil, fmt.Errorf("Failed to create guild application command: %v", err)
	}

	return
}

func GetGuildApplicationCommand(s *Session, applicationID Snowflake, guildID Snowflake, commandID Snowflake) (command *ApplicationCommand, err error) {
	endpoint := EndpointApplicationGuildCommand(applicationID.String(), guildID.String(), commandID.String())

	err = s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &command)
	if err != nil {
		return nil, fmt.Errorf("Failed to get guild application command: %v", err)
	}

	return
}

func EditGuildApplicationCommand(s *Session, applicationID Snowflake, guildID Snowflake, commandID Snowflake, commandArg ApplicationCommand) (command *ApplicationCommand, err error) {
	endpoint := EndpointApplicationGuildCommand(applicationID.String(), guildID.String(), commandID.String())

	err = s.Interface.FetchJJ(s, http.MethodPatch, endpoint, commandArg, nil, &command)
	if err != nil {
		return nil, fmt.Errorf("Failed to edit guild application command: %v", err)
	}

	return
}

func DeleteGuildApplicationCommand(s *Session, applicationID Snowflake, guildID Snowflake, commandID Snowflake) (err error) {
	endpoint := EndpointApplicationGuildCommand(applicationID.String(), guildID.String(), commandID.String())

	err = s.Interface.FetchJJ(s, http.MethodDelete, endpoint, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("Failed to delete guild application command: %v", err)
	}

	return
}

func BulkOverwriteGuildApplicationCommands(s *Session, applicationID Snowflake, guildID Snowflake, commandArgs []ApplicationCommand) (commands []*ApplicationCommand, err error) {
	endpoint := EndpointApplicationGuildCommands(applicationID.String(), guildID.String())

	err = s.Interface.FetchJJ(s, http.MethodPut, endpoint, commandArgs, nil, &commands)
	if err != nil {
		return nil, fmt.Errorf("Failed to bulk overwrite guild application commands: %v", err)
	}

	return
}

func GetGuildApplicationCommandPermissions(s *Session, applicationID Snowflake, guildID Snowflake) (permissions []GuildApplicationCommandPermissions, err error) {
	endpoint := EndpointApplicationGuildCommandsPermissions(applicationID.String(), guildID.String())

	err = s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &permissions)
	if err != nil {
		return nil, fmt.Errorf("Failed to get guild application command permissions: %v", err)
	}

	return
}

func GetApplicationCommandPermissions(s *Session, applicationID Snowflake, guildID Snowflake, commandID Snowflake) (permissions []GuildApplicationCommandPermissions, err error) {
	endpoint := EndpointApplicationGuildCommandPermissions(applicationID.String(), guildID.String(), commandID.String())

	err = s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &permissions)
	if err != nil {
		return nil, fmt.Errorf("Failed to get application command permissions: %v", err)
	}

	return
}

func EditApplicationCommandPermissions(s *Session, applicationID Snowflake, guildID Snowflake, commandID Snowflake, permissionsArg []GuildApplicationCommandPermissions) (permissions []*GuildApplicationCommandPermissions, err error) {
	endpoint := EndpointApplicationGuildCommandPermissions(applicationID.String(), guildID.String(), commandID.String())

	err = s.Interface.FetchJJ(s, http.MethodPut, endpoint, permissionsArg, nil, &permissions)
	if err != nil {
		return nil, fmt.Errorf("Failed to edit application command permissions: %v", err)
	}

	return
}

func BatchEditApplicationCommandPermissions(s *Session, applicationID Snowflake, guildID Snowflake, permissionsArg []GuildApplicationCommandPermissions) (permissions []*GuildApplicationCommandPermissions, err error) {
	endpoint := EndpointApplicationGuildCommandsPermissions(applicationID.String(), guildID.String())

	err = s.Interface.FetchJJ(s, http.MethodPut, endpoint, permissionsArg, nil, &permissions)
	if err != nil {
		return nil, fmt.Errorf("Failed to batch edit application command permissions: %v", err)
	}

	return
}
