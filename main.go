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
		http.Error(w, "Unable to retrieve player profile", 500)
		return
	}

	w.Header().Set("Content-type", "application/json")
	if err = json.NewEncoder(w).Encode(playerProfile); err != nil {
		http.Error(w, "Unable to build response", 500)
	}
}

func retrieveHeroProfile(w http.ResponseWriter, req *http.Request) {
	battleTag := req.URL.Query().Get(":battleTag")
	heroId, err := strconv.Atoi(req.URL.Query().Get(":heroId"))
	if err != nil {
		http.Error(w, "Invalid hero id", 500)
		return
	}

	heroProfile, err := GetHeroProfile("eu", battleTag, heroId)
	if err != nil {
		http.Error(w, "Unable to retrieve hero profile", 500)
		return
	}

	w.Header().Set("Content-type", "application/json")
	if err = json.NewEncoder(w).Encode(heroProfile); err != nil {
		http.Error(w, "Unable to build response", 500)
	}
}

func main() {
	m := pat.New()

	m.Get("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	m.Get("/profile/:battleTag/", http.HandlerFunc(retrievePlayerProfile))
	m.Get("/profile/:battleTag/hero/:heroId/", http.HandlerFunc(retrieveHeroProfile))

	http.Handle("/", m)
	if err := http.ListenAndServe(":3001", nil); err != nil {
		log.Fatal(err)
	}
}
