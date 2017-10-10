package main

type Config struct {
	FirstNames  map[int]FName  `json:"firstnames"`
	LastNames   map[int]string `json:"lastnames"`
	Occupations map[int]string `json:"occupations"`
	Adjectives  map[int]string `json:"adjectives"`
}

type FName struct {
	Name   string `json:"name"`
	Gender string `json:"gender"`
}
