package discord

import "time"

// AuthorizationInformation represents the current oauth authorization.
type AuthorizationInformation struct {
	Application *Application `json:"application"`
	Scopes      []string     `json:"scopes"`
	Expires     time.Time    `json:"expires"`
	User        *User        `json:"user"`
}
