package dtos

import (
	"encoding/json"
)

// KnnRequest is ...
type KnnRequest struct {
	Components []float64 `json:"components"`
	K          int       `json:"k"`
	Algorithm  int       `json:"algorithm"`
}

func (kmr KnnRequest) String() string {
	result, _ := json.Marshal(kmr)
	return string(string(result))
}
