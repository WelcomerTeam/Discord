package http

import (
	"net/http"
	"net/url"

	"github.com/WelcomerTeam/Discord/discord"
	"github.com/WelcomerTeam/Discord/structs"
	"golang.org/x/xerrors"
)

func (s *Session) CreateGuild(guildArg structs.Guild) (guild *structs.Guild, err error) {
	endpoint := EndpointGuilds

	err = s.Interface.FetchJJ(s, http.MethodPost, endpoint, guildArg, nil, &guild)
	if err != nil {
		return nil, xerrors.Errorf("Failed to create guild: %v", err)
	}

	return
}

func (s *Session) GetGuild(guildID discord.Snowflake) (guild *structs.Guild, err error) {
	endpoint := EndpointGuild(guildID.String())

	err = s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &guild)
	if err != nil {
		return nil, xerrors.Errorf("Failed to get guild: %v", err)
	}

	return
}

func (s *Session) GetGuildPreview(guildID discord.Snowflake) (guildPreview *structs.Guild, err error) {
	endpoint := EndpointGuildPreview(guildID.String())

	err = s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &guildPreview)
	if err != nil {
		return nil, xerrors.Errorf("Failed to get guild preview: %v", err)
	}

	return
}

func (s *Session) ModifyGuild(guildID discord.Snowflake, guildArg structs.Guild, reason *string) (guild *structs.Guild, err error) {
	endpoint := EndpointGuild(guildID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(structs.AuditLogReasonHeader, *reason)
	}

	err = s.Interface.FetchJJ(s, http.MethodPatch, endpoint, guildArg, headers, &guild)
	if err != nil {
		return nil, xerrors.Errorf("Failed to modify guild: %v", err)
	}

	return
}

func (s *Session) DeleteGuild(guildID discord.Snowflake) (err error) {
	endpoint := EndpointGuild(guildID.String())

	err = s.Interface.FetchJJ(s, http.MethodDelete, endpoint, nil, nil, nil)
	if err != nil {
		return xerrors.Errorf("Failed to delete guild: %v", err)
	}

	return
}

func (s *Session) GetGuildChannels(guildID discord.Snowflake) (channels []*structs.Channel, err error) {
	endpoint := EndpointGuildChannels(guildID.String())

	err = s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &channels)
	if err != nil {
		return nil, xerrors.Errorf("Failed to get guild channels: %v", err)
	}

	return
}

func (s *Session) CreateGuildChannel(guildID discord.Snowflake, channelArg structs.Channel, reason *string) (channel *structs.Channel, err error) {
	endpoint := EndpointGuildChannels(guildID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(structs.AuditLogReasonHeader, *reason)
	}

	err = s.Interface.FetchJJ(s, http.MethodPost, endpoint, channelArg, headers, &channel)
	if err != nil {
		return nil, xerrors.Errorf("Failed to create guild channel: %v", err)
	}

	return
}

func (s *Session) ModifyGuildChannelPositions(guildID discord.Snowflake, channelPermissionsArgs []structs.ChannelPermissionsParams, reason *string) (err error) {
	endpoint := EndpointGuildChannels(guildID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(structs.AuditLogReasonHeader, *reason)
	}

	err = s.Interface.FetchJJ(s, http.MethodPatch, endpoint, channelPermissionsArgs, headers, nil)
	if err != nil {
		return xerrors.Errorf("Failed to modify guild channel positions: %v", err)
	}

	return
}

func (s *Session) GetGuildMember(guildID discord.Snowflake, userID discord.Snowflake) (guildMember *structs.GuildMember, err error) {
	endpoint := EndpointGuildMember(guildID.String(), userID.String())

	err = s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, nil)
	if err != nil {
		return nil, xerrors.Errorf("Failed to get guild member: %v", err)
	}

	return
}

