package discord

import (
	"time"
)

// guild.go contains the structures to represent a guild.

// MessageNotificationLevel represents a guild's message notification level.
type MessageNotificationLevel int

// Message notification levels.
const (
	MessageNotificationsAllMessages MessageNotificationLevel = iota
	MessageNotificationsOnlyMentions
)

// ExplicitContentFilterLevel represents a guild's explicit content filter level.
type ExplicitContentFilterLevel int

// Explicit content filter levels.
const (
	ExplicitContentFilterDisabled ExplicitContentFilterLevel = iota
	ExplicitContentFilterMembersWithoutRoles
	ExplicitContentFilterAllMembers
)

// MFALevel represents a guild's MFA level.
type MFALevel int8

// MFA levels.
const (
	MFALevelNone MFALevel = iota
	MFALevelElevated
)

// VerificationLevel represents a guild's verification level.
type VerificationLevel int8

const (
	VerificationLevelNone VerificationLevel = iota
	VerificationLevelLow
	VerificationLevelMedium
	VerificationLevelHigh
	VerificationLevelVeryHigh
)

// SystemChannelFlags represents the flags of a system channel.
type SystemChannelFlags int8

const (
	SystemChannelFlagsSuppressJoin SystemChannelFlags = 1 << iota
	SystemChannelFlagsPremiumSubscriptions
)

// PremiumTier represents the current boosting tier of a guild.
type PremiumTier int8

const (
	PremiumTierNone PremiumTier = iota
	PremiumTier1
	PremiumTier2
	PremiumTier3
)

// GuildNSFWLevelType represents the level of the guild.
type GuildNSFWLevelType int8

const (
	GuildNSFWLevelTypDefault GuildNSFWLevelType = iota
	GuildNSFWLevelTypeExplicit
	GuildNSFWLevelTypeSafe
	GuildNSFWLevelTypeAgeRestricted
)

// Guild represents a guild on discord.
type Guild struct {
	ID              Snowflake `json:"id"`
	Name            string    `json:"name"`
	Icon            string    `json:"icon"`
	IconHash        string    `json:"icon_hash"`
	Splash          string    `json:"splash"`
	DiscoverySplash string    `json:"discovery_splash"`

	Owner       bool       `json:"owner"`
	OwnerID     *Snowflake `json:"owner_id,omitempty"`
	Permissions *Int64     `json:"permissions,omitempty"`
	Region      string     `json:"region"`

	AFKChannelID *Snowflake `json:"afk_channel_id,omitempty"`
	AFKTimeout   int32      `json:"afk_timeout"`

	WidgetEnabled   bool       `json:"widget_enabled"`
	WidgetChannelID *Snowflake `json:"widget_channel_id,omitempty"`

	VerificationLevel           VerificationLevel          `json:"verification_level"`
	DefaultMessageNotifications MessageNotificationLevel   `json:"default_message_notifications"`
	ExplicitContentFilter       ExplicitContentFilterLevel `json:"explicit_content_filter"`

	Roles    []*Role  `json:"roles"`
	Emojis   []*Emoji `json:"emojis"`
	Features []string `json:"features"`

	MFALevel           MFALevel            `json:"mfa_level"`
	ApplicationID      *Snowflake          `json:"application_id,omitempty"`
	SystemChannelID    *Snowflake          `json:"system_channel_id,omitempty"`
	SystemChannelFlags *SystemChannelFlags `json:"system_channel_flags,omitempty"`
	RulesChannelID     *Snowflake          `json:"rules_channel_id,omitempty"`

	JoinedAt    time.Time `json:"joined_at"`
	Large       bool      `json:"large"`
	Unavailable bool      `json:"unavailable"`
	MemberCount int32     `json:"member_count"`

	VoiceStates []*VoiceState  `json:"voice_states,omitempty"`
	Members     []*GuildMember `json:"members,omitempty"`
	Channels    []*Channel     `json:"channels,omitempty"`
	Presences   []*Activity    `json:"presences,omitempty"`

	MaxPresences  int32  `json:"max_presences"`
	MaxMembers    int32  `json:"max_members"`
	VanityURLCode string `json:"vanity_url_code"`
	Description   string `json:"description"`
	Banner        string `json:"banner"`

	PremiumTier              *PremiumTier        `json:"premium_tier,omitempty"`
	PremiumSubscriptionCount int32               `json:"premium_subscription_count"`
	PreferredLocale          string              `json:"preferred_locale"`
	PublicUpdatesChannelID   *Snowflake          `json:"public_updates_channel_id,omitempty"`
	MaxVideoChannelUsers     int32               `json:"max_video_channel_users"`
	ApproximateMemberCount   int32               `json:"approximate_member_count"`
	ApproximatePresenceCount int32               `json:"approximate_presence_count"`
	NSFWLevel                *GuildNSFWLevelType `json:"nsfw_level"`

	StageInstances            []*StageInstance  `json:"stage_instances,omitempty"`
	Stickers                  []*Sticker        `json:"stickers"`
	GuildScheduledEvents      []*ScheduledEvent `json:"guild_scheduled_events"`
	PremiumProgressBarEnabled bool              `json:"premium_progress_bar_enabled"`
}

