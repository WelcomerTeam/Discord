package discord

// soundboard.go contains all structures for Discord soundboard sounds.

// SoundboardSound represents a soundboard sound on Discord.
type SoundboardSound struct {
	User      *User      `json:"user,omitempty"`
	EmojiID   *Snowflake `json:"emoji_id,omitempty"`
	GuildID   *Snowflake `json:"guild_id,omitempty"`
	EmojiName string     `json:"emoji_name,omitempty"`
	Name      string     `json:"name"`
	SoundID   Snowflake  `json:"sound_id"`
	Volume    float64    `json:"volume"`
	Available bool       `json:"available"`
}

// SoundboardSoundSendParams represents the payload for sending a soundboard sound.
type SoundboardSoundSendParams struct {
	SoundID      Snowflake  `json:"sound_id"`
	SourceGuildID *Snowflake `json:"source_guild_id,omitempty"`
}

// CreateSoundboardSoundParams represents the payload for creating a soundboard sound.
type CreateSoundboardSoundParams struct {
	Name      string     `json:"name"`
	Sound     string     `json:"sound"`
	Volume    *float64   `json:"volume,omitempty"`
	EmojiID   *Snowflake `json:"emoji_id,omitempty"`
	EmojiName *string    `json:"emoji_name,omitempty"`
}

// ModifySoundboardSoundParams represents the payload for modifying a soundboard sound.
type ModifySoundboardSoundParams struct {
	Name      *string    `json:"name,omitempty"`
	Volume    *float64   `json:"volume,omitempty"`
	EmojiID   *Snowflake `json:"emoji_id,omitempty"`
	EmojiName *string    `json:"emoji_name,omitempty"`
}
