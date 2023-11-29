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
	err := json.NewEncoder(w).Encode(db.GetUser())
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	handleRequests()
}
