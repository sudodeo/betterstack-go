package logs

import (
	"log"
	"os"
	"testing"

	_ "github.com/joho/godotenv"
)

var bs *Betterstack

func TestMain(m *testing.M) {
	var err error
	bs, err = NewFromENV()
	if err != nil {
		log.Fatal("could not initialise test: ", err)
	}
	if bs.token == "" {
		log.Fatal("could not load token")
	}
	os.Exit(m.Run())
}
