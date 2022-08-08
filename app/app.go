package app

import (
	"html/template"
	"net/http"
	"time"

	"github.com/devasiajoseph/wemebox/website"
)

type PageData struct {
	Id             int                `storm:"id,increment" json:"Id" `
	Slug           string             `storm:"index,unique" json:"Slug"`
	Title          string             `json:"Title"`
	FileName       string             `storm:"index,unique" json:"FileName"`
	BaseFileName   string             `json:"BaseFileName"`
	HtmlContent    template.HTML      `json:"HtmlContent"`
	Markdown       string             `json:"Markdown"`
	PageType       string             `json:"PageType"`
	Meta           []website.MetaData `json:"Meta"`
	LastEdited     time.Time          `json:"LastEdited"`
	UserId         int                `storm:"index" json:"UserId"`
	Author         string             `storm:"index" json:"Author"`
	BasePageId     int                `storm:"index" json:"BasePageId"`
	TemplatePageId int                `storm:"index" json:"TemplatePageId"`
	Csrf           string
	StaticUrl      string
	KeyValue       string
	UAuthLoggedIn  bool
	LoggedInUser   string
	Message        string
	PageAuth       bool
	ItemCount      string
	RecapSiteKey   string
}

func RenderPage(w http.ResponseWriter, r *http.Request) {

}
