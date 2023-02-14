package main

import (
	"database/sql"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello World")
	simpleWebService()
}

func homeHandler(response http.ResponseWriter, request *http.Request){
	response.Write([]byte("My first simple API in Go Lang !"))
}

func simpleWebService(){
	http.HandleFunc("/",homeHandler)
	http.ListenAndServe(":3000",nil)
}

func webServiceWithEcho(){
	product := Product{
		ID: "1",
		Name: "Product 1",
		Price: 100.00,
	}

	saveProduct(product)
}

func saveProduct(product Product){
	db, err := sql.Open("sqlite3","./db.sqlite")

	if err !=nil{
		panic(err)
	}

	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO products(id, name, price) VALUES(?, ?, ?)")

	if err != nil{
		panic(err)
	}

	_,err = stmt.Exec(product.ID, product.Name,product.Price)

	if err != nil{
		panic(err)
	}
}

