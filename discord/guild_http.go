package discord

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

func CreateGuild(ctx context.Context, session *Session, guildArg Guild) (*Guild, error) {
	endpoint := EndpointGuilds

	var guild *Guild

	err := session.Interface.FetchJJ(ctx, session, http.MethodPost, endpoint, guildArg, nil, &guild)
	if err != nil {
		return nil, fmt.Errorf("failed to create guild: %w", err)
	}

	return guild, nil
}

func GetGuild(ctx context.Context, session *Session, guildID Snowflake) (*Guild, error) {
	endpoint := EndpointGuild(guildID.String())

	var guild *Guild

	err := session.Interface.FetchJJ(ctx, session, http.MethodGet, endpoint, nil, nil, &guild)
	if err != nil {
		return nil, fmt.Errorf("failed to get guild: %w", err)
	}

	return guild, nil
}

func GetGuildPreview(ctx context.Context, session *Session, guildID Snowflake) (*Guild, error) {
	endpoint := EndpointGuildPreview(guildID.String())

	var guildPreview *Guild

	err := session.Interface.FetchJJ(ctx, session, http.MethodGet, endpoint, nil, nil, &guildPreview)
	if err != nil {
		return nil, fmt.Errorf("failed to get guild preview: %w", err)
	}

	return guildPreview, nil
}

