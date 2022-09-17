package website

import (
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

type Website struct {
	DomainID   int    `db:"domain_id" json:"domain_id"`
	DomainName string `db:"domain_name" json:"domain_name"`
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

func GetUserWebsites(ownerID int) ([]Website, error) {
	db := postgres.Db
	var wb []Website
	sql := "select domain.domain_id,domain.domain_name,domain.active from domain left join domain_user_role" +
		" on domain.domain_id = domain_user_role.domain_id where domain_user_role.user_account_id=$1;"
	err := db.Select(&wb, sql, ownerID)
	if err != nil {
		log.Println("error getting user websites")
		log.Println(err)
	}
	return wb, err
}
