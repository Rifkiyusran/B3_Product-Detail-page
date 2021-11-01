package main

import (
	"encoding/json"
	_ "expvar"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Product struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Price     int    `json:"price"`
	Deskripsi string `json:"deskripsi"`
	Quantity  int    `json:"quantity"`
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Selamat datang di home page")
}

func allProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(Products)
}

func singleProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	id := params["id"]

	for _, product := range Products {
		if product.ID == id {
			json.NewEncoder(w).Encode(product)
			return
		}
	}
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var product Product

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	Products = append(Products, product)
	json.NewEncoder(w).Encode(product)
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	id := params["id"]

	var product Product

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for i, p := range Products {
		if p.ID == id {
			Products[i].Title = product.Title
			Products[i].Price = product.Price
			Products[i].Quantity = product.Quantity
			Products[i].Deskripsi = product.Deskripsi
			json.NewEncoder(w).Encode(Products[i])
			return
		}
	}
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	id := params["id"]

	for i, p := range Products {
		if p.ID == id {
			Products = append(Products[:i], Products[i+1:]...)
			json.NewEncoder(w).Encode(p)
			return
		}
	}
}

func handleRequest() {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/", home)
	r.HandleFunc("/products", allProducts).Methods("GET")
	r.HandleFunc("/products/{id}", singleProduct).Methods("GET")
	r.HandleFunc("/products", createProduct).Methods("POST")
	r.HandleFunc("/products/{id}", updateProduct).Methods("PUT")
	r.HandleFunc("/products/{id}", deleteProduct).Methods("DELETE")

	fmt.Println("Application running")
	//log.Fatal(http.ListenAndServe(":8080", r))
	log.Fatal(http.ListenAndServe(":8080", r))
}

func main() {
	Products = []Product{
		Product{ID: "100", Title: "Rockider MT 500", Price: 199000, Quantity: 20, Deskripsi: "Helm ini berfungsi untuk melindungi kepala dari benturan apabila mengalami kecelakaan "},
		Product{ID: "200", Title: "Rockbow WT-09", Price: 199000, Quantity: 20, Deskripsi: "Helm ini berfungsi untuk melindungi kepala dari benturan apabila mengalami kecelakaan "},
		Product{ID: "300", Title: "Rockbros WT-09", Price: 199000, Quantity: 20, Deskripsi: "Helm ini berfungsi untuk melindungi kepala dari benturan apabila mengalami kecelakaan "},
		Product{ID: "400", Title: "Rockbros BIke Googles", Price: 199000, Quantity: 20, Deskripsi: "Helm ini berfungsi untuk melindungi kepala dari benturan apabila mengalami kecelakaan "},
	}

	handleRequest()
}

// Products is a global variable to hold collection of products
var Products []Product

// func main() {

// 	logger := log.NewLogfmtLogger(os.Stdout)

// 	transport.RegisterHttpsServicesAndStartListener()

// 	port := os.Getenv("PORT")
// 	if port == "" {
// 		port = "8080"
// 	}

// 	logger.Log("listening-on", port)
// 	if err := http.ListenAndServe(":"+port, nil); err != nil {
// 		logger.Log("listen.error", err)
// 	}
// }
