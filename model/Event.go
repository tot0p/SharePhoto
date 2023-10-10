package model

type Event struct {
	UUID          string    `json:"uuid"`
	StartDateTime string    `json:"startDateTime"`
	EndDateTime   string    `json:"endDateTime"`
	Picture       []Picture `json:"picture,omitempty"`
}
