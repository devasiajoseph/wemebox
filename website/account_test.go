package website

import (
	"log"
	"testing"

	"github.com/devasiajoseph/wemebox/core"
	"github.com/devasiajoseph/wemebox/db/postgres"
)

func TestTrx(t *testing.T) {
	core.Start()
	postgres.InitDB()
	trx := Trx{
		DomainID: 1,
		Credit:   30,
	}

	err := trx.Execute()
	if err != nil {
		log.Println(err)
		t.Errorf("Error crediting domain")
	}

	trx = Trx{
		DomainID: 1,
		Debit:    30,
	}

	err = trx.Execute()
	if err != nil {
		log.Println(err)
		t.Errorf("Error debiting domain")
	}
}
