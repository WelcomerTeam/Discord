package discord

import (
	"context"
	"fmt"
	"net/http"
)

// ListDefaultSoundboardSounds returns a list of the default soundboard sounds.
func ListDefaultSoundboardSounds(ctx context.Context, session *Session) ([]SoundboardSound, error) {
	endpoint := EndpointDefaultSoundboardSounds

	var sounds []SoundboardSound

	err := session.Interface.FetchJJ(ctx, session, http.MethodGet, endpoint, nil, nil, &sounds)
	if err != nil {
		return nil, fmt.Errorf("failed to list default soundboard sounds: %w", err)
	}

	return sounds, nil
}

// ListGuildSoundboardSounds returns a list of all soundboard sounds in a guild.
func ListGuildSoundboardSounds(ctx context.Context, session *Session, guildID Snowflake) ([]SoundboardSound, error) {
	endpoint := EndpointGuildSoundboardSounds(guildID.String())

	var result struct {
		Items []SoundboardSound `json:"items"`
	}

	err := session.Interface.FetchJJ(ctx, session, http.MethodGet, endpoint, nil, nil, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to list guild soundboard sounds: %w", err)
	}

	return result.Items, nil
}

// GetGuildSoundboardSound returns a specific soundboard sound in a guild.
func GetGuildSoundboardSound(ctx context.Context, session *Session, guildID, soundID Snowflake) (*SoundboardSound, error) {
	endpoint := EndpointGuildSoundboardSound(guildID.String(), soundID.String())

	var sound *SoundboardSound

	err := session.Interface.FetchJJ(ctx, session, http.MethodGet, endpoint, nil, nil, &sound)
	if err != nil {
		return nil, fmt.Errorf("failed to get guild soundboard sound: %w", err)
	}

	return sound, nil
}

// CreateGuildSoundboardSound creates a new soundboard sound in a guild.
func CreateGuildSoundboardSound(ctx context.Context, session *Session, guildID Snowflake, params CreateSoundboardSoundParams, reason *string) (*SoundboardSound, error) {
	endpoint := EndpointGuildSoundboardSounds(guildID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	var sound *SoundboardSound

	err := session.Interface.FetchJJ(ctx, session, http.MethodPost, endpoint, params, headers, &sound)
	if err != nil {
		return nil, fmt.Errorf("failed to create guild soundboard sound: %w", err)
	}

	return sound, nil
}

// ModifyGuildSoundboardSound modifies an existing soundboard sound in a guild.
func ModifyGuildSoundboardSound(ctx context.Context, session *Session, guildID, soundID Snowflake, params ModifySoundboardSoundParams, reason *string) (*SoundboardSound, error) {
	endpoint := EndpointGuildSoundboardSound(guildID.String(), soundID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	var sound *SoundboardSound

	err := session.Interface.FetchJJ(ctx, session, http.MethodPatch, endpoint, params, headers, &sound)
	if err != nil {
		return nil, fmt.Errorf("failed to modify guild soundboard sound: %w", err)
	}

	return sound, nil
}

// DeleteGuildSoundboardSound deletes a soundboard sound from a guild.
func DeleteGuildSoundboardSound(ctx context.Context, session *Session, guildID, soundID Snowflake, reason *string) error {
	endpoint := EndpointGuildSoundboardSound(guildID.String(), soundID.String())

	headers := http.Header{}

	if reason != nil {
		headers.Add(AuditLogReasonHeader, *reason)
	}

	err := session.Interface.FetchJJ(ctx, session, http.MethodDelete, endpoint, nil, headers, nil)
	if err != nil {
		return fmt.Errorf("failed to delete guild soundboard sound: %w", err)
	}

	return nil
}
