package discord

import (
	"github.com/WelcomerTeam/Discord/discord"
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

// Message represents a message on Discord.
type Message struct {
	ID        discord.Snowflake  `json:"id"`
	ChannelID discord.Snowflake  `json:"channel_id"`
	GuildID   *discord.Snowflake `json:"guild_id,omitempty"`
	Author    *User              `json:"author"`
	Member    *GuildMember       `json:"member,omitempty"`

	Content         string `json:"content"`
	Timestamp       string `json:"timestamp"`
	EditedTimestamp string `json:"edited_timestamp"`
	TTS             bool   `json:"tts"`

	MentionEveryone bool                     `json:"mention_everyone"`
	Mentions        []*User                  `json:"mentions"`
	MentionRoles    []*discord.Snowflake     `json:"mention_roles"`
	MentionChannels []*MessageChannelMention `json:"mention_channels,omitempty"`

	Attachments []*MessageAttachment `json:"attachments"`
	Embeds      []*Embed             `json:"embeds"`
	Reactions   []*MessageReaction   `json:"reactions"`

	// Nonce          string                  `json:"nonce,omitempty"`
	Pinned            bool                    `json:"pinned"`
	WebhookID         *discord.Snowflake      `json:"webhook_id,omitempty"`
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

// MessageParams represents the structure for sending a message on Discord.
type MessageParams struct {
	Content          string                    `json:"content"`
	TTS              bool                      `json:"tts"`
	Embeds           []*Embed                  `json:"embeds"`
	AllowedMentions  []*MessageAllowedMentions `json:"allowed_mentions"`
	MessageReference *MessageReference         `json:"message_reference,omitempty"`
	Components       []*InteractionComponent   `json:"components,omitempty"`
	StickerIDs       []*discord.Snowflake      `json:"sticker_ids,omitempty"`
	Files            []*File                   `json:"files,omitempty"`
	PayloadJSON      *jsoniter.RawMessage      `json:"payload_json,omitempty"`
	Attachments      []*MessageAttachment      `json:"attachments,omitempty"`
}

// MessageInteraction represents an executed interaction.
type MessageInteraction struct {
	ID   discord.Snowflake `json:"id"`
	Type *InteractionType  `json:"type"`
	Name string            `json:"name"`
	User User              `json:"user"`
}

// MessageChannelMention represents a mentioned channel.
type MessageChannelMention struct {
	ID      discord.Snowflake `json:"id"`
	GuildID discord.Snowflake `json:"guild_id"`
	Type    ChannelType       `json:"type"`
	Name    string            `json:"name"`
}

// MessageReference represents crossposted messages or replys.
type MessageReference struct {
	ID              *discord.Snowflake `json:"message_id,omitempty"`
	ChannelID       *discord.Snowflake `json:"channel_id,omitempty"`
	GuildID         *discord.Snowflake `json:"guild_id,omitempty"`
	FailIfNotExists bool               `json:"fail_if_not_exists"`
}

// MessageReaction represents a reaction to a message on Discord.
type MessageReaction struct {
	Count int32  `json:"count"`
	Me    bool   `json:"me"`
	Emoji *Emoji `json:"emoji"`
}

// MessageAllowedMentions is the structure of the allowed mentions entry.
type MessageAllowedMentions struct {
	Parse       []MessageAllowedMentionsType `json:"parse"`
	Roles       []discord.Snowflake          `json:"roles"`
	Users       []discord.Snowflake          `json:"users"`
	RepliedUser bool                         `json:"replied_user"`
}

// MessageAttachment represents a message attachment on discord.
type MessageAttachment struct {
	ID        discord.Snowflake `json:"id"`
	Filename  string            `json:"filename"`
	Size      int32             `json:"size"`
	URL       string            `json:"url"`
	ProxyURL  string            `json:"proxy_url"`
	Height    int32             `json:"height"`
	Width     int32             `json:"width"`
	Ephemeral bool              `json:"ephemeral"`
}

// MessageActivity represents a message activity on Discord.
type MessageActivity struct {
	Type    MessageActivityType `json:"type"`
	PartyID string              `json:"party_id,omitempty"`
}
