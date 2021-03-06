package discord

import (
	jsoniter "github.com/json-iterator/go"
)

// gateway.go contains all structures for interacting with discord's gateway and contains
// all events and structures we send to

// GatewayOp represents the operation codes of a gateway message.
type GatewayOp uint8

const (
	GatewayOpDispatch GatewayOp = iota
	GatewayOpHeartbeat
	GatewayOpIdentify
	GatewayOpStatusUpdate
	GatewayOpVoiceStateUpdate
	_
	GatewayOpResume
	GatewayOpReconnect
	GatewayOpRequestGuildMembers
	GatewayOpInvalidSession
	GatewayOpHello
	GatewayOpHeartbeatACK
)

// IntentFlag represents a bitflag for intents.
type GatewayIntent uint32

const (
	IntentGuilds GatewayIntent = 1 << iota
	IntentGuildMembers
	IntentGuildBans
	IntentGuildEmojis
	IntentGuildIntegrations
	IntentGuildWebhooks
	IntentGuildInvites
	IntentGuildVoiceStates
	IntentGuildPresences
	IntentGuildMessages
	IntentGuildMessageReactions
	IntentGuildMessageTyping
	IntentDirectMessages
	IntentDirectMessageReactions
	IntentDirectMessageTyping
	IntentMessageContent
)

// Gateway close codes.
const (
	CloseUnknownError = 4000 + iota
	CloseUnknownOpCode
	CloseDecodeError
	CloseNotAuthenticated
	CloseAuthenticationFailed
	CloseAlreadyAuthenticated
	_
	CloseInvalidSeq
	CloseRateLimited
	CloseSessionTimeout
	CloseInvalidShard
	CloseShardingRequired
	CloseInvalidAPIVersion
	CloseInvalidIntents
	CloseDisallowedIntents
)

// GatewayPayload represents the base payload received from discord gateway.
type GatewayPayload struct {
	Op       GatewayOp           `json:"op"`
	Data     jsoniter.RawMessage `json:"d"`
	Sequence int32               `json:"s"`
	Type     string              `json:"t"`
}

// SentPayload represents the base payload we send to discords gateway.
type SentPayload struct {
	Op   GatewayOp   `json:"op"`
	Data interface{} `json:"d"`
}

// Gateway Commands

// Identify represents the initial handshake with the gateway.
type Identify struct {
	Token          string              `json:"token"`
	Properties     *IdentifyProperties `json:"properties"`
	Compress       bool                `json:"compress"`
	LargeThreshold int32               `json:"large_threshold"`
	Shard          [2]int32            `json:"shard,omitempty"`
	Presence       *UpdateStatus       `json:"presence,omitempty"`
	Intents        int32               `json:"intents"`
}

// IdentifyProperties are the extra properties sent in the identify packet.
type IdentifyProperties struct {
	OS      string `json:"$os"`
	Browser string `json:"$browser"`
	Device  string `json:"$device"`
}

// Resume resumes a dropped gateway connection.
type Resume struct {
	Token     string `json:"token"`
	SessionID string `json:"session_id"`
	Sequence  int32  `json:"seq"`
}

// Heartbeat represents the heartbeat packet.
type Heartbeat int

// Request guild members requests members for a guild.
type RequestGuildMembers struct {
	GuildID   Snowflake   `json:"guild_id"`
	Query     string      `json:"query"`
	Limit     int32       `json:"limit"`
	Presences bool        `json:"presences"`
	Nonce     string      `json:"nonce"`
	UserIDs   []Snowflake `json:"user_ids"`
}

// Update Presence updates a client's presence.
type UpdateStatus struct {
	Since      int32       `json:"since,omitempty"`
	Activities []*Activity `json:"activities,omitempty"`
	Status     string      `json:"status"`
	AFK        bool        `json:"afk"`
}