// GuildParams represents the parameters sent when modifying a guild.
type GuildParam struct {
	Name                        *string                     `json:"name,omitempty"`
	Region                      *string                     `json:"region,omitempty"`
	VerificationLevel           *VerificationLevel          `json:"verification_level,omitempty"`
	DefaultMessageNotifications *MessageNotificationLevel   `json:"default_message_notifications,omitempty"`
	ExplicitContentFilter       *ExplicitContentFilterLevel `json:"explicit_content_filter,omitempty"`
	Icon                        *string                     `json:"icon,omitempty"`
	OwnerID                     *Snowflake                  `json:"owner_id,omitempty"`
	Splash                      *string                     `json:"splash,omitempty"`
	Banner                      *string                     `json:"banner,omitempty"`
	DiscoverySplash             *string                     `json:"discovery_splash,omitempty"`
	AFKChannelID                *Snowflake                  `json:"afk_channel_id,omitempty"`
	AFKTimeout                  *int32                      `json:"afk_timeout,omitempty"`
	SystemChannelID             *Snowflake                  `json:"system_channel_id,omitempty"`
	SystemChannelFlags          *SystemChannelFlags         `json:"system_channel_flags,omitempty"`
	RulesChannelID              *Snowflake                  `json:"rules_channel_id,omitempty"`
	PublicUpdatesChannelID      *Snowflake                  `json:"public_updates_channel_id,omitempty"`
	PreferredLocale             *string                     `json:"preferred_locale,omitempty"`
	Features                    []string                    `json:"features,omitempty"`
	Description                 *string                     `json:"description,omitempty"`
	PremiumProgressBarEnabled   *bool                       `json:"premium_progress_bar_enabled,omitempty"`
}

// AuditLogs returns all audit logs matching query.
// userID: Filters audit logs by the userID provided.
// actionType: The action type to filter audit logs by.
// before: Only show audit logs before a certain snowflake.
// limit: Maximum number of audit log entries to return.
func (g *Guild) AuditLogs(s *Session, guildID Snowflake, userID *Snowflake, actionType *AuditLogActionType, before *Snowflake, limit *int32) ([]*AuditLogEntry, error) {
	return GetGuildAuditLog(s, g.ID, userID, actionType, before, limit)
}

// Ban bans a user.
// userID: ID of user that is getting banned.
// reason: Reason for ban.
func (g *Guild) Ban(s *Session, userID Snowflake, reason *string) error {
	return CreateGuildBan(s, g.ID, userID, reason)
}

// Bans returns a list of guild bans.
func (g *Guild) Bans(s *Session) ([]*GuildBan, error) {
	return GetGuildBans(s, g.ID)
}

// CloneChannel creates a copy of the target channel.
// reason: Reason for creating the channel.
func (g *Guild) CloneChannel(s *Session, c *Channel, reason *string) (*Channel, error) {
	return g.CreateChannel(s, ChannelParams{
		Name:                 c.Name,
		Type:                 c.Type,
		Topic:                c.Topic,
		Bitrate:              c.Bitrate,
		UserLimit:            c.UserLimit,
		RateLimitPerUser:     c.RateLimitPerUser,
		Position:             c.Position,
		PermissionOverwrites: c.PermissionOverwrites,
		ParentID:             c.ParentID,
		NSFW:                 c.NSFW,
	}, reason)
}

