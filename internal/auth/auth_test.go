package auth

import (
	"reblog/internal/model"
	"reblog/internal/query"
	"testing"
)

func TestSetKey(t *testing.T) {
	SetKey()
}

func TestVerifyToken(t *testing.T) {
	username := "testuser"
	nickname := "testnickname"
	email := "testemail"
	password := "testpassword"

	user := &model.User{
		Username: username,
		Nickname: nickname,
		Email:    email,
		Password: password,
	}

	err := query.User.Create(user)

	if err != nil {
		t.Errorf("创建用户失败: %v", err)
	}

	token := GetToken(username, password)
	
	tokenStatus := ValidToken(token)

	if !tokenStatus {
		t.Error("Token验证失败")
	}
}
