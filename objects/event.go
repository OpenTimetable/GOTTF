package objects

type Event struct {
	From     string   `json:"from"`
	To       string   `json:"to"`
	Title    string   `json:"title"`
	Location string   `json:"location"`
	Hosts    []string `json:"hosts"`
}
