package discord

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

func GetInvite(s *Session, inviteCode string, withCounts *bool, withExpiration *bool, guildScheduledEventID *Snowflake) (*Invite, error) {
	endpoint := EndpointInvite(inviteCode)

	var values url.Values

	if withCounts != nil {
		values.Set("with_counts", strconv.FormatBool(*withCounts))
	}

	if withExpiration != nil {
		values.Set("with_expiration", strconv.FormatBool(*withExpiration))
	}

	if guildScheduledEventID != nil {
		values.Set("guild_scheduled_event_id", guildScheduledEventID.String())
	}

	if len(values) > 0 {
		endpoint += "?" + values.Encode()
	}

	var invite *Invite

	err := s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &invite)
	if err != nil {
		return nil, fmt.Errorf("failed to get invite: %w", err)
	}

	return invite, nil
}

func DeleteInvite(s *Session, inviteCode string, reason *string) error {
	endpoint := EndpointInvite(inviteCode)

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err := s.Interface.FetchJJ(s, http.MethodDelete, endpoint, nil, headers, nil)
	if err != nil {
		return fmt.Errorf("failed to delete invite: %w", err)
	}

	return nil
}
