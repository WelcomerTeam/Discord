package discord

import (
	"net/http"
	"net/url"

	"golang.org/x/xerrors"
)

func CreateGuild(s *Session, guildArg Guild) (guild *Guild, err error) {
	endpoint := EndpointGuilds

	err = s.Interface.FetchJJ(s, http.MethodPost, endpoint, guildArg, nil, &guild)
	if err != nil {
		return nil, xerrors.Errorf("Failed to create guild: %v", err)
	}

	return
}

func GetGuild(s *Session, guildID Snowflake) (guild *Guild, err error) {
	endpoint := EndpointGuild(guildID.String())

	err = s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &guild)
	if err != nil {
		return nil, xerrors.Errorf("Failed to get guild: %v", err)
	}

	return
}

func GetGuildPreview(s *Session, guildID Snowflake) (guildPreview *Guild, err error) {
	endpoint := EndpointGuildPreview(guildID.String())

	err = s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &guildPreview)
	if err != nil {
		return nil, xerrors.Errorf("Failed to get guild preview: %v", err)
	}

	return
}

func ModifyGuild(s *Session, guildID Snowflake, guildArg GuildParam, reason *string) (guild *Guild, err error) {
	endpoint := EndpointGuild(guildID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err = s.Interface.FetchJJ(s, http.MethodPatch, endpoint, guildArg, headers, &guild)
	if err != nil {
		return nil, xerrors.Errorf("Failed to modify guild: %v", err)
	}

	return
}

func DeleteGuild(s *Session, guildID Snowflake) (err error) {
	endpoint := EndpointGuild(guildID.String())

	err = s.Interface.FetchJJ(s, http.MethodDelete, endpoint, nil, nil, nil)
	if err != nil {
		return xerrors.Errorf("Failed to delete guild: %v", err)
	}

	return
}

func GetGuildChannels(s *Session, guildID Snowflake) (channels []*Channel, err error) {
	endpoint := EndpointGuildChannels(guildID.String())

	err = s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &channels)
	if err != nil {
		return nil, xerrors.Errorf("Failed to get guild channels: %v", err)
	}

	return
}

