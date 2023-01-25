package testclient_test

import (
	"github.com/adamluzsi/testcase/assert"
	"github.com/kelseyhightower/envconfig"
	"github.com/mikejeuga/book_river_bookshop/blackboxtests/testclient"
	"github.com/mikejeuga/book_river_bookshop/models"
	"testing"
)

func TestName(t *testing.T) {

	var testConfig testclient.Config
	err := envconfig.Process("", &testConfig)
	if err != nil {
		t.Fatal("Could not load environment variables!")
	}
	testLibrarian := testclient.NewTestLibrarian(testConfig)

	book := models.Book{
		Title: "First To Go",
		Author: models.Author{
			FirstName: "Mike",
			LastName:  "Jeuga",
		},
		Edition: 2010,
	}

	uuid, err := testLibrarian.Add(book)
	assert.NoError(t, err)

	_, err = testLibrarian.GetBookBy(uuid)
	assert.NoError(t, err)
}
