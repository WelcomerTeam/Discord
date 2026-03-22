package discord

// automod.go contains all structures for Discord's Auto Moderation feature.

// AutoModerationEventType represents what triggers an AutoModeration rule.
type AutoModerationEventType int32

const (
	AutoModerationEventTypeMessageSend  AutoModerationEventType = 1
	AutoModerationEventTypeMemberUpdate AutoModerationEventType = 2
)

// AutoModerationTriggerType represents what content triggers an AutoModeration rule.
type AutoModerationTriggerType int32

const (
	AutoModerationTriggerTypeKeyword       AutoModerationTriggerType = 1
	AutoModerationTriggerTypeSpam          AutoModerationTriggerType = 2
	AutoModerationTriggerTypeKeywordPreset AutoModerationTriggerType = 3
	AutoModerationTriggerTypeMentionSpam   AutoModerationTriggerType = 4
	AutoModerationTriggerTypeMemberProfile AutoModerationTriggerType = 5
)

// AutoModerationActionType represents what action to take when a rule is triggered.
type AutoModerationActionType int32

const (
	AutoModerationActionTypeBlockMessage            AutoModerationActionType = 1
	AutoModerationActionTypeSendAlertMessage        AutoModerationActionType = 2
	AutoModerationActionTypeTimeout                 AutoModerationActionType = 3
	AutoModerationActionTypeBlockMemberInteraction  AutoModerationActionType = 4
)

// AutoModerationKeywordPresetType represents preset keyword lists.
type AutoModerationKeywordPresetType int32

const (
	AutoModerationKeywordPresetTypeProfanity     AutoModerationKeywordPresetType = 1
	AutoModerationKeywordPresetTypeSexualContent AutoModerationKeywordPresetType = 2
	AutoModerationKeywordPresetTypeSlurs         AutoModerationKeywordPresetType = 3
)

// AutoModerationTriggerMetadata represents additional metadata for a trigger.
type AutoModerationTriggerMetadata struct {
	KeywordFilter                []string                          `json:"keyword_filter,omitempty"`
	RegexPatterns                []string                          `json:"regex_patterns,omitempty"`
	Presets                      []AutoModerationKeywordPresetType `json:"presets,omitempty"`
	AllowList                    []string                          `json:"allow_list,omitempty"`
	MentionTotalLimit            *int32                            `json:"mention_total_limit,omitempty"`
	MentionRaidProtectionEnabled *bool                             `json:"mention_raid_protection_enabled,omitempty"`
}

// AutoModerationActionMetadata represents additional metadata for an action.
type AutoModerationActionMetadata struct {
	ChannelID       *Snowflake `json:"channel_id,omitempty"`
	DurationSeconds *int32     `json:"duration_seconds,omitempty"`
	CustomMessage   *string    `json:"custom_message,omitempty"`
}

// AutoModerationAction represents an action to be taken when a rule is triggered.
type AutoModerationAction struct {
	Type     AutoModerationActionType      `json:"type"`
	Metadata *AutoModerationActionMetadata `json:"metadata,omitempty"`
}

// AutoModerationRule represents an Auto Moderation rule.
type AutoModerationRule struct {
	ID              Snowflake                      `json:"id"`
	GuildID         Snowflake                      `json:"guild_id"`
	CreatorID       Snowflake                      `json:"creator_id"`
	Name            string                         `json:"name"`
	EventType       AutoModerationEventType        `json:"event_type"`
	TriggerType     AutoModerationTriggerType      `json:"trigger_type"`
	TriggerMetadata *AutoModerationTriggerMetadata `json:"trigger_metadata,omitempty"`
	Actions         []AutoModerationAction         `json:"actions"`
	Enabled         bool                           `json:"enabled"`
	ExemptRoles     []Snowflake                    `json:"exempt_roles"`
	ExemptChannels  []Snowflake                    `json:"exempt_channels"`
}
