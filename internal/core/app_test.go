package core

import (
	"fmt"
	"os"
	"testing"

	"github.com/ChuqiCloud/acmeidc/internal/config"
	"github.com/stretchr/testify/assert"
)

var TestPwd = os.Getenv("TESTPWD")

var TestConfig = &config.Config{
	DB: config.DBConfig{
		Type: "sqlite3",
		Name: fmt.Sprintf("%s/acmeidc_test.db", TestPwd),
	},
}

type MockService struct{}

func (s *MockService) Start() error {
	return nil
}

func (s *MockService) Stop() error {
	return nil
}

func NewMockService(app *App) *MockService {
	return &MockService{}
}

func TestNewApp(t *testing.T) {
	app := NewApp(TestConfig)

	assert.NotNil(t, app)
}

func TestAppInjectAndService(t *testing.T) {
	app := NewApp(TestConfig)

	mockservice := NewMockService(app)

	AppInject(app, mockservice)

	gotService, err := AppService[*MockService](app)
	if assert.NoError(t, err) {
		assert.NotNil(t, gotService)
		assert.Equal(t, mockservice, gotService)
	}

	t.Run("AccessNilService", func(t *testing.T) {
		app := NewApp(TestConfig)

		gotService, err := AppService[*MockService](app)

		assert.Nil(t, gotService)
		assert.Error(t, err)
	})
}

func TestBootstrap(t *testing.T) {
	app := NewApp(TestConfig)

	app.Bootstrap()
}
