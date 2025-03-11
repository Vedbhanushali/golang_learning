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

// go mod tidy - will remove unused dependencies from go.mod file
// go mod verify - will verify the dependencies in go.mod file
// go list -m all - will list all the dependencies in the project
// go mod why github.com/gorilla/mux - will list the reason why this dependency is used in the project
// go mod graph - will list the dependencies in the project in graph format
// "go mod edit -go 1.16" - will edit the go.mod file (in certain case changing through terminal on server) to change go version  -module to change module
// go mod vendor - will create a vendor folder in the project and copy all the dependencies in the vendor folder (can run entire project without internet)
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