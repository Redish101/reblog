package auth

import "testing"

func TestSetKey(t *testing.T) {
	SetKey()
}

func TestVerifyToken(t *testing.T) {
	token := GetToken("testuser", "testpassword")
	
	tokenStatus := ValidToken(token)

	if !tokenStatus {
		t.Error("Token验证失败")
	}
}