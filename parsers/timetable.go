package parsers

import "github.com/OpenTimetable/GOTTF/objects"

type Timetable struct {
	Metadata objects.Metadata `json:"metadata"`
	Cues    objects.Cues           `json:"cues"`
	Days    map[string]objects.Day `json:"days"`
}
