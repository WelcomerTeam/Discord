package discord

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

func CreateGuild(ctx context.Context, s *Session, guildArg Guild) (*Guild, error) {
	endpoint := EndpointGuilds

	var guild *Guild

	err := s.Interface.FetchJJ(ctx, s, http.MethodPost, endpoint, guildArg, nil, &guild)
	if err != nil {
		return nil, fmt.Errorf("failed to create guild: %w", err)
	}

	return guild, nil
}

func GetGuild(ctx context.Context, s *Session, guildID Snowflake) (*Guild, error) {
	endpoint := EndpointGuild(guildID.String())

	var guild *Guild

	err := s.Interface.FetchJJ(ctx, s, http.MethodGet, endpoint, nil, nil, &guild)
	if err != nil {
		return nil, fmt.Errorf("failed to get guild: %w", err)
	}

	return guild, nil
}

func GetGuildPreview(ctx context.Context, s *Session, guildID Snowflake) (*Guild, error) {
	endpoint := EndpointGuildPreview(guildID.String())

	var guildPreview *Guild

	err := s.Interface.FetchJJ(ctx, s, http.MethodGet, endpoint, nil, nil, &guildPreview)
	if err != nil {
		return nil, fmt.Errorf("failed to get guild preview: %w", err)
	}

	return guildPreview, nil
}

