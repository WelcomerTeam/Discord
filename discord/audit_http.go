package discord

import (
	"fmt"
	"net/http"
	"net/url"
)

func GetGuildAuditLog(s *Session, guildID Snowflake, userID *Snowflake, actionType *AuditLogActionType, before *Snowflake, limit *int32) ([]*AuditLogEntry, error) {
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

	var entries []*AuditLogEntry

	err := s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &entries)
	if err != nil {
		return nil, fmt.Errorf("failed to get guild audit log: %v", err)
	}

	return entries, nil
}
