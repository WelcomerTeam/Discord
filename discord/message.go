package discord

import (
	"time"

	jsoniter "github.com/json-iterator/go"
)

// message.go contains the structure that represents a discord message.

// MessageType represents the type of message that has been sent.
type MessageType uint8

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
type MessageFlags uint8

const (
	MessageFlagCrossposted MessageFlags = 1 << iota
	MessageFlagIsCrosspost
	MessageFlagSuppressEmbeds
	MessageFlagSourceMessageDeleted
	MessageFlagUrgent
	MessageFlagHasThread
	MessageFlaEphemeral
	MessageFlagLoading
)

// MessageAllowedMentionsType represents all the allowed mention types.
type MessageAllowedMentionsType string

const (
	MessageAllowedMentionsTypeRoles    MessageAllowedMentionsType = "roles"
	MessageAllowedMentionsTypeUsers    MessageAllowedMentionsType = "users"
	MessageAllowedMentionsTypeEveryone MessageAllowedMentionsType = "everyone"
)

// MessageActivityType represents the type of message activity.
type MessageActivityType uint8

const (
	MessageActivityTypeJoin MessageActivityType = 1 + iota
	MessageActivityTypeSpectate
	MessageActivityTypeListen
	MessageActivityTypeJoinRequest
)

// Message represents a message on discord.
type Message struct {
	ID        Snowflake    `json:"id"`
	ChannelID Snowflake    `json:"channel_id"`
	GuildID   *Snowflake   `json:"guild_id,omitempty"`
	Author    *User        `json:"author"`
	Member    *GuildMember `json:"member,omitempty"`

	Content         string    `json:"content"`
	Timestamp       time.Time `json:"timestamp"`
	EditedTimestamp time.Time `json:"edited_timestamp"`
	TTS             bool      `json:"tts"`

	MentionEveryone bool                     `json:"mention_everyone"`
	Mentions        []*User                  `json:"mentions"`
	MentionRoles    []*Snowflake             `json:"mention_roles"`
	MentionChannels []*MessageChannelMention `json:"mention_channels,omitempty"`

	Attachments []*MessageAttachment `json:"attachments"`
	Embeds      []*Embed             `json:"embeds"`
	Reactions   []*MessageReaction   `json:"reactions"`

	// Nonce          string                  `json:"nonce,omitempty"`
	Pinned            bool                    `json:"pinned"`
	WebhookID         *Snowflake              `json:"webhook_id,omitempty"`
	Type              MessageType             `json:"type"`
	Activity          *MessageActivity        `json:"activity,omitempty"`
	Application       *Application            `json:"application,omitempty"`
	MessageReference  []*MessageReference     `json:"message_referenced,omitempty"`
	Flags             *MessageFlags           `json:"flags,omitempty"`
	ReferencedMessage *Message                `json:"referenced_message,omitempty"`
	Interaction       *MessageInteraction     `json:"interaction,omitempty"`
	Thread            *Channel                `json:"thread,omitempty"`
	Components        []*InteractionComponent `json:"components,omitempty"`
	StickerItems      []*MessageSticker       `json:"sticker_items,omitempty"`
}

// AddReaction adds a reaction to a message
// emoji: unicode representation or emoji id.
func (m *Message) AddReaction(s *Session, emoji string) (err error) {
	return CreateReaction(s, m.ChannelID, m.ID, emoji)
}

// ClearReaction clears a specific reaction from a message.
// emoji: unicode representation or emoji id.
func (m *Message) ClearReaction(s *Session, emoji string) (err error) {
	return DeleteAllReactionsEmoji(s, m.ChannelID, m.ID, emoji)
}

// ClearReactions clears all reactions from a message.
func (m *Message) ClearReactions(s *Session) (err error) {
	return DeleteAllReactions(s, m.ChannelID, m.ID)
}

// Delete deletes a message.
// reason: reason for deleting the message.
func (m *Message) Delete(s *Session, reason *string) (err error) {
	return DeleteMessage(s, m.ChannelID, m.ID, reason)
}

