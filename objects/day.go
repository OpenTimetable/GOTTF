package objects

type Day struct {
	Date      string           `json:"date"`
	Classes   map[string]Class `json:"classes"`
	Events    []Event          `json:"events"`
	DayEvents []DayEvent       `json:"day_events"`
}
