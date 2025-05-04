package repo

// type UserRepo struct{}

// func NewUserRepo() *UserRepo {
// 	return &UserRepo{}
// }

// func (ur *UserRepo) GetInfoUser() string {
// 	return "UserRepository: GetInfoUser()"
// }

type IUserRepository interface {
	GetUserByEmail(email string) bool
}

type userRepository struct {
}

func NewUserRepository() IUserRepository {
	return &userRepository{}
}

func (ur *userRepository) GetUserByEmail(email string) bool {
	return true
}
