package util

import (
	"time"

	"golang.org/x/crypto/bcrypt"

	jwt "github.com/golang-jwt/jwt/v4"
	auth "github.com/nguyendhst/fx-template/domain/auth"
	user "github.com/nguyendhst/fx-template/domain/user"
)

func CreateAccessToken(user *user.User, secret string, expiration time.Duration) (accessToken string, err error) {
	exp := time.Now().Add(expiration)
	claims := &auth.JwtCustomClaims{
		Name: user.Name,
		ID:   user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: exp},
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return t, err
}

func CreateRefreshToken(user *user.User, secret string, expiration time.Duration) (refreshToken string, err error) {
	exp := time.Now().Add(expiration)
	claimsRefresh := &auth.JwtCustomRefreshClaims{
		ID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: exp},
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)
	rt, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return rt, err
}

func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
