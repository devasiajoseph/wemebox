package crud

import (
	"log"
	"testing"

	"github.com/devasiajoseph/wemebox/core"
	"github.com/devasiajoseph/wemebox/db/postgres"
)

func TestDelete(t *testing.T) {
	core.Start()
	postgres.InitDB()
	err := Delete("post", 1)
	if err != nil {
		t.Errorf("Error deleting object")
		log.Println(err)
	}
}
