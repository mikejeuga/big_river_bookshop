package web_test

import (
	"github.com/adamluzsi/testcase"
	"github.com/mikejeuga/book_river_bookshop/models"
	"github.com/mikejeuga/book_river_bookshop/src/bookshop"
	"github.com/mikejeuga/book_river_bookshop/src/web"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {
	s := testcase.NewSpec(t)

	var (
		stock    = make(models.Stock)
		bookShop = models.NewBookShop(stock)
		service  = bookshop.NewService(bookShop)
		server   = web.NewServer(service)
		res      = httptest.NewRecorder()
		req      = httptest.NewRequest(http.MethodGet, "/", nil)
	)

	act := func(t *testcase.T) {
		server.Handler.ServeHTTP(res, req)
	}

	s.Describe("the server is", func(s *testcase.Spec) {
		s.Test("healthy on '/'", func(t *testcase.T) {
			act(t)
			t.Must.Equal(http.StatusOK, res.Code)
		})
	})
}