// CreateChannel creates a channel.
// channelArg: Parameters passed for creating a channel.
// reason: Reason for creating the channel.
func (g *Guild) CreateChannel(s *Session, channelParams ChannelParams, reason *string) (*Channel, error) {
	return CreateGuildChannel(s, g.ID, channelParams, reason)
}

// CreateCustomEmojis creates an emoji for a guild.
// name: Name of the custom emoji.
// image: Bytes representing the image file to upload.
// roles: Roles that this emoji is limited to.
// reason: Reason for creating the emoji.
func (g *Guild) CreateCustomEmoji(s *Session, name string, image []byte, roles []*Snowflake, reason *string) (*Emoji, error) {
	params := EmojiParams{
		Name:  name,
		Roles: roles,
	}

	imageData, err := bytesToBase64Data(image)
	if err != nil {
		return nil, err
	}

	params.Image = imageData

	return CreateGuildEmoji(s, g.ID, params, reason)
}

// CreateRole creates a role.
// roleArg: Parameters passed for creating a role.
// reason: Reason for creating the role.
func (g *Guild) CreateRole(s *Session, roleParams RoleParams, reason *string) (*Role, error) {
	return CreateGuildRole(s, g.ID, roleParams, reason)
}

// Delete deletes a guild.
func (g *Guild) Delete(s *Session) error {
	return DeleteGuild(s, g.ID)
}

// Edit edits a guild.
// guildArg: Parameters passed for editing a guild.
// reason: Reason for editing the guild.
func (g *Guild) Edit(s *Session, guildArg GuildParam, reason *string) error {
	newGuild, err := ModifyGuild(s, g.ID, guildArg, reason)
	if err != nil {
		return err
	}

	*g = *newGuild

	return nil
}

// EditRolePositions edits role positions in a guild.
// guildRolePositionArgs: List of roles and their new role position.
func (g *Guild) EditRolePositions(s *Session, guildRolePositionArgs []ModifyGuildRolePosition, reason *string) ([]*Role, error) {
	return ModifyGuildRolePositions(s, g.ID, guildRolePositionArgs, reason)
}

// EstimatePrunedMembers returns an estimate of how many people will be pruned from a guild based on arguments.
// days: The number of days since speaking.
// includedRoles: By default pruning only removes users with no roles, any role in this list will be included.
func (g *Guild) EstimatePrunedMembers(s *Session, days *int32, includedRoles []Snowflake) (*int32, error) {
	return GetGuildPruneCount(s, g.ID, days, includedRoles)
}

// Integrations returns all guild integrations.
func (g *Guild) Integrations(s *Session) ([]*Integration, error) {
	return GetGuildIntegrations(s, g.ID)
}

// Invites returns all guild invites.
func (g *Guild) Invites(s *Session) ([]*Invite, error) {
	return GetGuildInvites(s, g.ID)
}

// Kick kicks a user from the guild.
// userID: ID of user to kick.
// reason: Reason for kicking the user.
func (g *Guild) Kick(s *Session, userID Snowflake, reason *string) error {
	return RemoveGuildMember(s, g.ID, userID, reason)
}

// Leave leaves a guild.
func (g *Guild) Leave(s *Session) error {
	return LeaveGuild(s, g.ID)
}

// PruneMembers prunes users from a guild based on arguments.
// days: The number of days since speaking.
// includedRoles: By default pruning only removes users with no roles, any role in this list will be included.
// computePruneCount: Returns how many users were pruned, usage on larger guilds is discouraged.
// reason: Reason for pruning members.
func (g *Guild) PruneMembers(s *Session, guildID Snowflake, days *int32, includedRoles []Snowflake, computePruneCount bool, reason *string) (*int32, error) {
	return BeginGuildPrune(s, g.ID, days, includedRoles, computePruneCount, reason)
}