func CreateGuildChannel(s *Session, guildID Snowflake, channelArg ChannelParams, reason *string) (channel *Channel, err error) {
	endpoint := EndpointGuildChannels(guildID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err = s.Interface.FetchJJ(s, http.MethodPost, endpoint, channelArg, headers, &channel)
	if err != nil {
		return nil, xerrors.Errorf("Failed to create guild channel: %v", err)
	}

	return
}

func ModifyGuildChannelPositions(s *Session, guildID Snowflake, channelPermissionsArgs []ChannelPermissionsParams, reason *string) (err error) {
	endpoint := EndpointGuildChannels(guildID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err = s.Interface.FetchJJ(s, http.MethodPatch, endpoint, channelPermissionsArgs, headers, nil)
	if err != nil {
		return xerrors.Errorf("Failed to modify guild channel positions: %v", err)
	}

	return
}

func GetGuildMember(s *Session, guildID Snowflake, userID Snowflake) (guildMember *GuildMember, err error) {
	endpoint := EndpointGuildMember(guildID.String(), userID.String())

	err = s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, nil)
	if err != nil {
		return nil, xerrors.Errorf("Failed to get guild member: %v", err)
	}

	return
}

func ListGuildMembers(s *Session, guildID Snowflake, limit *int32, after *Snowflake) (guildMembers []*GuildMember, err error) {
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

func SearchGuildMembers(s *Session, guildID Snowflake, query string, limit *int32) (guildMembers []*GuildMember, err error) {
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

func ModifyGuildMember(s *Session, guildID Snowflake, userID Snowflake, guildMemberArg GuildMember, reason *string) (guildMember *GuildMember, err error) {
	endpoint := EndpointGuildMember(guildID.String(), userID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err = s.Interface.FetchJJ(s, http.MethodPatch, endpoint, guildMemberArg, headers, &guildMember)
	if err != nil {
		return nil, xerrors.Errorf("Failed to modify guild member: %v", err)
	}

	return
}

func ModifyCurrentMember(s *Session, guildID Snowflake, userID Snowflake, guildMemberArg GuildMember, reason *string) (guildMember *GuildMember, err error) {
	endpoint := EndpointGuildMember(guildID.String(), "@me")

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err = s.Interface.FetchJJ(s, http.MethodPatch, endpoint, guildMemberArg, headers, &guildMember)
	if err != nil {
		return nil, xerrors.Errorf("Failed to modify current member: %v", err)
	}

	return
}

func AddGuildMemberRole(s *Session, guildID Snowflake, userID Snowflake, roleID Snowflake, reason *string) (err error) {
	endpoint := EndpointGuildMemberRole(guildID.String(), userID.String(), roleID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err = s.Interface.FetchJJ(s, http.MethodPut, endpoint, nil, headers, nil)
	if err != nil {
		return xerrors.Errorf("Failed to add guild member role: %v", err)
	}

	return
}

func RemoveGuildMemberRole(s *Session, guildID Snowflake, userID Snowflake, roleID Snowflake, reason *string) (err error) {
	endpoint := EndpointGuildMemberRole(guildID.String(), userID.String(), roleID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err = s.Interface.FetchJJ(s, http.MethodDelete, endpoint, nil, headers, nil)
	if err != nil {
		return xerrors.Errorf("Failed to remove guild member role: %v", err)
	}

	return
}

func RemoveGuildMember(s *Session, guildID Snowflake, userID Snowflake, reason *string) (err error) {
	endpoint := EndpointGuildMember(guildID.String(), userID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err = s.Interface.FetchJJ(s, http.MethodDelete, endpoint, nil, headers, nil)
	if err != nil {
		return xerrors.Errorf("Failed to remove guild member: %v", err)
	}

	return
}

func GetGuildBans(s *Session, guildID Snowflake) (bans []*GuildBan, err error) {
	endpoint := EndpointGuildBans(guildID.String())

	err = s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &bans)
	if err != nil {
		return nil, xerrors.Errorf("Failed to get guild bans: %v", err)
	}

	return
}

func GetGuildBan(s *Session, guildID Snowflake, userID Snowflake) (ban *GuildBan, err error) {
	endpoint := EndpointGuildBan(guildID.String(), userID.String())

	err = s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &ban)
	if err != nil {
		return nil, xerrors.Errorf("Failed to get guild ban: %v", err)
	}

	return
}

func CreateGuildBan(s *Session, guildID Snowflake, userID Snowflake, reason *string) (err error) {
	endpoint := EndpointGuildBan(guildID.String(), userID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err = s.Interface.FetchJJ(s, http.MethodPut, endpoint, nil, headers, nil)
	if err != nil {
		return xerrors.Errorf("Failed to create guild ban: %v", err)
	}

	return
}

func RemoveGuildBan(s *Session, guildID Snowflake, userID Snowflake, reason *string) (err error) {
	endpoint := EndpointGuildBan(guildID.String(), userID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err = s.Interface.FetchJJ(s, http.MethodPut, endpoint, nil, headers, nil)
	if err != nil {
		return xerrors.Errorf("Failed to create guild ban: %v", err)
	}

	return
}

func GetGuildRoles(s *Session, guildID Snowflake) (roles []*Role, err error) {
	endpoint := EndpointGuildRoles(guildID.String())

	err = s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &roles)
	if err != nil {
		return nil, xerrors.Errorf("Failed to get guild roles: %v", err)
	}

	return
}

func CreateGuildRole(s *Session, guildID Snowflake, roleArg RoleParams, reason *string) (role *Role, err error) {
	endpoint := EndpointGuildRoles(guildID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err = s.Interface.FetchJJ(s, http.MethodPost, endpoint, roleArg, headers, &role)
	if err != nil {
		return nil, xerrors.Errorf("Failed to create guild role: %v", err)
	}

	return
}

func ModifyGuildRolePositions(s *Session, guildID Snowflake, guildRolePositionArgs []ModifyGuildRolePosition, reason *string) (roles []*Role, err error) {
	endpoint := EndpointGuildRoles(guildID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err = s.Interface.FetchJJ(s, http.MethodPatch, endpoint, guildRolePositionArgs, headers, &roles)
	if err != nil {
		return nil, xerrors.Errorf("Failed to modify guild role positions: %v", err)
	}

	return
}

func ModifyGuildRole(s *Session, guildID Snowflake, roleID Snowflake, roleArg Role, reason *string) (role *Role, err error) {
	endpoint := EndpointGuildRole(guildID.String(), roleID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err = s.Interface.FetchJJ(s, http.MethodPatch, endpoint, roleArg, headers, &role)
	if err != nil {
		return nil, xerrors.Errorf("Failed to modify guild role: %v", err)
	}

	return
}

func DeleteGuildRole(s *Session, guildID Snowflake, roleID Snowflake, reason *string) (err error) {
	endpoint := EndpointGuildRole(guildID.String(), roleID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err = s.Interface.FetchJJ(s, http.MethodDelete, endpoint, nil, headers, nil)
	if err != nil {
		return xerrors.Errorf("Failed to delete guild role: %v", err)
	}

	return
}

func GetGuildPruneCount(s *Session, guildID Snowflake, days *int32, includedRoles []Snowflake) (pruned *int32, err error) {
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

	return pruned, nil
}

func BeginGuildPrune(s *Session, guildID Snowflake, days *int32, includedRoles []Snowflake, computePruneCount bool, reason *string) (pruned *int32, err error) {
	endpoint := EndpointGuildPrune(guildID.String())

	pruneArg := GuildPruneParam{
		ComputePruneCount: computePruneCount,
	}

	if days != nil {
		pruneArg.Days = days
	}

	for _, includedRole := range includedRoles {
		role := includedRole
		pruneArg.IncludeRoles = append(pruneArg.IncludeRoles, &role)
	}

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	prunedStruct := struct {
		Pruned int32 `json:"pruned"`
	}{}

	err = s.Interface.FetchJJ(s, http.MethodPost, endpoint, pruneArg, headers, &prunedStruct)
	if err != nil {
		return nil, xerrors.Errorf("Failed to begin guild prune: %v", err)
	}

	pruned = &prunedStruct.Pruned

	return pruned, nil
}

func GetGuildInvites(s *Session, guildID Snowflake) (invites []*Invite, err error) {
	endpoint := EndpointGuildInvites(guildID.String())

	err = s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &invites)
	if err != nil {
		return nil, xerrors.Errorf("Failed to get guild invites: %v", err)
	}

	return
}

func GetGuildIntegrations(s *Session, guildID Snowflake) (integrations []*Integration, err error) {
	endpoint := EndpointGuildIntegrations(guildID.String())

	err = s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &integrations)
	if err != nil {
		return nil, xerrors.Errorf("Failed to get guild integrations: %v", err)
	}

	return
}

func DeleteGuildIntegration(s *Session, guildID Snowflake, integrationID Snowflake, reason *string) (err error) {
	endpoint := EndpointGuildIntegration(guildID.String(), integrationID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err = s.Interface.FetchJJ(s, http.MethodDelete, endpoint, nil, headers, nil)
	if err != nil {
		return xerrors.Errorf("Failed to delete guild integration: %v", err)
	}

	return
}

func GetGuildVanityURL(s *Session, guildID Snowflake) (invite *Invite, err error) {
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
