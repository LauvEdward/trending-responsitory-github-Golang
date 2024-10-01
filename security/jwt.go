package security

import (
	"github.com/golang-jwt/jwt"
	"time"
	"trending-github-golang/model"
)

const SECRETKEY = "edwardlauv1234"

func GenerateToken(user model.User) (string, error) {
	clams := model.JwtCustomClams{
		Userid:   user.UserId,
		RoleUser: user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(0, 6, 0).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, clams)
	result, err := token.SignedString([]byte(SECRETKEY))
	if err != nil {
		return "", err
	}
	return result, nil
}