func (s *Session) ListGuildMembers(guildID discord.Snowflake, limit *int32, after *discord.Snowflake) (guildMembers []*structs.GuildMember, err error) {
	endpoint := EndpointGuildMembers(guildID.String())

	var values url.Values

	if limit != nil {
		values.Set("limit", string(*limit))
	}

	if after != nil {
		values.Set("after", after.String())
	}

	if len(values) > 0 {
		endpoint += "?" + values.Encode()
	}

	err = s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &guildMembers)
	if err != nil {
		return nil, xerrors.Errorf("Failed to list guild members: %v", err)
	}

	return
}

func (s *Session) SearchGuildMembers(guildID discord.Snowflake, query string, limit *int32) (guildMembers []*structs.GuildMember, err error) {
	endpoint := EndpointGuildMembersSearch(guildID.String())

	var values url.Values

	values.Set("query", query)

	if limit != nil {
		values.Set("limit", string(*limit))
	}

	endpoint += "?" + values.Encode()

	err = s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &guildMembers)
	if err != nil {
		return nil, xerrors.Errorf("Failed to search guild members: %v", err)
	}

	return
}

func (s *Session) ModifyGuildMember(guildID discord.Snowflake, userID discord.Snowflake, guildMemberArg structs.GuildMember, reason *string) (guildMember *structs.GuildMember, err error) {
	endpoint := EndpointGuildMember(guildID.String(), userID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(structs.AuditLogReasonHeader, *reason)
	}

	err = s.Interface.FetchJJ(s, http.MethodPatch, endpoint, guildMemberArg, headers, &guildMember)
	if err != nil {
		return nil, xerrors.Errorf("Failed to modify guild member: %v", err)
	}

	return
}

func (s *Session) ModifyCurrentMember(guildID discord.Snowflake, userID discord.Snowflake, guildMemberArg structs.GuildMember, reason *string) (guildMember *structs.GuildMember, err error) {
	endpoint := EndpointGuildMember(guildID.String(), "@me")

	headers := http.Header{}

	if reason != nil {
		headers.Add(structs.AuditLogReasonHeader, *reason)
	}

	err = s.Interface.FetchJJ(s, http.MethodPatch, endpoint, guildMemberArg, headers, &guildMember)
	if err != nil {
		return nil, xerrors.Errorf("Failed to modify current member: %v", err)
	}

	return
}

func (s *Session) AddGuildMemberRole(guildID discord.Snowflake, userID discord.Snowflake, roleID discord.Snowflake, reason *string) (err error) {
	endpoint := EndpointGuildMemberRole(guildID.String(), userID.String(), roleID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(structs.AuditLogReasonHeader, *reason)
	}

	err = s.Interface.FetchJJ(s, http.MethodPut, endpoint, nil, headers, nil)
	if err != nil {
		return xerrors.Errorf("Failed to add guild member role: %v", err)
	}

	return
}

func (s *Session) RemoveGuildMemberRole(guildID discord.Snowflake, userID discord.Snowflake, roleID discord.Snowflake, reason *string) (err error) {
	endpoint := EndpointGuildMemberRole(guildID.String(), userID.String(), roleID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(structs.AuditLogReasonHeader, *reason)
	}

	err = s.Interface.FetchJJ(s, http.MethodDelete, endpoint, nil, headers, nil)
	if err != nil {
		return xerrors.Errorf("Failed to remove guild member role: %v", err)
	}

	return
}

func (s *Session) RemoveGuildMember(guildID discord.Snowflake, userID discord.Snowflake, reason *string) (err error) {
	endpoint := EndpointGuildMember(guildID.String(), userID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(structs.AuditLogReasonHeader, *reason)
	}

	err = s.Interface.FetchJJ(s, http.MethodDelete, endpoint, nil, headers, nil)
	if err != nil {
		return xerrors.Errorf("Failed to remove guild member: %v", err)
	}

	return
}

func (s *Session) GetGuildBans(guildID discord.Snowflake) (bans []*structs.GuildBan, err error) {
	endpoint := EndpointGuildBans(guildID.String())

	err = s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &bans)
	if err != nil {
		return nil, xerrors.Errorf("Failed to get guild bans: %v", err)
	}

	return
}

