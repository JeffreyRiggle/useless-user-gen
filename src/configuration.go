package main

type Config struct {
	FirstNames  []FName  `json:"firstnames"`
	LastNames   []string `json:"lastnames"`
	Occupations []string `json:"occupations"`
	Adjectives  []string `json:"adjectives"`
}

type FName struct {
	Name   string `json:"name"`
	Gender string `json:"gender"`
}
