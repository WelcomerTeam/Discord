package http

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/WelcomerTeam/Discord/discord"
	"github.com/WelcomerTeam/Discord/structs"
	"golang.org/x/xerrors"
)

func (s *Session) GetInvite(inviteCode string, withCounts *bool, withExpiration *bool, guildScheduledEventID *discord.Snowflake) (invite *structs.Invite, err error) {
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

func (s *Session) DeleteInvite(inviteCode string, reason *string) (err error) {
	endpoint := EndpointInvite(inviteCode)

	headers := http.Header{}

	if reason != nil {
		headers.Add(structs.AuditLogReasonHeader, *reason)
	}

	err = s.Interface.FetchJJ(s, http.MethodDelete, endpoint, nil, headers, nil)
	if err != nil {
		return xerrors.Errorf("Failed to delete invite: %v", err)
	}

	return
}
