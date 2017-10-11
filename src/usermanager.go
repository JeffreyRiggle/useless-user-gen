package main

import "math/rand"

type UserManager struct {
	Users map[int]*User
}

func (manager UserManager) CreateUsers(numberOfUsers int) []*User {
	createdUsers := make([]*User, numberOfUsers)

	for i := 0; i < numberOfUsers; i++ {
		user := createUser(len(manager.Users) + 1)
		createdUsers = append(createdUsers, &user)
		manager.Users[user.Id] = &user
	}

	return createdUsers
}

func (manager UserManager) UpdateUsers(numberOfUsers int) []*User {
	updatedUsers := make([]*User, numberOfUsers)

	for i := 0; i < numberOfUsers; i++ {
		user := updateUser(*manager.Users[rand.Intn(len(manager.Users))])
		updatedUsers = append(updatedUsers, &user)
		manager.Users[user.Id] = &user
	}

	return updatedUsers
}

func (manager UserManager) GetUsers() []*User {
	retVal := make([]*User, len(manager.Users))

	for _, user := range manager.Users {
		retVal = append(retVal, user)
	}

	return retVal
}

func (manager UserManager) ClearUsers() {
	manager.Users = make(map[int]*User, 0)
}
