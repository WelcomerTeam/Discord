package discord

// user.go represents all structures for a discord user.

// UserFlags represents the flags on a user's account.
type UserFlags int

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
	ID            Snowflake        `json:"id"`
	Username      string           `json:"username"`
	Discriminator string           `json:"discriminator"`
	Avatar        string           `json:"avatar"`
	Bot           bool             `json:"bot"`
	System        bool             `json:"system,omitempty"`
	MFAEnabled    bool             `json:"mfa_enabled,omitempty"`
	Banner        string           `json:"banner,omitempty"`
	AccentColour  int32            `json:"accent_color"`
	Locale        string           `json:"locale,omitempty"`
	Verified      bool             `json:"verified,omitempty"`
	Email         string           `json:"email,omitempty"`
	Flags         *UserFlags       `json:"flags,omitempty"`
	PremiumType   *UserPremiumType `json:"premium_type,omitempty"`
	PublicFlags   *UserFlags       `json:"public_flags,omitempty"`
	DMChannelID   *Snowflake       `json:"dm_channel_id,omitempty"`
}

// CreateDM creates a DMChannel with a user. This should not need to be called as Send() transparently does this.
// If the user already has a DMChannel created, this will return a partial channel with just an ID set.
func (u *User) CreateDM(s *Session) (channel *Channel, err error) {
	if u.DMChannelID != nil {
		return &Channel{ID: *u.DMChannelID}, nil
	}

	channel, err = CreateDM(s, u.ID)
	if err != nil {
		return
	}

	u.DMChannelID = &channel.ID

	return
}

// Send sends a DM message to a user. This will create a DMChannel if one is not present.
// params: The message parameters used to send the message.
func (u *User) Send(s *Session, params MessageParams) (message *Message, err error) {
	dmChannel, err := u.CreateDM(s)
	if err != nil {
		return nil, err
	}

	return dmChannel.Send(s, params)
}

// ClientUser aliases User to provide current user specific methods.
type ClientUser User

// Edit modifies the current user.
// username: The new username to change to.
// avatar: File of new avatar to change to.
func (u *ClientUser) Edit(s *Session, username *string, avatar *[]byte) (err error) {
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

	newUser, err := ModifyCurrentUser(s, params)
	if err != nil {
		return
	}

	*u = *(*ClientUser)(newUser)

	return
}

// UserParam represents the payload sent to modify a user.
type UserParam struct {
	Username *string `json:"username,omitempty"`
	Avatar   *string `json:"avatar,omitempty"`
}