func (s *Session) GetGuildBan(guildID discord.Snowflake, userID discord.Snowflake) (ban *structs.GuildBan, err error) {
	endpoint := EndpointGuildBan(guildID.String(), userID.String())

	err = s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &ban)
	if err != nil {
		return nil, xerrors.Errorf("Failed to get guild ban: %v", err)
	}

	return
}

func (s *Session) CreateGuildBan(guildID discord.Snowflake, userID discord.Snowflake, reason *string) (err error) {
	endpoint := EndpointGuildBan(guildID.String(), userID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(structs.AuditLogReasonHeader, *reason)
	}

	err = s.Interface.FetchJJ(s, http.MethodPut, endpoint, nil, headers, nil)
	if err != nil {
		return xerrors.Errorf("Failed to create guild ban: %v", err)
	}

	return
}

func (s *Session) RemoveGuildBan(guildID discord.Snowflake, userID discord.Snowflake, reason *string) (err error) {
	endpoint := EndpointGuildBan(guildID.String(), userID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(structs.AuditLogReasonHeader, *reason)
	}

	err = s.Interface.FetchJJ(s, http.MethodPut, endpoint, nil, headers, nil)
	if err != nil {
		return xerrors.Errorf("Failed to create guild ban: %v", err)
	}

	return
}

func (s *Session) GetGuildRoles(guildID discord.Snowflake) (roles []*structs.Role, err error) {
	endpoint := EndpointGuildRoles(guildID.String())

	err = s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &roles)
	if err != nil {
		return nil, xerrors.Errorf("Failed to get guild roles: %v", err)
	}

	return
}

func (s *Session) CreateGuildRole(guildID discord.Snowflake, roleArg structs.Role, reason *string) (role *structs.Role, err error) {
	endpoint := EndpointGuildRoles(guildID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(structs.AuditLogReasonHeader, *reason)
	}

	err = s.Interface.FetchJJ(s, http.MethodPost, endpoint, roleArg, headers, &role)
	if err != nil {
		return nil, xerrors.Errorf("Failed to create guild role: %v", err)
	}

	return
}

func (s *Session) ModifyGuildRolePositions(guildID discord.Snowflake, guildRolePositionArgs []structs.ModifyGuildRolePosition, reason *string) (roles []*structs.Role, err error) {
	endpoint := EndpointGuildRoles(guildID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(structs.AuditLogReasonHeader, *reason)
	}

	err = s.Interface.FetchJJ(s, http.MethodPatch, endpoint, guildRolePositionArgs, headers, &roles)
	if err != nil {
		return nil, xerrors.Errorf("Failed to modify guild role positions: %v", err)
	}

	return
}

func (s *Session) ModifyGuildRole(guildID discord.Snowflake, roleID discord.Snowflake, roleArg structs.Role, reason *string) (role *structs.Role, err error) {
	endpoint := EndpointGuildRole(guildID.String(), roleID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(structs.AuditLogReasonHeader, *reason)
	}

	err = s.Interface.FetchJJ(s, http.MethodPatch, endpoint, roleArg, headers, &role)
	if err != nil {
		return nil, xerrors.Errorf("Failed to modify guild role: %v", err)
	}

	return
}

func (s *Session) DeleteGuildRole(guildID discord.Snowflake, roleID discord.Snowflake, reason *string) (err error) {
	endpoint := EndpointGuildRole(guildID.String(), roleID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(structs.AuditLogReasonHeader, *reason)
	}

	err = s.Interface.FetchJJ(s, http.MethodDelete, endpoint, nil, headers, nil)
	if err != nil {
		return xerrors.Errorf("Failed to delete guild role: %v", err)
	}

	return
}

