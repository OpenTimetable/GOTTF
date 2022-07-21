package objects

import "time"

type Span struct {
	From time.Time `json:"from"`
	To   time.Time `json:"to"`
}
