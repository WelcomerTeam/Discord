package discord

import "context"

// emoji.go contains all structures for emojis.

// Emoji represents an Emoji on discord.
type Emoji struct {
	GuildID       *Snowflake    `json:"guild_id,omitempty"`
	User          *User         `json:"user,omitempty"`
	Name          string        `json:"name"`
	Roles         SnowflakeList `json:"roles,omitempty"`
	ID            Snowflake     `json:"id"`
	RequireColons bool          `json:"require_colons"`
	Managed       bool          `json:"managed"`
	Animated      bool          `json:"animated"`
	Available     bool          `json:"available"`
}

// Delete deletes the emoji.
// reason: Reason for deleting the emoji.
func (e *Emoji) Delete(ctx context.Context, session *Session, reason *string) error {
	return DeleteGuildEmoji(ctx, session, *e.GuildID, e.ID, reason)
}

// Edit edits the emoji.
// name: The name of the emoji
// roles: Roles this emoji is limited to.
// reason: Reason for editing the emoji.
func (e *Emoji) Edit(ctx context.Context, session *Session, name string, roles []Snowflake, reason *string) error {
	params := EmojiParams{
		Name:  name,
		Roles: roles,
	}

	newEmoji, err := ModifyGuildEmoji(ctx, session, *e.GuildID, e.ID, params, reason)
	if err != nil {
		return err
	}

	*e = *newEmoji

	return nil
}

// EmojiParams represents the payload sent to discord.
type EmojiParams struct {
	Name  string        `json:"name"`
	Image string        `json:"image,omitempty"`
	Roles SnowflakeList `json:"roles"`
}
