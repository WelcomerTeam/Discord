package discord

import "github.com/WelcomerTeam/Discord/discord"

// role.go represents all structures for a discord guild role.

// Role represents a role on Discord.
type Role struct {
	ID           discord.Snowflake  `json:"id"`
	GuildID      *discord.Snowflake `json:"guild_id,omitempty"`
	Name         string             `json:"name"`
	Color        int32              `json:"color"`
	Hoist        bool               `json:"hoist"`
	Icon         string             `json:"icon,omitempty"`
	UnicodeEmoji string             `json:"unicode_emoji,omitempty"`
	Position     int32              `json:"position"`
	Permissions  discord.Int64      `json:"permissions"`
	Managed      bool               `json:"managed"`
	Mentionable  bool               `json:"mentionable"`
	Tags         *RoleTag           `json:"tags,omitempty"`
}

// RoleTag represents extra information about a role.
type RoleTag struct {
	PremiumSubscriber bool               `json:"premium_subscriber"`
	BotID             *discord.Snowflake `json:"bot_id,omitempty"`
	IntegrationID     *discord.Snowflake `json:"integration_id,omitempty"`
}
