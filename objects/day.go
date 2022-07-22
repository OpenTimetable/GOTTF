package objects

type Day struct {
	Classes   map[string][]Class `json:"classes"`
	Events    []Event            `json:"events"`
	DayEvents []DayEvent         `json:"dayevents"`
}
