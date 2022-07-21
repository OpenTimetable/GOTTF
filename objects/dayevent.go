package objects

type DayEvent struct {
	Title    string   `json:"title"`
	Location string   `json:"location"`
	Hosts    []string `json:"hosts"`
}
