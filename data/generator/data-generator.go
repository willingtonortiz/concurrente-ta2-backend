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

	for i := 0; i < 100; i++ {
		row := fmt.Sprintf("%d", i) + ","

		for j := 0; j < 5; j++ {

			if j == 4 {
				val := rand.Int() % 3
				row += fmt.Sprintf("%d", val) + "\n"
			} else {

				val := rand.Float64() * 100
				row += fmt.Sprintf("%f", val) + ","
			}
		}

		fo.WriteString(row)
	}

	fo.Close()
}
