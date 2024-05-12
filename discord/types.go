package discord

import (
	"bytes"
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

func (s *Snowflake) UnmarshalJSON(b []byte) error {
	if !bytes.Equal(b, null) {
		i, err := strconv.ParseInt(string(b[1:len(b)-1]), 10, 64)
		if err != nil {
			return fmt.Errorf("failed to unmarshal json: %v", err)
		}

		*s = Snowflake(i)
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
	if !bytes.Equal(b, null) {
		i, err := strconv.ParseInt(string(b[1:len(b)-1]), 10, 64)
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
