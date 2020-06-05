package main

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/willingtonortiz/knn-restapi/dtos"
	"github.com/willingtonortiz/knn-restapi/logic"
	"github.com/willingtonortiz/knn-restapi/models"
)

var elements models.Elements

func readData() models.Elements {
	file, err := os.Open("data/input.csv")

	if err != nil {
		log.Fatalln("No se pudo abrir el archivo", err)
	}

	r := csv.NewReader(file)

	elements := make([]models.Element, 0)

	for {
		record, err := r.Read()

		if err == io.EOF {
			break
		}

		// RECORD => ID, DATA[4] ,CLASS

		recordSize := len(record)
		components := make([]float64, recordSize-2)

		id, err := strconv.Atoi(record[0])

		for i := 0; i < recordSize-2; i++ {
			components[i], err = strconv.ParseFloat(record[i+1], 64)
		}

		class, err := strconv.Atoi(record[recordSize-1])

		elements = append(elements, models.Element{
			ID:         id,
			Components: components,
			Class:      class,
		})
	}

	return elements
}

func testRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to KNN API!")
}

// Retorna todos los items
func getAllItems(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(elements)
}

// Ejecute knn
func executeKnn(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Ocurrio un error en el servidor")
	}

	var request dtos.KnnRequest
	json.Unmarshal(reqBody, &request)

	// Seleccion de algoritmo
	algoID := request.Algorithm
	algorithm, err := selectAlgorithm(algoID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Creando item
	element := models.Element{Components: request.Components}

	knn := logic.Knn{}
	a, b, c := knn.FindKNearest(element, request.K, elements, algorithm)

	response := dtos.KMeansResponse{
		Element:         a,
		NearestElements: b,
		Elements:        c,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func selectAlgorithm(algoID int) (logic.DistanceCalculator, error) {
	if algoID == 1 {
		return logic.EuclideanDistanceCalculator{}, nil
	} else if algoID == 2 {
		return logic.ManhattanDistanceCalculator{}, nil
	} else {
		return nil, errors.New("Algoritmo no encontrado")
	}
}

func initialize() {
	elements = readData()
}

func startServer() {
	router := mux.NewRouter().StrictSlash(true)

	// Declaración de las rutas para el api
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Accept", "Accept-Language", "Content-Type", "Content-Language", "Origin"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})

	router.HandleFunc("/", testRoute)
	router.HandleFunc("/knn", executeKnn).Methods("POST")
	router.HandleFunc("/items", getAllItems).Methods("GET")

	fmt.Println("Server started on port 8080")

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}

func main() {
	// Leer los datos del archivo
	initialize()

	// Iniciar el servidor
	startServer()
}
