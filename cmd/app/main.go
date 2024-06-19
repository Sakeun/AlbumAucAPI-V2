package main

import (
	"encoding/json"
	"github.com/Sakeun/AlbumAucAPI-V2/cmd/db"
	"log"
	"net/http"
)

var Albums []Album

type Album struct {
	Id         int    `json:"id"`
	SellerId   int    `json:"sellerId"`
	Name       string `json:"name"`
	Genre      string `json:"genre"`
	Condition  string `json:"condition"`
	EndingTime string `json:"endingTime"`
	IsDone     bool   `json:"isDone"`
	Bids       []Bid  `json:"bids"`
}

type Bid struct {
	Id        int     `json:"id"`
	UserId    int     `json:"userId"`
	Amount    float64 `json:"amount"`
	BidPlaced string  `json:"bidPlaced"`
}

func handleRequests() {
	http.HandleFunc("/apiAuc/ViewAuc/getAllAuctions", returnAlbums)
	http.HandleFunc("/getUser", getOneUser)
	log.Fatal(http.ListenAndServe(":7044", nil))
}

func returnAlbums(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(Albums)
	if err != nil {
		log.Fatal(err)
	}
}

func getOneUser(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("user")

	if username != "" {
		user := db.GetUser(username)
		if user.Username != "" {
			err := json.NewEncoder(w).Encode(user)
			if err != nil {
				log.Fatal(err)
			}

			return
		}
	}

	http.Error(w, "No user found", http.StatusNotFound)
}

func main() {
	handleRequests()
}
