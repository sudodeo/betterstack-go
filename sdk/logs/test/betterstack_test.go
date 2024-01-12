package logs

import (
	"log"
	"os"
	"testing"

	"github.com/sudodeo/betterstack-go/sdk/logs"
)

var bs *logs.Betterstack

func TestMain(m *testing.M) {
	var err error
	bs, err = logs.NewFromENV()
	if err != nil {
		log.Fatal("could not initialise test: ", err)
	}
	if bs.Token == "" {
		log.Fatal("could not load token")
	}
	os.Exit(m.Run())
}
