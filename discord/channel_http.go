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

func GetChannel(ctx context.Context, session *Session, channelID Snowflake) (*Channel, error) {
	endpoint := EndpointChannel(channelID.String())

	var channel *Channel

	err := session.Interface.FetchJJ(ctx, session, http.MethodGet, endpoint, nil, nil, &channel)
	if err != nil {
		return nil, fmt.Errorf("failed to get channel: %w", err)
	}

	return channel, nil
}

func ModifyChannel(ctx context.Context, session *Session, channelID Snowflake, channelParams ChannelParams, reason *string) (*Channel, error) {
	endpoint := EndpointChannel(channelID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	var channel *Channel

	err := session.Interface.FetchJJ(ctx, session, http.MethodGet, endpoint, channelParams, headers, &channel)
	if err != nil {
		return nil, fmt.Errorf("failed to modify channel: %w", err)
	}

	return channel, nil
}

func DeleteChannel(ctx context.Context, session *Session, channelID Snowflake, reason *string) error {
	endpoint := EndpointChannel(channelID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err := session.Interface.FetchJJ(ctx, session, http.MethodDelete, endpoint, nil, headers, nil)
	if err != nil {
		return fmt.Errorf("failed to delete channel: %w", err)
	}

	return nil
}

func GetChannelMessages(ctx context.Context, session *Session, channelID Snowflake, around, before, after *Snowflake, limit *int32) ([]Message, error) {
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

	err := session.Interface.FetchJJ(ctx, session, http.MethodGet, endpoint, nil, nil, &messages)
	if err != nil {
		return nil, fmt.Errorf("failed to get channel messages: %w", err)
	}

	return messages, nil
}

func GetChannelMessage(ctx context.Context, session *Session, channelID, messageID Snowflake) (*Message, error) {
	endpoint := EndpointChannelMessage(channelID.String(), messageID.String())

	var message *Message

	err := session.Interface.FetchJJ(ctx, session, http.MethodGet, endpoint, nil, nil, &message)
	if err != nil {
		return nil, fmt.Errorf("failed to get channel message: %w", err)
	}

	return message, nil
}

func CreateMessage(ctx context.Context, session *Session, channelID Snowflake, messageParams MessageParams) (*Message, error) {
	endpoint := EndpointChannelMessages(channelID.String())

	var message *Message

	if len(messageParams.Files) > 0 {
		contentType, body, err := multipartBodyWithJSON(messageParams, messageParams.Files)
		if err != nil {
			return nil, err
		}

		err = session.Interface.FetchBJ(ctx, session, http.MethodPost, endpoint, contentType, body, nil, &message)
		if err != nil {
			return nil, fmt.Errorf("failed to create message: %w", err)
		}
	} else {
		err := session.Interface.FetchJJ(ctx, session, http.MethodPost, endpoint, messageParams, nil, &message)
		if err != nil {
			return nil, fmt.Errorf("failed to create message: %w", err)
		}
	}

	return message, nil
}

func CrosspostMessage(ctx context.Context, session *Session, channelID, messageID Snowflake) (*Message, error) {
	endpoint := EndpointChannelMessageCrosspost(channelID.String(), messageID.String())

	var message *Message

	err := session.Interface.FetchJJ(ctx, session, http.MethodPost, endpoint, nil, nil, &message)
	if err != nil {
		return nil, fmt.Errorf("failed to crosspost message: %w", err)
	}

	return message, nil
}

func CreateReaction(ctx context.Context, session *Session, channelID, messageID Snowflake, emoji string) error {
	endpoint := EndpointMessageReaction(channelID.String(), messageID.String(), emojiEscaper.Replace(emoji), "@me")

	err := session.Interface.FetchJJ(ctx, session, http.MethodPut, endpoint, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to create reaction: %w", err)
	}

	return nil
}

func DeleteOwnReaction(ctx context.Context, session *Session, channelID, messageID Snowflake, emoji string) error {
	endpoint := EndpointMessageReaction(channelID.String(), messageID.String(), emojiEscaper.Replace(emoji), "@me")

	err := session.Interface.FetchJJ(ctx, session, http.MethodDelete, endpoint, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete own reaction: %w", err)
	}

	return nil
}

func DeleteUserReaction(ctx context.Context, session *Session, channelID, messageID Snowflake, emoji string, userID Snowflake) error {
	endpoint := EndpointMessageReaction(channelID.String(), messageID.String(), emojiEscaper.Replace(emoji), userID.String())

	err := session.Interface.FetchJJ(ctx, session, http.MethodDelete, endpoint, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete user reaction: %w", err)
	}

	return nil
}

func GetReactions(ctx context.Context, session *Session, channelID, messageID Snowflake, emoji string, after *Snowflake, limit *int) ([]User, error) {
	endpoint := EndpointMessageReactions(channelID.String(), messageID.String(), emojiEscaper.Replace(emoji))

	values := url.Values{}

	if after != nil {
		values.Add("after", after.String())
	}

	if limit != nil {
		values.Add("limit", strconv.Itoa(*limit))
	}

	if len(values) > 0 {
		endpoint += "?" + values.Encode()
	}

	var users []User

	err := session.Interface.FetchJJ(ctx, session, http.MethodGet, endpoint, nil, nil, &users)
	if err != nil {
		return nil, fmt.Errorf("failed to get reactions: %w", err)
	}

	return users, nil
}

func DeleteAllReactions(ctx context.Context, session *Session, channelID, messageID Snowflake) error {
	endpoint := EndpointMessageReactionsAll(channelID.String(), messageID.String())

	err := session.Interface.FetchJJ(ctx, session, http.MethodDelete, endpoint, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete all reactions: %w", err)
	}

	return nil
}

func DeleteAllReactionsEmoji(ctx context.Context, session *Session, channelID, messageID Snowflake, emoji string) error {
	endpoint := EndpointMessageReactions(channelID.String(), messageID.String(), emojiEscaper.Replace(emoji))

	err := session.Interface.FetchJJ(ctx, session, http.MethodDelete, endpoint, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete all reactions emoji: %w", err)
	}

	return nil
}

func EditMessage(ctx context.Context, session *Session, channelID, messageID Snowflake, messageParams MessageParams) (*Message, error) {
	endpoint := EndpointChannelMessage(channelID.String(), messageID.String())

	var message *Message

	if len(messageParams.Files) > 0 {
		contentType, body, err := multipartBodyWithJSON(messageParams, messageParams.Files)
		if err != nil {
			return nil, err
		}

		err = session.Interface.FetchBJ(ctx, session, http.MethodPatch, endpoint, contentType, body, nil, &message)
		if err != nil {
			return nil, fmt.Errorf("failed to edit message: %w", err)
		}
	} else {
		err := session.Interface.FetchJJ(ctx, session, http.MethodPatch, endpoint, messageParams, nil, &message)
		if err != nil {
			return nil, fmt.Errorf("failed to edit message: %w", err)
		}
	}

	return message, nil
}

func DeleteMessage(ctx context.Context, session *Session, channelID, messageID Snowflake, reason *string) error {
	endpoint := EndpointChannelMessage(channelID.String(), messageID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err := session.Interface.FetchJJ(ctx, session, http.MethodDelete, endpoint, nil, headers, nil)
	if err != nil {
		return fmt.Errorf("failed to delet emessage: %w", err)
	}

	return nil
}

type bulkDeleteMessagesBody struct {
	Messages []Snowflake `json:"messages"`
}

func BulkDeleteMessages(ctx context.Context, session *Session, channelID Snowflake, messageIDs []Snowflake, reason *string) error {
	endpoint := EndpointChannelMessagesBulkDelete(channelID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err := session.Interface.FetchJJ(ctx, session, http.MethodPost, endpoint, bulkDeleteMessagesBody{messageIDs}, headers, nil)
	if err != nil {
		return fmt.Errorf("failed to bulk delete messages: %w", err)
	}

	return nil
}

func EditChannelPermissions(ctx context.Context, session *Session, channelID, overwriteID Snowflake, overwriteArg ChannelOverwrite, reason *string) error {
	endpoint := EndpointChannelPermission(channelID.String(), overwriteID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err := session.Interface.FetchJJ(ctx, session, http.MethodPut, endpoint, overwriteArg, headers, nil)
	if err != nil {
		return fmt.Errorf("failed to edit channel permissions: %w", err)
	}

	return nil
}

func GetChannelInvites(ctx context.Context, session *Session, channelID Snowflake) ([]Invite, error) {
	endpoint := EndpointChannelInvites(channelID.String())

	var invites []Invite

	err := session.Interface.FetchJJ(ctx, session, http.MethodGet, endpoint, nil, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get channel invites: %w", err)
	}

	return invites, nil
}

func CreateChannelInvite(ctx context.Context, session *Session, channelID Snowflake, inviteParams InviteParams, reason *string) (*Invite, error) {
	endpoint := EndpointChannelInvites(channelID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	var invite *Invite

	err := session.Interface.FetchJJ(ctx, session, http.MethodPost, endpoint, inviteParams, headers, &invite)
	if err != nil {
		return nil, fmt.Errorf("failed to create channel invite: %w", err)
	}

	return invite, nil
}

func DeleteChannelPermission(ctx context.Context, session *Session, channelID, overwriteID Snowflake, reason *string) error {
	endpoint := EndpointChannelPermission(channelID.String(), overwriteID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err := session.Interface.FetchJJ(ctx, session, http.MethodDelete, endpoint, nil, headers, nil)
	if err != nil {
		return fmt.Errorf("failed to delete channel permission: %w", err)
	}

	return nil
}

type FollowAnnouncementChannelBody struct {
	WebhookChannelID Snowflake `json:"webhook_channel_id"`
}

func FollowAnnouncementChannel(ctx context.Context, session *Session, channelID, webhookChannelID Snowflake) (*FollowedChannel, error) {
	endpoint := EndpointChannelFollow(channelID.String())

	var followedChannel *FollowedChannel

	err := session.Interface.FetchJJ(ctx, session, http.MethodPost, endpoint, FollowAnnouncementChannelBody{webhookChannelID}, nil, &followedChannel)
	if err != nil {
		return nil, fmt.Errorf("failed to follow announcement channel: %w", err)
	}

	return followedChannel, nil
}

func TriggerTypingIndicator(ctx context.Context, session *Session, channelID Snowflake) error {
	endpoint := EndpointChannelTyping(channelID.String())

	err := session.Interface.FetchJJ(ctx, session, http.MethodPost, endpoint, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to trigger typing indicator: %w", err)
	}

	return nil
}

func GetPinnedMessages(ctx context.Context, session *Session, channelID Snowflake) ([]Message, error) {
	endpoint := EndpointChannelMessagesPins(channelID.String())

	var pinnedMessages []Message

	err := session.Interface.FetchJJ(ctx, session, http.MethodGet, endpoint, nil, nil, &pinnedMessages)
	if err != nil {
		return nil, fmt.Errorf("failed to get pinned messages: %w", err)
	}

	return pinnedMessages, nil
}

func PinMessage(ctx context.Context, session *Session, channelID, messageID Snowflake, reason *string) error {
	endpoint := EndpointChannelMessagePin(channelID.String(), messageID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err := session.Interface.FetchJJ(ctx, session, http.MethodPut, endpoint, nil, headers, nil)
	if err != nil {
		return fmt.Errorf("failed to pin message: %w", err)
	}

	return nil
}

func UnpinMessage(ctx context.Context, session *Session, channelID, messageID Snowflake, reason *string) error {
	endpoint := EndpointChannelMessagePin(channelID.String(), messageID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err := session.Interface.FetchJJ(ctx, session, http.MethodDelete, endpoint, nil, headers, nil)
	if err != nil {
		return fmt.Errorf("failed to unpin message: %w", err)
	}

	return nil
}

// CreateThreadParams represents parameters for creating a thread.
type CreateThreadParams struct {
	Name                int32        `json:"name"`
	AutoArchiveDuration int32        `json:"auto_archive_duration,omitempty"`
	Type                *ChannelType `json:"type,omitempty"`
	Invitable           *bool        `json:"invitable,omitempty"`
}

// GroupDMAddRecipientParams represents parameters for adding a recipient to a group DM.
type GroupDMAddRecipientParams struct {
	AccessToken string `json:"access_token"`
	Nick        string `json:"nick,omitempty"`
}

// StartThreadWithMessage creates a thread from an existing message.
func StartThreadWithMessage(ctx context.Context, session *Session, channelID, messageID Snowflake, name string, autoArchiveDuration int32, reason *string) (*Channel, error) {
	endpoint := EndpointChannelMessage(channelID.String(), messageID.String()) + "/threads"

	params := struct {
		Name                string `json:"name"`
		AutoArchiveDuration int32  `json:"auto_archive_duration"`
	}{
		Name:                name,
		AutoArchiveDuration: autoArchiveDuration,
	}

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	var thread *Channel

	err := session.Interface.FetchJJ(ctx, session, http.MethodPost, endpoint, params, headers, &thread)
	if err != nil {
		return nil, fmt.Errorf("failed to start thread with message: %w", err)
	}

	return thread, nil
}

// StartThreadWithoutMessage creates a thread without an associated message.
func StartThreadWithoutMessage(ctx context.Context, session *Session, channelID Snowflake, params CreateThreadParams, reason *string) (*Channel, error) {
	endpoint := EndpointChannelThreads(channelID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	var thread *Channel

	err := session.Interface.FetchJJ(ctx, session, http.MethodPost, endpoint, params, nil, &thread)
	if err != nil {
		return nil, fmt.Errorf("failed to start thread without message: %w", err)
	}

	return thread, nil
}

// JoinThread adds the current user to a thread.
func JoinThread(ctx context.Context, session *Session, channelID Snowflake) error {
	endpoint := EndpointChannel(channelID.String()) + "/thread-members/@me"

	err := session.Interface.FetchJJ(ctx, session, http.MethodPut, endpoint, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to join thread: %w", err)
	}

	return nil
}

// LeaveThread removes the current user from a thread.
func LeaveThread(ctx context.Context, session *Session, channelID Snowflake) error {
	endpoint := EndpointChannel(channelID.String()) + "/thread-members/@me"

	err := session.Interface.FetchJJ(ctx, session, http.MethodDelete, endpoint, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to leave thread: %w", err)
	}

	return nil
}

// AddThreadMember adds a user to a thread.
func AddThreadMember(ctx context.Context, session *Session, channelID, userID Snowflake) error {
	endpoint := EndpointChannel(channelID.String()) + "/thread-members/" + userID.String()

	err := session.Interface.FetchJJ(ctx, session, http.MethodPut, endpoint, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to add thread member: %w", err)
	}

	return nil
}

// RemoveThreadMember removes a user from a thread.
func RemoveThreadMember(ctx context.Context, session *Session, channelID, userID Snowflake) error {
	endpoint := EndpointChannel(channelID.String()) + "/thread-members/" + userID.String()

	err := session.Interface.FetchJJ(ctx, session, http.MethodDelete, endpoint, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to remove thread member: %w", err)
	}

	return nil
}

// GetThreadMember retrieves information about a user in a thread.
func GetThreadMember(ctx context.Context, session *Session, channelID, userID Snowflake) (*ThreadMember, error) {
	endpoint := EndpointChannel(channelID.String()) + "/thread-members/" + userID.String()

	var threadMember *ThreadMember

	err := session.Interface.FetchJJ(ctx, session, http.MethodGet, endpoint, nil, nil, &threadMember)
	if err != nil {
		return nil, fmt.Errorf("failed to get thread member: %w", err)
	}

	return threadMember, nil
}

// ListThreadMembers lists all members in a thread.
func ListThreadMembers(ctx context.Context, session *Session, channelID Snowflake, after *Snowflake, limit *int32) ([]ThreadMember, error) {
	endpoint := EndpointChannel(channelID.String()) + "/thread-members"

	values := url.Values{}

	if after != nil {
		values.Add("after", after.String())
	}

	if limit != nil {
		values.Add("limit", strconv.FormatInt(int64(*limit), 10))
	}

	if len(values) > 0 {
		endpoint += "?" + values.Encode()
	}

	var threadMembers []ThreadMember

	err := session.Interface.FetchJJ(ctx, session, http.MethodGet, endpoint, nil, nil, &threadMembers)
	if err != nil {
		return nil, fmt.Errorf("failed to list thread members: %w", err)
	}

	return threadMembers, nil
}

// ListPublicArchivedThreads lists archived public threads in a channel.
func ListPublicArchivedThreads(ctx context.Context, session *Session, channelID Snowflake, before *string, limit *int32) (interface{}, error) {
	endpoint := EndpointChannel(channelID.String()) + "/threads/archived/public"

	values := url.Values{}

	if before != nil {
		values.Add("before", *before)
	}

	if limit != nil {
		values.Add("limit", strconv.FormatInt(int64(*limit), 10))
	}

	if len(values) > 0 {
		endpoint += "?" + values.Encode()
	}

	var threads interface{}

	err := session.Interface.FetchJJ(ctx, session, http.MethodGet, endpoint, nil, nil, &threads)
	if err != nil {
		return nil, fmt.Errorf("failed to list public archived threads: %w", err)
	}

	return threads, nil
}

// ListPrivateArchivedThreads lists archived private threads in a channel.
func ListPrivateArchivedThreads(ctx context.Context, session *Session, channelID Snowflake, before *string, limit *int32) (interface{}, error) {
	endpoint := EndpointChannel(channelID.String()) + "/threads/archived/private"

	values := url.Values{}

	if before != nil {
		values.Add("before", *before)
	}

	if limit != nil {
		values.Add("limit", strconv.FormatInt(int64(*limit), 10))
	}

	if len(values) > 0 {
		endpoint += "?" + values.Encode()
	}

	var threads interface{}

	err := session.Interface.FetchJJ(ctx, session, http.MethodGet, endpoint, nil, nil, &threads)
	if err != nil {
		return nil, fmt.Errorf("failed to list private archived threads: %w", err)
	}

	return threads, nil
}

// ListJoinedPrivateArchivedThreads lists archived private threads the current user is a member of.
func ListJoinedPrivateArchivedThreads(ctx context.Context, session *Session, channelID Snowflake, before *string, limit *int32) (interface{}, error) {
	endpoint := EndpointChannel(channelID.String()) + "/users/@me/threads/archived/private"

	values := url.Values{}

	if before != nil {
		values.Add("before", *before)
	}

	if limit != nil {
		values.Add("limit", strconv.FormatInt(int64(*limit), 10))
	}

	if len(values) > 0 {
		endpoint += "?" + values.Encode()
	}

	var threads interface{}

	err := session.Interface.FetchJJ(ctx, session, http.MethodGet, endpoint, nil, nil, &threads)
	if err != nil {
		return nil, fmt.Errorf("failed to list joined private archived threads: %w", err)
	}

	return threads, nil
}

// GroupDMAddRecipient adds a user to a group DM.
func GroupDMAddRecipient(ctx context.Context, session *Session, channelID, userID Snowflake, params GroupDMAddRecipientParams) error {
	endpoint := EndpointChannel(channelID.String()) + "/recipients/" + userID.String()

	err := session.Interface.FetchJJ(ctx, session, http.MethodPut, endpoint, params, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to add group dm recipient: %w", err)
	}

	return nil
}

// GroupDMRemoveRecipient removes a user from a group DM.
func GroupDMRemoveRecipient(ctx context.Context, session *Session, channelID, userID Snowflake) error {
	endpoint := EndpointChannel(channelID.String()) + "/recipients/" + userID.String()

	err := session.Interface.FetchJJ(ctx, session, http.MethodDelete, endpoint, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to remove group dm recipient: %w", err)
	}

	return nil
}
