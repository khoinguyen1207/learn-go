package repository

type UserRepository interface {
	FindAll() error
	Create() error
	FindById(id string) bool
	Update(id string) error
	Delete(id string) error
	FindByEmail(email string) bool
}
