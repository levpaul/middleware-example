package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-martini/martini"
	"github.com/stretchr/testify/mock"
)

// The Mock and related function(s) used to take place of the AuthManager
type AuthManagerMock struct{ mock.Mock }

func (a *AuthManagerMock) AuthenticateRequest(req *http.Request) bool {
	args := a.Mock.Called(req)
	return args.Bool(0)
}

// Struct to hold relevant objects for tests
type suite struct {
	aMock *AuthManagerMock
	req   *http.Request
	res   *httptest.ResponseRecorder
	m     *martini.Martini
}

// A function which is called before each test
func createSuite() *suite {
	s := new(suite)

	// Create a new Mock each time
	s.aMock = new(AuthManagerMock)

	s.req, _ = http.NewRequest("GET", "", nil)
	s.res = httptest.NewRecorder()

	// Setup a small martini with the AuthManagerMock and the Authentication middleware
	s.m = martini.New()
	s.m.Map(s.aMock)
	s.m.Use(Authentication())

	return s
}

// Actual Tests
func TestAuthenticationSuccess(t *testing.T) {
	s := createSuite()

	s.aMock.On("AuthenticateRequest", s.req).Return(true)
	s.m.ServeHTTP(s.res, s.req)

	expectedResponseCode := 200

	if s.res.Code != expectedResponseCode {
		t.Errorf(`Unexpected error code "%d", was expecting "%d"`, s.res.Code, expectedResponseCode)
	}
}

func TestAuthenticationFailure(t *testing.T) {
	s := createSuite()

	s.aMock.On("AuthenticateRequest", s.req).Return(false)
	s.m.ServeHTTP(s.res, s.req)

	expectedResponseCode := 401

	if s.res.Code != expectedResponseCode {
		t.Errorf(`Unexpected error code "%d", was expecting "%d"`, s.res.Code, expectedResponseCode)
	}
}
