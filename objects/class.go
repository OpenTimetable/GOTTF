package objects

type Class struct {
	Substitution bool     `json:"substitution"`
	Examination  bool     `json:"examination"`
	Canceled     bool     `json:"canceled"`
	Name         string   `json:"name"`
	Abbreviation string   `json:"abbreviation"`
	Location     string   `json:"location"`
	Hosts        []string `json:"hosts"`
}
