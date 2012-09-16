package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"log"
)

type Artisan struct {
	Slug        string
	Level       int
	StepCurrent int
	StepMax     int
}

type PlayerProfile struct {
	Heroes []struct {
		Name         string
		Id           int
		Level        int
		Hardcore     bool
		ParagonLevel int
		Gender       int
		Dead         bool
		Class        string
		LastUpdated  int `json:"last-updated"`
	}
	LastHeroPlayed   int
	LastUpdated      int
	Artisans         []Artisan
	HardcoreArtisans []Artisan
	Kills            struct {
		Monsters         int
		Elites           int
		HardcoreMonsters int
	}
	TimePlayed struct {
		Barbarian   float64
		DemonHunter float64 `json:"demon-hunter"`
		Monk        float64
		WitchDoctor float64 `json:"witch-doctor"`
		Wizard      float64
	}
	BattleTag string
}

func GetPlayerProfile(zone string, battleTag string) (*PlayerProfile, error) {
	url := fmt.Sprintf("http://%s.battle.net/api/d3/profile/%s/", zone, battleTag)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, errors.New(response.Status)
	}
	profile := new(PlayerProfile)
	err = json.NewDecoder(response.Body).Decode(profile)
	if err != nil {
		return nil, err
	}
	log.Print(profile)
	return profile, nil
}
