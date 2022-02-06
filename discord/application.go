package discord

import (
	jsoniter "github.com/json-iterator/go"
)

// application.go represents the application object and integrations.

// ApplicationTeamMemberState represents the state of a member in a team.
type ApplicationTeamMemberState uint8

const (
	ApplicationTeamMemberStateInvited ApplicationTeamMemberState = 1 + iota
	ApplicationTeamMemberStateAccepted
)

// ApplicationCommandType represents the different types of application command.
type ApplicationCommandType uint8

const (
	ApplicationCommandTypeChatInput ApplicationCommandType = 1 + iota
	ApplicationCommandTypeUser
	ApplicationCommandTypeMessage
)

// ApplicationCommandOptionType represents the different types of options.
type ApplicationCommandOptionType uint8

const (
	ApplicationCommandOptionTypeSubCommand ApplicationCommandOptionType = 1 + iota
	ApplicationCommandOptionTypeSubCommandGroup
	ApplicationCommandOptionTypeString
	ApplicationCommandOptionTypeInteger
	ApplicationCommandOptionTypeBoolean
	ApplicationCommandOptionTypeUser
	ApplicationCommandOptionTypeChannel
	ApplicationCommandOptionTypeRole
	ApplicationCommandOptionTypeMentionable
	ApplicationCommandOptionTypeNumber
)

// ApplicationCommandPermissionType represents the target for a command permission.
type ApplicationCommandPermissionType uint8

const (
	ApplicationCommandPermissionTypeRole ApplicationCommandPermissionType = 1 + iota
	ApplicationCommandPermissionTypeUser
)

// IntegrationType represents the type of integration.
type IntegrationType string

const (
	IntegrationTypeTwitch  IntegrationType = "twitch"
	IntegrationTypeYoutube IntegrationType = "youtube"
	IntegrationTypeDiscord IntegrationType = "discord"
)

// IntegrationExpireBehavior represents the integration expiration.
type IntegrationExpireBehavior uint8

const (
	IntegrationExpireBehaviorRemoveRole IntegrationExpireBehavior = iota
	IntegrationExpireBehaviorKick
)

// Application response from REST.
type Application struct {
	ID                  Snowflake        `json:"id"`
	Name                string           `json:"name"`
	Icon                string           `json:"icon,omitempty"`
	Description         string           `json:"description"`
	RPCOrigins          []string         `json:"rpc_origins,omitempty"`
	BotPublic           bool             `json:"bot_public"`
	BotRequireCodeGrant bool             `json:"bot_require_code_grant"`
	TermsOfServiceURL   string           `json:"terms_of_service,omitempty"`
	PrivacyPolicyURL    string           `json:"privacy_policy_url,omitempty"`
	Owner               *User            `json:"owner,omitempty"`
	Summary             string           `json:"summary"`
	VerifyKey           string           `json:"verify_key"`
	Team                *ApplicationTeam `json:"team,omitempty"`
	GuildID             *Snowflake       `json:"guild_id,omitempty"`
	PrimarySKUID        *Snowflake       `json:"primary_sku_id,omitempty"`
	Slug                string           `json:"slug,omitempty"`
	CoverImage          string           `json:"cover_image,omitempty"`
	Flags               int32            `json:"flags,omitempty"`
	Bot                 *User            `json:"bot,omitempty"`
}

// ApplicationTeam represents the team of an application.
type ApplicationTeam struct {
	Icon        string                   `json:"icon,omitempty"`
	ID          Snowflake                `json:"id"`
	Members     []*ApplicationTeamMember `json:"members"`
	Name        string                   `json:"name"`
	OwnerUserID Snowflake                `json:"owner_user_id"`
}

// ApplicationTeamMembers represents a member of a team.
type ApplicationTeamMember struct {
	MembershipState *ApplicationTeamMemberState `json:"membership_state"`
	Permissions     []string                    `json:"permissions"`
	TeamID          Snowflake                   `json:"team_id"`
	User            User                        `json:"user"`
}

// ApplicationCommand represents an application's command.
type ApplicationCommand struct {
	ID                Snowflake                   `json:"id"`
	Type              *ApplicationCommandType     `json:"type,omitempty"`
	ApplicationID     Snowflake                   `json:"application_id"`
	GuildID           *Snowflake                  `json:"guild_id,omitempty"`
	Name              string                      `json:"name"`
	Description       string                      `json:"description"`
	Options           []*ApplicationCommandOption `json:"options,omitempty"`
	DefaultPermission bool                        `json:"default_permission"`
	Version           int32                       `json:"version"`
}

// GuildApplicationCommandPermissions represent a guilds application permissions.
type GuildApplicationCommandPermissions struct {
	ID            Snowflake                        `json:"id"`
	ApplicationID Snowflake                        `json:"application_id"`
	GuildID       Snowflake                        `json:"guild_id"`
	Permissions   []*ApplicationCommandPermissions `json:"permissions"`
}

// ApplicationCommandPermissions represents the rules for enabling or disabling a command.
type ApplicationCommandPermissions struct {
	ID      Snowflake                        `json:"id"`
	Type    ApplicationCommandPermissionType `json:"type"`
	Allowed bool                             `json:"permission"`
}

// ApplicationCommandOption represents the options for an application command.
type ApplicationCommandOption struct {
	Type         ApplicationCommandOptionType      `json:"type"`
	Name         string                            `json:"name"`
	Description  string                            `json:"description"`
	Required     bool                              `json:"required"`
	Choices      []*ApplicationCommandOptionChoice `json:"choices,omitempty"`
	Options      []*ApplicationCommandOption       `json:"options,omitempty"`
	ChannelTypes []*ChannelType                    `json:"channel_types,omitempty"`
	MinValue     int32                             `json:"min_value,omitempty"`
	MaxValue     int32                             `json:"max_value,omitempty"`
	Autocomplete bool                              `json:"autocomplete"`
}

// ApplicationCommandOptionChoice represents the different choices.
type ApplicationCommandOptionChoice struct {
	Name  string              `json:"name"`
	Value jsoniter.RawMessage `json:"value"`
}

// ApplicationSelectOption represents the structure of select options.
type ApplicationSelectOption struct {
	Label       string `json:"label"`
	Value       string `json:"value"`
	Description string `json:"description,omitempty"`
	Emoji       *Emoji `json:"emoji,omitempty"`
	Default     bool   `json:"default"`
}

// Integration represents the structure of an integration.
type Integration struct {
	ID                Snowflake                  `json:"id"`
	GuildID           *Snowflake                 `json:"guild_id,omitempty"`
	Name              string                     `json:"name"`
	Type              IntegrationType            `json:"type"`
	Enabled           bool                       `json:"enabled"`
	Syncing           bool                       `json:"syncing"`
	RoleID            *Snowflake                 `json:"role_id,omitempty"`
	EnableEmoticons   bool                       `json:"enable_emoticons"`
	ExpireBehavior    *IntegrationExpireBehavior `json:"expire_behavior,omitempty"`
	ExpireGracePeriod int32                      `json:"expire_grace_period,omitempty"`
	User              *User                      `json:"user,omitempty"`
	Account           IntegrationAccount         `json:"account"`
	SyncedAt          string                     `json:"synced_at,omitempty"`
	SubscriberCount   int32                      `json:"subscriber_count,omitempty"`
	Revoked           bool                       `json:"revoked"`
	Application       *Application               `json:"application,omitempty"`
}

// IntegrationAccount represents the account of the integration.
type IntegrationAccount struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
