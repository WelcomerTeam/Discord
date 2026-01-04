package discord

import (
	"context"
	"fmt"
	"net/http"
)

// CreateStageInstanceParams represents the parameters for creating a stage instance.
type CreateStageInstanceParams struct {
	Topic                 string                    `json:"topic"`
	ChannelID             Snowflake                 `json:"channel_id"`
	PrivacyLevel          *StageChannelPrivacyLevel `json:"privacy_level,omitempty"`
	GuildScheduledEventID *Snowflake                `json:"guild_scheduled_event_id,omitempty"`
}

// ModifyStageInstanceParams represents the parameters for modifying a stage instance.
type ModifyStageInstanceParams struct {
	Topic        *string                   `json:"topic,omitempty"`
	PrivacyLevel *StageChannelPrivacyLevel `json:"privacy_level,omitempty"`
}

func CreateStageInstance(ctx context.Context, session *Session, params CreateStageInstanceParams, reason *string) (*StageInstance, error) {
	endpoint := "/stage-instances"

	headers := http.Header{}
	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	var stageInstance *StageInstance

	err := session.Interface.FetchJJ(ctx, session, http.MethodPost, endpoint, params, headers, &stageInstance)
	if err != nil {
		return nil, fmt.Errorf("failed to create stage instance: %w", err)
	}

	return stageInstance, nil
}

func GetStageInstance(ctx context.Context, session *Session, channelID Snowflake) (*StageInstance, error) {
	endpoint := "/stage-instances/" + channelID.String()

	var stageInstance *StageInstance

	err := session.Interface.FetchJJ(ctx, session, http.MethodGet, endpoint, nil, nil, &stageInstance)
	if err != nil {
		return nil, fmt.Errorf("failed to get stage instance: %w", err)
	}

	return stageInstance, nil
}

func ModifyStageInstance(ctx context.Context, session *Session, channelID Snowflake, params ModifyStageInstanceParams, reason *string) (*StageInstance, error) {
	endpoint := "/stage-instances/" + channelID.String()

	headers := http.Header{}
	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	var stageInstance *StageInstance

	err := session.Interface.FetchJJ(ctx, session, http.MethodPatch, endpoint, params, headers, &stageInstance)
	if err != nil {
		return nil, fmt.Errorf("failed to modify stage instance: %w", err)
	}

	return stageInstance, nil
}

func DeleteStageInstance(ctx context.Context, session *Session, channelID Snowflake, reason *string) error {
	endpoint := "/stage-instances/" + channelID.String()

	headers := http.Header{}
	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err := session.Interface.FetchJJ(ctx, session, http.MethodDelete, endpoint, nil, headers, nil)
	if err != nil {
		return fmt.Errorf("failed to delete stage instance: %w", err)
	}

	return nil
}
