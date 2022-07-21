package parsers

import "github.com/OpenTimetable/GOTTF/objects"

type Timetable struct {
	Version string        `json:"version"`
	Cues    objects.Cues  `json:"cues"`
	Days    []objects.Day `json:"days"`
}