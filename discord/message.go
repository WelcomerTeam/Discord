package discord

import (
	"context"
	"encoding/json"
	"time"
)

// message.go contains the structure that represents a discord message.

// MessageType represents the type of message that has been sent.
type MessageType uint16

const (
	MessageTypeDefault MessageType = iota
	MessageTypeRecipientAdd
	MessageTypeRecipientRemove
	MessageTypeCall
	MessageTypeChannelNameChange
	MessageTypeChannelIconChange
	MessageTypeChannelPinnedMessage
	MessageTypeGuildMemberJoin
	MessageTypeUserPremiumGuildSubscription
	MessageTypeUserPremiumGuildSubscriptionTier1
	MessageTypeUserPremiumGuildSubscriptionTier2
	MessageTypeUserPremiumGuildSubscriptionTier3
	MessageTypeChannelFollowAdd
	_
	MessageTypeGuildDiscoveryDisqualified
	MessageTypeGuildDiscoveryRequalified
	MessageTypeGuildDiscoveryGracePeriodInitialWarning
	MessageTypeGuildDiscoveryGracePeriodFinalWarning
	MessageTypeThreadCreated
	MessageTypeReply
	MessageTypeApplicationCommand
	MessageTypeThreadStarterMessage
	MessageTypeGuildInviteReminder
)

// MessageFlags represents the extra information on a message.
type MessageFlags uint16

const (
	MessageFlagCrossposted MessageFlags = 1 << iota
	MessageFlagIsCrosspost
	MessageFlagSuppressEmbeds
	MessageFlagSourceMessageDeleted
	MessageFlagUrgent
	MessageFlagHasThread
	MessageFlagEphemeral
	MessageFlagLoading
	MessageFlagFailedToMentionSomeRolesInThread
	_
	_
	_
	MessageFlagSuppressNotifications
	MessageFlagIsVoiceMessage
	MessageFlagHasSnapshot
	MessageFlagIsComponentsV2
)

// MessageAllowedMentionsType represents all the allowed mention types.
type MessageAllowedMentionsType string

const (
	MessageAllowedMentionsTypeRoles    MessageAllowedMentionsType = "roles"
	MessageAllowedMentionsTypeUsers    MessageAllowedMentionsType = "users"
	MessageAllowedMentionsTypeEveryone MessageAllowedMentionsType = "everyone"
)

// MessageActivityType represents the type of message activity.
type MessageActivityType uint16

const (
	MessageActivityTypeJoin MessageActivityType = 1 + iota
	MessageActivityTypeSpectate
	MessageActivityTypeListen
	MessageActivityTypeJoinRequest
)

// Message represents a message on discord.
type Message struct {
	Timestamp         time.Time               `json:"timestamp"`
	EditedTimestamp   time.Time               `json:"edited_timestamp"`
	Author            User                    `json:"author"`
	WebhookID         *Snowflake              `json:"webhook_id,omitempty"`
	Member            *GuildMember            `json:"member,omitempty"`
	GuildID           *Snowflake              `json:"guild_id,omitempty"`
	Thread            *Channel                `json:"thread,omitempty"`
	Interaction       *MessageInteraction     `json:"interaction,omitempty"`
	ReferencedMessage *Message                `json:"referenced_message,omitempty"`
	Flags             *MessageFlags           `json:"flags,omitempty"`
	Application       *Application            `json:"application,omitempty"`
	Activity          *MessageActivity        `json:"activity,omitempty"`
	Content           string                  `json:"content"`
	Embeds            []Embed                 `json:"embeds"`
	MentionRoles      []Snowflake             `json:"mention_roles"`
	Reactions         []MessageReaction       `json:"reactions"`
	StickerItems      []MessageSticker        `json:"sticker_items,omitempty"`
	Attachments       []MessageAttachment     `json:"attachments"`
	Components        []InteractionComponent  `json:"components,omitempty"`
	MentionChannels   []MessageChannelMention `json:"mention_channels,omitempty"`
	Mentions          []User                  `json:"mentions"`
	MessageReference  []MessageReference      `json:"message_referenced,omitempty"`
	ID                Snowflake               `json:"id"`
	ChannelID         Snowflake               `json:"channel_id"`
	MentionEveryone   bool                    `json:"mention_everyone"`
	TTS               bool                    `json:"tts"`
	Type              MessageType             `json:"type"`
	Pinned            bool                    `json:"pinned"`
}

// AddReaction adds a reaction to a message
// emoji: unicode representation or emoji id.
func (m *Message) AddReaction(ctx context.Context, session *Session, emoji string) error {
	return CreateReaction(ctx, session, m.ChannelID, m.ID, emoji)
}

// ClearReaction clears a specific reaction from a message.
// emoji: unicode representation or emoji id.
func (m *Message) ClearReaction(ctx context.Context, session *Session, emoji string) error {
	return DeleteAllReactionsEmoji(ctx, session, m.ChannelID, m.ID, emoji)
}

// ClearReactions clears all reactions from a message.
func (m *Message) ClearReactions(ctx context.Context, session *Session) error {
	return DeleteAllReactions(ctx, session, m.ChannelID, m.ID)
}

// Delete deletes a message.
// reason: reason for deleting the message.
func (m *Message) Delete(ctx context.Context, session *Session, reason *string) error {
	return DeleteMessage(ctx, session, m.ChannelID, m.ID, reason)
}

// Edit edits a message.
// messageArg: arguments for editing the message.
func (m *Message) Edit(ctx context.Context, session *Session, messageParams MessageParams) (*Message, error) {
	return EditMessage(ctx, session, m.ChannelID, m.ID, messageParams)
}

// Pin pins a message in a channel.
// reason: reason for pinning a message.
func (m *Message) Pin(ctx context.Context, session *Session, reason *string) error {
	return PinMessage(ctx, session, m.ChannelID, m.ID, reason)
}

