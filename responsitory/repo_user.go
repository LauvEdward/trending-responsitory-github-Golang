package responsitory

import (
	"context"
	"trending-github-golang/model"
	"trending-github-golang/model/req"
)

type UserRepo interface {
	CheckUser(c context.Context, user req.ReqSignIn) (model.User, error)
	SaveUser(c context.Context, user model.User) (model.User, error)
}
