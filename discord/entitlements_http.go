package discord

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// ListEntitlements returns a list of entitlements for an application.
func ListEntitlements(ctx context.Context, s *Session, applicationID Snowflake, userID *Snowflake, skuIDs []Snowflake, before, after *Snowflake, limit *Int64, guildID *Snowflake, exludeEnded *bool, exludeDeleted *bool) ([]Entitlement, error) {
	endpoint := EndpointEntitlements(applicationID.String())

	values := url.Values{}

	if userID != nil {
		values.Set("user_id", userID.String())
	}

	if len(skuIDs) > 0 {
		skuIDStrings := make([]string, len(skuIDs))

		for i, skuID := range skuIDs {
			skuIDStrings[i] = skuID.String()
		}

		values.Set("sku_ids", strings.Join(skuIDStrings, ","))
	}

	if before != nil {
		values.Set("before", before.String())
	}

	if after != nil {
		values.Set("after", after.String())
	}

	if limit != nil {
		values.Set("limit", strconv.Itoa(int(*limit)))
	}

	if guildID != nil {
		values.Set("guild_id", guildID.String())
	}

	if exludeEnded != nil {
		values.Set("exclude_ended", strconv.FormatBool(*exludeEnded))
	}

	if exludeDeleted != nil {
		values.Set("exclude_deleted", strconv.FormatBool(*exludeDeleted))
	}

	if len(values) > 0 {
		endpoint += "?" + values.Encode()
	}

	var entitlements []Entitlement

	err := s.Interface.FetchJJ(ctx, s, http.MethodGet, endpoint, nil, nil, &entitlements)
	if err != nil {
		return nil, fmt.Errorf("failed to list entitlements: %w", err)
	}

	return entitlements, nil
}

func CreateTestEntitlement(ctx context.Context, s *Session, applicationID Snowflake, entitlementParams EntitlementParams) (*Entitlement, error) {
	endpoint := EndpointEntitlements(applicationID.String())

	headers := http.Header{}

	var entitlement *Entitlement

	err := s.Interface.FetchJJ(ctx, s, http.MethodPost, endpoint, entitlementParams, headers, &entitlement)
	if err != nil {
		return nil, fmt.Errorf("failed to create test entitlement: %w", err)
	}

	return entitlement, nil
}

func DeleteTestEntitlement(ctx context.Context, s *Session, applicationID Snowflake, entitlementID Snowflake) error {
	endpoint := EndpointEntitlement(applicationID.String(), entitlementID.String())

	err := s.Interface.FetchJJ(ctx, s, http.MethodDelete, endpoint, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete test entitlement: %w", err)
	}

	return nil
}
