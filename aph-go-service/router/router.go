package router

import (
	"aph-go-service/transport"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	//product
	router.HandleFunc("/api/product", transport.GETlistproductEndpoint).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/product/{Product_id}", transport.GETproductEndpoint).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/product", transport.POSTproductEndpoint).Methods("POST", "OPTIONS")

	//carts
	router.HandleFunc("/api/cart", transport.GETlistcartEndpoint).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/cart/{Cart_id}", transport.GETcartEndpoint).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/cart", transport.POSTcartEndpoint).Methods("POST", "OPTIONS")

	return router
}