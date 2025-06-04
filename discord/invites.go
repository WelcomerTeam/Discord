package discord

import (
	"context"
	"time"
)

// invites.go contains all structures for invites.

// InviteTargetType represents the type of an invites target.
type InviteTargetType uint16

const (
	InviteTargetTypeStream InviteTargetType = 1 + iota
	InviteTargetTypeEmbeddedApplication
)

// EventStatus represents the status of an event.
type EventStatus uint16

const (
	EventStatusScheduled EventStatus = 1 + iota
	EventStatusActive
	EventStatusCompleted
	EventStatusCanceled
)

// ScheduledEntityType represents the type of event.
type ScheduledEntityType uint16

const (
	ScheduledEntityTypeStage ScheduledEntityType = 1 + iota
	ScheduledEntityTypeVoice
	ScheduledEntityTypeExternal
)

// Invite represents the structure of Invite data.
type Invite struct {
	ExpiresAt                time.Time            `json:"expires_at,omitempty"`
	CreatedAt                time.Time            `json:"created_at"`
	ScheduledEvent           *ScheduledEvent      `json:"guild_scheduled_event,omitempty"`
	StageInstance            *InviteStageInstance `json:"stage_instance,omitempty"`
	Inviter                  *User                `json:"inviter,omitempty"`
	TargetType               *InviteTargetType    `json:"target_type,omitempty"`
	TargetUser               *User                `json:"target_user,omitempty"`
	TargetApplication        Application          `json:"target_application"`
	Guild                    *Guild               `json:"guild,omitempty"`
	Channel                  *Channel             `json:"channel,omitempty"`
	GuildID                  *Snowflake           `json:"guild_id,omitempty"`
	Code                     string               `json:"code"`
	ApproximateMemberCount   int32                `json:"approximate_member_count,omitempty"`
	Uses                     int32                `json:"uses"`
	MaxUses                  int32                `json:"max_uses"`
	MaxAge                   int32                `json:"max_age"`
	ApproximatePresenceCount int32                `json:"approximate_presence_count,omitempty"`
	Temporary                bool                 `json:"temporary"`
}

// Delete deletes an invite.
// reason: Reason for deleting an invite.
func (i *Invite) Delete(ctx context.Context, session *Session, reason *string) error {
	return DeleteInvite(ctx, session, i.Code, reason)
}

// InviteStageInstance represents the structure of an invite stage instance.
type InviteStageInstance struct {
	Topic            string          `json:"topic"`
	Members          GuildMemberList `json:"members"`
	ParticipantCount int32           `json:"participant_count"`
	SpeakerCount     int32           `json:"speaker_count"`
}

// ScheduledEvent represents an scheduled event.
type ScheduledEvent struct {
	Status             EventStatus              `json:"status"`
	EntityType         ScheduledEntityType      `json:"entity_type"`
	ChannelID          *Snowflake               `json:"channel_id,omitempty"`
	CreatorID          *Snowflake               `json:"creator_id,omitempty"`
	Creator            *User                    `json:"creator,omitempty"`
	EntityMetadata     *EventMetadata           `json:"entity_metadata,omitempty"`
	EntityID           *Snowflake               `json:"entity_id,omitempty"`
	PrivacyLevel       StageChannelPrivacyLevel `json:"privacy_level"`
	ScheduledStartTime string                   `json:"scheduled_start_time"`
	ScheduledEndTime   string                   `json:"scheduled_end_time"`
	Description        string                   `json:"description,omitempty"`
	Name               string                   `json:"name"`
	ID                 Snowflake                `json:"id"`
	GuildID            Snowflake                `json:"guild_id"`
	UserCount          int32                    `json:"user_count,omitempty"`
}

// EventMetadata contains extra information about a scheduled event.
type EventMetadata struct {
	Location string `json:"location,omitempty"`
}

// ScheduledEventUser represents a user subscribed to an event.
type ScheduledEventUser struct {
	User    User         `json:"user"`
	Member  *GuildMember `json:"member,omitempty"`
	EventID Snowflake    `json:"guild_scheduled_event_id"`
}

// InviteParams represents the params to create an invite.
type InviteParams struct {
	MaxAge    int32 `json:"max_age"`
	MaxUses   int32 `json:"max_uses"`
	Temporary bool  `json:"temporary"`
	Unique    bool  `json:"unique"`
}
