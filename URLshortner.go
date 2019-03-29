package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Url struct {
	Title string `json:"title"`
}

func Geturl(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for key, item := range url {
		if item == params["id"] {

			json.NewEncoder(w).Encode(key)
			return
		}
	}
}

func Createurl(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	temp := rand.Intn(30)
	s := strconv.Itoa(temp)
	x := Url{"http://localhost:8000/" + s}
	url[params["id"]] = x
	json.NewEncoder(w).Encode(x)
}

var url map[string]Url

func main() {
	rand.Seed(time.Now().UnixNano())
	router := mux.NewRouter()
	router.HandleFunc("/url/{id}", Geturl).Methods("GET")
	router.HandleFunc("/url/{id}", Createurl).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))

}
