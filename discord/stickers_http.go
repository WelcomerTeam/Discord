package discord

import (
	"context"
	"fmt"
	"net/http"
)

// StickerParams represents the parameters for creating or modifying a sticker.
type StickerParams struct {
	Name        string  `json:"name"`
	Tags        string  `json:"tags"`
	Description *string `json:"description,omitempty"`
	Image       string  `json:"image"`
}

// ModifyStickerParams represents the parameters for modifying a sticker.
type ModifyStickerParams struct {
	Name        *string `json:"name,omitempty"`
	Tags        *string `json:"tags,omitempty"`
	Description *string `json:"description,omitempty"`
}

func ListNitroStickerPacks(ctx context.Context, session *Session) ([]interface{}, error) {
	endpoint := "/sticker-packs"

	var stickerPacks []interface{}

	err := session.Interface.FetchJJ(ctx, session, http.MethodGet, endpoint, nil, nil, &stickerPacks)
	if err != nil {
		return nil, fmt.Errorf("failed to list nitro sticker packs: %w", err)
	}

	return stickerPacks, nil
}

func ListGuildStickers(ctx context.Context, session *Session, guildID Snowflake) ([]Sticker, error) {
	endpoint := EndpointGuildStickers(guildID.String())

	var stickers []Sticker

	err := session.Interface.FetchJJ(ctx, session, http.MethodGet, endpoint, nil, nil, &stickers)
	if err != nil {
		return nil, fmt.Errorf("failed to list guild stickers: %w", err)
	}

	return stickers, nil
}

func GetGuildSticker(ctx context.Context, session *Session, guildID, stickerID Snowflake) (*Sticker, error) {
	endpoint := EndpointGuildSticker(guildID.String(), stickerID.String())

	var sticker *Sticker

	err := session.Interface.FetchJJ(ctx, session, http.MethodGet, endpoint, nil, nil, &sticker)
	if err != nil {
		return nil, fmt.Errorf("failed to get guild sticker: %w", err)
	}

	return sticker, nil
}

func CreateGuildSticker(ctx context.Context, session *Session, guildID Snowflake, params StickerParams, reason *string) (*Sticker, error) {
	endpoint := EndpointGuildStickers(guildID.String())

	headers := http.Header{}
	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	var sticker *Sticker

	err := session.Interface.FetchJJ(ctx, session, http.MethodPost, endpoint, params, headers, &sticker)
	if err != nil {
		return nil, fmt.Errorf("failed to create guild sticker: %w", err)
	}

	return sticker, nil
}

func ModifyGuildSticker(ctx context.Context, session *Session, guildID, stickerID Snowflake, params ModifyStickerParams, reason *string) (*Sticker, error) {
	endpoint := EndpointGuildSticker(guildID.String(), stickerID.String())

	headers := http.Header{}
	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	var sticker *Sticker

	err := session.Interface.FetchJJ(ctx, session, http.MethodPatch, endpoint, params, headers, &sticker)
	if err != nil {
		return nil, fmt.Errorf("failed to modify guild sticker: %w", err)
	}

	return sticker, nil
}

func DeleteGuildSticker(ctx context.Context, session *Session, guildID, stickerID Snowflake, reason *string) error {
	endpoint := EndpointGuildSticker(guildID.String(), stickerID.String())

	headers := http.Header{}
	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err := session.Interface.FetchJJ(ctx, session, http.MethodDelete, endpoint, nil, headers, nil)
	if err != nil {
		return fmt.Errorf("failed to delete guild sticker: %w", err)
	}

	return nil
}
