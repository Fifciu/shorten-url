package users

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyCreateUser struct {
	fullname string
	email    string
	password string
	called   bool
}

func (s *SpyCreateUser) CreateUser(fullname, email, password string) (*User, error) {
	s.fullname = fullname
	s.email = email
	s.password = password
	s.called = true
	return &User{0, fullname, email, password}, nil
}

func (s *SpyCreateUser) HashPassword(password string) (string, error) {
	return fmt.Sprintf("hashed-%s", password), nil
}

func (s *SpyCreateUser) GenerateJwtToken(user *User) (string, time.Time, error) {
	return "", time.Now(), nil
}

type SpyCreateUserError struct {
	errorMessage string
}

func (s *SpyCreateUserError) CreateUser(fullname, email, password string) (*User, error) {
	return nil, errors.New(s.errorMessage)
}

func (s *SpyCreateUserError) HashPassword(password string) (string, error) {
	return password, nil
}

func (s *SpyCreateUserError) GenerateJwtToken(user *User) (string, time.Time, error) {
	return "", time.Now(), nil
}

type SpyGenerateJwtError struct {
	errorMessage string
}

func (s *SpyGenerateJwtError) CreateUser(fullname, email, password string) (*User, error) {
	return &User{0, fullname, email, password}, nil
}

func (s *SpyGenerateJwtError) HashPassword(password string) (string, error) {
	return password, nil
}

func (s *SpyGenerateJwtError) GenerateJwtToken(user *User) (string, time.Time, error) {
	return "", time.Now(), errors.New(s.errorMessage)
}

type SpyGenerateToken struct {
	generatedToken string
	expiryTime     time.Time
}

func (s *SpyGenerateToken) CreateUser(fullname, email, password string) (*User, error) {
	return &User{0, fullname, email, password}, nil
}

func (s *SpyGenerateToken) HashPassword(password string) (string, error) {
	return fmt.Sprintf("hashed-%s", password), nil
}

func (s *SpyGenerateToken) GenerateJwtToken(user *User) (string, time.Time, error) {
	return s.generatedToken, s.expiryTime, nil
}

