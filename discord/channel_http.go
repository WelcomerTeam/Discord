package discord

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

var emojiEscaper = strings.NewReplacer("#", "%23")

func GetChannel(s *Session, channelID Snowflake) (*Channel, error) {
	endpoint := EndpointChannel(channelID.String())

	var channel *Channel

	err := s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &channel)
	if err != nil {
		return nil, fmt.Errorf("failed to get channel: %w", err)
	}

	return channel, nil
}

func ModifyChannel(s *Session, channelID Snowflake, channelParams ChannelParams, reason *string) (*Channel, error) {
	endpoint := EndpointChannel(channelID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	var channel *Channel

	err := s.Interface.FetchJJ(s, http.MethodGet, endpoint, channelParams, headers, &channel)
	if err != nil {
		return nil, fmt.Errorf("failed to modify channel: %w", err)
	}

	return channel, nil
}

func DeleteChannel(s *Session, channelID Snowflake, reason *string) error {
	endpoint := EndpointChannel(channelID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err := s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, headers, nil)
	if err != nil {
		return fmt.Errorf("failed to delete channel: %w", err)
	}

	return nil
}

func GetChannelMessages(s *Session, channelID Snowflake, around *Snowflake, before *Snowflake, after *Snowflake, limit *int32) ([]*Message, error) {
	endpoint := EndpointChannelMessages(channelID.String())

	values := url.Values{}

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
		values.Add("limit", strconv.Itoa(int(*limit)))
	}

	if len(values) > 0 {
		endpoint += "?" + values.Encode()
	}

	var messages []*Message

	err := s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &messages)
	if err != nil {
		return nil, fmt.Errorf("failed to get channel messages: %w", err)
	}

	return messages, nil
}

func GetChannelMessage(s *Session, channelID Snowflake, messageID Snowflake) (*Message, error) {
	endpoint := EndpointChannelMessage(channelID.String(), messageID.String())

	var message *Message

	err := s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &message)
	if err != nil {
		return nil, fmt.Errorf("failed to get channel message: %w", err)
	}

	return message, nil
}

func CreateMessage(s *Session, channelID Snowflake, messageParams MessageParams) (*Message, error) {
	endpoint := EndpointChannelMessages(channelID.String())

	var message *Message

	if len(messageParams.Files) > 0 {
		contentType, body, err := multipartBodyWithJSON(messageParams, messageParams.Files)
		if err != nil {
			return nil, err
		}

		err = s.Interface.FetchBJ(s, http.MethodPost, endpoint, contentType, body, nil, &message)
		if err != nil {
			return nil, fmt.Errorf("failed to create message: %w", err)
		}
	} else {
		err := s.Interface.FetchJJ(s, http.MethodPost, endpoint, messageParams, nil, &message)
		if err != nil {
			return nil, fmt.Errorf("failed to create message: %w", err)
		}
	}

	return message, nil
}

func CrosspostMessage(s *Session, channelID Snowflake, messageID Snowflake) (*Message, error) {
	endpoint := EndpointChannelMessageCrosspost(channelID.String(), messageID.String())

	var message *Message

	err := s.Interface.FetchJJ(s, http.MethodPost, endpoint, nil, nil, &message)
	if err != nil {
		return nil, fmt.Errorf("failed to crosspost message: %w", err)
	}

	return message, nil
}

func CreateReaction(s *Session, channelID Snowflake, messageID Snowflake, emoji string) error {
	endpoint := EndpointMessageReaction(channelID.String(), messageID.String(), emojiEscaper.Replace(emoji), "@me")

	err := s.Interface.FetchJJ(s, http.MethodPut, endpoint, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to create reaction: %w", err)
	}

	return nil
}

func DeleteOwnReaction(s *Session, channelID Snowflake, messageID Snowflake, emoji string) error {
	endpoint := EndpointMessageReaction(channelID.String(), messageID.String(), emojiEscaper.Replace(emoji), "@me")

	err := s.Interface.FetchJJ(s, http.MethodDelete, endpoint, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete own reaction: %w", err)
	}

	return nil
}

func DeleteUserReaction(s *Session, channelID Snowflake, messageID Snowflake, emoji string, userID Snowflake) error {
	endpoint := EndpointMessageReaction(channelID.String(), messageID.String(), emojiEscaper.Replace(emoji), userID.String())

	err := s.Interface.FetchJJ(s, http.MethodDelete, endpoint, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete user reaction: %w", err)
	}

	return nil
}

func GetReactions(s *Session, channelID Snowflake, messageID Snowflake, emoji string, after *Snowflake, limit *int) ([]*User, error) {
	endpoint := EndpointMessageReactions(channelID.String(), messageID.String(), emojiEscaper.Replace(emoji))

	var users []*User

	err := s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &users)
	if err != nil {
		return nil, fmt.Errorf("failed to get reactions: %w", err)
	}

	return users, nil
}

func DeleteAllReactions(s *Session, channelID Snowflake, messageID Snowflake) error {
	endpoint := EndpointMessageReactionsAll(channelID.String(), messageID.String())

	err := s.Interface.FetchJJ(s, http.MethodDelete, endpoint, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete all reactions: %w", err)
	}

	return nil
}

