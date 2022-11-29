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
	InteractionTypeModalSubmit
)

// InteractionCallbackType represents the type of interaction callbacks.
type InteractionCallbackType uint8

const (
	InteractionCallbackTypePong InteractionCallbackType = 1 + iota

	_
	_

	// InteractionCallbackTypeChannelMessageSource responds to an interaction with a message.
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

	// InteractionCallbackTypeModal responds to an interaction with a popup modal.
	InteractionCallbackTypeModal
)

// InteractionComponentType represents the type of component.
type InteractionComponentType uint8

const (
	InteractionComponentTypeActionRow InteractionComponentType = 1 + iota
	InteractionComponentTypeButton
	InteractionComponentTypeStringSelect
	InteractionComponentTypeTextInput
	InteractionComponentTypeUserInput
	InteractionComponentTypeRoleSelect
	InteractionComponentTypeMentionableSelect
	InteractionComponentTypeChannelSelect
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
	Type          InteractionType  `json:"type"`
	Data          *InteractionData `json:"data,omitempty"`

	GuildID        *Snowflake   `json:"guild_id,omitempty"`
	ChannelID      *Snowflake   `json:"channel_id,omitempty"`
	Member         *GuildMember `json:"member,omitempty"`
	User           *User        `json:"user,omitempty"`
	Token          string       `json:"token"`
	Version        int32        `json:"version"`
	Message        *Message     `json:"message,omitempty"`
	AppPermissions *Int64       `json:"app_permissions"`
	Locale         string       `json:"locale,omitempty"`
	GuildLocale    string       `json:"guild_locale,omitempty"`
}

// SendResponse sends an interacion response.
// interactionType: The type of interaction callback.
// messageArg: arguments for sending message.
// choices: optional autocomplete choices.
func (i *Interaction) SendResponse(s *Session, interactionType InteractionCallbackType, messageParams WebhookMessageParams, choices []*ApplicationCommandOptionChoice) error {
	return CreateInteractionResponse(s, i.ID, i.Token, InteractionResponse{
		Type: interactionType,
		Data: &InteractionCallbackData{
			WebhookMessageParams: messageParams,
			Choices:              choices,
		},
	})
}

// EditOriginalResponse edits the original interaction response.
// messageArg: arguments for editing message.
func (i *Interaction) EditOriginalResponse(s *Session, messageParams WebhookMessageParams) (*Message, error) {
	return EditOriginalInteractionResponse(s, i.ApplicationID, i.Token, messageParams)
}

// DeleteOriginalResponse deletes the original interaction response.
func (i *Interaction) DeleteOriginalResponse(s *Session) error {
	return DeleteOriginalInteractionResponse(s, i.ApplicationID, i.Token)
}

// SendFollowup sends a followup message.
// messageArg: arguments for sending message.
func (i *Interaction) SendFollowup(s *Session, messageParams WebhookMessageParams) (*InteractionFollowup, error) {
	message, err := CreateFollowupMessage(s, i.ApplicationID, i.Token, messageParams)
	if err != nil {
		return nil, err
	}

	return &InteractionFollowup{
		Message:     message,
		Interaction: i,
	}, nil
}

// InteractionFollowup represents a follow up message containing both message and the interaction parent.
type InteractionFollowup struct {
	*Message
	*Interaction
}

// EditFollowup edits the followup message.
// messageArg: arguments for editing message.
func (inf *InteractionFollowup) EditFollowup(s *Session, messageParams WebhookMessageParams) (*Message, error) {
	return EditFollowupMessage(s, inf.ApplicationID, inf.Token, inf.Message.ID, messageParams)
}

// DeleteFollowup deletes the followup message.
func (inf *InteractionFollowup) DeleteFollowup(s *Session) error {
	return DeleteFollowupMessage(s, inf.ApplicationID, inf.Token, inf.Message.ID)
}

// InteractionResponse represents the interaction response object.
type InteractionResponse struct {
	Type InteractionCallbackType  `json:"type"`
	Data *InteractionCallbackData `json:"data,omitempty"`
}

