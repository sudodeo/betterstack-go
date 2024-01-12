package uptime_test

import (
	"log"
	"os"
	"testing"

	"github.com/sudodeo/betterstack-go/sdk/uptime"
)

var bs *uptime.Betterstack

func TestMain(m *testing.M) {
	var err error
	bs, err = uptime.NewFromENV()
	if err != nil {
		log.Fatal("could not initialise test: ", err)
	}

	if bs.Token == "" {
		log.Fatal("could not load Token")
	}
	os.Exit(m.Run())
}
