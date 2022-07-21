package objects

type Cues struct {
	Periods  map[string]Span `json:"periods"`
	Recesses []Span          `json:"recesses"`
}
