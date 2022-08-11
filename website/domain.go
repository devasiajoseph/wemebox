package website

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/devasiajoseph/wemebox/db/postgres"
)

type Domain struct {
	DomainID   int    `db:"domain_id" json:"domain_id"`
	DomainName string `db:"domain_name" json:"domain_name"`
	DomainDir  string `db:"domain_dir" json:"domain_dir"`
	Active     bool   `db:"active" json:"active"`
}

func GetDomainName(r *http.Request) string {
	host := r.Host
	host = strings.Split(host, ":")[0]
	host = strings.Replace(host, "www.", "", -1)
	return host
}

func SaveWebsite(w http.ResponseWriter, r *http.Request) {

}

var sqlInsertDomain = "insert into domain (domain_name,domain_dir) values (:domain_name,:domain_dir) returning domain_id;"
var sqlAddDomainUser = "insert into domain_user_role (role,user_account_id,domain_id) values (:role,:user_account_id,:domain_id);"

func (d *Domain) Create(ownerID int) error {
	db := postgres.Db
	rows, err := db.NamedQuery(sqlInsertDomain, d)
	if err != nil {
		log.Println("Error creating domain")
		log.Println(err)
		return err
	}
	if rows.Next() {
		rows.Scan(d.DomainID)
	}
	defer rows.Close()

	dr := DomainUserRole{
		DomainID:      d.DomainID,
		Role:          roleOwner,
		UserAccountID: ownerID,
	}
	_, err = db.NamedExec(sqlAddDomainUser, dr)

	if err != nil {
		log.Println("Error adding user role")
		log.Println(err)
		return err
	}

	return err
}

func (d *Domain) Credit(amount int, notes string) error {
	if d.DomainID == 0 {
		return errors.New("invalid domain ID")
	}

	trx := Trx{
		DomainID: d.DomainID,
		Credit:   amount,
		Notes:    notes,
	}

	err := trx.Execute()
	if err != nil {
		log.Printf("error crediting domain => ID: %d", d.DomainID)
	}
	return err
}

func (d *Domain) Debit(amount int, notes string) error {
	if d.DomainID == 0 {
		return errors.New("invalid domain ID")
	}

	trx := Trx{
		DomainID: d.DomainID,
		Debit:    amount,
		Notes:    notes,
	}

	err := trx.Execute()
	if err != nil {
		log.Printf("error debiting domain => ID: %d", d.DomainID)
	}
	return err
}
