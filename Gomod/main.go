package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

/*go.sum -> at time pulling repo , any changes in repo correspond to our fail this clone
to assure security while cloning

Whenever pkg is installed ->
it installed in go folder original, cashe folder of go
Other time it get fetched from local not google


Some expensive Commands (Keep an Eye How they react)

go mod tidy -> tide all packages that we are using
go mod verify -> verify modules
go list -m all -> all package that application is dependent on
go list all -> all package that has been installed
go list  -> only main package
go list -m -versions github.com/gorilla/mux -> How many versions available
got mod why go list -m -versions github.com/gorilla/mux -> Why dependant on these module
go mod graph -> dependancy graph top to bottom (a dependent on b)
go mod edit -go 1.16 -> Change go version
go mod edit -module 1.16 -> Change module name
go mod vendor -> like node modules locally install
*/

func greet() {
	fmt.Println("Greeting from Mod")
}
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Sahil Shile</h1>"))
}
func main() {
	fmt.Println("Hello mod in golang")
	greet()
	// Routing in GO using gorilla mux
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	log.Fatal(http.ListenAndServe(":4000", r))
}
