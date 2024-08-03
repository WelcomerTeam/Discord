package discord

import (
	"bytes"
	"fmt"
	"strconv"
	"time"
)

// channel.go contains the information relating to channels

// ChannelType represents a channel's type.
type ChannelType uint16

const (
	ChannelTypeGuildText ChannelType = iota
	ChannelTypeDM
	ChannelTypeGuildVoice
	ChannelTypeGroupDM
	ChannelTypeGuildCategory
	ChannelTypeGuildNews
	ChannelTypeGuildStore
	_
	_
	_
	ChannelTypeGuildNewsThread
	ChannelTypeGuildPublicThread
	ChannelTypeGuildPrivateThread
	ChannelTypeGuildStageVoice
)

// VideoQualityMode represents the quality of the video.
type VideoQualityMode uint16

const (
	VideoQualityModeAuto VideoQualityMode = 1 + iota
	VideoqualityModeFull
)

// StageChannelPrivacyLevel represents the privacy level of a stage channel.
type StageChannelPrivacyLevel uint16

const (
	StageChannelPrivacyLevelPublic StageChannelPrivacyLevel = 1 + iota
	StageChannelPrivacyLevelGuildOnly
)

// Channel represents a Discord channel.
type Channel struct {
	OwnerID                    *Snowflake         `json:"owner_id,omitempty"`
	GuildID                    *Snowflake         `json:"guild_id,omitempty"`
	Permissions                *Int64             `json:"permissions,omitempty"`
	ThreadMember               *ThreadMember      `json:"member,omitempty"`
	ThreadMetadata             *ThreadMetadata    `json:"thread_metadata,omitempty"`
	VideoQualityMode           *VideoQualityMode  `json:"video_quality_mode,omitempty"`
	LastPinTimestamp           *time.Time         `json:"last_pin_timestamp,omitempty"`
	ParentID                   *Snowflake         `json:"parent_id,omitempty"`
	ApplicationID              *Snowflake         `json:"application_id,omitempty"`
	RTCRegion                  string             `json:"rtc_region,omitempty"`
	Topic                      string             `json:"topic,omitempty"`
	Icon                       string             `json:"icon,omitempty"`
	Name                       string             `json:"name,omitempty"`
	LastMessageID              string             `json:"last_message_id,omitempty"`
	PermissionOverwrites       []ChannelOverwrite `json:"permission_overwrites,omitempty"`
	Recipients                 []User             `json:"recipients,omitempty"`
	ID                         Snowflake          `json:"id"`
	UserLimit                  int32              `json:"user_limit,omitempty"`
	Bitrate                    int32              `json:"bitrate,omitempty"`
	MessageCount               int32              `json:"message_count,omitempty"`
	MemberCount                int32              `json:"member_count,omitempty"`
	RateLimitPerUser           int32              `json:"rate_limit_per_user,omitempty"`
	Position                   int32              `json:"position,omitempty"`
	DefaultAutoArchiveDuration int32              `json:"default_auto_archive_duration,omitempty"`
	NSFW                       bool               `json:"nsfw"`
	Type                       ChannelType        `json:"type"`
}

// ChannelParams the data that is provided when creating a channel.
type ChannelParams struct {
	ParentID             *Snowflake         `json:"parent_id,omitempty"`
	Name                 string             `json:"name"`
	Topic                string             `json:"topic,omitempty"`
	PermissionOverwrites []ChannelOverwrite `json:"permission_overwrites,omitempty"`
	Bitrate              int32              `json:"bitrate,omitempty"`
	UserLimit            int32              `json:"user_limit,omitempty"`
	RateLimitPerUser     int32              `json:"rate_limit_per_user,omitempty"`
	Position             int32              `json:"position,omitempty"`
	Type                 ChannelType        `json:"type"`
	NSFW                 bool               `json:"nsfw"`
}

// CreateInvite creates an invite to a channel.
// inviteArg: Parameters passed for creating an invite.
// reason: Reason for creating the invite.
func (c *Channel) CreateInvite(s *Session, inviteParams InviteParams, reason *string) (*Invite, error) {
	return CreateChannelInvite(s, c.ID, inviteParams, reason)
}

// CreateWebhook creates a webhook for a channel.
// webhookArg: Parameters passed for creating a webhook.
// reason: Reason for creating the webhook
func (c *Channel) CreateWebhook(s *Session, webhookArg WebhookParam, reason *string) (*Webhook, error) {
	return CreateWebhook(s, c.ID, webhookArg, reason)
}

// Delete deletes a webhook channel.
// reason: Reason for deleting the channel.
func (c *Channel) Delete(s *Session, reason *string) error {
	return DeleteChannel(s, c.ID, reason)
}

// DeleteMessages bulk deletes messages in a channel.
// messageIDs: List of message IDs to remove.
// reason: Reason for bulk delete.
func (c *Channel) DeleteMessages(s *Session, messageIDs []Snowflake, reason *string) error {
	return BulkDeleteMessages(s, c.ID, messageIDs, reason)
}

// Edit edits a channel.
// channelArg: Parameters passed for editing a channel.
// reason: Reason for editing the channel.
func (c *Channel) Edit(s *Session, channelParams ChannelParams, reason *string) error {
	newChannel, err := ModifyChannel(s, c.ID, channelParams, reason)
	if err != nil {
		return err
	}

	*c = *newChannel

	return nil
}

