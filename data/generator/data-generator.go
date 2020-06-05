package generator

import (
	"fmt"
	"math/rand"
	"os"
)

// GenerateData ...
func GenerateData() {
	fo, err := os.Create("data/input.csv")
	if err != nil {
		fmt.Println("Error al abrir el archivo")
	}

	rowsCount := 10000
	dataCount := 5
	classCount := 9
	maxNumber := 1000000.0

	for i := 0; i < rowsCount; i++ {
		row := fmt.Sprintf("%d", i) + ","

		for j := 0; j < dataCount+1; j++ {

			if j == dataCount {
				class := rand.Int() % classCount
				row += fmt.Sprintf("%d", class) + "\n"
			} else {

				val := rand.Float64() * maxNumber
				row += fmt.Sprintf("%f", val) + ","
			}
		}

		fo.WriteString(row)
	}

	fo.Close()
}
