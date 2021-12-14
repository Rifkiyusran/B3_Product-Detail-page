package datastruct

type HelloWorldRequest struct {
	NAME string `json:"name"`
}

type HelloWorldResponse struct {
	MESSAGE string `json:"message"`
}

//penulisan huruf awal harus besar agar menjadi publik (exported)
//penulisan huruf kecil berarti private (unexported)
type Tbl_Product struct {
	Product_id       	int64  		`json:"product_id"`
	Nama_product     	string 		`json:"nama_product"`
	Image_product 		string 		`json:"image_product"`
	Price_product 		string 		`json:"price_product"`
	Deskripsi_product   string 		`json:"deskripsi_product"`
	Rating_product   	string 		`json:"rating_product"`
}

type Tbl_Cart struct {
	Cart_id         int64  	`json:"cart_id"`
	Product_id     	int64  	`json:"product_id"`
	Qty     		int64	`json:"qty"`
}

type Response1 struct {
	Status   	string `json:"status"`
	Message 	string `json:"message,omitempty"`
	Product_id 	int64  `json:"product_id,omitempty"`
}

type Response2 struct {
	Status  string      	`json:"status"`
	Message string      	`json:"message"`
	Data    []Tbl_Product 	`json:"data"`
}

type Response3 struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    []Tbl_Cart 	`json:"data"`
}

// type Response5 struct {
// 	Status      string `json:"status"`
// 	Message     string `json:"message,omitempty"`
// 	Artikel_id2 int64  `json:"artikel2_id,omitempty"`
// }

type Response6 struct {
	Status   string `json:"status"`
	Message  string `json:"message,omitempty"`
	Cart_id  int64  `json:"cart_id,omitempty"`
}