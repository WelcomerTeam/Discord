package discord

import (
	"golang.org/x/xerrors"
	"net/http"
	"net/url"
	"strings"
)

var emojiEscaper = strings.NewReplacer("#", "%23")

func GetChannel(s *Session, channelID Snowflake) (channel *Channel, err error) {
	endpoint := EndpointChannel(channelID.String())

	err = s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &channel)
	if err != nil {
		return nil, xerrors.Errorf("Failed to get channel: %v", err)
	}

	return
}

func ModifyChannel(s *Session, channelID Snowflake, channelParams ChannelParams, reason *string) (channel *Channel, err error) {
	endpoint := EndpointChannel(channelID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err = s.Interface.FetchJJ(s, http.MethodGet, endpoint, channelParams, headers, &channel)
	if err != nil {
		return nil, xerrors.Errorf("Failed to modify channel: %v", err)
	}

	return
}

func DeleteChannel(s *Session, channelID Snowflake, reason *string) (err error) {
	endpoint := EndpointChannel(channelID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err = s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, headers, nil)
	if err != nil {
		return xerrors.Errorf("Failed to delete channel: %v", err)
	}

	return
}

func GetChannelMessages(s *Session, channelID Snowflake, around *Snowflake, before *Snowflake, after *Snowflake, limit *int32) (messages []*Message, err error) {
	endpoint := EndpointChannelMessages(channelID.String())

	var values url.Values

	if around != nil {
		values.Add("around", around.String())
	}

	if before != nil {
		values.Add("before", before.String())
	}

	if after != nil {
		values.Add("after", after.String())
	}

	if limit != nil {
		values.Add("limit", string(*limit))
	}

	if len(values) > 0 {
		endpoint += "?" + values.Encode()
	}

	err = s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &messages)
	if err != nil {
		return nil, xerrors.Errorf("Failed to get channel messages: %v", err)
	}

	return messages, nil
}

func GetChannelMessage(s *Session, channelID Snowflake, messageID Snowflake) (message *Message, err error) {
	endpoint := EndpointChannelMessage(channelID.String(), messageID.String())

	err = s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &message)
	if err != nil {
		return nil, xerrors.Errorf("Failed to get channel message: %v", err)
	}

	return
}

func CreateMessage(s *Session, channelID Snowflake, messageParams MessageParams) (message *Message, err error) {
	endpoint := EndpointChannelMessages(channelID.String())

	if len(messageParams.Files) > 0 {
		contentType, body, err := multipartBodyWithJSON(messageParams, messageParams.Files)
		if err != nil {
			return nil, err
		}

		err = s.Interface.FetchBJ(s, http.MethodPost, endpoint, contentType, body, nil, &message)
		if err != nil {
			return nil, xerrors.Errorf("Failed to create message: %v", err)
		}
	} else {
		err = s.Interface.FetchJJ(s, http.MethodPost, endpoint, messageParams, nil, &message)
		if err != nil {
			return nil, xerrors.Errorf("Failed to create message: %v", err)
		}
	}

	return
}

func CrosspostMessage(s *Session, channelID Snowflake, messageID Snowflake) (message *Message, err error) {
	endpoint := EndpointChannelMessageCrosspost(channelID.String(), messageID.String())

	err = s.Interface.FetchJJ(s, http.MethodPost, endpoint, nil, nil, &message)
	if err != nil {
		return nil, xerrors.Errorf("Failed to crosspost message: %v", err)
	}

	return
}

func CreateReaction(s *Session, channelID Snowflake, messageID Snowflake, emoji string) (err error) {
	endpoint := EndpointMessageReaction(channelID.String(), messageID.String(), emojiEscaper.Replace(emoji), "@me")

	err = s.Interface.FetchJJ(s, http.MethodPut, endpoint, nil, nil, nil)
	if err != nil {
		return xerrors.Errorf("Failed to create reaction: %v", err)
	}

	return
}

