package structs

import "github.com/WelcomerTeam/Discord/discord"

// emoji.go contains all structures for emojis.

// Emoji represents an Emoji on discord.
type Emoji struct {
	ID            discord.Snowflake   `json:"id"`
	GuildID       *discord.Snowflake  `json:"guild_id,omitempty"`
	Name          string              `json:"name"`
	Roles         []discord.Snowflake `json:"roles,omitempty"`
	User          *User               `json:"user,omitempty"`
	RequireColons bool                `json:"require_colons"`
	Managed       bool                `json:"managed"`
	Animated      bool                `json:"animated"`
	Available     bool                `json:"available"`
}

// EmojiParams represents the payload sent to discord.
type EmojiParams struct {
	Name  string               `json:"name"`
	Image string               `json:"image,omitempty"`
	Roles []*discord.Snowflake `json:"roles"`
}
