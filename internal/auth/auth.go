package auth

import (
	"fmt"
	"math/rand"
	"reblog/internal/hash"
	"reblog/internal/model"
	"reblog/internal/query"
	"regexp"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var key []byte

type TokenClaim struct {
	Username string `json:"usr"`
	Password string `json:"pwd"`

	jwt.RegisteredClaims
}

func SetKey() {
	key = []byte(hash.Hash("reblog_sign_key" + fmt.Sprint(time.Now().UnixMicro()+rand.Int63n(1000000000))))
}

func GetToken(username string, password string) string {
	u := query.User

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
			Issuer:    "reblog-server",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, _ := token.SignedString(key)

	return signedToken
}

func ValidToken(token string) bool {
	parsedToken, err := jwt.ParseWithClaims(token, &TokenClaim{}, func(t *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		fmt.Println(err)
		return false
	}

	if _, ok := parsedToken.Claims.(*TokenClaim); ok && parsedToken.Valid {
		return true
	}

	return false
}
