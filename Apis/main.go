package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// CRUD = Create, Read, Update, Delete in GO
// Model for coarse - file
type Coarse struct {
	CoarseID          string  `json:"id"`
	CoarseName        string  `json:"name"`
	CoarseDescription string  `json:"description"`
	CoarsePrice       string  `json:"price"`  //`json:"-"` hide price
	CoarseAuthor      *Author `json:"author"` // Making a reference
}
type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

// Fake db
var coarsesList []Coarse
var authorList []Author

// middleware helpers
func (c *Coarse) isEmpty() bool {
	// return c.CoarseID == "" && c.CoarseName == ""
	return c.CoarseName == ""
}

func getAllCoarse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("All coarses")
	w.Header().Set("Content-Type", "application/json") // set that data is json
	json.NewEncoder(w).Encode(coarsesList)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	// grab that course id
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json") // set that data in json

	params := mux.Vars(r) // get passed params
	// loop though courses and search id
	for _, course := range coarsesList {
		if course.CoarseID == params["id"] {
			json.NewEncoder(w).Encode(course) // return response
			return
		}
	}
	json.NewEncoder(w).Encode("No Course found with corresponding id")
}

func CreateCoarse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json") // set that data in json

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send a request body")
		return
	}
	var coarse Coarse
	json.NewDecoder(r.Body).Decode(&coarse) // decoding in struct format

	//encode -> convert to json
	//decode -> convert to struct

	if coarse.isEmpty() {
		json.NewEncoder(w).Encode("No data send in json")
		return
	}

	//We have the req body now in struct format
	// generate random number

	rand.Seed(time.Now().UnixNano())
	coarse.CoarseID = strconv.Itoa(rand.Intn(100))
	coarsesList = append(coarsesList, coarse)
	json.NewEncoder(w).Encode(coarse)
	return
}

func updateCoarse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json") // set that data in json
	params := mux.Vars(r)

	fmt.Println(params)
	for index, course := range coarsesList {
		fmt.Println(course.CoarseID, params["id"])
		if course.CoarseID == params["id"] {
			coarsesList = append(coarsesList[:index], coarsesList[index+1:]...) // remove index id
			var newcourse Coarse
			json.NewDecoder(r.Body).Decode(&newcourse)
			fmt.Println(newcourse)
			newcourse.CoarseID = params["id"]
			coarsesList = append(coarsesList, newcourse)
			json.NewEncoder(w).Encode(newcourse)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course found with corresponding id to update")
}
func deleteCoarse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json") // set that data in json
	params := mux.Vars(r)

	for index, course := range coarsesList {
		if course.CoarseID == params["id"] {
			coarsesList = append(coarsesList[:index], coarsesList[index+1:]...) // remove index id
			break
		}
	}
	json.NewEncoder(w).Encode("Course Deleted Successfully")
}

// PERFECT

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to the home page</h1>")) // send page instead data
}

func main() {
	fmt.Println("Routing for  CRUD operations")
	r := mux.NewRouter()

	// initializing some data to list

	coarse := Coarse{
		CoarseID:          "1",
		CoarseName:        "Go",
		CoarseDescription: "Go programming language",
		CoarsePrice:       "100",
		CoarseAuthor: &Author{
			FullName: "SahilShile",
			Website:  "sahilshile25.scom",
		},
	}
	coarsesList = append(coarsesList, coarse)

	// routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/coarse/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/courses", getAllCoarse).Methods("GET")

	r.HandleFunc("/create", CreateCoarse).Methods("POST")
	r.HandleFunc("/update/{id}", updateCoarse).Methods("PUT")
	r.HandleFunc("/delete/{id}", deleteCoarse).Methods("DELETE")
	// listen to port
	log.Fatal(http.ListenAndServe(":4000", r))
}

// Connect Mongo DB