func DeleteOwnReaction(s *Session, channelID Snowflake, messageID Snowflake, emoji string) (err error) {
	endpoint := EndpointMessageReaction(channelID.String(), messageID.String(), emojiEscaper.Replace(emoji), "@me")

	err = s.Interface.FetchJJ(s, http.MethodDelete, endpoint, nil, nil, nil)
	if err != nil {
		return xerrors.Errorf("Failed to delete own reaction: %v", err)
	}

	return
}

func DeleteUserReaction(s *Session, channelID Snowflake, messageID Snowflake, emoji string, userID Snowflake) (err error) {
	endpoint := EndpointMessageReaction(channelID.String(), messageID.String(), emojiEscaper.Replace(emoji), userID.String())

	err = s.Interface.FetchJJ(s, http.MethodDelete, endpoint, nil, nil, nil)
	if err != nil {
		return xerrors.Errorf("Failed to delete user reaction: %v", err)
	}

	return
}

func GetReactions(s *Session, channelID Snowflake, messageID Snowflake, emoji string, after *Snowflake, limit *int) (users []*User, err error) {
	endpoint := EndpointMessageReactions(channelID.String(), messageID.String(), emojiEscaper.Replace(emoji))

	err = s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &users)
	if err != nil {
		return nil, xerrors.Errorf("Failed to get reactions: %v", err)
	}

	return
}

func DeleteAllReactions(s *Session, channelID Snowflake, messageID Snowflake) (err error) {
	endpoint := EndpointMessageReactionsAll(channelID.String(), messageID.String())

	err = s.Interface.FetchJJ(s, http.MethodDelete, endpoint, nil, nil, nil)
	if err != nil {
		return xerrors.Errorf("Failed to delete all reactions: %v", err)
	}

	return
}

func DeleteAllReactionsEmoji(s *Session, channelID Snowflake, messageID Snowflake, emoji string) (err error) {
	endpoint := EndpointMessageReactions(channelID.String(), messageID.String(), emojiEscaper.Replace(emoji))

	err = s.Interface.FetchJJ(s, http.MethodDelete, endpoint, nil, nil, nil)
	if err != nil {
		return xerrors.Errorf("Failed to delete all reactions emoji: %v", err)
	}

	return
}

func EditMessage(s *Session, channelID Snowflake, messageID Snowflake, messageParams MessageParams) (message *Message, err error) {
	endpoint := EndpointChannelMessage(channelID.String(), messageID.String())

	if len(messageParams.Files) > 0 {
		contentType, body, err := multipartBodyWithJSON(messageParams, messageParams.Files)
		if err != nil {
			return nil, err
		}

		err = s.Interface.FetchBJ(s, http.MethodPatch, endpoint, contentType, body, nil, &message)
		if err != nil {
			return nil, xerrors.Errorf("Failed to edit message: %v", err)
		}
	} else {
		err = s.Interface.FetchJJ(s, http.MethodPatch, endpoint, messageParams, nil, &message)
		if err != nil {
			return nil, xerrors.Errorf("Failed to edit message: %v", err)
		}
	}

	return
}

