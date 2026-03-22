package discord

import (
	"context"
	"fmt"
	"net/http"
)

// GetGateway returns the gateway URL for connecting to Discord.
func GetGateway(ctx context.Context, session *Session) (*Gateway, error) {
	endpoint := EndpointGateway

	var gateway *Gateway

	err := session.Interface.FetchJJ(ctx, session, http.MethodGet, endpoint, nil, nil, &gateway)
	if err != nil {
		return nil, fmt.Errorf("failed to get gateway: %w", err)
	}

	return gateway, nil
}

// GetGatewayBot returns gateway connection information for a bot.
func GetGatewayBot(ctx context.Context, session *Session) (*GatewayBot, error) {
	endpoint := EndpointGatewayBot

	var gatewayBot *GatewayBot

	err := session.Interface.FetchJJ(ctx, session, http.MethodGet, endpoint, nil, nil, &gatewayBot)
	if err != nil {
		return nil, fmt.Errorf("failed to get gateway bot: %w", err)
	}

	return gatewayBot, nil
}
