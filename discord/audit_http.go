package discord

import (
	"golang.org/x/xerrors"
	"net/http"
	"net/url"
)

func GetGuildAuditLog(s *Session, guildID Snowflake, userID *Snowflake, actionType *AuditLogActionType, before *Snowflake, limit *int32) (entries []*AuditLogEntry, err error) {
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

	err = s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &entries)
	if err != nil {
		return nil, xerrors.Errorf("Failed to get guild audit log: %v", err)
	}

	return entries, nil
}
