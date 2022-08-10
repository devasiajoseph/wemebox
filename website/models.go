package website

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/devasiajoseph/wemebox/db/postgres"
)

var roleAdmin = "admin"
var roleOwner = "owner"
var roleEditor = "editor"

type DomainPage struct {
	PageFile     string `db:"page_file" json:"page_file"`
	BasePageFile string `db:"base_page_file" json:"base_page_file"`
	DomainDir    string `db:"domain_dir" json:"domain_dir"`
}

type Page struct {
	PageID       int    `db:"page_id" json:"page_id"`
	PageSlug     string `db:"page_slug" json:"page_slug"`
	PageFile     string `db:"page_file" json:"page_file"`
	BasePageFile string `db:"base_page_file" json:"base_page_file"`
	DomainID     int    `db:"domain_id" json:"domain_id"`
	CSRF         string
}

type Blog struct {
	BlogID        int       `db:"page_id" json:"page_id"`
	BlogTitle     string    `db:"page_title" json:"page_title"`
	BlogMarkdown  string    `db:"page_markdown" json:"page_markdown"`
	BlogHTML      string    `db:"page_html" json:"page_html"`
	BlogSlug      string    `db:"page_slug" json:"page_slug"`
	LastEdited    time.Time `db:"last_edited" json:"last_edited"`
	UserAccountID int       `db:"user_account_id" json:"UserAccountID"`
	//Author        uauth.UserAccount
}

type Domain struct {
	DomainID   int    `db:"domain_id" json:"domain_id"`
	DomainName string `db:"domain_name" json:"domain_name"`
	DomainDir  string `db:"domain_dir" json:"domain_dir"`
	Active     bool   `db:"active" json:"active"`
}

type DomainUserRole struct {
	RoleID        int    `db:"role_id" json:"role_id"`
	DomainID      int    `db:"domain_id" json:"domain_id"`
	UserAccountID int    `db:"user_account_id" json:"UserAccountID"`
	Role          string `db:"role" json:"role"`
}

type PageData struct {
	Title        string
	Content      string
	PageFile     string `db:"page_file" `
	BasePageFile string `db:"base_page_file" `
	DomainID     int    `db:"domain_id"`
	DomainDir    string `db:"domain_dir"`
	Domain       string
	//PageData      interface{}
	Csrf          string
	StaticUrl     string
	UAuthLoggedIn bool
	LoggedInUser  string
	Slug          string
}

type MetaData struct {
	Id      int    `storm:"id,increment" json:"id" `
	PageId  int    `storm:"index" json:"PageId"`
	Name    string `storm:"index" json:"Name"`
	Content string `json:"Content"`
}

func (d *Domain) Clean() {
	d.DomainName = strings.ReplaceAll(d.DomainName, " ", "")
}

func (d *Domain) Fetch() error {
	d.Clean()
	db := postgres.Db
	err := db.Get(d, "select * from domain where domain_name=$1", d.DomainName)
	if err != nil {
		log.Println("Error in getting domain from domain name")
		log.Println(err)
	}
	return err
}

var sqlDomainPage = "select domain.domain_dir,page.page_file,page.base_page_file " +
	"from domain left join page on domain.domain_id=page.domain_id " +
	"where domain.domain_name=$1 and page.page_slug=$2;"

func GetDomainPage(slug string, domain string) (DomainPage, error) {
	db := postgres.Db
	var page DomainPage
	err := db.Get(&page, sqlDomainPage, domain, slug)
	if err != nil {
		log.Println("Error getting page")
		log.Println(err)
	}
	return page, err
}

func (pd *PageData) DomainPage(r *http.Request) error {
	db := postgres.Db
	pd.Domain = GetDomain(r)
	err := db.Get(pd, sqlDomainPage, pd.Domain, pd.Slug)
	if err != nil {
		log.Println("Error getting page")
		log.Println(err)
	}

	return err
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
