package user

type UserHandler struct {
	UserService *UserService
}

func NewUserHandler(us *UserService) *UserHandler {
	return &UserHandler{UserService: us}
}
