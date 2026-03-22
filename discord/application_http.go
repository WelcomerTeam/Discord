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

// ListApplicationEmojis lists all emojis for an application.
func ListApplicationEmojis(ctx context.Context, session *Session, applicationID Snowflake) ([]Emoji, error) {
	endpoint := "/applications/" + applicationID.String() + "/emojis"

	var emojis []Emoji

	err := session.Interface.FetchJJ(ctx, session, http.MethodGet, endpoint, nil, nil, &emojis)
	if err != nil {
		return nil, fmt.Errorf("failed to list application emojis: %w", err)
	}

	return emojis, nil
}

// GetApplicationEmoji retrieves a single emoji for an application.
func GetApplicationEmoji(ctx context.Context, session *Session, applicationID, emojiID Snowflake) (*Emoji, error) {
	endpoint := "/applications/" + applicationID.String() + "/emojis/" + emojiID.String()

	var emoji *Emoji

	err := session.Interface.FetchJJ(ctx, session, http.MethodGet, endpoint, nil, nil, &emoji)
	if err != nil {
		return nil, fmt.Errorf("failed to get application emoji: %w", err)
	}

	return emoji, nil
}

// CreateApplicationEmoji creates an emoji for an application.
func CreateApplicationEmoji(ctx context.Context, session *Session, applicationID Snowflake, emojiParams EmojiParams) (*Emoji, error) {
	endpoint := "/applications/" + applicationID.String() + "/emojis"

	var emoji *Emoji

	err := session.Interface.FetchJJ(ctx, session, http.MethodPost, endpoint, emojiParams, nil, &emoji)
	if err != nil {
		return nil, fmt.Errorf("failed to create application emoji: %w", err)
	}

	return emoji, nil
}

// UpdateApplicationEmoji updates an emoji for an application.
func UpdateApplicationEmoji(ctx context.Context, session *Session, applicationID, emojiID Snowflake, emojiParams EmojiParams) (*Emoji, error) {
	endpoint := "/applications/" + applicationID.String() + "/emojis/" + emojiID.String()

	var emoji *Emoji

	err := session.Interface.FetchJJ(ctx, session, http.MethodPatch, endpoint, emojiParams, nil, &emoji)
	if err != nil {
		return nil, fmt.Errorf("failed to update application emoji: %w", err)
	}

	return emoji, nil
}

// DeleteApplicationEmoji deletes an emoji from an application.
func DeleteApplicationEmoji(ctx context.Context, session *Session, applicationID, emojiID Snowflake) error {
	endpoint := "/applications/" + applicationID.String() + "/emojis/" + emojiID.String()

	err := session.Interface.FetchJJ(ctx, session, http.MethodDelete, endpoint, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete application emoji: %w", err)
	}

	return nil
}

// GetApplication returns the application object for the given application ID.
func GetApplication(ctx context.Context, session *Session, applicationID Snowflake) (*Application, error) {
	endpoint := EndpointApplication(applicationID.String())

	var application *Application

	err := session.Interface.FetchJJ(ctx, session, http.MethodGet, endpoint, nil, nil, &application)
	if err != nil {
		return nil, fmt.Errorf("failed to get application: %w", err)
	}

	return application, nil
}

// ModifyApplication modifies an application.
func ModifyApplication(ctx context.Context, session *Session, applicationID Snowflake, params Application) (*Application, error) {
	endpoint := EndpointApplication(applicationID.String())

	var application *Application

	err := session.Interface.FetchJJ(ctx, session, http.MethodPatch, endpoint, params, nil, &application)
	if err != nil {
		return nil, fmt.Errorf("failed to modify application: %w", err)
	}

	return application, nil
}

// ModifyCurrentApplication modifies the current application.
func ModifyCurrentApplication(ctx context.Context, session *Session, params Application) (*Application, error) {
	endpoint := EndpointApplications + "/@me"

	var application *Application

	err := session.Interface.FetchJJ(ctx, session, http.MethodPatch, endpoint, params, nil, &application)
	if err != nil {
		return nil, fmt.Errorf("failed to modify current application: %w", err)
	}

	return application, nil
}

// GetApplicationRoleConnectionMetadataRecords returns a list of role connection metadata records for an application.
func GetApplicationRoleConnectionMetadataRecords(ctx context.Context, session *Session, applicationID Snowflake) ([]ApplicationRoleConnectionMetadata, error) {
	endpoint := EndpointApplicationRoleConnectionsMetadata(applicationID.String())

	var records []ApplicationRoleConnectionMetadata

	err := session.Interface.FetchJJ(ctx, session, http.MethodGet, endpoint, nil, nil, &records)
	if err != nil {
		return nil, fmt.Errorf("failed to get application role connection metadata records: %w", err)
	}

	return records, nil
}

// UpdateApplicationRoleConnectionMetadataRecords updates the role connection metadata records for an application.
func UpdateApplicationRoleConnectionMetadataRecords(ctx context.Context, session *Session, applicationID Snowflake, params []ApplicationRoleConnectionMetadata) ([]ApplicationRoleConnectionMetadata, error) {
	endpoint := EndpointApplicationRoleConnectionsMetadata(applicationID.String())

	var records []ApplicationRoleConnectionMetadata

	err := session.Interface.FetchJJ(ctx, session, http.MethodPut, endpoint, params, nil, &records)
	if err != nil {
		return nil, fmt.Errorf("failed to update application role connection metadata records: %w", err)
	}

	return records, nil
}
