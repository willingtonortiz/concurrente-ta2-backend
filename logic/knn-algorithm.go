package logic

import (
	"sort"

	"github.com/willingtonortiz/knn-restapi/models"
)

// Knn is ...
type Knn struct{}

// FindKNearest is ...
func (knn Knn) FindKNearest(
	it models.Element,
	k int,
	elements []models.Element,
	calc DistanceCalculator) (models.Element, []models.Element, []models.Element) {

	size := len(elements)

	currentElement := models.Point{
		Components: it.Components,
	}

	for i := 0; i < size; i++ {
		point := models.Point{
			Components: elements[i].Components,
		}
		elements[i].Distance = calc.Calculate(currentElement, point)
	}

	// for i := 0; i < size; i++ {
	// 	fmt.Printf("%f ", elements[i].Distance)
	// }

	sort.Sort(models.DistanceComparer(elements))

	// fmt.Println()

	// for i := 0; i < size; i++ {
	// 	fmt.Printf("%f ", items[i].Distance)
	// }

	/* ===== Encontrar los k elementos y las clases ===== */
	nearestElements := make([]models.Element, k)
	counters := make([]int, k)

	for i := 0; i < k; i++ {
		nearestElements[i] = elements[i]
		counters[elements[i].Class]++
	}

	class := 0
	max := 0
	for i := 0; i < k; i++ {
		if max < counters[i] {
			max = counters[i]
			class = i
		}

	}

	it.Class = class
	// fmt.Printf("\nClass: %d", it.Class)

	return it, nearestElements, elements
}
