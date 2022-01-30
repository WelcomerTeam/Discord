package structs

import (
	"time"

	"github.com/WelcomerTeam/Discord/discord"
)

// channel.go contains the information relating to channels

// ChannelType represents a channel's type.
type ChannelType uint8

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
type VideoQualityMode uint8

const (
	VideoQualityModeAuto VideoQualityMode = 1 + iota
	VideoqualityModeFull
)

// StageChannelPrivacyLevel represents the privacy level of a stage channel.
type StageChannelPrivacyLevel uint8

const (
	StageChannelPrivacyLevelPublic StageChannelPrivacyLevel = 1 + iota
	StageChannelPrivacyLevelGuildOnly
)

// Channel represents a Discord channel.
type Channel struct {
	ID                         discord.Snowflake   `json:"id"`
	GuildID                    *discord.Snowflake  `json:"guild_id,omitempty"`
	Type                       ChannelType         `json:"type"`
	Position                   int32               `json:"position,omitempty"`
	PermissionOverwrites       []*ChannelOverwrite `json:"permission_overwrites,omitempty"`
	Name                       string              `json:"name,omitempty"`
	Topic                      string              `json:"topic,omitempty"`
	NSFW                       bool                `json:"nsfw"`
	LastMessageID              string              `json:"last_message_id,omitempty"`
	Bitrate                    int32               `json:"bitrate,omitempty"`
	UserLimit                  int32               `json:"user_limit,omitempty"`
	RateLimitPerUser           int32               `json:"rate_limit_per_user,omitempty"`
	Recipients                 []*User             `json:"recipients,omitempty"`
	Icon                       string              `json:"icon,omitempty"`
	OwnerID                    *discord.Snowflake  `json:"owner_id,omitempty"`
	ApplicationID              *discord.Snowflake  `json:"application_id,omitempty"`
	ParentID                   *discord.Snowflake  `json:"parent_id,omitempty"`
	LastPinTimestamp           *time.Time          `json:"last_pin_timestamp,omitempty"`
	RTCRegion                  string              `json:"rtc_region,omitempty"`
	VideoQualityMode           *VideoQualityMode   `json:"video_quality_mode,omitempty"`
	MessageCount               int32               `json:"message_count,omitempty"`
	MemberCount                int32               `json:"member_count,omitempty"`
	ThreadMetadata             *ThreadMetadata     `json:"thread_metadata,omitempty"`
	ThreadMember               *ThreadMember       `json:"member,omitempty"`
	DefaultAutoArchiveDuration int32               `json:"default_auto_archive_duration,omitempty"`
	Permissions                *discord.Int64      `json:"permissions,omitempty"`
}

// ChannelOverwrite represents a permission overwrite for a channel.
type ChannelOverwrite struct {
	ID    discord.Snowflake    `json:"id"`
	Type  *ChannelOverrideType `json:"type"`
	Allow discord.Int64        `json:"allow"`
	Deny  discord.Int64        `json:"deny"`
}

// ChannelOverrideType represents the target of a channel override.
type ChannelOverrideType uint8

const (
	ChannelOverrideTypeRole ChannelOverrideType = iota
	ChannelOverrideTypeMember
)

// ThreadMetadata contains thread-specific channel fields.
type ThreadMetadata struct {
	Archived            bool      `json:"archived"`
	AutoArchiveDuration int32     `json:"auto_archive_duration"`
	ArchiveTimestamp    time.Time `json:"archive_timestamp"`
	Locked              bool      `json:"locked"`
}

// ThreadMember is used to indicate whether a user has joined a thread or not.
type ThreadMember struct {
	ID            *discord.Snowflake `json:"id,omitempty"`
	UserID        *discord.Snowflake `json:"user_id,omitempty"`
	GuildID       *discord.Snowflake `json:"guild_id,omitempty"`
	JoinTimestamp time.Time          `json:"join_timestamp"`
	Flags         int32              `json:"flags"`
}

// StageInstance represents a stage channel instance.
type StageInstance struct {
	ID                   discord.Snowflake         `json:"id"`
	GuildID              discord.Snowflake         `json:"guild_id"`
	ChannelID            discord.Snowflake         `json:"channel_id"`
	Topic                string                    `json:"topic"`
	PrivacyLabel         *StageChannelPrivacyLevel `json:"privacy_level"`
	DiscoverableDisabled bool                      `json:"discoverable_disabled"`
}

// FollowedChannel represents a followed channel.
type FollowedChannel struct {
	ChannelID discord.Snowflake `json:"channel_id"`
	WebhookID discord.Snowflake `json:"webhook_id"`
}