// History returns channel messages.
// around: Get messages around this message ID.
// before: Get messages before this message ID.
// after: Get messages after this message ID.
// limit: Maximum number of messages to return.
func (c *Channel) History(s *Session, around *Snowflake, before *Snowflake, after *Snowflake, limit *int32) ([]Message, error) {
	return GetChannelMessages(s, c.ID, around, before, after, limit)
}

// Invites returns all invites for this channel.
func (c *Channel) Invites(s *Session) ([]Invite, error) {
	return GetChannelInvites(s, c.ID)
}

// Pins returns all pinned messages in this channel.
func (c *Channel) Pins(s *Session) ([]Message, error) {
	return GetPinnedMessages(s, c.ID)
}

// Purge acts similar to History() however the resulting messages are deleted.
// around: Get messages around this message ID.
// before: Get messages before this message ID.
// after: Get messages after this message ID.
// limit: Maximum number of messages to return.
func (c *Channel) Purge(s *Session, around *Snowflake, before *Snowflake, after *Snowflake, limit *int32, reason *string) ([]Message, error) {
	messages, err := c.History(s, around, before, after, limit)
	if err != nil {
		return messages, err
	}

	messageIDs := make([]Snowflake, 0, len(messages))
	for _, message := range messages {
		messageIDs = append(messageIDs, message.ID)
	}

	err = BulkDeleteMessages(s, c.ID, messageIDs, reason)

	return messages, err
}

// Sends a message in a channel.
// messageArg: Parameters used to send a message.
func (c *Channel) Send(s *Session, messageParams MessageParams) (*Message, error) {
	return CreateMessage(s, c.ID, messageParams)
}

// SetPermissions sets permission overwrites.
// overwriteID: The role or user ID to overwrite permissions for.
// overwriteArg: Parameters used to to overwrite permissions.
// reason: Reason for setting permission overwrite.
func (c *Channel) SetPermissions(s *Session, overwriteID Snowflake, overwriteArg ChannelOverwrite, reason *string) error {
	return EditChannelPermissions(s, c.ID, overwriteID, overwriteArg, reason)
}

// TriggerTyping will show a typing indicator in the channel.
func (c *Channel) TriggerTyping(s *Session) error {
	return TriggerTypingIndicator(s, c.ID)
}

// Webhooks returns all webhooks for a channel.
func (c *Channel) Webhooks(s *Session) ([]Webhook, error) {
	return GetChannelWebhooks(s, c.ID)
}

// ChannelOverwrite represents a permission overwrite for a channel.
type ChannelOverwrite struct {
	Type  ChannelOverrideType `json:"type"`
	ID    Snowflake           `json:"id"`
	Allow Int64               `json:"allow"`
	Deny  Int64               `json:"deny"`
}

// ChannelOverrideType represents the target of a channel override.
type ChannelOverrideType Int64

func (in *ChannelOverrideType) UnmarshalJSON(b []byte) error {
	if !bytes.Equal(b, null) {
		// Discord will pass ChannelOverrideType as a string if it is in an audit log.
		if b[0] == '"' {
			i, err := strconv.ParseInt(string(b[1:len(b)-1]), 10, 64)
			if err != nil {
				return fmt.Errorf("failed to unmarshal json: %v", err)
			}

			*in = ChannelOverrideType(i)
		} else {
			i, err := strconv.ParseInt(string(b), 10, 64)
			if err != nil {
				return fmt.Errorf("failed to unmarshal json: %v", err)
			}

			*in = ChannelOverrideType(i)
		}
	}

	return nil
}

func (in ChannelOverrideType) MarshalJSON() ([]byte, error) {
	return int64ToStringBytes(int64(in)), nil
}

func (in ChannelOverrideType) String() string {
	return strconv.FormatInt(int64(in), 10)
}

const (
	ChannelOverrideTypeRole ChannelOverrideType = iota
	ChannelOverrideTypeMember
)

// ThreadMetadata contains thread-specific channel fields.
type ThreadMetadata struct {
	ArchiveTimestamp    time.Time `json:"archive_timestamp"`
	AutoArchiveDuration int32     `json:"auto_archive_duration"`
	Archived            bool      `json:"archived"`
	Locked              bool      `json:"locked"`
}

// ThreadMember is used to indicate whether a user has joined a thread or not.
type ThreadMember struct {
	ID            *Snowflake `json:"id,omitempty"`
	UserID        *Snowflake `json:"user_id,omitempty"`
	GuildID       *Snowflake `json:"guild_id,omitempty"`
	JoinTimestamp time.Time  `json:"join_timestamp"`
	Flags         int32      `json:"flags"`
}

// StageInstance represents a stage channel instance.
type StageInstance struct {
	PrivacyLabel         StageChannelPrivacyLevel `json:"privacy_level"`
	Topic                string                   `json:"topic"`
	ID                   Snowflake                `json:"id"`
	GuildID              Snowflake                `json:"guild_id"`
	ChannelID            Snowflake                `json:"channel_id"`
	DiscoverableDisabled bool                     `json:"discoverable_disabled"`
}

// FollowedChannel represents a followed channel.
type FollowedChannel struct {
	ChannelID Snowflake `json:"channel_id"`
	WebhookID Snowflake `json:"webhook_id"`
}

// ChannelPermissionsParams represents the arguments to modify guild channel permissions.
type ChannelPermissionsParams struct {
	ID              Snowflake `json:"id"`
	Position        int32     `json:"position"`
	LockPermissions bool      `json:"lock_permissions"`
	ParentID        Snowflake `json:"parent_id"`
}