func DeleteMessage(s *Session, channelID Snowflake, messageID Snowflake, reason *string) (err error) {
	endpoint := EndpointChannelMessage(channelID.String(), messageID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err = s.Interface.FetchJJ(s, http.MethodDelete, endpoint, nil, headers, nil)
	if err != nil {
		return xerrors.Errorf("Failed to delet emessage: %v", err)
	}

	return
}

func BulkDeleteMessages(s *Session, channelID Snowflake, messageIDs []Snowflake, reason *string) (err error) {
	endpoint := EndpointChannelMessagesBulkDelete(channelID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err = s.Interface.FetchJJ(s, http.MethodDelete, endpoint, nil, headers, nil)
	if err != nil {
		return xerrors.Errorf("Failed to bulk delete messages: %v", err)
	}

	return
}

func EditChannelPermissions(s *Session, channelID Snowflake, overwriteID Snowflake, overwriteArg ChannelOverwrite, reason *string) (err error) {
	endpoint := EndpointChannelPermission(channelID.String(), overwriteID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err = s.Interface.FetchJJ(s, http.MethodPut, endpoint, overwriteArg, headers, nil)
	if err != nil {
		return xerrors.Errorf("Failed to edit channel permissions: %v", err)
	}

	return
}

func GetChannelInvites(s *Session, channelID Snowflake) (invites []*Invite, err error) {
	endpoint := EndpointChannelInvites(channelID.String())

	err = s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, nil)
	if err != nil {
		return nil, xerrors.Errorf("Failed to get channel invites: %v", err)
	}

	return
}

func CreateChannelInvite(s *Session, channelID Snowflake, inviteParams InviteParams, reason *string) (invite *Invite, err error) {
	endpoint := EndpointChannelInvites(channelID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err = s.Interface.FetchJJ(s, http.MethodPost, endpoint, inviteParams, headers, &invite)
	if err != nil {
		return nil, xerrors.Errorf("Failed to create channel invite: %v", err)
	}

	return
}

func DeleteChannelPermission(s *Session, channelID Snowflake, overwriteID Snowflake, reason *string) (err error) {
	endpoint := EndpointChannelPermission(channelID.String(), overwriteID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err = s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, headers, nil)
	if err != nil {
		return xerrors.Errorf("Failed to delete channel permission: %v", err)
	}

	return
}

func FollowNewsChannel(s *Session, channelID Snowflake, webhookChannelID Snowflake) (followedChannel *FollowedChannel, err error) {
	endpoint := EndpointChannelFollow(channelID.String())

	err = s.Interface.FetchJJ(s, http.MethodPost, endpoint, nil, nil, &followedChannel)
	if err != nil {
		return nil, xerrors.Errorf("Failed to follow news channel: %v", err)
	}

	return
}

func TriggerTypingIndicator(s *Session, channelID Snowflake) (err error) {
	endpoint := EndpointChannelTyping(channelID.String())

	err = s.Interface.FetchJJ(s, http.MethodPost, endpoint, nil, nil, nil)
	if err != nil {
		return xerrors.Errorf("Failed to trigger typing indicator: %v", err)
	}

	return
}

func GetPinnedMessages(s *Session, channelID Snowflake) (pinnedMessages []*Message, err error) {
	endpoint := EndpointChannelMessagesPins(channelID.String())

	err = s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &pinnedMessages)
	if err != nil {
		return nil, xerrors.Errorf("Failed to get pinned messages: %v", err)
	}

	return
}

func PinMessage(s *Session, channelID Snowflake, messageID Snowflake, reason *string) (err error) {
	endpoint := EndpointChannelMessagePin(channelID.String(), messageID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err = s.Interface.FetchJJ(s, http.MethodPut, endpoint, nil, headers, nil)
	if err != nil {
		return xerrors.Errorf("Failed to pin message: %v", err)
	}

	return
}

func UnpinMessage(s *Session, channelID Snowflake, messageID Snowflake, reason *string) (err error) {
	endpoint := EndpointChannelMessagePin(channelID.String(), messageID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err = s.Interface.FetchJJ(s, http.MethodDelete, endpoint, nil, headers, nil)
	if err != nil {
		return xerrors.Errorf("Failed to unpin message: %v", err)
	}

	return
}

// TODO: GroupDMAddRecipient
// TODO: GroupDMRemoveRecipient
// TODO: StartThreadwithMessage
// TODO: StartThreadwithoutMessage
// TODO: JoinThread
// TODO: AddThreadMember
// TODO: LeaveThread
// TODO: RemoveThreadMember
// TODO: GetThreadMember
// TODO: ListThreadMembers
// TODO: ListPublicArchivedThreads
// TODO: ListPrivateArchivedThreads
// TODO: ListJoinedPrivateArchivedThreads