func DeleteAllReactionsEmoji(s *Session, channelID Snowflake, messageID Snowflake, emoji string) error {
	endpoint := EndpointMessageReactions(channelID.String(), messageID.String(), emojiEscaper.Replace(emoji))

	err := s.Interface.FetchJJ(s, http.MethodDelete, endpoint, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete all reactions emoji: %w", err)
	}

	return nil
}

func EditMessage(s *Session, channelID Snowflake, messageID Snowflake, messageParams MessageParams) (*Message, error) {
	endpoint := EndpointChannelMessage(channelID.String(), messageID.String())

	var message *Message

	if len(messageParams.Files) > 0 {
		contentType, body, err := multipartBodyWithJSON(messageParams, messageParams.Files)
		if err != nil {
			return nil, err
		}

		err = s.Interface.FetchBJ(s, http.MethodPatch, endpoint, contentType, body, nil, &message)
		if err != nil {
			return nil, fmt.Errorf("failed to edit message: %w", err)
		}
	} else {
		err := s.Interface.FetchJJ(s, http.MethodPatch, endpoint, messageParams, nil, &message)
		if err != nil {
			return nil, fmt.Errorf("failed to edit message: %w", err)
		}
	}

	return message, nil
}

func DeleteMessage(s *Session, channelID Snowflake, messageID Snowflake, reason *string) error {
	endpoint := EndpointChannelMessage(channelID.String(), messageID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err := s.Interface.FetchJJ(s, http.MethodDelete, endpoint, nil, headers, nil)
	if err != nil {
		return fmt.Errorf("failed to delet emessage: %w", err)
	}

	return nil
}

func BulkDeleteMessages(s *Session, channelID Snowflake, messageIDs []Snowflake, reason *string) error {
	endpoint := EndpointChannelMessagesBulkDelete(channelID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err := s.Interface.FetchJJ(s, http.MethodDelete, endpoint, nil, headers, nil)
	if err != nil {
		return fmt.Errorf("failed to bulk delete messages: %w", err)
	}

	return nil
}

func EditChannelPermissions(s *Session, channelID Snowflake, overwriteID Snowflake, overwriteArg ChannelOverwrite, reason *string) error {
	endpoint := EndpointChannelPermission(channelID.String(), overwriteID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err := s.Interface.FetchJJ(s, http.MethodPut, endpoint, overwriteArg, headers, nil)
	if err != nil {
		return fmt.Errorf("failed to edit channel permissions: %w", err)
	}

	return nil
}

func GetChannelInvites(s *Session, channelID Snowflake) ([]*Invite, error) {
	endpoint := EndpointChannelInvites(channelID.String())

	var invites []*Invite

	err := s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get channel invites: %w", err)
	}

	return invites, nil
}

func CreateChannelInvite(s *Session, channelID Snowflake, inviteParams InviteParams, reason *string) (*Invite, error) {
	endpoint := EndpointChannelInvites(channelID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	var invite *Invite

	err := s.Interface.FetchJJ(s, http.MethodPost, endpoint, inviteParams, headers, &invite)
	if err != nil {
		return nil, fmt.Errorf("failed to create channel invite: %w", err)
	}

	return invite, nil
}

func DeleteChannelPermission(s *Session, channelID Snowflake, overwriteID Snowflake, reason *string) error {
	endpoint := EndpointChannelPermission(channelID.String(), overwriteID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err := s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, headers, nil)
	if err != nil {
		return fmt.Errorf("failed to delete channel permission: %w", err)
	}

	return nil
}

func FollowNewsChannel(s *Session, channelID Snowflake, webhookChannelID Snowflake) (*FollowedChannel, error) {
	endpoint := EndpointChannelFollow(channelID.String())

	var followedChannel *FollowedChannel

	err := s.Interface.FetchJJ(s, http.MethodPost, endpoint, nil, nil, &followedChannel)
	if err != nil {
		return nil, fmt.Errorf("failed to follow news channel: %w", err)
	}

	return followedChannel, nil
}

func TriggerTypingIndicator(s *Session, channelID Snowflake) error {
	endpoint := EndpointChannelTyping(channelID.String())

	err := s.Interface.FetchJJ(s, http.MethodPost, endpoint, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to trigger typing indicator: %w", err)
	}

	return nil
}

func GetPinnedMessages(s *Session, channelID Snowflake) ([]*Message, error) {
	endpoint := EndpointChannelMessagesPins(channelID.String())

	var pinnedMessages []*Message

	err := s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &pinnedMessages)
	if err != nil {
		return nil, fmt.Errorf("failed to get pinned messages: %w", err)
	}

	return pinnedMessages, nil
}

func PinMessage(s *Session, channelID Snowflake, messageID Snowflake, reason *string) error {
	endpoint := EndpointChannelMessagePin(channelID.String(), messageID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err := s.Interface.FetchJJ(s, http.MethodPut, endpoint, nil, headers, nil)
	if err != nil {
		return fmt.Errorf("failed to pin message: %w", err)
	}

	return nil
}

func UnpinMessage(s *Session, channelID Snowflake, messageID Snowflake, reason *string) error {
	endpoint := EndpointChannelMessagePin(channelID.String(), messageID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err := s.Interface.FetchJJ(s, http.MethodDelete, endpoint, nil, headers, nil)
	if err != nil {
		return fmt.Errorf("failed to unpin message: %w", err)
	}

	return nil
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
