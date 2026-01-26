package service

type UserService interface {
	GetUsers()
	CreateUser()
	GetUserByID()
	UpdateUser()
	DeleteUser()
}