// QueryMembers returns guild members whose username or nickname matches query.
// query: Query string to match usernames and nicknames against.
// limit: Maximum number of members to return.
func (g *Guild) QueryMembers(s *Session, query string, limit *int32) ([]*GuildMember, error) {
	return SearchGuildMembers(s, g.ID, query, limit)
}

// Unban unbans a user from a guild.
// userID: ID of user to unban.
// reason: Reason for unbanning.
func (g *Guild) Unban(s *Session, userID Snowflake, reason *string) error {
	return RemoveGuildBan(s, g.ID, userID, reason)
}

// VanityInvite returns the vanity invite for a guild.
func (g *Guild) VanityInvite(s *Session) (*Invite, error) {
	invite, err := GetGuildVanityURL(s, g.ID)
	if err != nil {
		return invite, err
	}

	g.VanityURLCode = invite.Code

	return invite, nil
}

// Webhooks returns all webhooks for a guild.
func (g *Guild) Webhooks(s *Session) ([]*Webhook, error) {
	return GetGuildWebhooks(s, g.ID)
}

// UnavailableGuild represents an unavailable guild.
type UnavailableGuild struct {
	ID          Snowflake `json:"id"`
	Unavailable bool      `json:"unavailable"`
}

// GuildMember represents a guild member on discord.
type GuildMember struct {
	User                       *User       `json:"user,omitempty"`
	GuildID                    *Snowflake  `json:"guild_id,omitempty"`
	Nick                       string      `json:"nick,omitempty"`
	Avatar                     string      `json:"avatar,omitempty"`
	Roles                      []Snowflake `json:"roles"`
	JoinedAt                   time.Time   `json:"joined_at"`
	PremiumSince               string      `json:"premium_since,omitempty"`
	Deaf                       bool        `json:"deaf"`
	Mute                       bool        `json:"mute"`
	Pending                    bool        `json:"pending"`
	Permissions                *Int64      `json:"permissions"`
	CommunicationDisabledUntil string      `json:"communication_disabled_until,omitempty"`
}

// GuildMemberParams represents the arguments used to modify a guild member.
type GuildMemberParams struct {
	Nick                       *string     `json:"nick,omitempty"`
	Roles                      []Snowflake `json:"roles,omitempty"`
	Deaf                       *bool       `json:"deaf,omitempty"`
	Mute                       *bool       `json:"mute,omitempty"`
	ChannelID                  *Snowflake  `json:"channel_id,omitempty"`
	CommunicationDisabledUntil *string     `json:"communication_disabled_until,omitempty"`
}

// AddRoles adds roles to a guild member.
// roles: List of roles to add to the guild member.
// reason: Reason for adding the roles to the guild member.
// atomic: When true, will send multiple AddGuildMemberRole requests instead of at once.
func (gm *GuildMember) AddRoles(s *Session, roles []Snowflake, reason *string, atomic bool) error {
	guildMemberRoles := make(map[Snowflake]bool)

	for _, guildMemberRole := range gm.Roles {
		guildMemberRoles[guildMemberRole] = true
	}

	if atomic {
		for _, roleID := range roles {
			if _, ok := guildMemberRoles[roleID]; !ok {
				err := AddGuildMemberRole(s, *gm.GuildID, gm.User.ID, roleID, reason)
				if err != nil {
					return err
				}

				gm.Roles = append(gm.Roles, roleID)
			}
		}

		return nil
	}

	for _, addedRoleID := range roles {
		guildMemberRoles[addedRoleID] = true
	}

	newRoles := make([]Snowflake, 0, len(guildMemberRoles))

	for roleID := range guildMemberRoles {
		newRoles = append(newRoles, roleID)
	}

	return gm.Edit(s, GuildMemberParams{Roles: newRoles}, reason)
}

// Ban bans the guild member from the guild.
// reason: Reason for banning the guild member.
func (gm *GuildMember) Ban(s *Session, reason *string) error {
	return CreateGuildBan(s, *gm.GuildID, gm.User.ID, reason)
}