func (s *Session) GetGuildPruneCount(guildID discord.Snowflake, days *int32, includedRoles []discord.Snowflake) (pruned *int32, err error) {
	endpoint := EndpointGuildPrune(guildID.String())

	// Construct includedRoles query argument.
	// Comma delimited array of snowflakes.
	rolesString := ""
	for index, roleID := range includedRoles {
		rolesString += roleID.String()
		if index < len(includedRoles)-1 {
			rolesString += ","
		}
	}

	var values url.Values

	if days != nil {
		values.Set("days", string(*days))
	}

	if len(rolesString) > 0 {
		values.Set("include_roles", rolesString)
	}

	if len(values) > 0 {
		endpoint += "?" + values.Encode()
	}

	prunedStruct := struct {
		Pruned int32 `json:"pruned"`
	}{}

	err = s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &prunedStruct)
	if err != nil {
		return nil, xerrors.Errorf("Failed to get guild prune count: %v", err)
	}

	pruned = &prunedStruct.Pruned

	return
}

func (s *Session) BeginGuildPrune(guildID discord.Snowflake, pruneArg structs.GuildPruneParam, reason *string) (pruned *int32, err error) {
	endpoint := EndpointGuildPrune(guildID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(structs.AuditLogReasonHeader, *reason)
	}

	prunedStruct := struct {
		Pruned int32 `json:"pruned"`
	}{}

	err = s.Interface.FetchJJ(s, http.MethodPost, endpoint, pruneArg, headers, &prunedStruct)
	if err != nil {
		return nil, xerrors.Errorf("Failed to begin guild prune: %v", err)
	}

	pruned = &prunedStruct.Pruned

	return
}

func (s *Session) GetGuildInvites(guildID discord.Snowflake) (invites []*structs.Invite, err error) {
	endpoint := EndpointGuildInvites(guildID.String())

	err = s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &invites)
	if err != nil {
		return nil, xerrors.Errorf("Failed to get guild invites: %v", err)
	}

	return
}

func (s *Session) GetGuildIntegrations(guildID discord.Snowflake) (integrations []*structs.Integration, err error) {
	endpoint := EndpointGuildIntegrations(guildID.String())

	err = s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &integrations)
	if err != nil {
		return nil, xerrors.Errorf("Failed to get guild integrations: %v", err)
	}

	return
}

func (s *Session) DeleteGuildIntegration(guildID discord.Snowflake, integrationID discord.Snowflake, reason *string) (err error) {
	endpoint := EndpointGuildIntegration(guildID.String(), integrationID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(structs.AuditLogReasonHeader, *reason)
	}

	err = s.Interface.FetchJJ(s, http.MethodDelete, endpoint, nil, headers, nil)
	if err != nil {
		return xerrors.Errorf("Failed to delete guild integration: %v", err)
	}

	return
}

func (s *Session) GetGuildVanityURL(guildID discord.Snowflake) (invite *structs.Invite, err error) {
	endpoint := EndpointGuildVanityURL(guildID.String())

	err = s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &invite)
	if err != nil {
		return nil, xerrors.Errorf("Failed to get guild vanity url: %v", err)
	}

	return
}

// TODO: GetGuildWidgetImage
// TODO: GetGuildWelcomeScreen
// TODO: ModifyGuildWelcomeScreen
// TODO: GetGuildWidget
// TODO: ModifyGuildWidget
// TODO: GetGuildWidgetSettings
// TODO: AddGuildMember
// TODO: GetGuildVoiceRegions
// TODO: ListActiveThreads

// TODO: ModifyCurrentUserVoiceState
// TODO: ModifyUserVoiceState

// TODO: ListScheduledEventsforGuild
// TODO: CreateGuildScheduledEvent
// TODO: GetGuildScheduledEvent
// TODO: ModifyGuildScheduledEvent
// TODO: DeleteGuildScheduledEvent
// TODO: GetGuildScheduledEventUsers
// TODO: GuildScheduledEventStatusUpdateAutomation
// TODO: GuildScheduledEventPermissionsRequirements

// TODO: GetGuildTemplate
// TODO: CreateGuildfromGuildTemplate
// TODO: GetGuildTemplates
// TODO: CreateGuildTemplate
// TODO: SyncGuildTemplate
// TODO: ModifyGuildTemplate
// TODO: DeleteGuildTemplate
