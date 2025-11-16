package user

type UserService struct {
	userRepository *UserRepository
}

func NewUserService(ur *UserRepository) *UserService {
	return &UserService{userRepository: ur}
}
