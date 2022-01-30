package http

import (
	"context"
	"net/http"
	"net/url"

	"github.com/WelcomerTeam/Discord/discord"
	"github.com/WelcomerTeam/Discord/structs"
	"golang.org/x/xerrors"
)

func (s *Session) GetGuildAuditLog(ctx context.Context, guildID discord.Snowflake, userID *discord.Snowflake, actionType *structs.AuditLogActionType, before *discord.Snowflake, limit *int32) (entries []*structs.AuditLogEntry, err error) {
	endpoint := EndpointGuildAuditLogs(guildID.String())

	values := url.Values{}

	if userID != nil {
		values.Set("user_id", userID.String())
	}

	if actionType != nil {
		values.Set("action_type", string(*actionType))
	}

	if before != nil {
		values.Set("before", before.String())
	}

	if limit != nil {
		values.Set("limit", string(*limit))
	}

	if len(values) > 0 {
		endpoint += "?" + values.Encode()
	}

	err = s.Interface.FetchJJ(ctx, http.MethodGet, endpoint, nil, nil, &entries)
	if err != nil {
		return nil, xerrors.Errorf("Failed to get guild audit log: %v", err)
	}

	return entries, nil
}
