package discord

import (
	"net/http"
	"net/url"
	"strconv"

	"golang.org/x/xerrors"
)

func GetInvite(s *Session, inviteCode string, withCounts *bool, withExpiration *bool, guildScheduledEventID *Snowflake) (invite *Invite, err error) {
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

	err = s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &invite)
	if err != nil {
		return nil, xerrors.Errorf("Failed to get invite: %v", err)
	}

	return
}

func DeleteInvite(s *Session, inviteCode string, reason *string) (err error) {
	endpoint := EndpointInvite(inviteCode)

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err = s.Interface.FetchJJ(s, http.MethodDelete, endpoint, nil, headers, nil)
	if err != nil {
		return xerrors.Errorf("Failed to delete invite: %v", err)
	}

	return
}
