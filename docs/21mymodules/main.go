package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// go get -u github.com/gorilla/mux
// downloaded module is located at
// ${GOPATH}/pkg/mod/github.com/cache/packagename it will using from local copy
const PORT = 4000
func main() {
	Start()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), r))
}

func Start() {
	fmt.Println("Server started listening on", PORT)
}

func serveHome(w http.ResponseWriter,r *http.Request){
	fmt.Println("Request received for home page")
	w.Write([]byte("<h1>Welcome to my home page</h1>"))
}