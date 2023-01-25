package main

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/mikejeuga/book_river_bookshop/models"
	"github.com/mikejeuga/book_river_bookshop/src/bookshop"
	"github.com/mikejeuga/book_river_bookshop/src/web"
	"github.com/mikejeuga/book_river_bookshop/src/web/auth"
	"log"
)

func main() {
	var c auth.Config
	err := envconfig.Process("", &c)
	if err != nil {
		log.Fatal("Could not load environment variables!")
	}

	stock := make(models.Stock)
	bookShop := models.NewBookShop(stock)
	service := bookshop.NewService(bookShop)
	server := web.NewServer(c, service)

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}
