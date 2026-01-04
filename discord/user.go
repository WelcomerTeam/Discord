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
)

// User represents a user on discord.
type User struct {
	DMChannelID   *Snowflake      `json:"dm_channel_id,omitempty"`
	Banner        string          `json:"banner,omitempty"`
	GlobalName    string          `json:"global_name"`
	Avatar        string          `json:"avatar"`
	Username      string          `json:"username"`
	Discriminator string          `json:"discriminator"`
	Locale        string          `json:"locale,omitempty"`
	Email         string          `json:"email,omitempty"`
	ID            Snowflake       `json:"id"`
	PremiumType   UserPremiumType `json:"premium_type,omitempty"`
	Flags         UserFlags       `json:"flags,omitempty"`
	AccentColor   int32           `json:"accent_color"`
	PublicFlags   UserFlags       `json:"public_flags,omitempty"`
	MFAEnabled    bool            `json:"mfa_enabled,omitempty"`
	Verified      bool            `json:"verified,omitempty"`
	Bot           bool            `json:"bot"`
	System        bool            `json:"system,omitempty"`
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