// CreateDM creates a DMChannel with a user. This should not need to be called as Send() transparently does this.
// If the user already has a DMChannel created, this will return a partial channel with just an ID set.
func (gm *GuildMember) CreateDM(s *Session) (*Channel, error) {
	return gm.User.CreateDM(s)
}

// Edit edits a guild member.
// guildMemberArg: Parameters used to update a guild member.
// reason: Reason for editing the guild member.
func (gm *GuildMember) Edit(s *Session, guildMemberParams GuildMemberParams, reason *string) error {
	newMember, err := ModifyGuildMember(s, *gm.GuildID, gm.User.ID, guildMemberParams, reason)
	if err != nil {
		return err
	}

	*gm = *newMember

	return nil
}

// Kick kicks the guild member.
// reason: Reason for kicking the guild member.
func (gm *GuildMember) Kick(s *Session, reason *string) error {
	return RemoveGuildMember(s, *gm.GuildID, gm.User.ID, reason)
}

// MoveTo moves the guild member to a different voice channel.
// channelID: Channel to move the user to, if nil they are removed from voice.
// reason: Reason for moving the guild member
func (gm *GuildMember) MoveTo(s *Session, channelID *Snowflake, reason *string) error {
	return gm.Edit(s, GuildMemberParams{ChannelID: channelID}, reason)
}

// RemoveRoles removes roles from a guild member.
func (gm *GuildMember) RemoveRoles(s *Session, roles []Snowflake, reason *string, atomic bool) error {
	guildMemberRoles := make(map[Snowflake]bool)

	for _, guildMemberRole := range gm.Roles {
		guildMemberRoles[guildMemberRole] = true
	}

	if atomic {
		for _, roleID := range roles {
			if _, ok := guildMemberRoles[roleID]; ok {
				err := RemoveGuildMemberRole(s, *gm.GuildID, gm.User.ID, roleID, reason)
				if err != nil {
					return err
				}

				delete(guildMemberRoles, roleID)

				// Remove role from guild member roles.
				// Whilst inefficient, we reconstruct the roles
				// on every pass incase one errors.

				newRoles := make([]Snowflake, 0, len(guildMemberRoles))

				for roleID := range guildMemberRoles {
					newRoles = append(newRoles, roleID)
				}

				gm.Roles = newRoles
			}
		}

		return nil
	}

	for _, removedRoleID := range roles {
		delete(guildMemberRoles, removedRoleID)
	}

	newRoles := make([]Snowflake, 0, len(guildMemberRoles))

	for roleID := range guildMemberRoles {
		newRoles = append(newRoles, roleID)
	}

	return gm.Edit(s, GuildMemberParams{Roles: newRoles}, reason)
}

// Send sends a DM message to a user. This will create a DMChannel if one is not present.
// params: The message parameters used to send the message.
func (gm *GuildMember) Send(s *Session, params MessageParams) (*Message, error) {
	return gm.User.Send(s, params)
}

// VoiceState represents the voice state on discord.
type VoiceState struct {
	UserID                  Snowflake    `json:"user_id"`
	ChannelID               Snowflake    `json:"channel_id"`
	GuildID                 *Snowflake   `json:"guild_id,omitempty"`
	Member                  *GuildMember `json:"member,omitempty"`
	SessionID               string       `json:"session_id"`
	Deaf                    bool         `json:"deaf"`
	Mute                    bool         `json:"mute"`
	SelfDeaf                bool         `json:"self_deaf"`
	SelfMute                bool         `json:"self_mute"`
	SelfStream              bool         `json:"self_stream"`
	SelfVideo               bool         `json:"self_video"`
	Suppress                bool         `json:"suppress"`
	RequestToSpeakTimestamp time.Time    `json:"request_to_speak_timestamp"`
}

// GuildBan represents a ban entry.
type GuildBan struct {
	GuildID *Snowflake `json:"guild_id,omitempty"`
	Reason  string
	User    *User `json:"user"`
}

// GuildPruneParam represents the arguments for a guild prune.
type GuildPruneParam struct {
	Days              *int32       `json:"days,omitempty"`
	ComputePruneCount bool         `json:"compute_prune_count"`
	IncludeRoles      []*Snowflake `json:"include_roles"`
}
