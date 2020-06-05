package dtos

import (
	"encoding/json"

	"github.com/willingtonortiz/knn-restapi/models"
)

// KMeansResponse is ...
type KMeansResponse struct {
	Element         models.Element   `json:"element"`
	NearestElements []models.Element `json:"nearestElements"`
	Elements        []models.Element `json:"elements"`
}

func (kmr KMeansResponse) String() string {
	result, _ := json.Marshal(kmr)
	return string(string(result))
}
