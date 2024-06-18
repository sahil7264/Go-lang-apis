package main

import (
	"fmt"
	"main/routes"
	"net/http"
)

func main() {
	fmt.Println("Server is starting soon...")
	r := routes.Router()
	http.ListenAndServe(":4000", r)
	fmt.Println("Listening to port :", 4000)
}
