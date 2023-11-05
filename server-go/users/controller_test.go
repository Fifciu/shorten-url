package users

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type SpyCreateUser struct {
	name     string
	email    string
	password string
	called   bool
}

func (s *SpyCreateUser) CreateUser(name, email, password string) (*User, error) {
	s.name = name
	s.email = email
	s.password = password
	s.called = true
	return &User{name, email, password}, nil
}

func (s *SpyCreateUser) HashPassword(password string) (string, error) {
	return fmt.Sprintf("hashed-%s", password), nil
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
					"name":  "John",
					"email": "john@doe.com",
				},
				http.StatusBadRequest,
			},
			{
				map[string]string{
					"name":     "John",
					"email":    "johndoe.com",
					"password": "zaq1@WSX",
				},
				http.StatusBadRequest,
			},
			{
				map[string]string{
					"name":     "John",
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

	t.Run("creates an user with expected email and name", func(t *testing.T) {
		spyCreateUser := &SpyCreateUser{}
		authContoller := NewAuthController(spyCreateUser)
		expectedName := "John"
		expectedEmail := "john@doe.com"
		expectedPassword := "zaq1@WSX"
		body := newPayload(t, map[string]string{
			"name":     expectedName,
			"email":    expectedEmail,
			"password": expectedPassword,
		})

		request, _ := http.NewRequest(http.MethodPost, "/authentication/register", body)
		response := httptest.NewRecorder()

		authContoller.Endpoints["/register"](response, request)

		if spyCreateUser.name != expectedName {
			t.Errorf("called CreateUser with name %q but wanted %q", spyCreateUser.name, expectedName)
		}

		if spyCreateUser.email != expectedEmail {
			t.Errorf("called CreateUser with email %q but wanted %q", spyCreateUser.email, expectedEmail)
		}
	})

	t.Run("creates an user with hashed password", func(t *testing.T) {
		spyCreateUser := &SpyCreateUser{}
		authContoller := NewAuthController(spyCreateUser)
		expectedName := "John"
		expectedEmail := "john@doe.com"
		password := "zaq1@WSX"
		expectedPassword := fmt.Sprintf("hashed-%s", password)
		body := newPayload(t, map[string]string{
			"name":     expectedName,
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
