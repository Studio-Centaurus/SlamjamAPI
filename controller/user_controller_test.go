package controller_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Studio-Centaurus/SlamjamAPI/controller"
	"github.com/Studio-Centaurus/SlamjamAPI/mocks"
	"github.com/Studio-Centaurus/SlamjamAPI/models"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSignup(t *testing.T) {
	app := fiber.New()

	// Create a mock repository
	mockRepo := new(mocks.MockUserRepository)
	userController := &controller.UserController{
		Repo: mockRepo,
	}

	app.Post("/user/signup", userController.Signup)

	user := models.User{
		UserName: "testuser",
		Password: "testpassword",
	}

	mockRepo.On("CreateUser", mock.AnythingOfType("models.User")).Return(nil)

	userJSON, _ := json.Marshal(user)
	req := httptest.NewRequest(http.MethodPost, "/user/signup", bytes.NewBuffer(userJSON))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req, -1)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	var responseBody models.User
	json.NewDecoder(resp.Body).Decode(&responseBody)

	assert.Equal(t, user.UserName, responseBody.UserName)

	mockRepo.AssertExpectations(t)
}

func TestLogin(t *testing.T) {
	app := fiber.New()

	// Create a mock repository
	mockRepo := new(mocks.MockUserRepository)
	userController := &controller.UserController{
		Repo: mockRepo,
	}

	app.Post("/user/login", userController.Login)

	loginRequest := models.LoginRequest{
		Username: "testuser",
		Password: "testpassword",
	}

	user := models.User{
		UserName: "testuser",
		Password: "$2a$10$CwTycUXWue0Thq9StjUM0uEox8Bt9fG1OeEsmY1T2iJPi2sFQ5Pjm", // hashed password for "testpassword"
	}

	mockRepo.On("FindByCredentials", loginRequest.Username, loginRequest.Password).Return(&user, nil)

	loginRequestJSON, _ := json.Marshal(loginRequest)
	req := httptest.NewRequest(http.MethodPost, "/user/login", bytes.NewBuffer(loginRequestJSON))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req, -1)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	var responseBody map[string]string
	json.NewDecoder(resp.Body).Decode(&responseBody)

	assert.Equal(t, "Login successful", responseBody["message"])
	assert.NotEmpty(t, responseBody["token"])

	mockRepo.AssertExpectations(t)
}
