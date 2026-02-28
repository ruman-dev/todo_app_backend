package users

var users = []User{}

type User struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Create(newUser User) *User {
	// if newUser.Id == 0 {
	// 	return nil
	// }

	newUser.Id = len(users) + 1
	users = append(users, newUser)
	return &newUser
}

func LoginUser(loginUser Login) *User {
	for _, user := range users {
		if user.Email == loginUser.Email && user.Password == loginUser.Password {
			return &user
		}
	}
	return nil
}