func TestRegister(t *testing.T) {
	t.Run("rejects additional fields", func(t *testing.T) {
		authContoller := NewAuthController(&SpyCreateUser{})
		body := newPayload(t, map[string]string{
			"notSupported": "field",
		})
		request, _ := http.NewRequest(http.MethodPost, "/authentication/register", body)
		response := httptest.NewRecorder()

		authContoller.Endpoints["/register"](response, request)

		assertStatus(t, response.Code, http.StatusBadRequest)
	})

	t.Run("checks for required fields", func(t *testing.T) {
		requiredFieldsTests := []struct {
			body       map[string]string
			statusCode int
		}{
			{
				map[string]string{},
				http.StatusBadRequest,
			},
			{
				map[string]string{
					"fullname": "John",
					"email":    "john@doe.com",
				},
				http.StatusBadRequest,
			},
			{
				map[string]string{
					"fullname": "John",
					"email":    "johndoe.com",
					"password": "zaq1@WSX",
				},
				http.StatusBadRequest,
			},
			{
				map[string]string{
					"fullname": "John",
					"email":    "john@doe.com",
					"password": "zaq1@WSX",
				},
				http.StatusOK,
			},
		}
		authContoller := NewAuthController(&SpyCreateUser{})

		for _, test := range requiredFieldsTests {
			body := newPayload(t, test.body)
			request, _ := http.NewRequest(http.MethodPost, "/authentication/register", body)
			response := httptest.NewRecorder()

			authContoller.Endpoints["/register"](response, request)

			assertStatus(t, response.Code, test.statusCode)
		}
	})

	t.Run("creates an user with expected email and fullname", func(t *testing.T) {
		spyCreateUser := &SpyCreateUser{}
		authContoller := NewAuthController(spyCreateUser)
		expectedFullname := "John"
		expectedEmail := "john@doe.com"
		expectedPassword := "zaq1@WSX"
		body := newPayload(t, map[string]string{
			"fullname": expectedFullname,
			"email":    expectedEmail,
			"password": expectedPassword,
		})

		request, _ := http.NewRequest(http.MethodPost, "/authentication/register", body)
		response := httptest.NewRecorder()

		authContoller.Endpoints["/register"](response, request)

		if spyCreateUser.fullname != expectedFullname {
			t.Errorf("called CreateUser with name %q but wanted %q", spyCreateUser.fullname, expectedFullname)
		}

		if spyCreateUser.email != expectedEmail {
			t.Errorf("called CreateUser with email %q but wanted %q", spyCreateUser.email, expectedEmail)
		}
	})

	t.Run("creates an user with hashed password", func(t *testing.T) {
		spyCreateUser := &SpyCreateUser{}
		authContoller := NewAuthController(spyCreateUser)
		expectedFullname := "John"
		expectedEmail := "john@doe.com"
		password := "zaq1@WSX"
		expectedPassword := fmt.Sprintf("hashed-%s", password)
		body := newPayload(t, map[string]string{
			"fullname": expectedFullname,
			"email":    expectedEmail,
			"password": password,
		})

		request, _ := http.NewRequest(http.MethodPost, "/authentication/register", body)
		response := httptest.NewRecorder()

		authContoller.Endpoints["/register"](response, request)

		if spyCreateUser.password != expectedPassword {
			t.Errorf("called CreateUser with hashed password equal %q but wanted %q", spyCreateUser.password, expectedPassword)
		}
	})

	t.Run("returns HTTP 400 if error different than 500 occured while creating an user", func(t *testing.T) {
		errMsg := "some-err"
		SpyError := &SpyCreateUserError{errorMessage: errMsg}
		authController := NewAuthController(SpyError)
		body := newPayload(t, map[string]string{
			"fullname": "John",
			"email":    "abc@local.local",
			"password": "zaq1@WSX",
		})

		request, _ := http.NewRequest(http.MethodPost, "/authentication/register", body)
		response := httptest.NewRecorder()

		authController.Endpoints["/register"](response, request)

		result := response.Result()

		if result.StatusCode != http.StatusBadRequest {
			t.Errorf("Wrong error HTTP Code! Got %d, wanted %d", result.StatusCode, http.StatusBadRequest)
		}

		var parsedResponse struct {
			Message string
		}
		defer result.Body.Close()
		json.NewDecoder(result.Body).Decode(&parsedResponse)

		fmt.Println(parsedResponse)
		if parsedResponse.Message != errMsg {
			t.Errorf("Wrong error message, got %s but expected %s", parsedResponse.Message, errMsg)
		}
	})

	t.Run("returns HTTP 500 if generating JWT failed", func(t *testing.T) {
		errMsg := "some-err"
		SpyError := &SpyGenerateJwtError{errorMessage: errMsg}
		authController := NewAuthController(SpyError)
		body := newPayload(t, map[string]string{
			"fullname": "John",
			"email":    "abc@local.local",
			"password": "zaq1@WSX",
		})

		request, _ := http.NewRequest(http.MethodPost, "/authentication/register", body)
		response := httptest.NewRecorder()

		authController.Endpoints["/register"](response, request)

		result := response.Result()

		if result.StatusCode != http.StatusInternalServerError {
			t.Errorf("Wrong error HTTP Code! Got %d, wanted %d", result.StatusCode, http.StatusBadRequest)
		}

		var parsedResponse struct {
			Message string
		}
		defer result.Body.Close()
		json.NewDecoder(result.Body).Decode(&parsedResponse)

		fmt.Println(parsedResponse)
		if parsedResponse.Message != errMsg {
			t.Errorf("Wrong error message, got %s but expected %s", parsedResponse.Message, errMsg)
		}
	})

	t.Run("sets session cookie with JWT", func(t *testing.T) {
		token := "abc"
		expiryTime, _ := time.Parse("DD-MM-YYYY", "10-10-2010")
		spyGenerateToken := &SpyGenerateToken{generatedToken: token, expiryTime: expiryTime}
		authController := NewAuthController(spyGenerateToken)
		body := newPayload(t, map[string]string{
			"fullname": "John",
			"email":    "abc@local.local",
			"password": "zaq1@WSX",
		})

		request, _ := http.NewRequest(http.MethodPost, "/authentication/register", body)
		response := httptest.NewRecorder()

		authController.Endpoints["/register"](response, request)

		result := response.Result()

		if result.StatusCode != http.StatusOK {
			t.Errorf("Wrong HTTP Code! Got %d, wanted %d", result.StatusCode, http.StatusBadRequest)
		}

		cookie := result.Cookies()[0]
		if cookie == nil {
			t.Errorf("Session cookie hasn't been set")
		}
		if cookie.Value != token {
			t.Errorf("Incorrect token in the cookie. Got %q wanted %q", cookie.Value, token)
		}
		if cookie.Path != "/" {
			t.Errorf("Incorrect path in the cookie")
		}
		if expiryTime.Sub(cookie.Expires) != 0 {
			t.Errorf("Incorrect expiry time")
		}
	})
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("Got %d but wanted %d HTTP Code", got, want)
	}
}

func newPayload(t testing.TB, body map[string]string) *bytes.Buffer {
	jsonBody, _ := json.Marshal(body)
	return bytes.NewBuffer(jsonBody)
}
