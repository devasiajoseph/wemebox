package page

import (
	"time"

	"github.com/devasiajoseph/wemebox/uauth"
)

type Page struct {
	PageID       int       `db:"page_id" json:"page_id"`
	PageType     string    `db:"page_type" json:"page_type"`
	PageTitle    string    `db:"page_title" json:"page_title"`
	PageMarkdown string    `db:"page_markdown" json:"page_markdown"`
	PageHTML     string    `db:"page_html" json:"page_html"`
	PageSlug     string    `db:"page_slug" json:"page_slug"`
	LastEdited   time.Time `db:"last_edited" json:"last_edited"`
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
