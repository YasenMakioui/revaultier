package user

type User struct {
	username       string
	hashedPassword string
	SessionToken   string
	CSRFToken      string
}