// Publish publishes a message. This must be in an announcement channel.
func (m *Message) Publish(ctx context.Context, session *Session) (*Message, error) {
	return CrosspostMessage(ctx, session, m.ChannelID, m.ID)
}

// RemoveReaction removes a specific reaction from a specific user.
// emoji: unicode representation or emoji id.
// user: The user to remove the reaction from.
func (m *Message) RemoveReaction(ctx context.Context, session *Session, emoji string, user User) error {
	return DeleteUserReaction(ctx, session, m.ChannelID, m.ID, emoji, user.ID)
}

// Reply will send a new message in the same channel as the target message and references the target.
// This is the same as using Send() and setting the message as the MessageReference.
// messageArg: arguments for sending a message.
func (m *Message) Reply(ctx context.Context, session *Session, messageParams MessageParams) (*Message, error) {
	messageParams.MessageReference = &MessageReference{
		ID:              &m.ID,
		ChannelID:       &m.ChannelID,
		GuildID:         m.GuildID,
		FailIfNotExists: true,
	}

	channel := &Channel{ID: m.ChannelID}

	return channel.Send(ctx, session, messageParams)
}

// Unpin unpins a message.
// reason: Reason for unpinning.
func (m *Message) Unpin(ctx context.Context, session *Session, reason *string) error {
	return UnpinMessage(ctx, session, m.ChannelID, m.ID, reason)
}

// MessageParams represents the structure for sending a message on discord.
type MessageParams struct {
	MessageReference *MessageReference        `json:"message_reference,omitempty"`
	PayloadJSON      *json.RawMessage         `json:"payload_json,omitempty"`
	Content          string                   `json:"content"`
	Embeds           []Embed                  `json:"embeds,omitempty"`
	AllowedMentions  []MessageAllowedMentions `json:"allowed_mentions,omitempty"`
	Components       []InteractionComponent   `json:"components,omitempty"`
	StickerIDs       []Snowflake              `json:"sticker_ids,omitempty"`
	Files            []File                   `json:"-"`
	Attachments      []MessageAttachment      `json:"attachments,omitempty"`
	Flags            MessageFlags             `json:"flags,omitempty"`
	TTS              bool                     `json:"tts"`
}

func NewMessage(content string) *MessageParams {
	return &MessageParams{
		Content: content,
	}
}

func (m *MessageParams) SetTTS(tts bool) *MessageParams {
	m.TTS = tts

	return m
}

func (m *MessageParams) AddEmbed(embed Embed) *MessageParams {
	m.Embeds = append(m.Embeds, embed)

	return m
}

func (m *MessageParams) AddAllowedMention(allowedMention MessageAllowedMentions) *MessageParams {
	m.AllowedMentions = append(m.AllowedMentions, allowedMention)

	return m
}

func (m *MessageParams) AddComponent(component InteractionComponent) *MessageParams {
	m.Components = append(m.Components, component)

	return m
}

func (m *MessageParams) AddFile(file File) *MessageParams {
	m.Files = append(m.Files, file)

	return m
}

// MessageInteraction represents an executed interaction.
type MessageInteraction struct {
	User User            `json:"user"`
	Type InteractionType `json:"type"`
	Name string          `json:"name"`
	ID   Snowflake       `json:"id"`
}

// MessageChannelMention represents a mentioned channel.
type MessageChannelMention struct {
	Name    string      `json:"name"`
	ID      Snowflake   `json:"id"`
	GuildID Snowflake   `json:"guild_id"`
	Type    ChannelType `json:"type"`
}

// MessageReference represents crossposted messages or replys.
type MessageReference struct {
	ID              *Snowflake `json:"message_id,omitempty"`
	ChannelID       *Snowflake `json:"channel_id,omitempty"`
	GuildID         *Snowflake `json:"guild_id,omitempty"`
	FailIfNotExists bool       `json:"fail_if_not_exists"`
}

// MessageReaction represents a reaction to a message on discord.
type MessageReaction struct {
	Emoji        Emoji                       `json:"emoji"`
	BurstColors  []string                    `json:"burst_colors"`
	CountDetails MessageReactionCountDetails `json:"count_details"`
	Count        int32                       `json:"count"`
	BurstCount   int32                       `json:"burst_count"`
	MeBurst      bool                        `json:"me_burst"`
	BurstMe      bool                        `json:"burst_me"`
	Me           bool                        `json:"me"`
}

// MessageReactionCountDetails represents the count details of a message reaction.
type MessageReactionCountDetails struct {
	Burst  int32 `json:"burst"`
	Normal int32 `json:"normal"`
}

// MessageAllowedMentions is the structure of the allowed mentions entry.
type MessageAllowedMentions struct {
	Parse       []MessageAllowedMentionsType `json:"parse"`
	Roles       []Snowflake                  `json:"roles"`
	Users       []Snowflake                  `json:"users"`
	RepliedUser bool                         `json:"replied_user"`
}

// MediaItem represents an image in an embed or any media type.
type MediaItem struct {
	URL         string `json:"url"`
	ProxyURL    string `json:"proxy_url,omitempty"`
	Height      int32  `json:"height,omitempty"`
	Width       int32  `json:"width,omitempty"`
	ContentType string `json:"content_type,omitempty"`
}

// MessageAttachment represents a message attachment on discord.
type MessageAttachment struct {
	MediaItem
	ID        Snowflake `json:"id"`
	Size      int32     `json:"size"`
	Ephemeral bool      `json:"ephemeral"`
	Filename  string    `json:"filename"`
}

// MessageActivity represents a message activity on discord.
type MessageActivity struct {
	PartyID string              `json:"party_id,omitempty"`
	Type    MessageActivityType `json:"type"`
}
