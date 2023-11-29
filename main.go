package main

import (
	"encoding/json"
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
	log.Fatal(http.ListenAndServe(":7044", nil))
}

func returnAlbums(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(Albums)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	Albums = []Album{
		{
			Id:         1,
			SellerId:   1,
			Name:       "Antifragile",
			Genre:      "K-pop",
			Condition:  "As new",
			EndingTime: "2023-11-30 22:00:00",
			IsDone:     false,
			Bids: []Bid{
				{
					Id:        1,
					UserId:    1,
					Amount:    30,
					BidPlaced: "2023-11-23 21:30:00",
				},
			},
		},
	}
	handleRequests()
}
