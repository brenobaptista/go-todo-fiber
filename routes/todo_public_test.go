package routes

import (
	"net/http/httptest"
	"testing"

	"github.com/brenobaptista/go-todo-fiber/configs"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestTodoPublicRoutes(t *testing.T) {
	err := godotenv.Load("../.env.test")

	if err != nil {
		panic(err)
	}

	tests := []struct {
		description   string
		route         string
		method        string
		expectedError bool
		expectedCode  int
	}{
		{
			description:   "get all to-dos",
			route:         "/api/todos",
			method:        "GET",
			expectedError: false,
			expectedCode:  200,
		},
		{
			description:   "get a to-do by given ID",
			route:         "/api/todos/60e6725c46311dad3e62e5fb",
			method:        "GET",
			expectedError: false,
			expectedCode:  200,
		},
		{
			description:   "try to get a to-do using unknown ID",
			route:         "/api/todos/123456",
			method:        "GET",
			expectedError: false,
			expectedCode:  400,
		},
	}

	app := fiber.New()

	configs.ConnectDB()

	TodoPublicRoutes(app)

	for _, test := range tests {
		req := httptest.NewRequest(test.method, test.route, nil)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req, -1)

		assert.Equalf(t, test.expectedError, err != nil, test.description)

		if test.expectedError {
			continue
		}

		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}
