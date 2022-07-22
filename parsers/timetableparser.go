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
	if err != nil {
		return err
	}
	return nil
}

func (t *Timetable) ComposeV1() (string, error) {
	res, err := json.Marshal(t)
	if err != nil {
		return string(res), err
	}
	return string(res), nil
}
