package structs

import "github.com/WelcomerTeam/Discord/discord"

// user.go represents all structures for a discord user.

// UserFlags represents the flags on a user's account.
type UserFlags int

// User flags.
const UserFlagsNone UserFlags = 0

const (
	UserFlagsDiscordEmployee UserFlags = 1 << iota
	UserFlagsPartneredServerOwner
	UserFlagsHypeSquadEvents
	UserFlagsBugHunterLevel1
	_
	_
	UserFlagsHouseBravery
	UserFlagsHouseBrilliance
	UserFlagsHouseBalance
	UserFlagsEarlySupporter
	UserFlagsTeamUser
	_
	_
	_
	UserFlagsBugHunterLevel2
	_
	UserFlagsVerifiedBot
	UserFlagsVerifiedDeveloper
	UserFlagsCertifiedModerator
	UserFlagsBotHTTPInteractions
)

// UserPremiumType represents the type of Nitro on a user's account.
type UserPremiumType int

// User premium type.
const (
	UserPremiumTypeNone UserPremiumType = iota
	UserPremiumTypeNitroClassic
	UserPremiumTypeNitro
)

// User represents a user on Discord.
type User struct {
	ID            discord.Snowflake  `json:"id"`
	Username      string             `json:"username"`
	Discriminator string             `json:"discriminator"`
	Avatar        string             `json:"avatar"`
	Bot           bool               `json:"bot"`
	System        bool               `json:"system,omitempty"`
	MFAEnabled    bool               `json:"mfa_enabled,omitempty"`
	Banner        string             `json:"banner,omitempty"`
	AccentColour  int32              `json:"accent_color"`
	Locale        string             `json:"locale,omitempty"`
	Verified      bool               `json:"verified,omitempty"`
	Email         string             `json:"email,omitempty"`
	Flags         *UserFlags         `json:"flags,omitempty"`
	PremiumType   *UserPremiumType   `json:"premium_type,omitempty"`
	PublicFlags   *UserFlags         `json:"public_flags,omitempty"`
	DMChannelID   *discord.Snowflake `json:"dm_channel_id,omitempty"`
}

// UserParam represents the payload sent to modify a user.
type UserParam struct {
	Username string `json:"username"`
	Avatar   string `json:"avatar,omitempty"`
}
