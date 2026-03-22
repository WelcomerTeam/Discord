package discord

import "context"

// role.go represents all structures for a discord guild role.

// Role represents a role on discord.
type Role struct {
	GuildID      *Snowflake  `json:"guild_id,omitempty"`
	Tags         *RoleTag    `json:"tags,omitempty"`
	Colors       *RoleColors `json:"colors,omitempty"`
	Name         string      `json:"name"`
	Icon         string      `json:"icon,omitempty"`
	UnicodeEmoji string      `json:"unicode_emoji,omitempty"`
	Description  string      `json:"description,omitempty"`
	ID           Snowflake   `json:"id"`
	Permissions  Int64       `json:"permissions"`
	Color        int32       `json:"color"`
	Position     int32       `json:"position"`
	Flags        int32       `json:"flags,omitempty"`
	Hoist        bool        `json:"hoist"`
	Managed      bool        `json:"managed"`
	Mentionable  bool        `json:"mentionable"`
}

// RoleParams represents the structure used to create a role.
type RoleParams struct {
	Name         *string `json:"name,omitempty"`
	Permissions  *Int64  `json:"permissions,omitempty"`
	Color        *int32  `json:"color,omitempty"`
	Hoist        *bool   `json:"hoist,omitempty"`
	Icon         *string `json:"icon,omitempty"`
	UnicodeEmoji *string `json:"unicode_emoji,omitempty"`
	Mentionable  *bool   `json:"mentionable,omitempty"`
}

// Delete deletes a guild role.
// reason: Reason for deleting a guild role.
func (r *Role) Delete(ctx context.Context, session *Session, reason *string) error {
	return DeleteGuildRole(ctx, session, *r.GuildID, r.ID, reason)
}

// Edit edits a guild role.
// params: The role parameters to update the role with.
// reason: Reason for editing a guild role.
func (r *Role) Edit(ctx context.Context, session *Session, params Role, reason *string) error {
	newRole, err := ModifyGuildRole(ctx, session, *r.GuildID, r.ID, params, reason)
	if err != nil {
		return err
	}

	*r = *newRole

	return nil
}

// RoleTag represents extra information about a role.
type RoleTag struct {
	BotID                 *Snowflake `json:"bot_id,omitempty"`
	IntegrationID         *Snowflake `json:"integration_id,omitempty"`
	SubscriptionListingID *Snowflake `json:"subscription_listing_id,omitempty"`
	PremiumSubscriber     bool       `json:"premium_subscriber"`
	AvailableForPurchase  *bool      `json:"available_for_purchase,omitempty"`
	GuildConnections      *bool      `json:"guild_connections,omitempty"`
}

// RoleColors represents extra color information about a role.
type RoleColors struct {
	PrimaryColor   int32  `json:"primary_color"`
	SecondaryColor *int32 `json:"secondary_color,omitempty"`
	TertiaryColor  *int32 `json:"tertiary_color,omitempty"`
}

// ModifyGuildRolePosition represents the argument for modifying guild role positions.
type ModifyGuildRolePosition struct {
	Position *int32    `json:"position,omitempty"`
	ID       Snowflake `json:"id"`
}