// Edit edits a message.
// messageArg: arguments for editing the message.
func (m *Message) Edit(s *Session, messageArg MessageParams) (message *Message, err error) {
	return EditMessage(s, m.ChannelID, m.ID, messageArg)
}

// Pin pins a message in a channel.
// reason: reason for pinning a message.
func (m *Message) Pin(s *Session, reason *string) (err error) {
	return PinMessage(s, m.ChannelID, m.ID, reason)
}

// Publish publishes a message. This must be in an announcement channel.
func (m *Message) Publish(s *Session) (message *Message, err error) {
	return CrosspostMessage(s, m.ChannelID, m.ID)
}

// RemoveReaction removes a specific reaction from a specific user.
// emoji: unicode representation or emoji id.
// user: The user to remove the reaction from.
func (m *Message) RemoveReaction(s *Session, emoji string, user User) (err error) {
	return DeleteUserReaction(s, m.ChannelID, m.ID, emoji, user.ID)
}

// Reply will send a new message in the same channel as the target message and references the target.
// This is the same as using Send() and setting the message as the MessageReference.
// messageArg: arguments for sending a message.
func (m *Message) Reply(s *Session, messageArg MessageParams) (message *Message, err error) {
	messageArg.MessageReference = &MessageReference{
		ID:              &m.ID,
		ChannelID:       &m.ChannelID,
		GuildID:         m.GuildID,
		FailIfNotExists: true,
	}

	channel := &Channel{ID: m.ChannelID}

	return channel.Send(s, messageArg)
}

// Unpin unpins a message.
// reason: Reason for unpinning.
func (m *Message) Unpin(s *Session, reason *string) (err error) {
	return UnpinMessage(s, m.ChannelID, m.ID, reason)
}

// MessageParams represents the structure for sending a message on discord.
type MessageParams struct {
	Content          string                    `json:"content"`
	TTS              bool                      `json:"tts"`
	Embeds           []*Embed                  `json:"embeds,omitempty"`
	AllowedMentions  []*MessageAllowedMentions `json:"allowed_mentions,omitempty"`
	MessageReference *MessageReference         `json:"message_reference,omitempty"`
	Components       []*InteractionComponent   `json:"components,omitempty"`
	StickerIDs       []*Snowflake              `json:"sticker_ids,omitempty"`
	Files            []*File                   `json:"files,omitempty"`
	PayloadJSON      *jsoniter.RawMessage      `json:"payload_json,omitempty"`
	Attachments      []*MessageAttachment      `json:"attachments,omitempty"`
}

// MessageInteraction represents an executed interaction.
type MessageInteraction struct {
	ID   Snowflake        `json:"id"`
	Type *InteractionType `json:"type"`
	Name string           `json:"name"`
	User User             `json:"user"`
}

// MessageChannelMention represents a mentioned channel.
type MessageChannelMention struct {
	ID      Snowflake   `json:"id"`
	GuildID Snowflake   `json:"guild_id"`
	Type    ChannelType `json:"type"`
	Name    string      `json:"name"`
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
	Count int32  `json:"count"`
	Me    bool   `json:"me"`
	Emoji *Emoji `json:"emoji"`
}

// MessageAllowedMentions is the structure of the allowed mentions entry.
type MessageAllowedMentions struct {
	Parse       []MessageAllowedMentionsType `json:"parse"`
	Roles       []Snowflake                  `json:"roles"`
	Users       []Snowflake                  `json:"users"`
	RepliedUser bool                         `json:"replied_user"`
}

// MessageAttachment represents a message attachment on discord.
type MessageAttachment struct {
	ID        Snowflake `json:"id"`
	Filename  string    `json:"filename"`
	Size      int32     `json:"size"`
	URL       string    `json:"url"`
	ProxyURL  string    `json:"proxy_url"`
	Height    int32     `json:"height"`
	Width     int32     `json:"width"`
	Ephemeral bool      `json:"ephemeral"`
}

// MessageActivity represents a message activity on discord.
type MessageActivity struct {
	Type    MessageActivityType `json:"type"`
	PartyID string              `json:"party_id,omitempty"`
}
