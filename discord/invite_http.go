package discord

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

func GetInvite(ctx context.Context, session *Session, inviteCode string, withCounts, withExpiration *bool, guildScheduledEventID *Snowflake) (*Invite, error) {
	endpoint := EndpointInvite(inviteCode)

	values := url.Values{}

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

	err := session.Interface.FetchJJ(ctx, session, http.MethodGet, endpoint, nil, nil, &invite)
	if err != nil {
		return nil, fmt.Errorf("failed to get invite: %w", err)
	}

	return invite, nil
}

func DeleteInvite(ctx context.Context, session *Session, inviteCode string, reason *string) error {
	endpoint := EndpointInvite(inviteCode)

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err := session.Interface.FetchJJ(ctx, session, http.MethodDelete, endpoint, nil, headers, nil)
	if err != nil {
		return fmt.Errorf("failed to delete invite: %w", err)
	}

	return nil
}
