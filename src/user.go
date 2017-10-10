package main

import "math/rand"

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
	return User{user.Id, "", "", "", ""}
}

func generateGender() string {
	g := rand.Intn(1)

	if g == 1 {
		return "m"
	}

	return "f"
}

func generateFN(gender string) string {
	for {
		ran := rand.Intn(len(config.FirstNames))

		if config.FirstNames[ran].Gender == gender {
			return config.FirstNames[ran].Name
		}
	}
}

func generateLN() string {
	ran := rand.Intn(len(config.LastNames))
	return config.LastNames[ran]
}

func generateOccupation() string {
	ran1 := rand.Intn(len(config.Occupations))
	ran2 := rand.Intn(len(config.Adjectives))
	return config.Adjectives[ran2] + " " + config.Occupations[ran1]
}
