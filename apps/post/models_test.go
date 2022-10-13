package post

import (
	"fmt"
	"log"
	"testing"

	"github.com/devasiajoseph/wemebox/core"
	"github.com/devasiajoseph/wemebox/db/postgres"
)

func TestList(t *testing.T) {
	core.Start()
	postgres.InitDB()
	l := List{
		DomainID: 1,
		Page:     1,
		Limit:    50,
	}

	err := l.Fetch()
	if err != nil {
		t.Errorf("Error fetching posts")
		log.Println(err)
	}

	fmt.Printf("%d records fetched \n", len(l.Data))
}
