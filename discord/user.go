package discord

import (
	"context"
	"fmt"
)

// user.go represents all structures for a discord user.

// UserFlags represents the flags on a user's account.
type UserFlags uint32

// User flags.
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
	_
	_
	UserFlagsActiveDeveloper
)

// UserPremiumType represents the type of Nitro on a user's account.
type UserPremiumType int

// User premium type.
const (
	UserPremiumTypeNone UserPremiumType = iota
	UserPremiumTypeNitroClassic
	UserPremiumTypeNitro
	UserPremiumTypeNitroBasic
)

// User represents a user on discord.
type User struct {
	ID                   Snowflake             `json:"id"`
	DMChannelID          *Snowflake            `json:"dm_channel_id,omitempty"`
	AvatarDecorationData *AvatarDecorationData `json:"avatar_decoration_data,omitempty"`
	PrimaryGuild         *UserPrimaryGuild     `json:"primary_guild,omitempty"`
	Username             string                `json:"username"`
	Discriminator        string                `json:"discriminator"`
	Locale               string                `json:"locale,omitempty"`
	GlobalName           *string               `json:"global_name"`
	Avatar               *string               `json:"avatar"`
	Banner               *string               `json:"banner,omitempty"`
	Email                *string               `json:"email,omitempty"`
	AccentColor          *int32                `json:"accent_color,omitempty"`
	Flags                UserFlags             `json:"flags,omitempty"`
	PublicFlags          UserFlags             `json:"public_flags,omitempty"`
	PremiumType          UserPremiumType       `json:"premium_type,omitempty"`
	Bot                  bool                  `json:"bot,omitempty"`
	System               bool                  `json:"system,omitempty"`
	MFAEnabled           bool                  `json:"mfa_enabled,omitempty"`
	Verified             bool                  `json:"verified,omitempty"`
}

// AvatarDecorationData represents the avatar decoration data for a user.
type AvatarDecorationData struct {
	Asset string    `json:"asset"`
	SKUID Snowflake `json:"sku_id"`
}

// UserPrimaryGuild represents the primary guild information for a user.
type UserPrimaryGuild struct {
	IdentityGuildID *Snowflake `json:"identity_guild_id,omitempty"`
	IdentityEnabled *bool      `json:"identity_enabled,omitempty"`
	Tag             *string    `json:"tag,omitempty"`
	Badge           *string    `json:"badge,omitempty"`
}

// CreateDM creates a DMChannel with a user. This should not need to be called as Send() transparently does this.
// If the user already has a DMChannel created, this will return a partial channel with just an ID set.
func (u *User) CreateDM(ctx context.Context, session *Session) (*Channel, error) {
	if u.DMChannelID != nil {
		return &Channel{ID: *u.DMChannelID}, nil
	}

	var channel *Channel

	channel, err := CreateDM(ctx, session, u.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to create DM channel: %w", err)
	}

	u.DMChannelID = &channel.ID

	return channel, nil
}

// Send sends a DM message to a user. This will create a DMChannel if one is not present.
// params: The message parameters used to send the message.
func (u *User) Send(ctx context.Context, session *Session, params MessageParams) (*Message, error) {
	dmChannel, err := u.CreateDM(ctx, session)
	if err != nil {
		return nil, fmt.Errorf("failed to create DM channel: %w", err)
	}

	return dmChannel.Send(ctx, session, params)
}

// ClientUser aliases User to provide current user specific methods.
type ClientUser User

// Edit modifies the current user.
// username: The new username to change to.
// avatar: File of new avatar to change to.
func (u *ClientUser) Edit(ctx context.Context, session *Session, username *string, avatar *[]byte) error {
	params := UserParam{
		Username: username,
	}

	if avatar != nil {
		avatarBase64, err := bytesToBase64Data(*avatar)
		if err != nil {
			return err
		}

		params.Avatar = &avatarBase64
	}

	newUser, err := ModifyCurrentUser(ctx, session, params)
	if err != nil {
		return err
	}

	*u = *(*ClientUser)(newUser)

	return nil
}

// UserParam represents the payload sent to modify a user.
type UserParam struct {
	Username *string `json:"username,omitempty"`
	Avatar   *string `json:"avatar,omitempty"`
}
