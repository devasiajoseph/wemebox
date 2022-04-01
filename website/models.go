package website

import (
	"log"
	"time"

	"github.com/devasiajoseph/wemebox/db/postgres"
	"github.com/devasiajoseph/wemebox/uauth"
)

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
	Author        uauth.UserAccount
}

type Domain struct {
	DomainID   int    `db:"domain_id" json:"domain_id"`
	DomainName string `db:"domain_name" json:"domain_name"`
	DomainDir  string `db:"domain_dir" json:"domain_dir"`
}

type PageData struct {
	Title         string
	Content       string
	PageFile      string
	BasePageFile  string
	PageData      interface{}
	Csrf          string
	StaticUrl     string
	UAuthLoggedIn bool
	LoggedInUser  string
}

func GetDomainPage(slug string, domain string) (Page, error) {
	db := postgres.Db
	var page Page
	err := db.Get(&page, "select * from page where url_slug=$1", slug)
	if err != nil {
		log.Println("Error getting page")
		log.Println(err)
	}
	return page, err
}
