package core

import (
	"testing"

	"github.com/redish101/reblog/internal/hash"
	"github.com/redish101/reblog/internal/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/gen/field"
)

func TestGetServiceName(t *testing.T) {
	gotName := getServiceName[*MockService]()

	assert.Equal(t, "*core.MockService", gotName)
}

func TestAuthService(t *testing.T) {
	app := NewApp(TestConfig)

	app.Bootstrap()

	authService, err := AppService[*AuthService](app)
	assert.IsType(t, &AuthService{}, authService)
	assert.NoError(t, err)

	user, err := app.Query().User.Attrs(field.Attrs(&model.User{
		Username: "reblog_test_user",
		Password: hash.Hash("reblog_test_password"),
	})).FirstOrCreate()
	assert.NotNil(t, user)
	assert.NoError(t, err)

	token := authService.GetToken("reblog_test_user", "reblog_test_password")
	assert.NotNil(t, token)
	assert.True(t, authService.VerifyToken(token))
}
