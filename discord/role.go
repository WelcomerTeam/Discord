package discord

// role.go represents all structures for a discord guild role.

// Role represents a role on discord.
type Role struct {
	ID           Snowflake  `json:"id"`
	GuildID      *Snowflake `json:"guild_id,omitempty"`
	Name         string     `json:"name"`
	Color        int32      `json:"color"`
	Hoist        bool       `json:"hoist"`
	Icon         string     `json:"icon,omitempty"`
	UnicodeEmoji string     `json:"unicode_emoji,omitempty"`
	Position     int32      `json:"position"`
	Permissions  Int64      `json:"permissions"`
	Managed      bool       `json:"managed"`
	Mentionable  bool       `json:"mentionable"`
	Tags         *RoleTag   `json:"tags,omitempty"`
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
func (r *Role) Delete(s *Session, reason *string) error {
	return DeleteGuildRole(s, *r.GuildID, r.ID, reason)
}

// Edit edits a guild role.
// params: The role parameters to update the role with.
// reason: Reason for editing a guild role.
func (r *Role) Edit(s *Session, params Role, reason *string) error {
	newRole, err := ModifyGuildRole(s, *r.GuildID, r.ID, params, reason)
	if err != nil {
		return err
	}

	*r = *newRole

	return nil
}

// RoleTag represents extra information about a role.
type RoleTag struct {
	PremiumSubscriber bool       `json:"premium_subscriber"`
	BotID             *Snowflake `json:"bot_id,omitempty"`
	IntegrationID     *Snowflake `json:"integration_id,omitempty"`
}

// ModifyGuildRolePosition represents the argument for modifying guild role positions.
type ModifyGuildRolePosition struct {
	ID       Snowflake `json:"id"`
	Position *int32    `json:"position,omitempty"`
}
