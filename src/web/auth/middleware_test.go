package auth_test

import (
	"github.com/adamluzsi/testcase/assert"
	"github.com/kelseyhightower/envconfig"
	"github.com/mikejeuga/book_river_bookshop/src/web/auth"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMiddleware(t *testing.T) {
	var littleSecret auth.Config
	err := envconfig.Process("", &littleSecret)
	if err != nil {
		t.Fatal("Could not load environment variables!")
	}

	for _, tc := range []struct {
		description      string
		res              *httptest.ResponseRecorder
		req              *http.Request
		authHeader       string
		expectCode       int
		expectNextCalled bool
	}{
		{
			description:      "when the request has the correct headers",
			res:              httptest.NewRecorder(),
			req:              httptest.NewRequest(http.MethodGet, "/", nil),
			authHeader:       littleSecret.BigSecret,
			expectCode:       http.StatusOK,
			expectNextCalled: true,
		},
		{
			description:      "when the request does not have the correct headers",
			res:              httptest.NewRecorder(),
			req:              httptest.NewRequest(http.MethodGet, "/", nil),
			authHeader:       "veryBad",
			expectCode:       http.StatusUnauthorized,
			expectNextCalled: false,
		},
	} {
		t.Run(tc.description, func(t *testing.T) {
			tc.req.Header.Set("X-API-KEY", tc.authHeader)
			nextCalled := false
			testHandler := func(w http.ResponseWriter, r *http.Request) {
				nextCalled = true
			}

			authMiddleware := auth.FOMW(littleSecret.BigSecret)

			authMiddleware(http.HandlerFunc(testHandler)).ServeHTTP(tc.res, tc.req)

			assert.Equal(t, tc.expectCode, tc.res.Code)
			assert.Equal(t, tc.expectNextCalled, nextCalled)
		})
	}
}
