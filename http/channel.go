package http

import (
	"context"
	"net/http"

	"github.com/WelcomerTeam/Discord/discord"
	"github.com/WelcomerTeam/Discord/structs"
	"golang.org/x/xerrors"
)

func (s *Session) GetChannel(ctx context.Context, channelID discord.Snowflake) (channel *structs.Channel, err error) {
	endpoint := EndpointChannel(channelID.String())

	err = s.Interface.FetchJJ(ctx, http.MethodGet, endpoint, nil, nil, &channel)
	if err != nil {
		return nil, xerrors.Errorf("Failed to get channel: %v", err)
	}

	return
}

func (s *Session) ModifyChannel(ctx context.Context, channelID discord.Snowflake, channelArg structs.Channel, reason *string) (channel *structs.Channel, err error) {
	endpoint := EndpointChannel(channelID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(structs.AuditLogReasonHeader, *reason)
	}

	err = s.Interface.FetchJJ(ctx, http.MethodGet, endpoint, channelArg, headers, &channel)
	if err != nil {
		return nil, xerrors.Errorf("Failed to modify channel: %v", err)
	}

	return
}

func (s *Session) DeleteChannel(ctx context.Context, channelID discord.Snowflake, reason *string) (err error) {
	endpoint := EndpointChannel(channelID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(structs.AuditLogReasonHeader, *reason)
	}

	err = s.Interface.FetchJJ(ctx, http.MethodGet, endpoint, nil, headers, nil)
	if err != nil {
		return xerrors.Errorf("Failed to delete channel: %v", err)
	}

	return
}

func (s *Session) GetChannelMessages(ctx context.Context, channelID discord.Snowflake) (messages []*structs.Message, err error) {
	endpoint := EndpointChannelMessages(channelID.String())

	err = s.Interface.FetchJJ(ctx, http.MethodGet, endpoint, nil, nil, &messages)
	if err != nil {
		return nil, xerrors.Errorf("Failed to get channel messages: %v", err)
	}

	return
}

func (s *Session) GetChannelMessage(ctx context.Context, channelID discord.Snowflake, messageID discord.Snowflake) (message *structs.Message, err error) {
	endpoint := EndpointChannelMessage(channelID.String(), messageID.String())

	err = s.Interface.FetchJJ(ctx, http.MethodGet, endpoint, nil, nil, &message)
	if err != nil {
		return nil, xerrors.Errorf("Failed to get channel message: %v", err)
	}

	return
}

func (s *Session) CreateMessage(ctx context.Context, channelID discord.Snowflake, messageArg structs.Message) (message *structs.Message, err error) {
	endpoint := EndpointChannelMessages(channelID.String())

	err = s.Interface.FetchJJ(ctx, http.MethodPost, endpoint, messageArg, nil, &message)
	if err != nil {
		return nil, xerrors.Errorf("Failed to create message: %v", err)
	}

	return
}

func (s *Session) CrosspostMessage(ctx context.Context, channelID discord.Snowflake, messageID discord.Snowflake) (message *structs.Message, err error) {
	endpoint := EndpointChannelMessageCrosspost(channelID.String(), messageID.String())

	err = s.Interface.FetchJJ(ctx, http.MethodPost, endpoint, nil, nil, &message)
	if err != nil {
		return nil, xerrors.Errorf("Failed to crosspost message: %v", err)
	}

	return
}

func (s *Session) CreateReaction(ctx context.Context, channelID discord.Snowflake, messageID discord.Snowflake, emoji string) (err error) {
	endpoint := EndpointMessageReaction(channelID.String(), messageID.String(), emoji, "@me")

	err = s.Interface.FetchJJ(ctx, http.MethodPut, endpoint, nil, nil, nil)
	if err != nil {
		return xerrors.Errorf("Failed to create reaction: %v", err)
	}

	return
}

func (s *Session) DeleteOwnReaction(ctx context.Context, channelID discord.Snowflake, messageID discord.Snowflake, emoji string) (err error) {
	endpoint := EndpointMessageReaction(channelID.String(), messageID.String(), emoji, "@me")

	err = s.Interface.FetchJJ(ctx, http.MethodDelete, endpoint, nil, nil, nil)
	if err != nil {
		return xerrors.Errorf("Failed to delete own reaction: %v", err)
	}

	return
}

