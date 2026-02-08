package service

type UserService interface {
	GetUsers(search string, page, limit int) error
	CreateUser() error
	GetUserByID(id string) error
	UpdateUser(id string) error
	DeleteUser(id string) error
}
