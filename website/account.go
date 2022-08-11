package website

import (
	"log"

	"github.com/devasiajoseph/wemebox/db/postgres"
)

type Account struct {
	AccountID int `db:"account_id" json:"account_id"`
	DomainID  int `db:"domain_id" json:"domain_id"`
	Balance   int `db:"balance" json:"balance"`
}

type Trx struct {
	TrxID    int    `db:"account_id" json:"account_id"`
	DomainID int    `db:"domain_id" json:"domain_id"`
	Credit   int    `db:"credit" json:"credit"`
	Debit    int    `db:"debit" json:"debit"`
	Notes    string `db:"notes" json:"notes"`
}

func updateBalance(dID int) error {
	return nil
}

func (trx *Trx) Execute() error {
	db := postgres.Db
	_, err := db.NamedExec("insert into trx (domain_id,credit,debit,notes) values (:domain_id,:credit,:debit,:notes)", trx)

	if err != nil {
		log.Println(err)
	}
	err = updateBalance(trx.DomainID)
	return err
}
