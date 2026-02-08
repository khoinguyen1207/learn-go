package repository

type userRepository struct {
}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (ur *userRepository) FindAll() error {
	return nil
}

func (ur *userRepository) Create() error {
	return nil
}

func (ur *userRepository) FindById(id string) bool {
	return false
}

func (ur *userRepository) Update(id string) error {
	return nil
}

func (ur *userRepository) Delete(id string) error {
	return nil
}

func (ur *userRepository) FindByEmail(email string) bool {
	return false
}
