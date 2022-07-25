package gottf

import (
	"encoding/json"
	"fmt"

	"github.com/OpenTimetable/GOTTF/parsers"
)

func getVersionFromString(s string) string {
	var result map[string]interface{}
	err := json.Unmarshal([]byte(s), &result)
	if err != nil {
		return ""
	}

	meta := result["metadata"].(map[string]interface{})

	return meta["version"].(string)
}

func ParseTimetable(s string) (parsers.Timetable, error) {
	version := getVersionFromString(s)
	result := parsers.Timetable{}
	var err error

	switch version {
	case "1.0":
		err = result.ParseV1(s)
	default:
		err = fmt.Errorf("unsupported timetable version: %s", version)
	}

	return result, err
}

func ComposeTimetable(timetable parsers.Timetable, version string) (string, error) {
	switch version {
	case "1.0":
		return timetable.ComposeV1()
	}

	return "", fmt.Errorf("unsupported timetable version: %s", version)
}
