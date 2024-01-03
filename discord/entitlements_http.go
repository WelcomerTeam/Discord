package discord

import (
	"fmt"
	"net/http"
)

func ListEntitlements(s *Session, applicationID Snowflake) ([]*Entitlement, error) {
	endpoint := EndpointEntitlements(applicationID.String())

	var entitlements []*Entitlement

	err := s.Interface.FetchJJ(s, http.MethodGet, endpoint, nil, nil, &entitlements)
	if err != nil {
		return nil, fmt.Errorf("failed to list entitlements: %w", err)
	}

	return entitlements, nil
}

func CreateTestEntitlement(s *Session, applicationID Snowflake, entitlementParams EntitlementParams) (*Entitlement, error) {
	endpoint := EndpointEntitlements(applicationID.String())

	headers := http.Header{}

	var entitlement *Entitlement

	err := s.Interface.FetchJJ(s, http.MethodPost, endpoint, entitlementParams, headers, &entitlement)
	if err != nil {
		return nil, fmt.Errorf("failed to create test entitlement: %w", err)
	}

	return entitlement, nil
}

func DeleteTestEntitlement(s *Session, applicationID Snowflake, entitlementID Snowflake) error {
	endpoint := EndpointEntitlement(applicationID.String(), entitlementID.String())

	err := s.Interface.FetchJJ(s, http.MethodDelete, endpoint, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete test entitlement: %w", err)
	}

	return nil
}