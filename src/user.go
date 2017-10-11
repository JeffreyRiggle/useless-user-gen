package main

import "math/rand"

type UserConfigRepository struct {
	FirstNames  map[int]FName  `json:"firstnames"`
	LastNames   map[int]string `json:"lastnames"`
	Occupations map[int]string `json:"occupations"`
	Adjectives  map[int]string `json:"adjectives"`
}

type User struct {
	Id         int    `json:"id"`
	FirstName  string `json:"firstname"`
	LastName   string `json:"lastname"`
	Occupation string `json:"occupation"`
	Gender     string `json:"gender"`
}

func createUser(id int) User {
	gender := generateGender()

	return User{id, generateFN(gender), generateLN(), generateOccupation(), gender}
}

func updateUser(user User) User {
	return User{user.Id, generateFN(user.Gender), generateLN(), generateOccupation(), user.Gender}
}

func generateGender() string {
	g := rand.Intn(2)

	if g == 1 {
		return "m"
	}

	return "f"
}

func generateFN(gender string) string {
	for {
		ran := rand.Intn(len(repo.FirstNames))

		if repo.FirstNames[ran].Gender == gender {
			return repo.FirstNames[ran].Name
		}
	}
}

func generateLN() string {
	ran := rand.Intn(len(repo.LastNames))
	return repo.LastNames[ran]
}

func generateOccupation() string {
	ran1 := rand.Intn(len(repo.Occupations))
	ran2 := rand.Intn(len(repo.Adjectives))
	return repo.Adjectives[ran2] + " " + repo.Occupations[ran1]
}
