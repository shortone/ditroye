package main

import (
	"encoding/json"
	"github.com/bmizerany/pat"
	"log"
	"net/http"
	"strconv"
	"html/template"
)

type Profile struct {
	Player PlayerProfile
	Heroes []HeroProfile
}

var homePage = template.Must(template.ParseFiles(
	"templates/_base.html",
	"templates/index.html",
))

func home(w http.ResponseWriter, req *http.Request) {
	battleTag := req.URL.Query().Get("battleTag")
	if battleTag != "" {
		profile, err := GetPlayerProfile("eu", battleTag)
		if err != nil {
			http.Error(w, "Unable to retrieve player profile", 404)
			return
		}
		if err := homePage.Execute(w, profile); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		if err := homePage.Execute(w, nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func hero(w http.ResponseWriter, req *http.Request) {
	battleTag := req.URL.Query().Get("battleTag")
	heroId, err := strconv.Atoi(req.URL.Query().Get("heroId"))
	if err != nil {
		log.Print(err)
		heroId = -1
	}

	if battleTag != "" && heroId != -1 {
		profile, err := GetHeroProfile("eu", battleTag, heroId)
		if err != nil {
			http.Error(w, "Unable to retrieve hero profile", 404)
			return
		}
		w.Header().Set("Content-type", "application/json")
		if err = json.NewEncoder(w).Encode(profile); err != nil {
			http.Error(w, "Unable to build response", 500)
		}
	} else {
		http.Redirect(w, req, "/", 301)
	}
}

func main() {
	m := pat.New()

	m.Get("/", http.HandlerFunc(home))
	m.Get("/hero", http.HandlerFunc(hero))
	m.Get("/app/", http.StripPrefix("/app/", http.FileServer(http.Dir("./app/"))))

	http.Handle("/", m)
	if err := http.ListenAndServe(":3001", nil); err != nil {
		log.Fatal(err)
	}
}
