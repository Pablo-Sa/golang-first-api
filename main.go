package main

import (
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

}