func ModifyGuild(ctx context.Context, s *Session, guildID Snowflake, guildArg GuildParam, reason *string) (*Guild, error) {
	endpoint := EndpointGuild(guildID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	var guild *Guild

	err := s.Interface.FetchJJ(ctx, s, http.MethodPatch, endpoint, guildArg, headers, &guild)
	if err != nil {
		return nil, fmt.Errorf("failed to modify guild: %w", err)
	}

	return guild, nil
}

func DeleteGuild(ctx context.Context, s *Session, guildID Snowflake) error {
	endpoint := EndpointGuild(guildID.String())

	err := s.Interface.FetchJJ(ctx, s, http.MethodDelete, endpoint, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete guild: %w", err)
	}

	return nil
}

func GetGuildChannels(ctx context.Context, s *Session, guildID Snowflake) ([]Channel, error) {
	endpoint := EndpointGuildChannels(guildID.String())

	var channels []Channel

	err := s.Interface.FetchJJ(ctx, s, http.MethodGet, endpoint, nil, nil, &channels)
	if err != nil {
		return nil, fmt.Errorf("failed to get guild channels: %w", err)
	}

	return channels, nil
}

func CreateGuildChannel(ctx context.Context, s *Session, guildID Snowflake, channelParams ChannelParams, reason *string) (*Channel, error) {
	endpoint := EndpointGuildChannels(guildID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	var channel *Channel

	err := s.Interface.FetchJJ(ctx, s, http.MethodPost, endpoint, channelParams, headers, &channel)
	if err != nil {
		return nil, fmt.Errorf("failed to create guild channel: %w", err)
	}

	return channel, nil
}

func ModifyGuildChannelPositions(ctx context.Context, s *Session, guildID Snowflake, channelPermissionsArgs []ChannelPermissionsParams, reason *string) error {
	endpoint := EndpointGuildChannels(guildID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err := s.Interface.FetchJJ(ctx, s, http.MethodPatch, endpoint, channelPermissionsArgs, headers, nil)
	if err != nil {
		return fmt.Errorf("failed to modify guild channel positions: %w", err)
	}

	return nil
}

func GetGuildMember(ctx context.Context, s *Session, guildID Snowflake, userID Snowflake) (*GuildMember, error) {
	endpoint := EndpointGuildMember(guildID.String(), userID.String())

	var guildMember *GuildMember

	err := s.Interface.FetchJJ(ctx, s, http.MethodGet, endpoint, nil, nil, &guildMember)
	if err != nil {
		return nil, fmt.Errorf("failed to get guild member: %w", err)
	}

	return guildMember, nil
}

func ListGuildMembers(ctx context.Context, s *Session, guildID Snowflake, limit *int32, after *Snowflake) ([]GuildMember, error) {
	endpoint := EndpointGuildMembers(guildID.String())

	values := url.Values{}

	if limit != nil {
		values.Set("limit", strconv.Itoa(int(*limit)))
	}

	if after != nil {
		values.Set("after", after.String())
	}

	if len(values) > 0 {
		endpoint += "?" + values.Encode()
	}

	var guildMembers []GuildMember

	err := s.Interface.FetchJJ(ctx, s, http.MethodGet, endpoint, nil, nil, &guildMembers)
	if err != nil {
		return nil, fmt.Errorf("failed to list guild members: %w", err)
	}

	return guildMembers, nil
}

func SearchGuildMembers(ctx context.Context, s *Session, guildID Snowflake, query string, limit *int32) ([]GuildMember, error) {
	endpoint := EndpointGuildMembersSearch(guildID.String())

	values := url.Values{}

	values.Set("query", query)

	if limit != nil {
		values.Set("limit", strconv.Itoa(int(*limit)))
	}

	endpoint += "?" + values.Encode()

	var guildMembers []GuildMember

	err := s.Interface.FetchJJ(ctx, s, http.MethodGet, endpoint, nil, nil, &guildMembers)
	if err != nil {
		return nil, fmt.Errorf("failed to search guild members: %w", err)
	}

	return guildMembers, nil
}

func ModifyGuildMember(ctx context.Context, s *Session, guildID Snowflake, userID Snowflake, guildMemberParams GuildMemberParams, reason *string) (*GuildMember, error) {
	endpoint := EndpointGuildMember(guildID.String(), userID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	var guildMember *GuildMember

	err := s.Interface.FetchJJ(ctx, s, http.MethodPatch, endpoint, guildMemberParams, headers, &guildMember)
	if err != nil {
		return nil, fmt.Errorf("failed to modify guild member: %w", err)
	}

	return guildMember, nil
}

func ModifyCurrentMember(ctx context.Context, s *Session, guildID Snowflake, guildMemberArg GuildMember, reason *string) (*GuildMember, error) {
	endpoint := EndpointGuildMember(guildID.String(), "@me")

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	var guildMember *GuildMember

	err := s.Interface.FetchJJ(ctx, s, http.MethodPatch, endpoint, guildMemberArg, headers, &guildMember)
	if err != nil {
		return nil, fmt.Errorf("failed to modify current member: %w", err)
	}

	return guildMember, nil
}

func AddGuildMemberRole(ctx context.Context, s *Session, guildID Snowflake, userID Snowflake, roleID Snowflake, reason *string) error {
	endpoint := EndpointGuildMemberRole(guildID.String(), userID.String(), roleID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err := s.Interface.FetchJJ(ctx, s, http.MethodPut, endpoint, nil, headers, nil)
	if err != nil {
		return fmt.Errorf("failed to add guild member role: %w", err)
	}

	return nil
}

func RemoveGuildMemberRole(ctx context.Context, s *Session, guildID Snowflake, userID Snowflake, roleID Snowflake, reason *string) error {
	endpoint := EndpointGuildMemberRole(guildID.String(), userID.String(), roleID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err := s.Interface.FetchJJ(ctx, s, http.MethodDelete, endpoint, nil, headers, nil)
	if err != nil {
		return fmt.Errorf("failed to remove guild member role: %w", err)
	}

	return nil
}

func RemoveGuildMember(ctx context.Context, s *Session, guildID Snowflake, userID Snowflake, reason *string) error {
	endpoint := EndpointGuildMember(guildID.String(), userID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err := s.Interface.FetchJJ(ctx, s, http.MethodDelete, endpoint, nil, headers, nil)
	if err != nil {
		return fmt.Errorf("failed to remove guild member: %w", err)
	}

	return nil
}

func GetGuildBans(ctx context.Context, s *Session, guildID Snowflake) ([]GuildBan, error) {
	endpoint := EndpointGuildBans(guildID.String())

	var bans []GuildBan

	err := s.Interface.FetchJJ(ctx, s, http.MethodGet, endpoint, nil, nil, &bans)
	if err != nil {
		return nil, fmt.Errorf("failed to get guild bans: %w", err)
	}

	return bans, nil
}

func GetGuildBan(ctx context.Context, s *Session, guildID Snowflake, userID Snowflake) (*GuildBan, error) {
	endpoint := EndpointGuildBan(guildID.String(), userID.String())

	var ban *GuildBan

	err := s.Interface.FetchJJ(ctx, s, http.MethodGet, endpoint, nil, nil, &ban)
	if err != nil {
		return nil, fmt.Errorf("failed to get guild ban: %w", err)
	}

	return ban, nil
}

func CreateGuildBan(ctx context.Context, s *Session, guildID Snowflake, userID Snowflake, reason *string) error {
	endpoint := EndpointGuildBan(guildID.String(), userID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err := s.Interface.FetchJJ(ctx, s, http.MethodPut, endpoint, nil, headers, nil)
	if err != nil {
		return fmt.Errorf("failed to create guild ban: %w", err)
	}

	return nil
}

func RemoveGuildBan(ctx context.Context, s *Session, guildID Snowflake, userID Snowflake, reason *string) error {
	endpoint := EndpointGuildBan(guildID.String(), userID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err := s.Interface.FetchJJ(ctx, s, http.MethodPut, endpoint, nil, headers, nil)
	if err != nil {
		return fmt.Errorf("failed to create guild ban: %w", err)
	}

	return nil
}

func GetGuildRoles(ctx context.Context, s *Session, guildID Snowflake) ([]Role, error) {
	endpoint := EndpointGuildRoles(guildID.String())

	var roles []Role

	err := s.Interface.FetchJJ(ctx, s, http.MethodGet, endpoint, nil, nil, &roles)
	if err != nil {
		return nil, fmt.Errorf("failed to get guild roles: %w", err)
	}

	return roles, nil
}

func CreateGuildRole(ctx context.Context, s *Session, guildID Snowflake, roleParams RoleParams, reason *string) (*Role, error) {
	endpoint := EndpointGuildRoles(guildID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	var role *Role

	err := s.Interface.FetchJJ(ctx, s, http.MethodPost, endpoint, roleParams, headers, &role)
	if err != nil {
		return nil, fmt.Errorf("failed to create guild role: %w", err)
	}

	return role, nil
}

func ModifyGuildRolePositions(ctx context.Context, s *Session, guildID Snowflake, guildRolePositionArgs []ModifyGuildRolePosition, reason *string) ([]Role, error) {
	endpoint := EndpointGuildRoles(guildID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	var roles []Role

	err := s.Interface.FetchJJ(ctx, s, http.MethodPatch, endpoint, guildRolePositionArgs, headers, &roles)
	if err != nil {
		return nil, fmt.Errorf("failed to modify guild role positions: %w", err)
	}

	return roles, nil
}

func ModifyGuildRole(ctx context.Context, s *Session, guildID Snowflake, roleID Snowflake, roleArg Role, reason *string) (*Role, error) {
	endpoint := EndpointGuildRole(guildID.String(), roleID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	var role *Role

	err := s.Interface.FetchJJ(ctx, s, http.MethodPatch, endpoint, roleArg, headers, &role)
	if err != nil {
		return nil, fmt.Errorf("failed to modify guild role: %w", err)
	}

	return role, nil
}

func DeleteGuildRole(ctx context.Context, s *Session, guildID Snowflake, roleID Snowflake, reason *string) error {
	endpoint := EndpointGuildRole(guildID.String(), roleID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err := s.Interface.FetchJJ(ctx, s, http.MethodDelete, endpoint, nil, headers, nil)
	if err != nil {
		return fmt.Errorf("failed to delete guild role: %w", err)
	}

	return nil
}

func GetGuildPruneCount(ctx context.Context, s *Session, guildID Snowflake, days *int32, includedRoles []Snowflake) (*int32, error) {
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

	values := url.Values{}

	if days != nil {
		values.Set("days", strconv.Itoa(int(*days)))
	}

	if len(rolesString) > 0 {
		values.Set("include_roles", rolesString)
	}

	if len(values) > 0 {
		endpoint += "?" + values.Encode()
	}

	var prunedStruct struct {
		Pruned int32 `json:"pruned"`
	}

	err := s.Interface.FetchJJ(ctx, s, http.MethodGet, endpoint, nil, nil, &prunedStruct)
	if err != nil {
		return nil, fmt.Errorf("failed to get guild prune count: %w", err)
	}

	return &prunedStruct.Pruned, nil
}

func BeginGuildPrune(ctx context.Context, s *Session, guildID Snowflake, days *int32, includedRoles []Snowflake, computePruneCount bool, reason *string) (*int32, error) {
	endpoint := EndpointGuildPrune(guildID.String())

	pruneArg := GuildPruneParam{
		ComputePruneCount: computePruneCount,
	}

	if days != nil {
		pruneArg.Days = days
	}

	for _, includedRole := range includedRoles {
		role := includedRole
		pruneArg.IncludeRoles = append(pruneArg.IncludeRoles, role)
	}

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	prunedStruct := struct {
		Pruned int32 `json:"pruned"`
	}{}

	err := s.Interface.FetchJJ(ctx, s, http.MethodPost, endpoint, pruneArg, headers, &prunedStruct)
	if err != nil {
		return nil, fmt.Errorf("failed to begin guild prune: %w", err)
	}

	return &prunedStruct.Pruned, nil
}

func GetGuildInvites(ctx context.Context, s *Session, guildID Snowflake) ([]Invite, error) {
	endpoint := EndpointGuildInvites(guildID.String())

	var invites []Invite

	err := s.Interface.FetchJJ(ctx, s, http.MethodGet, endpoint, nil, nil, &invites)
	if err != nil {
		return nil, fmt.Errorf("failed to get guild invites: %w", err)
	}

	return invites, nil
}

func GetGuildIntegrations(ctx context.Context, s *Session, guildID Snowflake) ([]Integration, error) {
	endpoint := EndpointGuildIntegrations(guildID.String())

	var integrations []Integration

	err := s.Interface.FetchJJ(ctx, s, http.MethodGet, endpoint, nil, nil, &integrations)
	if err != nil {
		return nil, fmt.Errorf("failed to get guild integrations: %w", err)
	}

	return integrations, nil
}

func DeleteGuildIntegration(ctx context.Context, s *Session, guildID Snowflake, integrationID Snowflake, reason *string) error {
	endpoint := EndpointGuildIntegration(guildID.String(), integrationID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err := s.Interface.FetchJJ(ctx, s, http.MethodDelete, endpoint, nil, headers, nil)
	if err != nil {
		return fmt.Errorf("failed to delete guild integration: %w", err)
	}

	return nil
}

func GetGuildVanityURL(ctx context.Context, s *Session, guildID Snowflake) (*Invite, error) {
	endpoint := EndpointGuildVanityURL(guildID.String())

	var invite *Invite

	err := s.Interface.FetchJJ(ctx, s, http.MethodGet, endpoint, nil, nil, &invite)
	if err != nil {
		return nil, fmt.Errorf("failed to get guild vanity url: %w", err)
	}

	return invite, nil
}

// TODO: GetGuildWidgetImage
// TODO: GetGuildWelcomeScreen
// TODO: ModifyGuildWelcomeScreen
// TODO: GetGuildWidget
// TODO: ModifyGuildWidget
// TODO: GetGuildWidgetSettings
// TODO: AddGuildMember
// TODO: GetGuildVoiceRegions

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
