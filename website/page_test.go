package website

import (
	"fmt"
	"log"
	"testing"

	"github.com/devasiajoseph/wemebox/core"
	"github.com/devasiajoseph/wemebox/db/postgres"
)

func TestDomainPage(t *testing.T) {
	core.Start()
	postgres.InitDB()

	dp, err := GetDomainPage("about-us", "localhost")
	if err != nil {
		log.Println(err)
		t.Error("Error getting domain page")
	}

	fmt.Printf("%+v\n", dp)
}
