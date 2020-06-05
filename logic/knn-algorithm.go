package logic

import (
	"fmt"
	"sort"
	"time"

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

	// Cantidad de cores de mi PC
	var coresCount int = 8
	end := make(chan bool)

	//Utilización de goroutines para hacer los cálculos
	startTime := time.Now()
	for i := 0; i < coresCount; i++ {
		go func(id int) {
			for j := id; j < size; j += coresCount {
				point := models.Point{Components: elements[j].Components}
				elements[j].Distance = calc.Calculate(currentElement, point)
			}
			end <- true
		}(i)
	}

	for i := 0; i < coresCount; i++ {
		<-end
	}
	printElapsedTime(startTime)

	/* Sin utilizar goroutines */
	//startTime := time.Now()
	//for j := 0; j < size; j++ {
	//	point := models.Point{Components: elements[j].Components}
	//	elements[j].Distance = calc.Calculate(currentElement, point)
	//}
	//elapsed := time.Since(startTime)
	//fmt.Printf("Tiempo de ejecución: %s\n", elapsed.String())

	// Ordenando los elementos
	sort.Sort(models.DistanceComparer(elements))

	/* ===== Encontrar los k elementos y las clases ===== */
	classesCount := 9
	nearestElements := make([]models.Element, k)
	counters := make([]int, classesCount)

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

func printElapsedTime(startTime time.Time) {
	elapsed := time.Since(startTime)

	fmt.Printf("segundos=%f, milisegundos=%d, nanosegundos=%d\n", elapsed.Seconds(), elapsed.Milliseconds(), elapsed.Nanoseconds())
}
