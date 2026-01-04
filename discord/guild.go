package discord

import (
	"context"
	"fmt"
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
type MFALevel uint16

// MFA levels.
const (
	MFALevelNone MFALevel = iota
	MFALevelElevated
)

// VerificationLevel represents a guild's verification level.
type VerificationLevel uint16

const (
	VerificationLevelNone VerificationLevel = iota
	VerificationLevelLow
	VerificationLevelMedium
	VerificationLevelHigh
	VerificationLevelVeryHigh
)

// SystemChannelFlags represents the flags of a system channel.
type SystemChannelFlags uint16

const (
	SystemChannelFlagsSuppressJoin SystemChannelFlags = 1 << iota
	SystemChannelFlagsPremiumSubscriptions
	SystemChannelFlagsSuppressSetupTips
	SystemChannelFlagsHideMemberJoinStickerReplyButtons
	SystemChannelFlagsSuppressSubscriptionNotifications
	SystemChannelFlagsHideRoleSubscriptionReplyButtons
	_
	_
)

// PremiumTier represents the current boosting tier of a guild.
type PremiumTier uint16

const (
	PremiumTierNone PremiumTier = iota
	PremiumTier1
	PremiumTier2
	PremiumTier3
)

// GuildNSFWLevelType represents the level of the guild.
type GuildNSFWLevelType uint16

const (
	GuildNSFWLevelTypDefault GuildNSFWLevelType = iota
	GuildNSFWLevelTypeExplicit
	GuildNSFWLevelTypeSafe
	GuildNSFWLevelTypeAgeRestricted
)

// Guild represents a guild on discord.
type Guild struct {
	JoinedAt                    time.Time                  `json:"joined_at"`
	WidgetChannelID             *Snowflake                 `json:"widget_channel_id,omitempty"`
	NSFWLevel                   GuildNSFWLevelType         `json:"nsfw_level"`
	PublicUpdatesChannelID      *Snowflake                 `json:"public_updates_channel_id,omitempty"`
	PremiumTier                 *PremiumTier               `json:"premium_tier,omitempty"`
	RulesChannelID              *Snowflake                 `json:"rules_channel_id,omitempty"`
	SystemChannelFlags          *SystemChannelFlags        `json:"system_channel_flags,omitempty"`
	OwnerID                     *Snowflake                 `json:"owner_id,omitempty"`
	Permissions                 *Int64                     `json:"permissions,omitempty"`
	SystemChannelID             *Snowflake                 `json:"system_channel_id,omitempty"`
	AFKChannelID                *Snowflake                 `json:"afk_channel_id,omitempty"`
	ApplicationID               *Snowflake                 `json:"application_id,omitempty"`
	Description                 string                     `json:"description"`
	PreferredLocale             string                     `json:"preferred_locale"`
	Name                        string                     `json:"name"`
	Icon                        string                     `json:"icon"`
	IconHash                    string                     `json:"icon_hash"`
	Banner                      string                     `json:"banner"`
	VanityURLCode               string                     `json:"vanity_url_code"`
	Splash                      string                     `json:"splash"`
	DiscoverySplash             string                     `json:"discovery_splash"`
	Region                      string                     `json:"region"`
	Presences                   ActivityList               `json:"presences,omitempty"`
	GuildScheduledEvents        ScheduledEventList         `json:"guild_scheduled_events"`
	Stickers                    StickerList                `json:"stickers"`
	Features                    StringList                 `json:"features"`
	StageInstances              StageInstanceList          `json:"stage_instances,omitempty"`
	Roles                       RoleList                   `json:"roles"`
	Emojis                      EmojiList                  `json:"emojis"`
	VoiceStates                 VoiceStateList             `json:"voice_states,omitempty"`
	Members                     GuildMemberList            `json:"members,omitempty"`
	Channels                    ChannelList                `json:"channels,omitempty"`
	ID                          Snowflake                  `json:"id"`
	ExplicitContentFilter       ExplicitContentFilterLevel `json:"explicit_content_filter"`
	DefaultMessageNotifications MessageNotificationLevel   `json:"default_message_notifications"`
	ApproximateMemberCount      int32                      `json:"approximate_member_count"`
	MaxMembers                  int32                      `json:"max_members"`
	MemberCount                 int32                      `json:"member_count"`
	AFKTimeout                  int32                      `json:"afk_timeout"`
	MaxPresences                int32                      `json:"max_presences"`
	PremiumSubscriptionCount    int32                      `json:"premium_subscription_count"`
	ApproximatePresenceCount    int32                      `json:"approximate_presence_count"`
	MaxVideoChannelUsers        int32                      `json:"max_video_channel_users"`
	Unavailable                 bool                       `json:"unavailable"`
	WidgetEnabled               bool                       `json:"widget_enabled"`
	VerificationLevel           VerificationLevel          `json:"verification_level"`
	Large                       bool                       `json:"large"`
	MFALevel                    MFALevel                   `json:"mfa_level"`
	Owner                       bool                       `json:"owner"`
	PremiumProgressBarEnabled   bool                       `json:"premium_progress_bar_enabled"`
}

// GuildParams represents the parameters sent when modifying a guild.
type GuildParam struct {
	AFKChannelID                *Snowflake                  `json:"afk_channel_id,omitempty"`
	SystemChannelFlags          *SystemChannelFlags         `json:"system_channel_flags,omitempty"`
	VerificationLevel           *VerificationLevel          `json:"verification_level,omitempty"`
	DefaultMessageNotifications *MessageNotificationLevel   `json:"default_message_notifications,omitempty"`
	ExplicitContentFilter       *ExplicitContentFilterLevel `json:"explicit_content_filter,omitempty"`
	Icon                        *string                     `json:"icon,omitempty"`
	OwnerID                     *Snowflake                  `json:"owner_id,omitempty"`
	Name                        *string                     `json:"name,omitempty"`
	PremiumProgressBarEnabled   *bool                       `json:"premium_progress_bar_enabled,omitempty"`
	Banner                      *string                     `json:"banner,omitempty"`
	Splash                      *string                     `json:"splash,omitempty"`
	AFKTimeout                  *int32                      `json:"afk_timeout,omitempty"`
	SystemChannelID             *Snowflake                  `json:"system_channel_id,omitempty"`
	Region                      *string                     `json:"region,omitempty"`
	RulesChannelID              *Snowflake                  `json:"rules_channel_id,omitempty"`
	PublicUpdatesChannelID      *Snowflake                  `json:"public_updates_channel_id,omitempty"`
	PreferredLocale             *string                     `json:"preferred_locale,omitempty"`
	DiscoverySplash             *string                     `json:"discovery_splash,omitempty"`
	Description                 *string                     `json:"description,omitempty"`
	Features                    StringList                  `json:"features,omitempty"`
}

// AuditLogs returns all audit logs matching query.
// userID: Filters audit logs by the userID provided.
// actionType: The action type to filter audit logs by.
// before: Only show audit logs before a certain snowflake.
// limit: Maximum number of audit log entries to return.
func (g *Guild) AuditLogs(ctx context.Context, session *Session, userID *Snowflake, actionType *AuditLogActionType, before *Snowflake, limit *int32) ([]AuditLogEntry, error) {
	return GetGuildAuditLog(ctx, session, g.ID, userID, actionType, before, limit)
}

// Ban bans a user.
// userID: ID of user that is getting banned.
// reason: Reason for ban.
func (g *Guild) Ban(ctx context.Context, session *Session, userID Snowflake, reason *string) error {
	return CreateGuildBan(ctx, session, g.ID, userID, reason)
}

// Bans returns a list of guild bans.
func (g *Guild) Bans(ctx context.Context, session *Session) ([]GuildBan, error) {
	return GetGuildBans(ctx, session, g.ID)
}

// CloneChannel creates a copy of the target channel.
// reason: Reason for creating the channel.
func (g *Guild) CloneChannel(ctx context.Context, session *Session, channel *Channel, reason *string) (*Channel, error) {
	return g.CreateChannel(ctx, session, ChannelParams{
		Name:                 channel.Name,
		Type:                 channel.Type,
		Topic:                channel.Topic,
		Bitrate:              channel.Bitrate,
		UserLimit:            channel.UserLimit,
		RateLimitPerUser:     channel.RateLimitPerUser,
		Position:             channel.Position,
		PermissionOverwrites: channel.PermissionOverwrites,
		ParentID:             channel.ParentID,
		NSFW:                 channel.NSFW,
	}, reason)
}

// CreateChannel creates a channel.
// channelArg: Parameters passed for creating a channel.
// reason: Reason for creating the channel.
func (g *Guild) CreateChannel(ctx context.Context, session *Session, channelParams ChannelParams, reason *string) (*Channel, error) {
	return CreateGuildChannel(ctx, session, g.ID, channelParams, reason)
}

// CreateCustomEmojis creates an emoji for a guild.
// name: Name of the custom emoji.
// image: Bytes representing the image file to upload.
// roles: Roles that this emoji is limited to.
// reason: Reason for creating the emoji.
func (g *Guild) CreateCustomEmoji(ctx context.Context, session *Session, name string, image []byte, roles []Snowflake, reason *string) (*Emoji, error) {
	params := EmojiParams{
		Name:  name,
		Roles: roles,
	}

	imageData, err := bytesToBase64Data(image)
	if err != nil {
		return nil, fmt.Errorf("failed to encode emoji: %w", err)
	}

	params.Image = imageData

	return CreateGuildEmoji(ctx, session, g.ID, params, reason)
}

// CreateRole creates a role.
// roleArg: Parameters passed for creating a role.
// reason: Reason for creating the role.
func (g *Guild) CreateRole(ctx context.Context, session *Session, roleParams RoleParams, reason *string) (*Role, error) {
	return CreateGuildRole(ctx, session, g.ID, roleParams, reason)
}

// Delete deletes a guild.
func (g *Guild) Delete(ctx context.Context, session *Session) error {
	return DeleteGuild(ctx, session, g.ID)
}

// Edit edits a guild.
// guildArg: Parameters passed for editing a guild.
// reason: Reason for editing the guild.
func (g *Guild) Edit(ctx context.Context, session *Session, guildArg GuildParam, reason *string) error {
	newGuild, err := ModifyGuild(ctx, session, g.ID, guildArg, reason)
	if err != nil {
		return err
	}

	*g = *newGuild

	return nil
}

// EditRolePositions edits role positions in a guild.
// guildRolePositionArgs: List of roles and their new role position.
func (g *Guild) EditRolePositions(ctx context.Context, session *Session, guildRolePositionArgs []ModifyGuildRolePosition, reason *string) ([]Role, error) {
	return ModifyGuildRolePositions(ctx, session, g.ID, guildRolePositionArgs, reason)
}

// EstimatePrunedMembers returns an estimate of how many people will be pruned from a guild based on arguments.
// days: The number of days since speaking.
// includedRoles: By default pruning only removes users with no roles, any role in this list will be included.
func (g *Guild) EstimatePrunedMembers(ctx context.Context, session *Session, days *int32, includedRoles []Snowflake) (*int32, error) {
	return GetGuildPruneCount(ctx, session, g.ID, days, includedRoles)
}

// Integrations returns all guild integrations.
func (g *Guild) Integrations(ctx context.Context, session *Session) ([]Integration, error) {
	return GetGuildIntegrations(ctx, session, g.ID)
}

// Invites returns all guild invites.
func (g *Guild) Invites(ctx context.Context, session *Session) ([]Invite, error) {
	return GetGuildInvites(ctx, session, g.ID)
}

// Kick kicks a user from the guild.
// userID: ID of user to kick.
// reason: Reason for kicking the user.
func (g *Guild) Kick(ctx context.Context, session *Session, userID Snowflake, reason *string) error {
	return RemoveGuildMember(ctx, session, g.ID, userID, reason)
}

// Leave leaves a guild.
func (g *Guild) Leave(ctx context.Context, session *Session) error {
	return LeaveGuild(ctx, session, g.ID)
}

// PruneMembers prunes users from a guild based on arguments.
// days: The number of days since speaking.
// includedRoles: By default pruning only removes users with no roles, any role in this list will be included.
// computePruneCount: Returns how many users were pruned, usage on larger guilds is discouraged.
// reason: Reason for pruning members.
func (g *Guild) PruneMembers(ctx context.Context, session *Session, days *int32, includedRoles []Snowflake, computePruneCount bool, reason *string) (*int32, error) {
	return BeginGuildPrune(ctx, session, g.ID, days, includedRoles, computePruneCount, reason)
}

// QueryMembers returns guild members whose username or nickname matches query.
// query: Query string to match usernames and nicknames against.
// limit: Maximum number of members to return.
func (g *Guild) QueryMembers(ctx context.Context, session *Session, query string, limit *int32) ([]GuildMember, error) {
	return SearchGuildMembers(ctx, session, g.ID, query, limit)
}

// Unban unbans a user from a guild.
// userID: ID of user to unban.
// reason: Reason for unbanning.
func (g *Guild) Unban(ctx context.Context, session *Session, userID Snowflake, reason *string) error {
	return RemoveGuildBan(ctx, session, g.ID, userID, reason)
}

// VanityInvite returns the vanity invite for a guild.
func (g *Guild) VanityInvite(ctx context.Context, session *Session) (*Invite, error) {
	invite, err := GetGuildVanityURL(ctx, session, g.ID)
	if err != nil {
		return invite, err
	}

	g.VanityURLCode = invite.Code

	return invite, nil
}

// Webhooks returns all webhooks for a guild.
func (g *Guild) Webhooks(ctx context.Context, session *Session) ([]Webhook, error) {
	return GetGuildWebhooks(ctx, session, g.ID)
}

// UnavailableGuild represents an unavailable guild.
type UnavailableGuild struct {
	ID          Snowflake `json:"id"`
	Unavailable bool      `json:"unavailable"`
}

// GuildMember represents a guild member on discord.
type GuildMember struct {
	JoinedAt                   time.Time     `json:"joined_at"`
	CommunicationDisabledUntil *time.Time    `json:"communication_disabled_until,omitempty"`
	PremiumSince               *time.Time    `json:"premium_since,omitempty"`
	User                       *User         `json:"user,omitempty"`
	GuildID                    *Snowflake    `json:"guild_id,omitempty"`
	Permissions                *Int64        `json:"permissions"`
	Nick                       string        `json:"nick,omitempty"`
	Avatar                     string        `json:"avatar,omitempty"`
	Roles                      SnowflakeList `json:"roles"`
	Deaf                       bool          `json:"deaf"`
	Mute                       bool          `json:"mute"`
	Pending                    bool          `json:"pending"`
}

// GuildMemberParams represents the arguments used to modify a guild member.
type GuildMemberParams struct {
	Nick                       *string     `json:"nick,omitempty"`
	Deaf                       *bool       `json:"deaf,omitempty"`
	Mute                       *bool       `json:"mute,omitempty"`
	ChannelID                  *Snowflake  `json:"channel_id,omitempty"`
	CommunicationDisabledUntil *time.Time  `json:"communication_disabled_until,omitempty"`
	Roles                      []Snowflake `json:"roles,omitempty"`
}

// AddRoles adds roles to a guild member.
// roles: List of roles to add to the guild member.
// reason: Reason for adding the roles to the guild member.
// atomic: When true, will send multiple AddGuildMemberRole requests instead of at once.
func (gm *GuildMember) AddRoles(ctx context.Context, session *Session, roles []Snowflake, reason *string, atomic bool) error {
	guildMemberRoles := make(map[Snowflake]bool)

	for _, guildMemberRole := range gm.Roles {
		guildMemberRoles[guildMemberRole] = true
	}

	if atomic {
		for _, roleID := range roles {
			if _, ok := guildMemberRoles[roleID]; !ok {
				err := AddGuildMemberRole(ctx, session, *gm.GuildID, gm.User.ID, roleID, reason)
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

	return gm.Edit(ctx, session, GuildMemberParams{Roles: newRoles}, reason)
}

// Ban bans the guild member from the guild.
// reason: Reason for banning the guild member.
func (gm *GuildMember) Ban(ctx context.Context, session *Session, reason *string) error {
	return CreateGuildBan(ctx, session, *gm.GuildID, gm.User.ID, reason)
}

// CreateDM creates a DMChannel with a user. This should not need to be called as Send() transparently does this.
// If the user already has a DMChannel created, this will return a partial channel with just an ID set.
func (gm *GuildMember) CreateDM(ctx context.Context, session *Session) (*Channel, error) {
	return gm.User.CreateDM(ctx, session)
}

// Edit edits a guild member.
// guildMemberArg: Parameters used to update a guild member.
// reason: Reason for editing the guild member.
func (gm *GuildMember) Edit(ctx context.Context, session *Session, guildMemberParams GuildMemberParams, reason *string) error {
	newMember, err := ModifyGuildMember(ctx, session, *gm.GuildID, gm.User.ID, guildMemberParams, reason)
	if err != nil {
		return err
	}

	*gm = *newMember

	return nil
}

// Kick kicks the guild member.
// reason: Reason for kicking the guild member.
func (gm *GuildMember) Kick(ctx context.Context, session *Session, reason *string) error {
	return RemoveGuildMember(ctx, session, *gm.GuildID, gm.User.ID, reason)
}

// MoveTo moves the guild member to a different voice channel.
// channelID: Channel to move the user to, if nil they are removed from voice.
// reason: Reason for moving the guild member
func (gm *GuildMember) MoveTo(ctx context.Context, session *Session, channelID *Snowflake, reason *string) error {
	return gm.Edit(ctx, session, GuildMemberParams{ChannelID: channelID}, reason)
}

// RemoveRoles removes roles from a guild member.
func (gm *GuildMember) RemoveRoles(ctx context.Context, session *Session, roles []Snowflake, reason *string, atomic bool) error {
	guildMemberRoles := make(map[Snowflake]bool)

	for _, guildMemberRole := range gm.Roles {
		guildMemberRoles[guildMemberRole] = true
	}

	if atomic {
		for _, roleID := range roles {
			if _, ok := guildMemberRoles[roleID]; ok {
				err := RemoveGuildMemberRole(ctx, session, *gm.GuildID, gm.User.ID, roleID, reason)
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

	return gm.Edit(ctx, session, GuildMemberParams{Roles: newRoles}, reason)
}

// Send sends a DM message to a user. This will create a DMChannel if one is not present.
// params: The message parameters used to send the message.
func (gm *GuildMember) Send(ctx context.Context, session *Session, params MessageParams) (*Message, error) {
	return gm.User.Send(ctx, session, params)
}

// VoiceState represents the voice state on discord.
type VoiceState struct {
	RequestToSpeakTimestamp time.Time    `json:"request_to_speak_timestamp"`
	GuildID                 *Snowflake   `json:"guild_id,omitempty"`
	Member                  *GuildMember `json:"member,omitempty"`
	SessionID               string       `json:"session_id"`
	UserID                  Snowflake    `json:"user_id"`
	ChannelID               Snowflake    `json:"channel_id"`
	Mute                    bool         `json:"mute"`
	SelfDeaf                bool         `json:"self_deaf"`
	SelfMute                bool         `json:"self_mute"`
	SelfStream              bool         `json:"self_stream"`
	SelfVideo               bool         `json:"self_video"`
	Suppress                bool         `json:"suppress"`
	Deaf                    bool         `json:"deaf"`
}

// GuildBan represents a ban entry.
type GuildBan struct {
	GuildID *Snowflake `json:"guild_id,omitempty"`
	User    User       `json:"user"`
	Reason  string
}

// GuildPruneParam represents the arguments for a guild prune.
type GuildPruneParam struct {
	Days              *int32      `json:"days,omitempty"`
	IncludeRoles      []Snowflake `json:"include_roles"`
	ComputePruneCount bool        `json:"compute_prune_count"`
}
