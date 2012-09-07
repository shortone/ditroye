package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Skill struct {
	Skill struct {
		Slug              string
		Name              string
		Icon              string
		TooltipParams     string
		Description       string
		SimpleDescription string
	}
	Rune struct {
		Slug              string
		Name              string
		Icon              string
		TooltipParams     string
		Description       string
		SimpleDescription string
		OrderIndex        int
	}
}

type HeroProfile struct {
	Id       int
	Name     string
	Class    string
	Gender   int
	Level    int
	Hardcore bool
	//	Skills struct {
	//		Active []Skill
	//	}
	Stats struct {
		DamageIncrease  float64
		DamageReduction float64
		CritChance      float64
		Life            int
		Strength        int
	}
}

func GetHeroProfile(zone string, battleTag string, heroId int) (*HeroProfile, error) {
	url := fmt.Sprintf("http://%s.battle.net/api/d3/profile/%s/hero/%d", zone, battleTag, heroId)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, errors.New(response.Status)
	}
	profile := new(HeroProfile)
	err = json.NewDecoder(response.Body).Decode(profile)
	if err != nil {
		return nil, err
	}
	return profile, nil
}
