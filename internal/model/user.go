package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name        string
	Surname     string
	TC          string
	PhoneNumber string
	Email       string
}

type Service struct {
	DB *gorm.DB
}

type UserService struct {
	//GetUserByID
	//GetUsers
	//PostUser
	//PutUser
	//DeleteUser
}

//GetUserByID - retrieves comments by their ID from the database
func GetUserByID(ID uint) (User, error) {
	return User{}, nil

}

//GetUser - Get all users from the database
func GetUsers() (User, error) {
	return User{}, nil

}

//PostUser - add a new user to the database
func PostUser(user User) (User, error) {
	return User{}, nil

}

//PutUser - Update a row from the database
func PutUser(ID uint, newuser User) (User, error) {
	return User{}, nil

}

//DeleteUser - Delete a row from database
func DeleteUser(ID uint) (User, error) {
	return User{}, nil

}

func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}
