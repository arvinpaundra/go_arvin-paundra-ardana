package main

/*
	BEFORE
*/

// type user struct {
// 	id       int
// 	username int
// 	password int
// }

// type userservice struct {
// 	t []user
// }

// func (u userservice) getallusers() []user {
// 	return u.t
// }

// func (u userservice) getuserbyid(id int) user {
// 	for _, r := range u.t {
// 		if id == r.id {
// 			return r
// 		}
// 	}

// 	return user{}
// }

/*
	AFTER
*/

type User struct {
	ID       int
	Username string
	Password string
}

type UserService struct {
	Users []User
}

func (u UserService) GetAllUsers() []User {
	return u.Users
}

func (u UserService) GetUserByID(id int) User {
	for _, user := range u.Users {
		if id == user.ID {
			return user
		}
	}

	return User{}
}
