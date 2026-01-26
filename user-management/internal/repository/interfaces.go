package repository

type UserRepository interface {
	FindAll()
	Create()
	FindById()
	Update()
	Delete()
}
