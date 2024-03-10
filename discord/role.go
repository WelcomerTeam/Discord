package discord

// role.go represents all structures for a discord guild role.

// Role represents a role on discord.
type Role struct {
	GuildID      *Snowflake `json:"guild_id,omitempty"`
	Tags         *RoleTag   `json:"tags,omitempty"`
	Name         string     `json:"name"`
	Icon         string     `json:"icon,omitempty"`
	UnicodeEmoji string     `json:"unicode_emoji,omitempty"`
	ID           Snowflake  `json:"id"`
	Permissions  Int64      `json:"permissions"`
	Color        int32      `json:"color"`
	Position     int32      `json:"position"`
	Hoist        bool       `json:"hoist"`
	Managed      bool       `json:"managed"`
	Mentionable  bool       `json:"mentionable"`
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
	BotID             *Snowflake `json:"bot_id,omitempty"`
	IntegrationID     *Snowflake `json:"integration_id,omitempty"`
	PremiumSubscriber bool       `json:"premium_subscriber"`
}

// ModifyGuildRolePosition represents the argument for modifying guild role positions.
type ModifyGuildRolePosition struct {
	Position *int32    `json:"position,omitempty"`
	ID       Snowflake `json:"id"`
}
