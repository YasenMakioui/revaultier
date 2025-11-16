package auth

import "errors"

// type UsernameTakenError struct{}

// func (u *UsernameTakenError) Error() string {
// 	return "username is already taken"
// }

var ErrUsernameTaken = errors.New("username is already taken")
