package main

import (
	"encoding/json"

	"log"

	"net/http"

	"github.com/gorilla/mux"
)

//ID ....
type ID struct {
	namee string
	numb  int
}

var ids []ID

func main() {
	// abc := ID{namee: "Suyash", numb: 8400660624}
	// ids = append(ids, abc)
	r := mux.NewRouter()
	r.HandleFunc("/display", getAll).Methods("GET")
	r.HandleFunc("/put", addOne).Methods("POST")
	r.HandleFunc("/post", updateOne).Methods("PUT")
	r.HandleFunc("/delete", deleteOne).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":9999", r))
}

func getAll(w http.ResponseWriter, r *http.Request) {
	abc := ID{namee: "Suyash", numb: 8400660624}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(abc)
}

func addOne(w http.ResponseWriter, r *http.Request) {

}

func updateOne(w http.ResponseWriter, r *http.Request) {

}

func deleteOne(w http.ResponseWriter, r *http.Request) {

}
