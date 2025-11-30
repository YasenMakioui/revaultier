package user

import "revaultier/configuration"

type UserService struct {
	cfg            *configuration.Config
	userRepository *UserRepository
}

func NewUserService(cfg *configuration.Config, ur *UserRepository) *UserService {
	return &UserService{cfg: cfg, userRepository: ur}
}
