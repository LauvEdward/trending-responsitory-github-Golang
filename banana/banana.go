package banana

import "github.com/pkg/errors"

type Banana struct {
}

var (
	UserConflict   = errors.New("User already exists")
	UserNotFound   = errors.New("Can not find user")
	SignUpFailed   = errors.New("Sign up failed")
	PasswordFailed = errors.New("Wrong password")
	EmailFailed    = errors.New("Wrong email")
)
