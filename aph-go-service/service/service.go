package service

import (
	"aph-go-service/config"
	"aph-go-service/datastruct"
	"database/sql"
	fmt "fmt"
	"log"
)

func HelloWorld(name string) string {
	var helloOutput string
	helloOutput = fmt.Sprintf("Hello, %s ", name)
	return helloOutput
}

//product parameternya ya
func POSTproduct(product datastruct.Tbl_Product) int64 {
	db := config.CreateConnection()

	defer db.Close()

	sqlStatement := `INSERT INTO tbl_product (product_id, nama_product, image_product, price_product, deskripsi_product, rating_product) VALUES ($1, $2, $3, $4, $5,$6)`

	var Product_id int64

	err := db.QueryRow(sqlStatement, product.Product_id, product.Nama_product, product.Image_product, product.Price_product, product.Deskripsi_product, product.Rating_product).Scan(&Product_id)

	if err != nil {
		log.Fatalf("Tidak bisa mengeksekusi query. %v", err)
	}

	fmt.Printf("insert data single record %v", Product_id)

	return Product_id
}

//mengambil semua data list product
func GETlistproduct() ([]datastruct.Tbl_Product, error) {
	db := config.CreateConnection()

	defer db.Close()

	var products []datastruct.Tbl_Product

	sqlStatement := `SELECT * FROM tbl_product`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Tidak bisa mengeksekusi Query %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var product datastruct.Tbl_Product
		err = rows.Scan(&product.Product_id, &product.Nama_product, &product.Image_product, &product.Price_product, &product.Deskripsi_product, &product.Rating_product)

		if err != nil {
			log.Fatalf("tidak bisa mengambil data %v", err)
		}

		products = append(products, product)
	}

	return products, err
}

//mengambil data product persatuan
func GETproduct(Product_id int64) (datastruct.Tbl_Product, error) {
	db := config.CreateConnection()

	defer db.Close()

	var product datastruct.Tbl_Product

	sqlStatement := `SELECT * FROM tbl_product WHERE product_id=$1`

	row := db.QueryRow(sqlStatement, product.Product_id)

	err := row.Scan(&product.Product_id, &product.Nama_product, &product.Image_product, &product.Price_product, &product.Deskripsi_product, &product.Rating_product)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("Tidak ada data yang dicari!")
		return product, nil
	case nil:
		return product, nil
	default:
		log.Fatalf("tidak bisa mengambil data %v", err)
	}

	return product, err
}

//Post cart
func POSTcart(cart datastruct.Tbl_Cart) int64 {
	db := config.CreateConnection()

	defer db.Close()

	sqlStatement := `INSERT INTO tbl_cart (cart_id, product_id, qty) VALUES ($1, $2, $3)`

	var Cart_id int64

	err := db.QueryRow(sqlStatement, cart.Cart_id, cart.Product_id, cart.Qty).Scan(&Cart_id)

	if err != nil {
		log.Fatalf("Tidak bisa mengeksekusi query. %v", err)
	}

	fmt.Printf("insert data single record %v", Cart_id)

	return Cart_id
}

//mengambil semua data list product
func GETlistcart() ([]datastruct.Tbl_Cart, error) {
	db := config.CreateConnection()

	defer db.Close()

	var list_cart []datastruct.Tbl_Cart

	sqlStatement := `SELECT * FROM tbl_cart`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Tidak bisa mengeksekusi Query %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var cart datastruct.Tbl_Cart
		err = rows.Scan(&cart.Cart_id, &cart.Product_id, &cart.Qty)

		if err != nil {
			log.Fatalf("tidak bisa mengambil data %v", err)
		}

		list_cart = append(list_cart, cart)
	}

	return list_cart, err
}

//mengambil data artikel per-id
func GETcart(Cart_id int64) (datastruct.Tbl_Cart, error) {
	db := config.CreateConnection()

	defer db.Close()

	var cart datastruct.Tbl_Cart

	sqlStatement := `SELECT * FROM tbl_cart WHERE cart_id=$1`

	row := db.QueryRow(sqlStatement, cart.Cart_id)

	err := row.Scan(&cart.Cart_id, &cart.Product_id, &cart.Qty)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("Tidak ada data yang dicari!")
		return cart, nil
	case nil:
		return cart, nil
	default:
		log.Fatalf("tidak bisa mengambil data %v", err)
	}

	return cart, err
}