func ModifyGuild(ctx context.Context, session *Session, guildID Snowflake, guildArg GuildParam, reason *string) (*Guild, error) {
	endpoint := EndpointGuild(guildID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	var guild *Guild

	err := session.Interface.FetchJJ(ctx, session, http.MethodPatch, endpoint, guildArg, headers, &guild)
	if err != nil {
		return nil, fmt.Errorf("failed to modify guild: %w", err)
	}

	return guild, nil
}

func DeleteGuild(ctx context.Context, session *Session, guildID Snowflake) error {
	endpoint := EndpointGuild(guildID.String())

	err := session.Interface.FetchJJ(ctx, session, http.MethodDelete, endpoint, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete guild: %w", err)
	}

	return nil
}

func GetGuildChannels(ctx context.Context, session *Session, guildID Snowflake) ([]Channel, error) {
	endpoint := EndpointGuildChannels(guildID.String())

	var channels []Channel

	err := session.Interface.FetchJJ(ctx, session, http.MethodGet, endpoint, nil, nil, &channels)
	if err != nil {
		return nil, fmt.Errorf("failed to get guild channels: %w", err)
	}

	return channels, nil
}

func CreateGuildChannel(ctx context.Context, session *Session, guildID Snowflake, channelParams ChannelParams, reason *string) (*Channel, error) {
	endpoint := EndpointGuildChannels(guildID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	var channel *Channel

	err := session.Interface.FetchJJ(ctx, session, http.MethodPost, endpoint, channelParams, headers, &channel)
	if err != nil {
		return nil, fmt.Errorf("failed to create guild channel: %w", err)
	}

	return channel, nil
}

func ModifyGuildChannelPositions(ctx context.Context, session *Session, guildID Snowflake, channelPermissionsArgs []ChannelPermissionsParams, reason *string) error {
	endpoint := EndpointGuildChannels(guildID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err := session.Interface.FetchJJ(ctx, session, http.MethodPatch, endpoint, channelPermissionsArgs, headers, nil)
	if err != nil {
		return fmt.Errorf("failed to modify guild channel positions: %w", err)
	}

	return nil
}

func GetGuildMember(ctx context.Context, session *Session, guildID, userID Snowflake) (*GuildMember, error) {
	endpoint := EndpointGuildMember(guildID.String(), userID.String())

	var guildMember *GuildMember

	err := session.Interface.FetchJJ(ctx, session, http.MethodGet, endpoint, nil, nil, &guildMember)
	if err != nil {
		return nil, fmt.Errorf("failed to get guild member: %w", err)
	}

	return guildMember, nil
}

func ListGuildMembers(ctx context.Context, session *Session, guildID Snowflake, limit *int32, after *Snowflake) ([]GuildMember, error) {
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

	err := session.Interface.FetchJJ(ctx, session, http.MethodGet, endpoint, nil, nil, &guildMembers)
	if err != nil {
		return nil, fmt.Errorf("failed to list guild members: %w", err)
	}

	return guildMembers, nil
}

func SearchGuildMembers(ctx context.Context, session *Session, guildID Snowflake, query string, limit *int32) ([]GuildMember, error) {
	endpoint := EndpointGuildMembersSearch(guildID.String())

	values := url.Values{}

	values.Set("query", query)

	if limit != nil {
		values.Set("limit", strconv.Itoa(int(*limit)))
	}

	endpoint += "?" + values.Encode()

	var guildMembers []GuildMember

	err := session.Interface.FetchJJ(ctx, session, http.MethodGet, endpoint, nil, nil, &guildMembers)
	if err != nil {
		return nil, fmt.Errorf("failed to search guild members: %w", err)
	}

	return guildMembers, nil
}

func ModifyGuildMember(ctx context.Context, session *Session, guildID, userID Snowflake, guildMemberParams GuildMemberParams, reason *string) (*GuildMember, error) {
	endpoint := EndpointGuildMember(guildID.String(), userID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	var guildMember *GuildMember

	err := session.Interface.FetchJJ(ctx, session, http.MethodPatch, endpoint, guildMemberParams, headers, &guildMember)
	if err != nil {
		return nil, fmt.Errorf("failed to modify guild member: %w", err)
	}

	return guildMember, nil
}

func ModifyCurrentMember(ctx context.Context, session *Session, guildID Snowflake, params ModifyCurrentMemberParams, reason *string) (*GuildMember, error) {
	endpoint := EndpointGuildMember(guildID.String(), "@me")

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	var guildMember *GuildMember

	err := session.Interface.FetchJJ(ctx, session, http.MethodPatch, endpoint, params, headers, &guildMember)
	if err != nil {
		return nil, fmt.Errorf("failed to modify current member: %w", err)
	}

	return guildMember, nil
}

func AddGuildMemberRole(ctx context.Context, session *Session, guildID, userID, roleID Snowflake, reason *string) error {
	endpoint := EndpointGuildMemberRole(guildID.String(), userID.String(), roleID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err := session.Interface.FetchJJ(ctx, session, http.MethodPut, endpoint, nil, headers, nil)
	if err != nil {
		return fmt.Errorf("failed to add guild member role: %w", err)
	}

	return nil
}

func RemoveGuildMemberRole(ctx context.Context, session *Session, guildID, userID, roleID Snowflake, reason *string) error {
	endpoint := EndpointGuildMemberRole(guildID.String(), userID.String(), roleID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err := session.Interface.FetchJJ(ctx, session, http.MethodDelete, endpoint, nil, headers, nil)
	if err != nil {
		return fmt.Errorf("failed to remove guild member role: %w", err)
	}

	return nil
}

func RemoveGuildMember(ctx context.Context, session *Session, guildID, userID Snowflake, reason *string) error {
	endpoint := EndpointGuildMember(guildID.String(), userID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err := session.Interface.FetchJJ(ctx, session, http.MethodDelete, endpoint, nil, headers, nil)
	if err != nil {
		return fmt.Errorf("failed to remove guild member: %w", err)
	}

	return nil
}

func GetGuildBans(ctx context.Context, session *Session, guildID Snowflake) ([]GuildBan, error) {
	endpoint := EndpointGuildBans(guildID.String())

	var bans []GuildBan

	err := session.Interface.FetchJJ(ctx, session, http.MethodGet, endpoint, nil, nil, &bans)
	if err != nil {
		return nil, fmt.Errorf("failed to get guild bans: %w", err)
	}

	return bans, nil
}

func GetGuildBan(ctx context.Context, session *Session, guildID, userID Snowflake) (*GuildBan, error) {
	endpoint := EndpointGuildBan(guildID.String(), userID.String())

	var ban *GuildBan

	err := session.Interface.FetchJJ(ctx, session, http.MethodGet, endpoint, nil, nil, &ban)
	if err != nil {
		return nil, fmt.Errorf("failed to get guild ban: %w", err)
	}

	return ban, nil
}

func CreateGuildBan(ctx context.Context, session *Session, guildID, userID Snowflake, reason *string) error {
	endpoint := EndpointGuildBan(guildID.String(), userID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err := session.Interface.FetchJJ(ctx, session, http.MethodPut, endpoint, nil, headers, nil)
	if err != nil {
		return fmt.Errorf("failed to create guild ban: %w", err)
	}

	return nil
}

func RemoveGuildBan(ctx context.Context, session *Session, guildID, userID Snowflake, reason *string) error {
	endpoint := EndpointGuildBan(guildID.String(), userID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err := session.Interface.FetchJJ(ctx, session, http.MethodPut, endpoint, nil, headers, nil)
	if err != nil {
		return fmt.Errorf("failed to create guild ban: %w", err)
	}

	return nil
}

// BulkBanUserParams represents a single user to ban in bulk ban operation.
type BulkBanUserParams struct {
	UserID            Snowflake `json:"user_id"`
	DeleteMessageDays int32     `json:"delete_message_days,omitempty"`
}

// BulkBanUsersParams represents the parameters for bulk banning users.
type BulkBanUsersParams struct {
	UserIDs           []Snowflake `json:"user_ids"`
	DeleteMessageDays int32       `json:"delete_message_days,omitempty"`
	AuditLogReason    string      `json:"-"`
}

// BulkBanUsersResponse represents the response from bulk banning users.
type BulkBanUsersResponse struct {
	BannedUserIDs []Snowflake `json:"banned_user_ids"`
}

// BulkBanUsers bans multiple users from a guild.
func BulkBanUsers(ctx context.Context, session *Session, guildID Snowflake, params BulkBanUsersParams, reason *string) (*BulkBanUsersResponse, error) {
	endpoint := "/guilds/" + guildID.String() + "/bulk-ban"

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	var response *BulkBanUsersResponse

	err := session.Interface.FetchJJ(ctx, session, http.MethodPost, endpoint, params, headers, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to bulk ban users: %w", err)
	}

	return response, nil
}

func GetGuildRoles(ctx context.Context, session *Session, guildID Snowflake) ([]Role, error) {
	endpoint := EndpointGuildRoles(guildID.String())

	var roles []Role

	err := session.Interface.FetchJJ(ctx, session, http.MethodGet, endpoint, nil, nil, &roles)
	if err != nil {
		return nil, fmt.Errorf("failed to get guild roles: %w", err)
	}

	return roles, nil
}

func CreateGuildRole(ctx context.Context, session *Session, guildID Snowflake, roleParams RoleParams, reason *string) (*Role, error) {
	endpoint := EndpointGuildRoles(guildID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	var role *Role

	err := session.Interface.FetchJJ(ctx, session, http.MethodPost, endpoint, roleParams, headers, &role)
	if err != nil {
		return nil, fmt.Errorf("failed to create guild role: %w", err)
	}

	return role, nil
}

func ModifyGuildRolePositions(ctx context.Context, session *Session, guildID Snowflake, guildRolePositionArgs []ModifyGuildRolePosition, reason *string) ([]Role, error) {
	endpoint := EndpointGuildRoles(guildID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	var roles []Role

	err := session.Interface.FetchJJ(ctx, session, http.MethodPatch, endpoint, guildRolePositionArgs, headers, &roles)
	if err != nil {
		return nil, fmt.Errorf("failed to modify guild role positions: %w", err)
	}

	return roles, nil
}

func ModifyGuildRole(ctx context.Context, session *Session, guildID, roleID Snowflake, roleArg Role, reason *string) (*Role, error) {
	endpoint := EndpointGuildRole(guildID.String(), roleID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	var role *Role

	err := session.Interface.FetchJJ(ctx, session, http.MethodPatch, endpoint, roleArg, headers, &role)
	if err != nil {
		return nil, fmt.Errorf("failed to modify guild role: %w", err)
	}

	return role, nil
}

func DeleteGuildRole(ctx context.Context, session *Session, guildID, roleID Snowflake, reason *string) error {
	endpoint := EndpointGuildRole(guildID.String(), roleID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err := session.Interface.FetchJJ(ctx, session, http.MethodDelete, endpoint, nil, headers, nil)
	if err != nil {
		return fmt.Errorf("failed to delete guild role: %w", err)
	}

	return nil
}

func GetGuildPruneCount(ctx context.Context, session *Session, guildID Snowflake, days *int32, includedRoles []Snowflake) (*int32, error) {
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

	err := session.Interface.FetchJJ(ctx, session, http.MethodGet, endpoint, nil, nil, &prunedStruct)
	if err != nil {
		return nil, fmt.Errorf("failed to get guild prune count: %w", err)
	}

	return &prunedStruct.Pruned, nil
}

func BeginGuildPrune(ctx context.Context, session *Session, guildID Snowflake, days *int32, includedRoles []Snowflake, computePruneCount bool, reason *string) (*int32, error) {
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

	err := session.Interface.FetchJJ(ctx, session, http.MethodPost, endpoint, pruneArg, headers, &prunedStruct)
	if err != nil {
		return nil, fmt.Errorf("failed to begin guild prune: %w", err)
	}

	return &prunedStruct.Pruned, nil
}

func GetGuildInvites(ctx context.Context, session *Session, guildID Snowflake) ([]Invite, error) {
	endpoint := EndpointGuildInvites(guildID.String())

	var invites []Invite

	err := session.Interface.FetchJJ(ctx, session, http.MethodGet, endpoint, nil, nil, &invites)
	if err != nil {
		return nil, fmt.Errorf("failed to get guild invites: %w", err)
	}

	return invites, nil
}

func GetGuildIntegrations(ctx context.Context, session *Session, guildID Snowflake) ([]Integration, error) {
	endpoint := EndpointGuildIntegrations(guildID.String())

	var integrations []Integration

	err := session.Interface.FetchJJ(ctx, session, http.MethodGet, endpoint, nil, nil, &integrations)
	if err != nil {
		return nil, fmt.Errorf("failed to get guild integrations: %w", err)
	}

	return integrations, nil
}

func DeleteGuildIntegration(ctx context.Context, session *Session, guildID, integrationID Snowflake, reason *string) error {
	endpoint := EndpointGuildIntegration(guildID.String(), integrationID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err := session.Interface.FetchJJ(ctx, session, http.MethodDelete, endpoint, nil, headers, nil)
	if err != nil {
		return fmt.Errorf("failed to delete guild integration: %w", err)
	}

	return nil
}

func GetGuildVanityURL(ctx context.Context, session *Session, guildID Snowflake) (*Invite, error) {
	endpoint := EndpointGuildVanityURL(guildID.String())

	var invite *Invite

	err := session.Interface.FetchJJ(ctx, session, http.MethodGet, endpoint, nil, nil, &invite)
	if err != nil {
		return nil, fmt.Errorf("failed to get guild vanity url: %w", err)
	}

	return invite, nil
}

// GetGuildWidgetImage gets the guild widget image.
func GetGuildWidgetImage(ctx context.Context, session *Session, guildID Snowflake, style *string) ([]byte, error) {
	endpoint := EndpointGuildWidgetImage(guildID.String())

	params := url.Values{}
	if style != nil {
		params.Add("style", *style)
	}

	if len(params) > 0 {
		endpoint += "?" + params.Encode()
	}

	var imageData []byte

	err := session.Interface.FetchBJ(ctx, session, http.MethodGet, endpoint, "image/png", nil, nil, &imageData)
	if err != nil {
		return nil, fmt.Errorf("failed to get guild widget image: %w", err)
	}

	return imageData, nil
}

// GetGuildWelcomeScreen gets the guild's welcome screen.
func GetGuildWelcomeScreen(ctx context.Context, session *Session, guildID Snowflake) (interface{}, error) {
	endpoint := EndpointGuildWelcomeScreen(guildID.String())

	var welcomeScreen interface{}

	err := session.Interface.FetchJJ(ctx, session, http.MethodGet, endpoint, nil, nil, &welcomeScreen)
	if err != nil {
		return nil, fmt.Errorf("failed to get guild welcome screen: %w", err)
	}

	return welcomeScreen, nil
}

// ModifyGuildWelcomeScreen modifies the guild's welcome screen.
func ModifyGuildWelcomeScreen(ctx context.Context, session *Session, guildID Snowflake, params interface{}, reason *string) (interface{}, error) {
	endpoint := EndpointGuildWelcomeScreen(guildID.String())

	headers := http.Header{}
	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	var welcomeScreen interface{}

	err := session.Interface.FetchJJ(ctx, session, http.MethodPatch, endpoint, params, headers, &welcomeScreen)
	if err != nil {
		return nil, fmt.Errorf("failed to modify guild welcome screen: %w", err)
	}

	return welcomeScreen, nil
}

// GetGuildWidget gets the guild widget.
func GetGuildWidget(ctx context.Context, session *Session, guildID Snowflake) (interface{}, error) {
	endpoint := EndpointGuildWidgetJSON(guildID.String())

	var widget interface{}

	err := session.Interface.FetchJJ(ctx, session, http.MethodGet, endpoint, nil, nil, &widget)
	if err != nil {
		return nil, fmt.Errorf("failed to get guild widget: %w", err)
	}

	return widget, nil
}

// ModifyGuildWidget modifies the guild widget.
func ModifyGuildWidget(ctx context.Context, session *Session, guildID Snowflake, params interface{}, reason *string) (interface{}, error) {
	endpoint := EndpointGuildWidget(guildID.String())

	headers := http.Header{}
	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	var widget interface{}

	err := session.Interface.FetchJJ(ctx, session, http.MethodPatch, endpoint, params, headers, &widget)
	if err != nil {
		return nil, fmt.Errorf("failed to modify guild widget: %w", err)
	}

	return widget, nil
}

// GetGuildWidgetSettings gets the guild widget settings.
func GetGuildWidgetSettings(ctx context.Context, session *Session, guildID Snowflake) (interface{}, error) {
	endpoint := EndpointGuildWidget(guildID.String())

	var settings interface{}

	err := session.Interface.FetchJJ(ctx, session, http.MethodGet, endpoint, nil, nil, &settings)
	if err != nil {
		return nil, fmt.Errorf("failed to get guild widget settings: %w", err)
	}

	return settings, nil
}

// AddGuildMember adds a guild member using an OAuth2 access token.
func AddGuildMember(ctx context.Context, session *Session, guildID, userID Snowflake, accessToken string, nick *string, roles []Snowflake, mute, deaf *bool) (interface{}, error) {
	endpoint := EndpointGuildMember(guildID.String(), userID.String())

	params := struct {
		AccessToken string      `json:"access_token"`
		Nick        *string     `json:"nick,omitempty"`
		Roles       []Snowflake `json:"roles,omitempty"`
		Mute        *bool       `json:"mute,omitempty"`
		Deaf        *bool       `json:"deaf,omitempty"`
	}{
		AccessToken: accessToken,
		Nick:        nick,
		Roles:       roles,
		Mute:        mute,
		Deaf:        deaf,
	}

	var member interface{}

	err := session.Interface.FetchJJ(ctx, session, http.MethodPut, endpoint, params, nil, &member)
	if err != nil {
		return nil, fmt.Errorf("failed to add guild member: %w", err)
	}

	return member, nil
}

// ListVoiceRegions lists voice regions available to guilds.
func ListVoiceRegions(ctx context.Context, session *Session) ([]interface{}, error) {
	endpoint := "/voice/regions"

	var regions []interface{}

	err := session.Interface.FetchJJ(ctx, session, http.MethodGet, endpoint, nil, nil, &regions)
	if err != nil {
		return nil, fmt.Errorf("failed to list voice regions: %w", err)
	}

	return regions, nil
}

// GetGuildVoiceRegions lists voice regions available to a specific guild.
func GetGuildVoiceRegions(ctx context.Context, session *Session, guildID Snowflake) ([]interface{}, error) {
	endpoint := EndpointGuildVoiceRegions(guildID.String())

	var regions []interface{}

	err := session.Interface.FetchJJ(ctx, session, http.MethodGet, endpoint, nil, nil, &regions)
	if err != nil {
		return nil, fmt.Errorf("failed to get guild voice regions: %w", err)
	}

	return regions, nil
}

// ModifyCurrentUserVoiceState modifies the voice state of the current user.
func ModifyCurrentUserVoiceState(ctx context.Context, session *Session, guildID Snowflake, channelID *Snowflake, suppress *bool) error {
	endpoint := EndpointGuildVoiceStateSelf(guildID.String())

	params := struct {
		ChannelID *Snowflake `json:"channel_id,omitempty"`
		Suppress  *bool      `json:"suppress,omitempty"`
	}{
		ChannelID: channelID,
		Suppress:  suppress,
	}

	err := session.Interface.FetchJJ(ctx, session, http.MethodPatch, endpoint, params, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to modify current user voice state: %w", err)
	}

	return nil
}

// ModifyUserVoiceState modifies the voice state of a user.
func ModifyUserVoiceState(ctx context.Context, session *Session, guildID, userID Snowflake, channelID *Snowflake, suppress *bool) error {
	endpoint := EndpointGuildVoiceState(guildID.String(), userID.String())

	params := struct {
		ChannelID *Snowflake `json:"channel_id,omitempty"`
		Suppress  *bool      `json:"suppress,omitempty"`
	}{
		ChannelID: channelID,
		Suppress:  suppress,
	}

	err := session.Interface.FetchJJ(ctx, session, http.MethodPatch, endpoint, params, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to modify user voice state: %w", err)
	}

	return nil
}

// ListScheduledEventsForGuild lists scheduled events for a guild.
func ListScheduledEventsForGuild(ctx context.Context, session *Session, guildID Snowflake, withUserCount *bool) ([]interface{}, error) {
	endpoint := EndpointGuildScheduledEvents(guildID.String())

	params := url.Values{}
	if withUserCount != nil {
		params.Add("with_user_count", strconv.FormatBool(*withUserCount))
	}

	if len(params) > 0 {
		endpoint += "?" + params.Encode()
	}

	var events []interface{}

	err := session.Interface.FetchJJ(ctx, session, http.MethodGet, endpoint, nil, nil, &events)
	if err != nil {
		return nil, fmt.Errorf("failed to list scheduled events for guild: %w", err)
	}

	return events, nil
}

// CreateGuildScheduledEvent creates a scheduled event for a guild.
func CreateGuildScheduledEvent(ctx context.Context, session *Session, guildID Snowflake, params interface{}, reason *string) (interface{}, error) {
	endpoint := EndpointGuildScheduledEvents(guildID.String())

	headers := http.Header{}
	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	var event interface{}

	err := session.Interface.FetchJJ(ctx, session, http.MethodPost, endpoint, params, headers, &event)
	if err != nil {
		return nil, fmt.Errorf("failed to create guild scheduled event: %w", err)
	}

	return event, nil
}

// GetGuildScheduledEvent gets a scheduled event for a guild.
func GetGuildScheduledEvent(ctx context.Context, session *Session, guildID, eventID Snowflake, withUserCount *bool) (interface{}, error) {
	endpoint := EndpointGuildScheduledEvent(guildID.String(), eventID.String())

	params := url.Values{}
	if withUserCount != nil {
		params.Add("with_user_count", strconv.FormatBool(*withUserCount))
	}

	if len(params) > 0 {
		endpoint += "?" + params.Encode()
	}

	var event interface{}

	err := session.Interface.FetchJJ(ctx, session, http.MethodGet, endpoint, nil, nil, &event)
	if err != nil {
		return nil, fmt.Errorf("failed to get guild scheduled event: %w", err)
	}

	return event, nil
}

// ModifyGuildScheduledEvent modifies a scheduled event for a guild.
func ModifyGuildScheduledEvent(ctx context.Context, session *Session, guildID, eventID Snowflake, params interface{}, reason *string) (interface{}, error) {
	endpoint := EndpointGuildScheduledEvent(guildID.String(), eventID.String())

	headers := http.Header{}
	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	var event interface{}

	err := session.Interface.FetchJJ(ctx, session, http.MethodPatch, endpoint, params, headers, &event)
	if err != nil {
		return nil, fmt.Errorf("failed to modify guild scheduled event: %w", err)
	}

	return event, nil
}

// DeleteGuildScheduledEvent deletes a scheduled event for a guild.
func DeleteGuildScheduledEvent(ctx context.Context, session *Session, guildID, eventID Snowflake) error {
	endpoint := EndpointGuildScheduledEvent(guildID.String(), eventID.String())

	err := session.Interface.FetchJJ(ctx, session, http.MethodDelete, endpoint, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete guild scheduled event: %w", err)
	}

	return nil
}

// GetGuildScheduledEventUsers gets users who have responded to a scheduled event.
func GetGuildScheduledEventUsers(ctx context.Context, session *Session, guildID, eventID Snowflake, limit *int, before, after *Snowflake) ([]interface{}, error) {
	endpoint := EndpointGuildScheduledEventUsers(guildID.String(), eventID.String())

	params := url.Values{}
	if limit != nil {
		params.Add("limit", strconv.Itoa(*limit))
	}

	if before != nil {
		params.Add("before", before.String())
	}

	if after != nil {
		params.Add("after", after.String())
	}

	if len(params) > 0 {
		endpoint += "?" + params.Encode()
	}

	var users []interface{}

	err := session.Interface.FetchJJ(ctx, session, http.MethodGet, endpoint, nil, nil, &users)
	if err != nil {
		return nil, fmt.Errorf("failed to get guild scheduled event users: %w", err)
	}

	return users, nil
}

// GetGuildTemplate gets a guild template.
func GetGuildTemplate(ctx context.Context, session *Session, templateCode string) (interface{}, error) {
	endpoint := "/guilds/templates/" + templateCode

	var template interface{}

	err := session.Interface.FetchJJ(ctx, session, http.MethodGet, endpoint, nil, nil, &template)
	if err != nil {
		return nil, fmt.Errorf("failed to get guild template: %w", err)
	}

	return template, nil
}

// CreateGuildFromTemplate creates a guild from a template.
func CreateGuildFromTemplate(ctx context.Context, session *Session, templateCode string, params interface{}) (interface{}, error) {
	endpoint := "/guilds/templates/" + templateCode

	var guild interface{}

	err := session.Interface.FetchJJ(ctx, session, http.MethodPost, endpoint, params, nil, &guild)
	if err != nil {
		return nil, fmt.Errorf("failed to create guild from template: %w", err)
	}

	return guild, nil
}

// GetGuildTemplates gets the guild templates for a guild.
func GetGuildTemplates(ctx context.Context, session *Session, guildID Snowflake) ([]interface{}, error) {
	endpoint := EndpointGuildTemplates(guildID.String())

	var templates []interface{}

	err := session.Interface.FetchJJ(ctx, session, http.MethodGet, endpoint, nil, nil, &templates)
	if err != nil {
		return nil, fmt.Errorf("failed to get guild templates: %w", err)
	}

	return templates, nil
}

// CreateGuildTemplate creates a template for a guild.
func CreateGuildTemplate(ctx context.Context, session *Session, guildID Snowflake, params interface{}, reason *string) (interface{}, error) {
	endpoint := EndpointGuildTemplates(guildID.String())

	headers := http.Header{}
	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	var template interface{}

	err := session.Interface.FetchJJ(ctx, session, http.MethodPost, endpoint, params, headers, &template)
	if err != nil {
		return nil, fmt.Errorf("failed to create guild template: %w", err)
	}

	return template, nil
}

// SyncGuildTemplate syncs a guild template.
func SyncGuildTemplate(ctx context.Context, session *Session, guildID Snowflake, templateCode string, reason *string) (interface{}, error) {
	endpoint := EndpointGuildTemplate(guildID.String(), templateCode)

	headers := http.Header{}
	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	var template interface{}

	err := session.Interface.FetchJJ(ctx, session, http.MethodPut, endpoint, nil, headers, &template)
	if err != nil {
		return nil, fmt.Errorf("failed to sync guild template: %w", err)
	}

	return template, nil
}

// ModifyGuildTemplate modifies a guild template.
func ModifyGuildTemplate(ctx context.Context, session *Session, guildID Snowflake, templateCode string, params interface{}, reason *string) (interface{}, error) {
	endpoint := EndpointGuildTemplate(guildID.String(), templateCode)

	headers := http.Header{}
	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	var template interface{}

	err := session.Interface.FetchJJ(ctx, session, http.MethodPatch, endpoint, params, headers, &template)
	if err != nil {
		return nil, fmt.Errorf("failed to modify guild template: %w", err)
	}

	return template, nil
}

// DeleteGuildTemplate deletes a guild template.
func DeleteGuildTemplate(ctx context.Context, session *Session, guildID Snowflake, templateCode string, reason *string) error {
	endpoint := EndpointGuildTemplate(guildID.String(), templateCode)

	headers := http.Header{}
	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err := session.Interface.FetchJJ(ctx, session, http.MethodDelete, endpoint, nil, headers, nil)
	if err != nil {
		return fmt.Errorf("failed to delete guild template: %w", err)
	}

	return nil
}
