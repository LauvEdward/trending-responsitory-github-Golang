package repo_impl

import (
	"errors"
	"github.com/lib/pq"
	"golang.org/x/net/context"
	"trending-github-golang/banana"
	"trending-github-golang/db"
	"trending-github-golang/model"
)

type UserRepoImpl struct {
	Db *db.Sql
}

func (u *UserRepoImpl) SaveUser(c context.Context, user model.User) (model.User, error) {
	sqlQuery := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3)`
	_, err := u.Db.Db.NamedExecContext(c, sqlQuery, user)
	if err != nil {
		var err *pq.Error
		if errors.As(err, &err) {
			if err.Code.Name() == "unique_violation" {
				return user, banana.UserConflict
			}
		}
		return user, banana.SignUpFailed
	}
	return user, nil
}
