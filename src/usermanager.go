package main

import "math/rand"

type UserManager struct {
	Users map[int]*User
}

func (manager UserManager) CreateUsers(numberOfUsers int) []*User {
	createdUsers := make([]*User, 0)

	for i := 0; i < numberOfUsers; i++ {
		user := createUser(len(manager.Users))
		createdUsers = append(createdUsers, &user)
		manager.Users[user.Id] = &user
	}

	return createdUsers
}

func (manager UserManager) UpdateUsers(numberOfUsers int) []*User {
	updatedUsers := make([]*User, 0)

	for i := 0; i < numberOfUsers; i++ {
		user := updateUser(*manager.Users[rand.Intn(len(manager.Users)-1)])
		updatedUsers = append(updatedUsers, &user)
		manager.Users[user.Id] = &user
	}

	return updatedUsers
}

func (manager UserManager) GetUsers() []*User {
	retVal := make([]*User, 0)

	for _, user := range manager.Users {
		retVal = append(retVal, user)
	}

	return retVal
}

func (manager UserManager) ClearUsers() {
	for k := range manager.Users {
		delete(manager.Users, k)
	}
}
