package parsers

import (
	"encoding/json"
)

type TimetableParser interface {
	ParseV1(string) error
	ComposeV1() (string, error)
}

func (t *Timetable) ParseV1(s string) error {
	err := json.Unmarshal([]byte(s), &t)
	return err
}

func (t *Timetable) ComposeV1() (string, error) {
	res, err := json.Marshal(t)
	return string(res), err
}
