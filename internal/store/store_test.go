package store

import (
	"os"
	"testing"
)

var dataBaseURL string

//MainTest вызывается один раз перед всеми тестами в конкретном пакете
func TestMain(m *testing.M){
	dataBaseURL = os.Getenv("TEST_DB_URL")
	if dataBaseURL == ""{
		dataBaseURL = "user=testuser password=qawsed345rf dbname=testdb sslmode=disable"
	}

	os.Exit(m.Run())
}
