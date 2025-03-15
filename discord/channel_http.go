package discord

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

var emojiEscaper = strings.NewReplacer("#", "%23")

func GetChannel(ctx context.Context, s *Session, channelID Snowflake) (*Channel, error) {
	endpoint := EndpointChannel(channelID.String())

	var channel *Channel

	err := s.Interface.FetchJJ(ctx, s, http.MethodGet, endpoint, nil, nil, &channel)
	if err != nil {
		return nil, fmt.Errorf("failed to get channel: %w", err)
	}

	return channel, nil
}

func ModifyChannel(ctx context.Context, s *Session, channelID Snowflake, channelParams ChannelParams, reason *string) (*Channel, error) {
	endpoint := EndpointChannel(channelID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	var channel *Channel

	err := s.Interface.FetchJJ(ctx, s, http.MethodGet, endpoint, channelParams, headers, &channel)
	if err != nil {
		return nil, fmt.Errorf("failed to modify channel: %w", err)
	}

	return channel, nil
}

func DeleteChannel(ctx context.Context, s *Session, channelID Snowflake, reason *string) error {
	endpoint := EndpointChannel(channelID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err := s.Interface.FetchJJ(ctx, s, http.MethodDelete, endpoint, nil, headers, nil)
	if err != nil {
		return fmt.Errorf("failed to delete channel: %w", err)
	}

	return nil
}

func GetChannelMessages(ctx context.Context, s *Session, channelID Snowflake, around *Snowflake, before *Snowflake, after *Snowflake, limit *int32) ([]Message, error) {
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

	var messages []Message

	err := s.Interface.FetchJJ(ctx, s, http.MethodGet, endpoint, nil, nil, &messages)
	if err != nil {
		return nil, fmt.Errorf("failed to get channel messages: %w", err)
	}

	return messages, nil
}

func GetChannelMessage(ctx context.Context, s *Session, channelID Snowflake, messageID Snowflake) (*Message, error) {
	endpoint := EndpointChannelMessage(channelID.String(), messageID.String())

	var message *Message

	err := s.Interface.FetchJJ(ctx, s, http.MethodGet, endpoint, nil, nil, &message)
	if err != nil {
		return nil, fmt.Errorf("failed to get channel message: %w", err)
	}

	return message, nil
}

func CreateMessage(ctx context.Context, s *Session, channelID Snowflake, messageParams MessageParams) (*Message, error) {
	endpoint := EndpointChannelMessages(channelID.String())

	var message *Message

	if len(messageParams.Files) > 0 {
		contentType, body, err := multipartBodyWithJSON(messageParams, messageParams.Files)
		if err != nil {
			return nil, err
		}

		err = s.Interface.FetchBJ(ctx, s, http.MethodPost, endpoint, contentType, body, nil, &message)
		if err != nil {
			return nil, fmt.Errorf("failed to create message: %w", err)
		}
	} else {
		err := s.Interface.FetchJJ(ctx, s, http.MethodPost, endpoint, messageParams, nil, &message)
		if err != nil {
			return nil, fmt.Errorf("failed to create message: %w", err)
		}
	}

	return message, nil
}

func CrosspostMessage(ctx context.Context, s *Session, channelID Snowflake, messageID Snowflake) (*Message, error) {
	endpoint := EndpointChannelMessageCrosspost(channelID.String(), messageID.String())

	var message *Message

	err := s.Interface.FetchJJ(ctx, s, http.MethodPost, endpoint, nil, nil, &message)
	if err != nil {
		return nil, fmt.Errorf("failed to crosspost message: %w", err)
	}

	return message, nil
}

func CreateReaction(ctx context.Context, s *Session, channelID Snowflake, messageID Snowflake, emoji string) error {
	endpoint := EndpointMessageReaction(channelID.String(), messageID.String(), emojiEscaper.Replace(emoji), "@me")

	err := s.Interface.FetchJJ(ctx, s, http.MethodPut, endpoint, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to create reaction: %w", err)
	}

	return nil
}

func DeleteOwnReaction(ctx context.Context, s *Session, channelID Snowflake, messageID Snowflake, emoji string) error {
	endpoint := EndpointMessageReaction(channelID.String(), messageID.String(), emojiEscaper.Replace(emoji), "@me")

	err := s.Interface.FetchJJ(ctx, s, http.MethodDelete, endpoint, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete own reaction: %w", err)
	}

	return nil
}

func DeleteUserReaction(ctx context.Context, s *Session, channelID Snowflake, messageID Snowflake, emoji string, userID Snowflake) error {
	endpoint := EndpointMessageReaction(channelID.String(), messageID.String(), emojiEscaper.Replace(emoji), userID.String())

	err := s.Interface.FetchJJ(ctx, s, http.MethodDelete, endpoint, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete user reaction: %w", err)
	}

	return nil
}

func GetReactions(ctx context.Context, s *Session, channelID Snowflake, messageID Snowflake, emoji string, after *Snowflake, limit *int) ([]User, error) {
	endpoint := EndpointMessageReactions(channelID.String(), messageID.String(), emojiEscaper.Replace(emoji))

	values := url.Values{}

	if after != nil {
		values.Add("after", after.String())
	}

	if limit != nil {
		values.Add("limit", strconv.Itoa(int(*limit)))
	}

	if len(values) > 0 {
		endpoint += "?" + values.Encode()
	}

	var users []User

	err := s.Interface.FetchJJ(ctx, s, http.MethodGet, endpoint, nil, nil, &users)
	if err != nil {
		return nil, fmt.Errorf("failed to get reactions: %w", err)
	}

	return users, nil
}

func DeleteAllReactions(ctx context.Context, s *Session, channelID Snowflake, messageID Snowflake) error {
	endpoint := EndpointMessageReactionsAll(channelID.String(), messageID.String())

	err := s.Interface.FetchJJ(ctx, s, http.MethodDelete, endpoint, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete all reactions: %w", err)
	}

	return nil
}

func DeleteAllReactionsEmoji(ctx context.Context, s *Session, channelID Snowflake, messageID Snowflake, emoji string) error {
	endpoint := EndpointMessageReactions(channelID.String(), messageID.String(), emojiEscaper.Replace(emoji))

	err := s.Interface.FetchJJ(ctx, s, http.MethodDelete, endpoint, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete all reactions emoji: %w", err)
	}

	return nil
}

func EditMessage(ctx context.Context, s *Session, channelID Snowflake, messageID Snowflake, messageParams MessageParams) (*Message, error) {
	endpoint := EndpointChannelMessage(channelID.String(), messageID.String())

	var message *Message

	if len(messageParams.Files) > 0 {
		contentType, body, err := multipartBodyWithJSON(messageParams, messageParams.Files)
		if err != nil {
			return nil, err
		}

		err = s.Interface.FetchBJ(ctx, s, http.MethodPatch, endpoint, contentType, body, nil, &message)
		if err != nil {
			return nil, fmt.Errorf("failed to edit message: %w", err)
		}
	} else {
		err := s.Interface.FetchJJ(ctx, s, http.MethodPatch, endpoint, messageParams, nil, &message)
		if err != nil {
			return nil, fmt.Errorf("failed to edit message: %w", err)
		}
	}

	return message, nil
}

func DeleteMessage(ctx context.Context, s *Session, channelID Snowflake, messageID Snowflake, reason *string) error {
	endpoint := EndpointChannelMessage(channelID.String(), messageID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err := s.Interface.FetchJJ(ctx, s, http.MethodDelete, endpoint, nil, headers, nil)
	if err != nil {
		return fmt.Errorf("failed to delet emessage: %w", err)
	}

	return nil
}

type bulkDeleteMessagesBody struct {
	Messages []Snowflake `json:"messages"`
}

func BulkDeleteMessages(ctx context.Context, s *Session, channelID Snowflake, messageIDs []Snowflake, reason *string) error {
	endpoint := EndpointChannelMessagesBulkDelete(channelID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err := s.Interface.FetchJJ(ctx, s, http.MethodPost, endpoint, bulkDeleteMessagesBody{messageIDs}, headers, nil)
	if err != nil {
		return fmt.Errorf("failed to bulk delete messages: %w", err)
	}

	return nil
}

func EditChannelPermissions(ctx context.Context, s *Session, channelID Snowflake, overwriteID Snowflake, overwriteArg ChannelOverwrite, reason *string) error {
	endpoint := EndpointChannelPermission(channelID.String(), overwriteID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err := s.Interface.FetchJJ(ctx, s, http.MethodPut, endpoint, overwriteArg, headers, nil)
	if err != nil {
		return fmt.Errorf("failed to edit channel permissions: %w", err)
	}

	return nil
}

func GetChannelInvites(ctx context.Context, s *Session, channelID Snowflake) ([]Invite, error) {
	endpoint := EndpointChannelInvites(channelID.String())

	var invites []Invite

	err := s.Interface.FetchJJ(ctx, s, http.MethodGet, endpoint, nil, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get channel invites: %w", err)
	}

	return invites, nil
}

func CreateChannelInvite(ctx context.Context, s *Session, channelID Snowflake, inviteParams InviteParams, reason *string) (*Invite, error) {
	endpoint := EndpointChannelInvites(channelID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	var invite *Invite

	err := s.Interface.FetchJJ(ctx, s, http.MethodPost, endpoint, inviteParams, headers, &invite)
	if err != nil {
		return nil, fmt.Errorf("failed to create channel invite: %w", err)
	}

	return invite, nil
}

func DeleteChannelPermission(ctx context.Context, s *Session, channelID Snowflake, overwriteID Snowflake, reason *string) error {
	endpoint := EndpointChannelPermission(channelID.String(), overwriteID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err := s.Interface.FetchJJ(ctx, s, http.MethodDelete, endpoint, nil, headers, nil)
	if err != nil {
		return fmt.Errorf("failed to delete channel permission: %w", err)
	}

	return nil
}

type FollowAnnouncementChannelBody struct {
	WebhookChannelID Snowflake `json:"webhook_channel_id"`
}

func FollowAnnouncementChannel(ctx context.Context, s *Session, channelID Snowflake, webhookChannelID Snowflake) (*FollowedChannel, error) {
	endpoint := EndpointChannelFollow(channelID.String())

	var followedChannel *FollowedChannel

	err := s.Interface.FetchJJ(ctx, s, http.MethodPost, endpoint, FollowAnnouncementChannelBody{webhookChannelID}, nil, &followedChannel)
	if err != nil {
		return nil, fmt.Errorf("failed to follow announcement channel: %w", err)
	}

	return followedChannel, nil
}

func TriggerTypingIndicator(ctx context.Context, s *Session, channelID Snowflake) error {
	endpoint := EndpointChannelTyping(channelID.String())

	err := s.Interface.FetchJJ(ctx, s, http.MethodPost, endpoint, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to trigger typing indicator: %w", err)
	}

	return nil
}

func GetPinnedMessages(ctx context.Context, s *Session, channelID Snowflake) ([]Message, error) {
	endpoint := EndpointChannelMessagesPins(channelID.String())

	var pinnedMessages []Message

	err := s.Interface.FetchJJ(ctx, s, http.MethodGet, endpoint, nil, nil, &pinnedMessages)
	if err != nil {
		return nil, fmt.Errorf("failed to get pinned messages: %w", err)
	}

	return pinnedMessages, nil
}

func PinMessage(ctx context.Context, s *Session, channelID Snowflake, messageID Snowflake, reason *string) error {
	endpoint := EndpointChannelMessagePin(channelID.String(), messageID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err := s.Interface.FetchJJ(ctx, s, http.MethodPut, endpoint, nil, headers, nil)
	if err != nil {
		return fmt.Errorf("failed to pin message: %w", err)
	}

	return nil
}

func UnpinMessage(ctx context.Context, s *Session, channelID Snowflake, messageID Snowflake, reason *string) error {
	endpoint := EndpointChannelMessagePin(channelID.String(), messageID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err := s.Interface.FetchJJ(ctx, s, http.MethodDelete, endpoint, nil, headers, nil)
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
