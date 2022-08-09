package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name        string
	Surname     string
	TC          string
	PhoneNumber string
	Email       string
}

type userService struct {
	DB *gorm.DB
}

type UserService interface {
	GetUserByID(ID uint) (User, error)
	GetUsers() (*[]User, error)
	PostUser(user User) (User, error)
	PutUser(ID uint, newUser User) (User, error)
	DeleteUser(ID uint) (User, error)
}

//GetUserByID - retrieves comments by their ID from the database
func (s *userService) GetUserByID(ID uint) (User, error) {
	var user User
	if result := s.DB.First(&user); result.Error != nil {

	}
	return user, nil

}

//GetUser - Get all users from the database
func (s *userService) GetUsers() (*[]User, error) {
	var user []User
	if result := s.DB.Find(&user); result.Error != nil {

	}
	return &user, nil

}

//PostUser - add a new user to the database
func (s *userService) PostUser(user User) (User, error) {
	if result := s.DB.Save(&user); result.Error != nil {

	}
	return user, nil

}

//PutUser - Update a row from the database
func (s *userService) PutUser(ID uint, newUser User) (User, error) {
	var user User
	user, err := s.GetUserByID(ID)
	if err != nil {
		return User{}, err
	}
	if result := s.DB.Model(&user).Updates(newUser); result.Error != nil {

		return User{}, err
	}
	return user, nil

}

//DeleteUser - Delete a row from database
func (s *userService) DeleteUser(ID uint) (User, error) {
	var user User
	if result := s.DB.Delete(&user, ID); result.Error != nil {

	}
	return user, nil

}

func NewUserService(db *gorm.DB) *userService {
	return &userService{
		DB: db,
	}
}
