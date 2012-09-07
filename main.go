package main

import (
	"encoding/json"
	"github.com/bmizerany/pat"
	"log"
	"net/http"
	"strconv"
)

func retrievePlayerProfile(w http.ResponseWriter, req *http.Request) {
	battleTag := req.URL.Query().Get(":battleTag")

	playerProfile, err := GetPlayerProfile("eu", battleTag)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-type", "application/json")
	if err = json.NewEncoder(w).Encode(playerProfile); err != nil {
		log.Fatal(err)
	}
}

func retrieveHeroProfile(w http.ResponseWriter, req *http.Request) {
	battleTag := req.URL.Query().Get(":battleTag")
	heroId, _ := strconv.Atoi(req.URL.Query().Get(":heroId"))

	heroProfile, err := GetHeroProfile("eu", battleTag, heroId)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-type", "application/json")
	if err = json.NewEncoder(w).Encode(heroProfile); err != nil {
		log.Fatal(err)
	}
}

func main() {
	m := pat.New()

	//	m.Get("/", http.FileServer(http.Dir("/static")))

	m.Get("/profile/:battleTag/", http.HandlerFunc(retrievePlayerProfile))
	m.Get("/profile/:battleTag/hero/:heroId/", http.HandlerFunc(retrieveHeroProfile))

	http.Handle("/", m)
	if err := http.ListenAndServe(":3001", nil); err != nil {
		log.Fatal(err)
	}
}
