package gottf_test

import (
	"io/ioutil"
	"testing"

	gottf "github.com/OpenTimetable/GOTTF"
	"github.com/nsf/jsondiff"
)

func TestV1(t *testing.T) {
	t.Log("Testing timetable parsing with version 1.0")
	t.Log("-----------------------------------------")

	file, err := ioutil.ReadFile("example-1.0.ottf")
	if err != nil {
		t.Errorf("Failed to read example timetable: %s", err)
	}

	v1_0_timetable := string(file)
	// t.Logf("Parsing timetable: %s", v1_0_timetable)

	timetable, err := gottf.ParseTimetable(v1_0_timetable)
	if err != nil {
		t.Errorf("Error parsing timetable: %s", err)
	}

	composed, err := gottf.ComposeTimetable(timetable, "1.0")
	if err != nil {
		t.Errorf("Error composing timetable: %s", err)
	}

	diffOpts := jsondiff.DefaultConsoleOptions()
	res, diff := jsondiff.Compare([]byte(v1_0_timetable), []byte(composed), &diffOpts)

	if res != jsondiff.FullMatch {
		t.Errorf("the expected result is not equal to what we have: %s", diff)
	}

	t.Logf("Timetable: %+v", timetable)
	t.Logf("Composed timetable: %s", composed)
	t.Log("-----------------------------------------")
}
