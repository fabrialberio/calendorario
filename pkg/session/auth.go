package session

import (
	"calendorario/pkg/database"
	"errors"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrCookieNotFound = errors.New("cookie not found")
	ErrCookieExpired  = errors.New("cookie is expired")
)

const sessionCookie = "jwt"

func HashPassword(password string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return []byte{}, err
	}

	return hash, nil
}

func AuthenticatedUser(r *http.Request) (*database.User, error) {
	cookie, err := r.Cookie(sessionCookie)
	if err != nil {
		return nil, ErrCookieNotFound
	}

	claims, err := validateJWT(cookie.Value)
	if errors.Is(err, jwt.ErrTokenExpired) {
		return nil, ErrCookieExpired
	} else if err != nil {
		return nil, err
	}

	return &claims.User, nil
}

func SetAuthenticatedUser(w http.ResponseWriter, user *database.User, password []byte) error {
	err := bcrypt.CompareHashAndPassword(user.PasswordHash, password)
	if err != nil {
		return err
	}

	token, err := generateJWT(user.ID, user.Role)
	if err != nil {
		return err
	}

	setCookie(w, sessionCookie, token)

	return nil
}

func UnsetAuthenticatedUser(w http.ResponseWriter) {
	unsetCookie(w, sessionCookie)
}
