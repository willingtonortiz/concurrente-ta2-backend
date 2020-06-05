package models

// Item is ...
type Item struct {
	ID       int     `json:"id"`
	Point    Point   `json:"point"`
	Class    int     `json:"class"`
	Distance float64 `json:"distance"`
}

// Items is ...
type Items []Item

// ByDistance is ...
type ByDistance []Item

func (a ByDistance) Len() int           { return len(a) }
func (a ByDistance) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByDistance) Less(i, j int) bool { return a[i].Distance < a[j].Distance }
