package web_test

import (
	"github.com/adamluzsi/testcase"
	"github.com/kelseyhightower/envconfig"
	"github.com/mikejeuga/book_river_bookshop/models"
	"github.com/mikejeuga/book_river_bookshop/src/bookshop"
	"github.com/mikejeuga/book_river_bookshop/src/web"
	"github.com/mikejeuga/book_river_bookshop/src/web/auth"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {
	s := testcase.NewSpec(t)

	var testConfig auth.Config
	err := envconfig.Process("", &testConfig)
	if err != nil {
		t.Fatal("Could not load environment variables!")
	}

	var (
		stock    = make(models.Stock)
		bookShop = models.NewBookShop(stock)
		service  = bookshop.NewService(bookShop)
		server   = web.NewServer(testConfig, service)
		res      = httptest.NewRecorder()
		req      = httptest.NewRequest(http.MethodGet, "/", nil)
	)

	act := func(t *testcase.T) {
		req.Header.Set("X-API-KEY", testConfig.BigSecret)
		server.Handler.ServeHTTP(res, req)
	}

	s.Describe("the server is", func(s *testcase.Spec) {
		s.Test("healthy on '/'", func(t *testcase.T) {
			act(t)
			t.Must.Equal(http.StatusOK, res.Code)
		})
	})
}
