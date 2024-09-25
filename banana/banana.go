package banana

import "github.com/pkg/errors"

type Banana struct {
}

var (
	UserConflict  = errors.New("user already exists")
	EmailConflict = errors.New("email already exists")
	SignUpFailed  = errors.New("sign up failed")
)
