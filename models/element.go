package models

// Element ...
type Element struct {
	ID         int       `json:"id"`
	Components []float64 `json:"components"`
	Class      int       `json:"class"`
	Distance   float64   `json:"distance"`
}

// Elements ...
type Elements []Element

// DistanceComparer ...
type DistanceComparer Elements

func (a DistanceComparer) Len() int           { return len(a) }
func (a DistanceComparer) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a DistanceComparer) Less(i, j int) bool { return a[i].Distance < a[j].Distance }
