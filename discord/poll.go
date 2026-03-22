package discord

import "time"

// poll.go contains all structures for Discord polls.

// PollLayoutType represents the layout type of a poll.
type PollLayoutType uint16

const (
	PollLayoutTypeDefault PollLayoutType = 1
)

// Poll represents a poll attached to a message.
type Poll struct {
	Question         PollMedia      `json:"question"`
	Answers          []PollAnswer   `json:"answers"`
	Expiry           time.Time      `json:"expiry"`
	Results          *PollResults   `json:"results,omitempty"`
	LayoutType       PollLayoutType `json:"layout_type"`
	AllowMultiselect bool           `json:"allow_multiselect"`
}

// PollAnswer represents an answer in a poll.
type PollAnswer struct {
	PollMedia PollMedia `json:"poll_media"`
	AnswerID  int32     `json:"answer_id"`
}

// PollMedia represents the media content for a poll question or answer.
type PollMedia struct {
	Emoji *Emoji `json:"emoji,omitempty"`
	Text  string `json:"text,omitempty"`
}

// PollResults represents the results of a poll.
type PollResults struct {
	AnswerCounts []PollResultsEntry `json:"answer_counts"`
	IsFinalized  bool               `json:"is_finalized"`
}

// PollResultsEntry represents the vote count for a single poll answer.
type PollResultsEntry struct {
	ID      int32 `json:"id"`
	Count   int32 `json:"count"`
	MeVoted bool  `json:"me_voted"`
}