func (s *Session) DeleteUserReaction(ctx context.Context, channelID discord.Snowflake, messageID discord.Snowflake, emoji string, userID discord.Snowflake) (err error) {
	endpoint := EndpointMessageReaction(channelID.String(), messageID.String(), emoji, userID.String())

	err = s.Interface.FetchJJ(ctx, http.MethodDelete, endpoint, nil, nil, nil)
	if err != nil {
		return xerrors.Errorf("Failed to delete user reaction: %v", err)
	}

	return
}

func (s *Session) GetReactions(ctx context.Context, channelID discord.Snowflake, messageID discord.Snowflake, emoji string, after *discord.Snowflake, limit *int) (users []*structs.User, err error) {
	endpoint := EndpointMessageReactions(channelID.String(), messageID.String(), emoji)

	err = s.Interface.FetchJJ(ctx, http.MethodGet, endpoint, nil, nil, &users)
	if err != nil {
		return nil, xerrors.Errorf("Failed to get reactions: %v", err)
	}

	return
}

func (s *Session) DeleteAllReactions(ctx context.Context, channelID discord.Snowflake, messageID discord.Snowflake) (err error) {
	endpoint := EndpointMessageReactionsAll(channelID.String(), messageID.String())

	err = s.Interface.FetchJJ(ctx, http.MethodDelete, endpoint, nil, nil, nil)
	if err != nil {
		return xerrors.Errorf("Failed to delete all reactions: %v", err)
	}

	return
}

func (s *Session) DeleteAllReactionsEmoji(ctx context.Context, channelID discord.Snowflake, messageID discord.Snowflake, emoji string) (err error) {
	endpoint := EndpointMessageReactions(channelID.String(), messageID.String(), emoji)

	err = s.Interface.FetchJJ(ctx, http.MethodDelete, endpoint, nil, nil, nil)
	if err != nil {
		return xerrors.Errorf("Failed to delete all reactions emoji: %v", err)
	}

	return
}

func (s *Session) EditMessage(ctx context.Context, channelID discord.Snowflake, messageID discord.Snowflake, messageArg structs.Message) (message *structs.Message, err error) {
	endpoint := EndpointChannelMessage(channelID.String(), messageID.String())

	err = s.Interface.FetchJJ(ctx, http.MethodPatch, endpoint, nil, nil, &message)
	if err != nil {
		return nil, xerrors.Errorf("Failed to edit message: %v", err)
	}

	return
}

func (s *Session) DeleteMessage(ctx context.Context, channelID discord.Snowflake, messageID discord.Snowflake, reason *string) (err error) {
	endpoint := EndpointChannelMessage(channelID.String(), messageID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(structs.AuditLogReasonHeader, *reason)
	}

	err = s.Interface.FetchJJ(ctx, http.MethodDelete, endpoint, nil, headers, nil)
	if err != nil {
		return xerrors.Errorf("Failed to delet emessage: %v", err)
	}

	return
}

func (s *Session) BulkDeleteMessages(ctx context.Context, channelID discord.Snowflake, messageIDs []discord.Snowflake, reason *string) (err error) {
	endpoint := EndpointChannelMessagesBulkDelete(channelID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(structs.AuditLogReasonHeader, *reason)
	}

	err = s.Interface.FetchJJ(ctx, http.MethodDelete, endpoint, nil, headers, nil)
	if err != nil {
		return xerrors.Errorf("Failed to bulk delete messages: %v", err)
	}

	return
}

func (s *Session) EditChannelPermissions(ctx context.Context, channelID discord.Snowflake, overwriteID discord.Snowflake, overwriteArg structs.ChannelOverwrite, reason *string) (err error) {
	endpoint := EndpointChannelPermission(channelID.String(), overwriteID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(structs.AuditLogReasonHeader, *reason)
	}

	err = s.Interface.FetchJJ(ctx, http.MethodPut, endpoint, nil, headers, nil)
	if err != nil {
		return xerrors.Errorf("Failed to edit channel permissions: %v", err)
	}

	return
}

