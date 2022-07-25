package objects

type Metadata struct {
  Version string `json:"version"`
  Timezone string `json:"timezone"`
  Author string `json:"author"`
  Timestamp uint64 `json:"timestamp"`
}
