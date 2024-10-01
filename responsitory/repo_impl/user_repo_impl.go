package repo_impl

import (
	"database/sql"
	"github.com/lib/pq"
	"golang.org/x/net/context"
	"time"
	"trending-github-golang/banana"
	"trending-github-golang/db"
	"trending-github-golang/log"
	"trending-github-golang/model"
	"trending-github-golang/model/req"
	"trending-github-golang/responsitory"
	"trending-github-golang/security"
)

type UserRepoImpl struct {
	SqlU *db.Sql
}

func NewUserRepoImpl(sqlU *db.Sql) responsitory.UserRepo {
	return &UserRepoImpl{SqlU: sqlU}
}

func (u *UserRepoImpl) CheckUser(c context.Context, user req.ReqSignIn) (model.User, error) {
	sqlQuery := `SELECT * FROM users WHERE email=$1`

	var foundUser model.User
	err := u.SqlU.Db.GetContext(c, &foundUser, sqlQuery, user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.User{}, banana.UserNotFound // Handle user not found
		}
		return model.User{}, banana.UserConflict // Handle other errors
	}
	isVerifyPass := security.VerifyPassword(user.Password, foundUser.Password)
	if isVerifyPass {
		return foundUser, nil
	}

	return model.User{}, banana.PasswordFailed
}

func (u UserRepoImpl) SaveUser(c context.Context, user model.User) (model.User, error) {
	sqlQuery := `
		INSERT INTO users(user_id, email, password, role, full_name, created_at, updated_at)
		VALUES(:user_id, :email, :password, :role, :full_name, :created_at, :updated_at)
	`
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	_, err := u.SqlU.Db.NamedExecContext(c, sqlQuery, user)
	if err != nil {
		if e, ok := err.(*pq.Error); ok {
			println(e.Code.Name())
			if e.Code.Name() == "unique_violation" {
				return user, banana.UserConflict
			}
		}
		log.Error(err)
		return user, banana.SignUpFailed
	}
	return user, nil
}
