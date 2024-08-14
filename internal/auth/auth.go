package auth

import (
	"fmt"
	"math/rand"
	"regexp"
	"time"

	"github.com/ChuqiCloud/acmeidc/internal/hash"
	"github.com/ChuqiCloud/acmeidc/internal/model"
	"github.com/ChuqiCloud/acmeidc/internal/query"

	"github.com/golang-jwt/jwt/v5"
)

type Auth struct {
	key   []byte
	query *query.Query
}

func NewAuth(q *query.Query) *Auth {
	return &Auth{
		key:   NewKey(),
		query: q,
	}
}

type TokenClaim struct {
	Username string `json:"usr"`
	Password string `json:"pwd"`

	jwt.RegisteredClaims
}

func NewKey() []byte {
	return []byte(hash.Hash("acmeidc_sign_key" + fmt.Sprint(time.Now().UnixMicro()+rand.Int63n(1000000000))))
}

func (a *Auth) GetToken(username string, password string) string {
	u := a.query.User

	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	emailMatched, _ := regexp.MatchString(emailRegex, username)

	var user *model.User

	if emailMatched {
		user, _ = u.Where(u.Email.Eq(username)).First()
	} else {
		user, _ = u.Where(u.Username.Eq(username)).First()
	}

	if user == nil {
		return ""
	}

	if hash.Hash(password) != user.Password {
		return ""
	}

	claims := TokenClaim{
		user.Username,
		user.Password,
		jwt.RegisteredClaims{
			Issuer:    "acmeidc-server",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, _ := token.SignedString(a.key)

	return signedToken
}

func (a *Auth) VerifyToken(token string) bool {
	parsedToken, err := jwt.ParseWithClaims(token, &TokenClaim{}, func(t *jwt.Token) (interface{}, error) {
		return a.key, nil
	})

	if err != nil {
		return false
	}

	if _, ok := parsedToken.Claims.(*TokenClaim); ok && parsedToken.Valid {
		return true
	}

	return false
}