// InteractionData represents the structure of interaction data.
type InteractionData struct {
	ID       Snowflake                `json:"id"`
	Name     string                   `json:"name"`
	Type     ApplicationCommandType   `json:"type"`
	Resolved *InteractionResolvedData `json:"resolved,omitempty"`
	Options  []*InteractionDataOption `json:"options,omitempty"`
	GuildID  *Snowflake               `json:"guild_id,omitempty"`
	TargetID *Snowflake               `json:"target_id,omitempty"`

	CustomID      string                     `json:"custom_id,omitempty"`
	ComponentType *InteractionComponentType  `json:"component_type,omitempty"`
	Values        []*ApplicationSelectOption `json:"values,omitempty"`
	Components    []*InteractionComponent    `json:"components,omitempty"`
}

// InteractionData represents the structure of the interaction callback data.
// Not all message fields are supported, allowed fields are: tts, content
// embeds, allowed_mentions, flags, components and attachments.
type InteractionCallbackData struct {
	WebhookMessageParams

	Title      string                            `json:"title,omitempty"`
	CustomID   string                            `json:"custom_id,omitempty"`
	Components []*InteractionComponent           `json:"components,omitempty"`
	Choices    []*ApplicationCommandOptionChoice `json:"choices,omitempty"`
}

// InteractionDataOption represents the structure of an interaction option.
type InteractionDataOption struct {
	Name    string                       `json:"name"`
	Type    ApplicationCommandOptionType `json:"type"`
	Value   jsoniter.RawMessage          `json:"value,omitempty"`
	Options []*InteractionDataOption     `json:"options,omitempty"`
	Focused bool                         `json:"focused,omitempty"`
}

// InteractionResolvedData represents any extra payload data for an interaction.
type InteractionResolvedData struct {
	Users       map[Snowflake]*User              `json:"users,omitempty"`
	Members     map[Snowflake]*GuildMember       `json:"members,omitempty"`
	Roles       map[Snowflake]*Role              `json:"roles,omitempty"`
	Channels    map[Snowflake]*Channel           `json:"channels,omitempty"`
	Messages    map[Snowflake]*Message           `json:"messages,omitempty"`
	Attachments map[Snowflake]*MessageAttachment `json:"attachments,omitempty"`
}

// InteractionComponent represents the structure of a component.
type InteractionComponent struct {
	Type     InteractionComponentType  `json:"type"`
	Style    InteractionComponentStyle `json:"style,omitempty"`
	Label    string                    `json:"label,omitempty"`
	Emoji    *Emoji                    `json:"emoji,omitempty"`
	CustomID string                    `json:"custom_id,omitempty"`
	URL      string                    `json:"url,omitempty"`
	Disabled bool                      `json:"disabled"`

	Options      []*ApplicationSelectOption `json:"options,omitempty"`
	ChannelTypes []*ChannelType             `json:"channel_types,omitempty"`
	Placeholder  string                     `json:"placeholder,omitempty"`
	MinValues    *int32                     `json:"min_values,omitempty"`
	MaxValues    *int32                     `json:"max_values,omitempty"`

	Components []*InteractionComponent `json:"components,omitempty"`
}

func NewInteractionComponent(componentType InteractionComponentType) *InteractionComponent {
	return &InteractionComponent{
		Type: componentType,
	}
}

func (ic *InteractionComponent) SetCustomID(customID string) *InteractionComponent {
	ic.CustomID = customID

	return ic
}

func (ic *InteractionComponent) SetDisabled(disabled bool) *InteractionComponent {
	ic.Disabled = disabled

	return ic
}

func (ic *InteractionComponent) SetStyle(style InteractionComponentStyle) *InteractionComponent {
	ic.Style = style

	return ic
}

func (ic *InteractionComponent) SetLabel(label string) *InteractionComponent {
	ic.Label = label

	return ic
}

func (ic *InteractionComponent) SetEmoji(emoji *Emoji) *InteractionComponent {
	ic.Emoji = emoji

	return ic
}

func (ic *InteractionComponent) SetURL(url string) *InteractionComponent {
	ic.URL = url

	return ic
}

func (ic *InteractionComponent) AddOption(option *ApplicationSelectOption) *InteractionComponent {
	ic.Options = append(ic.Options, option)

	return ic
}

func (ic *InteractionComponent) SetPlaceholder(placeholder string) *InteractionComponent {
	ic.Placeholder = placeholder

	return ic
}

func (ic *InteractionComponent) SetMinMaxValues(minValue *int32, maxValue *int32) *InteractionComponent {
	ic.MinValues = minValue
	ic.MaxValues = maxValue

	return ic
}

func (ic *InteractionComponent) AddComponent(component *InteractionComponent) *InteractionComponent {
	ic.Components = append(ic.Components, component)

	return ic
}
