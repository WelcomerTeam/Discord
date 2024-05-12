package discord_test

import (
	"bytes"
	"math"
	"testing"
	"time"

	"github.com/WelcomerTeam/Discord/discord"
)

func TestSnowflake_MarshalJSON(t *testing.T) {
	snowflake := discord.Snowflake(1234567890)
	expected := []byte(`"1234567890"`)

	result, err := snowflake.MarshalJSON()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !bytes.Equal(result, expected) {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

func TestSnowflake_MarshalJSON_Empty(t *testing.T) {
	snowflake := discord.Snowflake(0)
	expected := []byte(`"0"`)

	result, err := snowflake.MarshalJSON()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !bytes.Equal(result, expected) {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

func TestSnowflake_MarshalJSON_MaxInt64(t *testing.T) {
	snowflake := discord.Snowflake(math.MaxInt64)
	expected := []byte(`"9223372036854775807"`)

	result, err := snowflake.MarshalJSON()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !bytes.Equal(result, expected) {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

func TestSnowflake_UnmarshalJSON(t *testing.T) {
	snowflake := discord.Snowflake(0)
	data := []byte(`"1234567890"`)
	expected := discord.Snowflake(1234567890)

	err := snowflake.UnmarshalJSON(data)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if snowflake != expected {
		t.Errorf("Expected %d, but got %d", expected, snowflake)
	}
}

func TestSnowflake_UnmarshalJSON_Null(t *testing.T) {
	snowflake := discord.Snowflake(1234567890)
	data := []byte(`null`)
	expected := discord.Snowflake(0)

	err := snowflake.UnmarshalJSON(data)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if snowflake != expected {
		t.Errorf("Expected %d, but got %d", expected, snowflake)
	}
}

func TestSnowflake_UnmarshalJSON_InvalidData(t *testing.T) {
	snowflake := discord.Snowflake(0)
	data := []byte(`"invalid"`)
	expected := discord.Snowflake(0)

	err := snowflake.UnmarshalJSON(data)
	if err == nil {
		t.Error("Expected an error, but got nil")
	}

	if snowflake != expected {
		t.Errorf("Expected %d, but got %d", expected, snowflake)
	}
}
func TestSnowflake_String(t *testing.T) {
	snowflake := discord.Snowflake(1234567890)
	expected := "1234567890"

	result := snowflake.String()

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

func TestSnowflake_String_Zero(t *testing.T) {
	snowflake := discord.Snowflake(0)
	expected := "0"

	result := snowflake.String()

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

func TestSnowflake_String_MaxInt64(t *testing.T) {
	snowflake := discord.Snowflake(math.MaxInt64)
	expected := "9223372036854775807"

	result := snowflake.String()

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
func TestSnowflake_Time(t *testing.T) {
	snowflake := discord.Snowflake(1234567890)
	expected := time.Unix(0, ((int64(snowflake)>>22)+discord.DiscordCreation)*1000000)

	result := snowflake.Time()

	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestSnowflake_Time_Zero(t *testing.T) {
	snowflake := discord.Snowflake(0)
	expected := time.Unix(0, ((int64(snowflake)>>22)+discord.DiscordCreation)*1000000)

	result := snowflake.Time()

	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestSnowflake_Time_MaxInt64(t *testing.T) {
	snowflake := discord.Snowflake(math.MaxInt64)
	expected := time.Unix(0, ((int64(snowflake)>>22)+discord.DiscordCreation)*1000000)

	result := snowflake.Time()

	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestSnowflake_Time_Fixed(t *testing.T) {
	snowflake := discord.Snowflake(143090142360371200)
	expected := time.Unix(1454185748, 425000000)

	result := snowflake.Time()

	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}

}
