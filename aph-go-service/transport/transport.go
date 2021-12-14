package transport

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"aph-go-service/datastruct"
	"aph-go-service/logging"
	"aph-go-service/service"

	"github.com/gorilla/mux"
)

//product
func POSTproductEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var product datastruct.Tbl_Product

	err := json.NewDecoder(r.Body).Decode(&product)

	if err != nil {
		log.Fatalf("tidak bisa mendecode dari request body. %v", err)
	}

	insertProduct_id := service.POSTproduct(product)

	logging.Log(fmt.Sprintf("menambahkan product %d", product.Product_id))

	res := datastruct.Response1{
		Product_id: insertProduct_id,
		Status:   "Berhasil",
		Message:  "Data product telah ditambahkan",
	}

	json.NewEncoder(w).Encode(res)
}

func GETlistproductEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-WWW-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	products, err := service.GETlistproduct()

	if err != nil {
		log.Fatalf("Tidak bisa mengambil data. %v", err)
	}

	logging.Log(fmt.Sprintf("menampilkan seluruh product"))

	res := datastruct.Response2{
		Status:  "Berhasil",
		Message: "Data keseluruhan product",
		Data:    products,
	}

	json.NewEncoder(w).Encode(res)
}

func GETproductEndpoint(W http.ResponseWriter, r *http.Request) {
	W.Header().Set("Context-Type", "application/x-WWW-form-urlencoded")
	W.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)

	Product_id, err := strconv.Atoi(params["Product_id"])

	if err != nil {
		log.Fatalf("Tidak bisa mengubah dari string ke int. %v", err)
	}

	product, err := service.GETproduct(int64(Product_id))

	if err != nil {
		log.Fatalf("Tidak bisa mengambil data product. %v", err)
	}

	logging.Log(fmt.Sprintf("menampilkan product %d", product.Product_id))

	res := datastruct.Response1{
		Status:   "Berhasil",
		Message:  "Data product",
		Product_id: product.Product_id,
	}

	json.NewEncoder(W).Encode(res)
}

//carts
func POSTcartEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var cart datastruct.Tbl_Cart

	err := json.NewDecoder(r.Body).Decode(&cart)

	if err != nil {
		log.Fatalf("tidak bisa mendecode dari request body. %v", err)
	}

	insertCart_id := service.POSTcart(cart)

	logging.Log(fmt.Sprintf("menambahkan cart %d", cart.Cart_id))

	res := datastruct.Response6{
		Status:   "Berhasil",
		Cart_id: 	insertCart_id,
		Message:  "Data cart telah ditambahkan",
	}

	json.NewEncoder(w).Encode(res)
}

func GETlistcartEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-WWW-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	list_cart, err := service.GETlistcart()

	if err != nil {
		log.Fatalf("Tidak bisa mengambil data. %v", err)
	}

	logging.Log(fmt.Sprintf("menampilkan seluruh cart"))

	res := datastruct.Response3{
		Status:  "Berhasil",
		Message: "Data keseluruhan video",
		Data:    list_cart,
	}

	json.NewEncoder(w).Encode(res)
}

func GETcartEndpoint(W http.ResponseWriter, r *http.Request) {
	W.Header().Set("Context-Type", "application/x-WWW-form-urlencoded")
	W.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)

	Cart_id, err := strconv.Atoi(params["Cart_id"])

	if err != nil {
		log.Fatalf("Tidak bisa mengubah dari string ke int. %v", err)
	}

	cart, err := service.GETcart(int64(Cart_id))

	if err != nil {
		log.Fatalf("Tidak bisa mengambil data artikel %v", err)
	}

	logging.Log(fmt.Sprintf("menampilkan cart %d", cart.Cart_id))

	res := datastruct.Response6{
		Status:   "Berhasil",
		Cart_id: cart.Cart_id,
		Message:  "Data cart",
	}

	json.NewEncoder(W).Encode(res)
}