func (s *Session) GetChannelInvites(ctx context.Context, channelID discord.Snowflake) (invites []*structs.Invite, err error) {
	endpoint := EndpointChannelInvites(channelID.String())

	err = s.Interface.FetchJJ(ctx, http.MethodGet, endpoint, nil, nil, nil)
	if err != nil {
		return nil, xerrors.Errorf("Failed to get channel invites: %v", err)
	}

	return
}

func (s *Session) CreateChannelInvite(ctx context.Context, channelID discord.Snowflake, inviteParams structs.InviteParams, reason *string) (invite *structs.Invite, err error) {
	endpoint := EndpointChannelInvites(channelID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(structs.AuditLogReasonHeader, *reason)
	}

	err = s.Interface.FetchJJ(ctx, http.MethodPost, endpoint, nil, headers, &invite)
	if err != nil {
		return nil, xerrors.Errorf("Failed to create channel invite: %v", err)
	}

	return
}

func (s *Session) DeleteChannelPermission(ctx context.Context, channelID discord.Snowflake, overwriteID discord.Snowflake, reason *string) (err error) {
	endpoint := EndpointChannelPermission(channelID.String(), overwriteID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(structs.AuditLogReasonHeader, *reason)
	}

	err = s.Interface.FetchJJ(ctx, http.MethodGet, endpoint, nil, headers, nil)
	if err != nil {
		return xerrors.Errorf("Failed to delete channel permission: %v", err)
	}

	return
}

func (s *Session) FollowNewsChannel(ctx context.Context, channelID discord.Snowflake, webhookChannelID discord.Snowflake) (followedChannel *structs.FollowedChannel, err error) {
	endpoint := EndpointChannelFollow(channelID.String())

	err = s.Interface.FetchJJ(ctx, http.MethodPost, endpoint, nil, nil, &followedChannel)
	if err != nil {
		return nil, xerrors.Errorf("Failed to follow news channel: %v", err)
	}

	return
}

func (s *Session) TriggerTypingIndicator(ctx context.Context, channelID discord.Snowflake) (err error) {
	endpoint := EndpointChannelTyping(channelID.String())

	err = s.Interface.FetchJJ(ctx, http.MethodPost, endpoint, nil, nil, nil)
	if err != nil {
		return xerrors.Errorf("Failed to trigger typing indicator: %v", err)
	}

	return
}

func (s *Session) GetPinnedMessages(ctx context.Context, channelID discord.Snowflake) (pinnedMessages []*structs.Message, err error) {
	endpoint := EndpointChannelMessagesPins(channelID.String())

	err = s.Interface.FetchJJ(ctx, http.MethodGet, endpoint, nil, nil, &pinnedMessages)
	if err != nil {
		return nil, xerrors.Errorf("Failed to get pinned messages: %v", err)
	}

	return
}

func (s *Session) PinMessage(ctx context.Context, channelID discord.Snowflake, messageID discord.Snowflake, reason *string) (err error) {
	endpoint := EndpointChannelMessagePin(channelID.String(), messageID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(structs.AuditLogReasonHeader, *reason)
	}

	err = s.Interface.FetchJJ(ctx, http.MethodPut, endpoint, nil, headers, nil)
	if err != nil {
		return xerrors.Errorf("Failed to pin message: %v", err)
	}

	return
}

func (s *Session) UnpinMessage(ctx context.Context, channelID discord.Snowflake, messageID discord.Snowflake, reason *string) (err error) {
	endpoint := EndpointChannelMessagePin(channelID.String(), messageID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(structs.AuditLogReasonHeader, *reason)
	}

	err = s.Interface.FetchJJ(ctx, http.MethodDelete, endpoint, nil, headers, nil)
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
// TODO: ListActiveThreads
// TODO: ListPublicArchivedThreads
// TODO: ListPrivateArchivedThreads
// TODO: ListJoinedPrivateArchivedThreads
