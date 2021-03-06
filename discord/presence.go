package discord

// PresenceStatus represents a presence's status.
type PresenceStatus string

// Presence statuses.
const (
	PresenceStatusIdle    PresenceStatus = "idle"
	PresenceStatusDND     PresenceStatus = "dnd"
	PresenceStatusOnline  PresenceStatus = "online"
	PresenceStatusOffline PresenceStatus = "offline"
)

// ActivityType represents an activity's type.
type ActivityType int

// Activity types.
const (
	ActivityTypeGame ActivityType = iota
	ActivityTypeStreaming
	ActivityTypeListening
)

// ActivityFlag represents an activity's flags.
type ActivityFlag int

// Activity flags.
const (
	ActivityFlagInstance ActivityFlag = 1 << iota
	ActivityFlagJoin
	ActivityFlagSpectate
	ActivityFlagJoinRequest
	ActivityFlagSync
	ActivityFlagPlay
)

// Activity represents an activity as sent as part of other packets.
type Activity struct {
	Name          string        `json:"name"`
	Type          ActivityType  `json:"type"`
	URL           string        `json:"url"`
	Timestamps    *Timestamps   `json:"timestamps,omitempty"`
	ApplicationID *Snowflake    `json:"application_id"`
	Details       string        `json:"details"`
	State         string        `json:"state"`
	Party         *Party        `json:"party,omitempty"`
	Assets        *Assets       `json:"assets,omitempty"`
	Secrets       *Secrets      `json:"secrets,omitempty"`
	Instance      bool          `json:"instance"`
	Flags         *ActivityFlag `json:"flags,omitempty"`
}

// Timestamps represents the starting and ending timestamp of an activity.
type Timestamps struct {
	Start int32 `json:"start"`
	End   int32 `json:"end"`
}

// Party represents an activity's current party information.
type Party struct {
	ID   string  `json:"id,omitempty"`
	Size []int32 `json:"size,omitempty"`
}

// Assets represents an activity's images and their hover texts.
type Assets struct {
	LargeImage string `json:"large_image"`
	LargeText  string `json:"large_text"`
	SmallImage string `json:"small_image"`
	SmallText  string `json:"small_text"`
}

// Secrets represents an activity's secrets for Rich Presence joining and spectating.
type Secrets struct {
	Join     string `json:"join"`
	Spectate string `json:"spectate"`
	Match    string `json:"match"`
}

// ClientStatus represent's the status of a client.
type ClientStatus struct {
	Desktop string `json:"desktop"`
	Mobile  string `json:"mobile"`
	Web     string `json:"web"`
}
