package parsers

import (
	"encoding/json"
)

type TimetableParser interface {
	ParseV1(string) error
	Compose() (string, error)
}

func (t *Timetable) ParseV1(s string) error {
	err := json.Unmarshal([]byte(s), &t)
	if err != nil {
		return err
	}
	return nil
}

func (t *Timetable) Compose() (string, error) {
	return "", nil
}
