package model

import "strconv"

type CountryRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (c *CountryRequest) IsValid() bool {
	if _, err := strconv.Atoi(c.ID); err != nil {
		return false
	}
	if len(c.Name) == 0 {
		return false
	}
	return true
}

type TownRequest struct {
	Name     string `json:"name"`
	CountryID string `json:"country_id"`
}

func (c *TownRequest) IsValid() bool {
	if _, err := strconv.Atoi(c.CountryID); err != nil {
		return false
	}
	if len(c.Name) == 0 {
		return false
	}
	return true
}

type DistrictRequest struct {
	Name   string `json:"name"`
	TownID string `json:"town_id"`
}

func (c *DistrictRequest) IsValid() bool {
	return len(c.Name) > 0 && len(c.TownID) > 0
}
