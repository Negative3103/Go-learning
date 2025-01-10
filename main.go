package main

type User struct {
	name   string
	number int
	userId int
}

func (user User) getName() string {
	return user.name
}

func createUser(name string, number int, userId int) User {
	return User{
		name:   name,
		number: number,
		userId: userId,
	}
}

func (user *User) changeName(newName string) bool {
	if newName != "" {
		user.name = newName
		return true
	} else {
		return false
	}
}
