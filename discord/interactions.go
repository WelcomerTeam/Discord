package discord

import (
	jsoniter "github.com/json-iterator/go"
)

// interactions.go represents the interaction objects.

// InteractionType represents the type of interaction.
type InteractionType uint8

const (
	InteractionTypePing InteractionType = 1 + iota
	InteractionTypeApplicationCommand
	InteractionTypeMessageComponent
	InteractionTypeApplicationCommandAutocomplete
)

// InteractionCallbackType represents the type of interaction callbacks.
type InteractionCallbackType uint8

const (
	InteractionCallbackTypePong InteractionCallbackType = 1 + iota

	// InteractionCallbackTypeChannelMessageSource responds to an interaction with a message
	InteractionCallbackTypeChannelMessageSource

	// InteractionCallbackTypeDeferredChannelMessageSource acknowledges an interaction and
	// edits a response later, users see a loading state.
	InteractionCallbackTypeDeferredChannelMessageSource

	// InteractionCallbackTypeDeferredUpdateMessage acknowledges an interaction and edits
	// a response later, users do not see a loading state.
	InteractionCallbackTypeDeferredUpdateMessage

	// InteractionCallbackTypeUpdateMessage edits the message the component was attached to.
	InteractionCallbackTypeUpdateMessage

	// InteractionCallbackTypeAutocompleteResult responds to an autocomplete interaction.
	InteractionCallbackTypeAutocompleteResult
)

// InteractionComponentType represents the type of component.
type InteractionComponentType uint8

const (
	InteractionComponentTypeActionRow InteractionComponentType = 1 + iota
	InteractionComponentTypeButton
	InteractionComponentTypeSelectMenu
)

// InteractionComponentStyle represents the style of a component.
type InteractionComponentStyle uint8

const (
	InteractionComponentStylePrimary InteractionComponentStyle = 1 + iota
	InteractionComponentStyleSecondary
	InteractionComponentStyleSuccess
	InteractionComponentStyleDanger
	InteractionComponentStyleLink
)

// Interaction represents the structure of an interaction.
type Interaction struct {
	ID            Snowflake        `json:"id"`
	ApplicationID Snowflake        `json:"application_id"`
	Type          *InteractionType `json:"type"`
	Data          *InteractionData `json:"data,omitempty"`

	GuildID     *Snowflake   `json:"guild_id,omitempty"`
	ChannelID   *Snowflake   `json:"channel_id,omitempty"`
	Member      *GuildMember `json:"member,omitempty"`
	User        *User        `json:"user,omitempty"`
	Token       string       `json:"token"`
	Version     int32        `json:"version"`
	Message     *Message     `json:"message,omitempty"`
	Locale      string       `json:"locale,omitempty"`
	GuildLocale string       `json:"guild_locale,omitempty"`
}

// InteractionResponse represents the interaction response object.
type InteractionResponse struct {
	Type *InteractionCallbackType `json:"type"`
	Data *InteractionCallbackData `json:"data,omitempty"`
}

// InteractionData represents the structure of interaction data.
type InteractionData struct {
	ID            Snowflake                  `json:"id"`
	Name          string                     `json:"name"`
	Type          ApplicationCommandType     `json:"type"`
	Resolved      *InteractionResolvedData   `json:"resolved,omitempty"`
	Options       []*InteractionDataOption   `json:"option,omitempty"`
	CustomID      string                     `json:"custom_id,omitempty"`
	ComponentType *InteractionComponentType  `json:"component_type,omitempty"`
	Values        []*ApplicationSelectOption `json:"values,omitempty"`
	TargetID      *Snowflake                 `json:"target_id,omitempty"`
}

// InteractionData represents the structure of the interaction callback data.
// Not all message fields are supported, allowed fields are: tts, content
// embeds, allowed_mentions, flags, components and attachments.
type InteractionCallbackData struct {
	*Message
	Choices []*ApplicationCommandOptionChoice `json:"choices,omitempty"`
}

// InteractionDataOption represents the structure of an interaction option.
type InteractionDataOption struct {
	Name    string                       `json:"name"`
	Type    ApplicationCommandOptionType `json:"type"`
	Value   jsoniter.RawMessage          `json:"value,omitempty"`
	Options []*InteractionDataOption     `json:"options,omitempty"`
	Focused bool                         `json:"focused"`
}

// InteractionResolvedData represents any extra payload data for an interaction.
type InteractionResolvedData struct {
	Users    []*User        `json:"users,omitempty"`
	Members  []*GuildMember `json:"members,omitempty"`
	Roles    []*Role        `json:"roles,omitempty"`
	Channels []*Channel     `json:"channels,omitempty"`
	Messages []*Message     `json:"messages,omitempty"`
}

// InteractionComponent represents the structure of a component.
type InteractionComponent struct {
	Type        InteractionComponentType   `json:"type"`
	CustomID    string                     `json:"custom_id,omitempty"`
	Disabled    bool                       `json:"disabled"`
	Style       InteractionComponentStyle  `json:"style,omitempty"`
	Label       string                     `json:"label,omitempty"`
	Emoji       *Emoji                     `json:"emoji,omitempty"`
	URL         string                     `json:"url,omitempty"`
	Options     []*ApplicationSelectOption `json:"options,omitempty"`
	Placeholder string                     `json:"placeholder,omitempty"`
	MinValues   *int32                     `json:"min_values,omitempty"`
	MaxValues   *int32                     `json:"max_values,omitempty"`
	Components  []*InteractionComponent    `json:"components,omitempty"`
}
