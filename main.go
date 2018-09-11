package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Friend struct {
	ID          string `json:"id,omitempty"`
	Email       string `json:"email,omitempty"`
	Friends     string `json:"friends,omitempty"`
	Blocks      string `json:"blocks,omitempty"`
	Subscribers string `json:"subscribers,omitempty"`
}

var friends []Friend

func GetFriendsAPIStatus(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(friends)
}

func main() {
	router := mux.NewRouter()
	friends = append(friends, Friend{ID: "1", Email: "John", Friends: "Doe", Blocks: "Blocks", Subscribers: "Subscribed"})
	router.HandleFunc("/friends", GetFriendsAPIStatus).Methods("GET")
	log.Fatal(http.ListenAndServe(":8670", router))
}
