package objects

import "time"

type Event struct {
	From     time.Time `json:"from"`
	To       time.Time `json:"to"`
	Title    string    `json:"title"`
	Location string    `json:"location"`
	Hosts    []string  `json:"hosts"`
}
