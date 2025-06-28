package discord

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

const (
	DiscordCreation = 1420070400000
)

var null = []byte("null")

// Placeholder type for easy identification.
type Snowflake int64

func (s *Snowflake) IsNil() bool {
	return *s == 0
}

func (s *Snowflake) UnmarshalJSON(b []byte) error {
	if !bytes.Equal(b, null) {
		if len(b) >= 2 && b[0] == '"' {
			i, err := strconv.ParseInt(string(b[1:len(b)-1]), 10, 64)
			if err != nil {
				return fmt.Errorf("failed to unmarshal json: %v", err)
			}

			*s = Snowflake(i)
		} else {
			i, err := strconv.ParseInt(string(b), 10, 64)
			if err != nil {
				return fmt.Errorf("failed to unmarshal json: %v", err)
			}

			*s = Snowflake(i)
		}
	} else {
		*s = 0
	}

	return nil
}

func (s Snowflake) MarshalJSON() ([]byte, error) {
	return int64ToStringBytes(int64(s)), nil
}

func (s Snowflake) String() string {
	return strconv.FormatInt(int64(s), 10)
}

// Time returns the creation time of the Snowflake.
func (s Snowflake) Time() time.Time {
	nsec := (int64(s) >> 22) + DiscordCreation

	return time.Unix(0, nsec*1000000)
}

// int64 to allow for marshalling support.
type Int64 int64

func (in *Int64) UnmarshalJSON(b []byte) error {
	if b[0] == '"' {
		i, err := strconv.ParseInt(string(b[1:len(b)-1]), 10, 64)
		if err != nil {
			return fmt.Errorf("failed to unmarshal json: %v", err)
		}

		*in = Int64(i)
	} else {
		i, err := strconv.ParseInt(string(b), 10, 64)
		if err != nil {
			return fmt.Errorf("failed to unmarshal json: %v", err)
		}

		*in = Int64(i)
	}

	return nil
}

func (in Int64) MarshalJSON() ([]byte, error) {
	return int64ToStringBytes(int64(in)), nil
}

func (in Int64) String() string {
	return strconv.FormatInt(int64(in), 10)
}

func int64ToStringBytes(s int64) []byte {
	buf := make([]byte, 0, 24) // maxInt64JsonLength + 2

	buf = append(buf, '"')
	buf = strconv.AppendInt(buf, s, 10)
	buf = append(buf, '"')

	return buf
}

// Ensure empty lists are marshalled as []
type List[T any] []T

func (l List[T]) MarshalJSON() ([]byte, error) {
	if len(l) == 0 {
		return []byte("[]"), nil
	}

	return json.Marshal([]T(l))
}

type SnowflakeList = List[Snowflake]
type StringList = List[string]
type Int64List = List[Int64]
type StageInstanceList = List[StageInstance]
type StickerList = List[Sticker]
type ScheduledEventList = List[ScheduledEvent]
type RoleList = List[Role]
type EmojiList = List[Emoji]
type VoiceStateList = List[VoiceState]
type GuildMemberList = List[GuildMember]
type ChannelList = List[Channel]
type ActivityList = List[Activity]
type PresenceUpdateList = List[PresenceUpdate]
type ChannelOverwriteList = List[ChannelOverwrite]
type UserList = List[User]
type AuditLogEntryList = List[AuditLogEntry]
type AuditLogChangesList = List[AuditLogChanges]
type IntegrationList = List[Integration]
type WebhookList = List[Webhook]
type EmbedFieldList = List[EmbedField]
type EmbedList = List[Embed]
type UnavailableGuildList = List[UnavailableGuild]
type ThreadMemberList = List[ThreadMember]
type ApplicationCommandOptionList = List[ApplicationCommandOption]
type ApplicationCommandPermissionsList = List[ApplicationCommandPermissions]
type ChannelTypeList = List[ChannelType]
type PresenceStatusList = List[PresenceStatus]
type MessageReactionList = List[MessageReaction]
type MessageAttachmentList = List[MessageAttachment]
