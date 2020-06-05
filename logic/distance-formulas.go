package logic

import (
	"fmt"
	"math"

	"github.com/willingtonortiz/knn-restapi/models"
)

// Distance interface

// DistanceCalculator is ...
type DistanceCalculator interface {
	Calculate(a, b models.Point) float64
}

// Manhattan distance

// EuclideanDistanceCalculator is ...
type EuclideanDistanceCalculator struct{}

// Calculate is ...
func (ed EuclideanDistanceCalculator) Calculate(a, b models.Point) float64 {
	size := len(a.Components)

	sum := 0.0

	for i := 0; i < size; i++ {
		sum += math.Pow(a.Components[i]-b.Components[i], 2)
	}

	return math.Sqrt(sum)
}

func testEuclideanDistance() {
	var calculator DistanceCalculator = EuclideanDistanceCalculator{}

	a := models.Point{[]float64{0, 0}}
	b := models.Point{[]float64{3, 4}}

	c := calculator.Calculate(a, b)

	if c == 5 {
		fmt.Println("Success")
	} else {

		fmt.Println("Failed")
	}
}

// Manhattan distance

// ManhattanDistanceCalculator is ...
type ManhattanDistanceCalculator struct{}

// Calculate is ...
func (md ManhattanDistanceCalculator) Calculate(a, b models.Point) float64 {
	size := len(a.Components)

	sum := 0.0

	for i := 0; i < size; i++ {
		sum += math.Abs(a.Components[i] - b.Components[i])
	}

	return sum
}

func testManhattanDistance() {
	var calculator DistanceCalculator = ManhattanDistanceCalculator{}

	a := models.Point{[]float64{0, 0}}
	b := models.Point{[]float64{3, 4}}

	c := calculator.Calculate(a, b)

	if c == 7 {
		fmt.Println("Success")
	} else {

		fmt.Println("Failed")
	}
}
