package discord

// onboarding.go contains all structures for Discord guild onboarding.

// GuildOnboardingMode represents the criteria used to satisfy constraints.
type GuildOnboardingMode int32

const (
	GuildOnboardingModeDefault  GuildOnboardingMode = 0
	GuildOnboardingModeAdvanced GuildOnboardingMode = 1
)

// OnboardingPromptType represents the type of an onboarding prompt.
type OnboardingPromptType int32

const (
	OnboardingPromptTypeMultipleChoice OnboardingPromptType = 0
	OnboardingPromptTypeDropdown       OnboardingPromptType = 1
)

// GuildOnboarding represents the onboarding configuration for a guild.
type GuildOnboarding struct {
	GuildID          Snowflake            `json:"guild_id"`
	Prompts          []OnboardingPrompt   `json:"prompts"`
	DefaultChannelIDs []Snowflake         `json:"default_channel_ids"`
	Enabled          bool                 `json:"enabled"`
	Mode             *GuildOnboardingMode `json:"mode,omitempty"`
}

// OnboardingPrompt represents a prompt shown during guild onboarding.
type OnboardingPrompt struct {
	ID           Snowflake              `json:"id"`
	Title        string                 `json:"title"`
	Options      []OnboardingPromptOption `json:"options"`
	SingleSelect bool                   `json:"single_select"`
	Required     bool                   `json:"required"`
	InOnboarding bool                   `json:"in_onboarding"`
	Type         OnboardingPromptType   `json:"type"`
}

// OnboardingPromptOption represents an option in an onboarding prompt.
type OnboardingPromptOption struct {
	ID          Snowflake   `json:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	EmojiID     *Snowflake  `json:"emoji_id,omitempty"`
	EmojiName   *string     `json:"emoji_name,omitempty"`
	RoleIDs     []Snowflake `json:"role_ids"`
	ChannelIDs  []Snowflake `json:"channel_ids"`
}
