package discord

import (
	"context"
	"fmt"
	"net/http"
)

// ListAutoModerationRules returns a list of all auto-moderation rules for a guild.
func ListAutoModerationRules(ctx context.Context, session *Session, guildID Snowflake) ([]AutoModerationRule, error) {
	endpoint := EndpointGuildAutoModerationRules(guildID.String())

	var rules []AutoModerationRule

	err := session.Interface.FetchJJ(ctx, session, http.MethodGet, endpoint, nil, nil, &rules)
	if err != nil {
		return nil, fmt.Errorf("failed to list auto moderation rules: %w", err)
	}

	return rules, nil
}

// GetAutoModerationRule returns a specific auto-moderation rule for a guild.
func GetAutoModerationRule(ctx context.Context, session *Session, guildID, ruleID Snowflake) (*AutoModerationRule, error) {
	endpoint := EndpointGuildAutoModerationRule(guildID.String(), ruleID.String())

	var rule *AutoModerationRule

	err := session.Interface.FetchJJ(ctx, session, http.MethodGet, endpoint, nil, nil, &rule)
	if err != nil {
		return nil, fmt.Errorf("failed to get auto moderation rule: %w", err)
	}

	return rule, nil
}

// CreateAutoModerationRule creates a new auto-moderation rule for a guild.
func CreateAutoModerationRule(ctx context.Context, session *Session, guildID Snowflake, params AutoModerationRule, reason *string) (*AutoModerationRule, error) {
	endpoint := EndpointGuildAutoModerationRules(guildID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	var rule *AutoModerationRule

	err := session.Interface.FetchJJ(ctx, session, http.MethodPost, endpoint, params, headers, &rule)
	if err != nil {
		return nil, fmt.Errorf("failed to create auto moderation rule: %w", err)
	}

	return rule, nil
}

// ModifyAutoModerationRule modifies an existing auto-moderation rule for a guild.
func ModifyAutoModerationRule(ctx context.Context, session *Session, guildID, ruleID Snowflake, params AutoModerationRule, reason *string) (*AutoModerationRule, error) {
	endpoint := EndpointGuildAutoModerationRule(guildID.String(), ruleID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	var rule *AutoModerationRule

	err := session.Interface.FetchJJ(ctx, session, http.MethodPatch, endpoint, params, headers, &rule)
	if err != nil {
		return nil, fmt.Errorf("failed to modify auto moderation rule: %w", err)
	}

	return rule, nil
}

// DeleteAutoModerationRule deletes an auto-moderation rule from a guild.
func DeleteAutoModerationRule(ctx context.Context, session *Session, guildID, ruleID Snowflake, reason *string) error {
	endpoint := EndpointGuildAutoModerationRule(guildID.String(), ruleID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err := session.Interface.FetchJJ(ctx, session, http.MethodDelete, endpoint, nil, headers, nil)
	if err != nil {
		return fmt.Errorf("failed to delete auto moderation rule: %w", err)
	}

	return nil
}
