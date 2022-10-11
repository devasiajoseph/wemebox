package post

import (
	"log"
	"time"

	"github.com/devasiajoseph/wemebox/apps/crud"
	"github.com/devasiajoseph/wemebox/cpmath"
	"github.com/devasiajoseph/wemebox/db/postgres"
)

type Object struct {
	PostID   int       `db:"post_id" json:"post_id"`
	PostRaw  string    `db:"post_raw" json:"post_raw"`
	Created  time.Time `db:"created" json:"created"`
	DomainID int       `db:"domain_id" json:"domain_id"`
}

type PostImage struct {
	PostImageID        int `db:"post_image_id" json:"post_image_id"`
	PostImage          int `db:"post_image" json:"post_image"`
	PostImageThumbnail int `db:"post_image_thumbnail" json:"post_image_thumbnail"`
	PostID             int `db:"post_id" json:"post_id"`
}

type List struct {
	Data     []Object `json:"data"`
	Total    int      `json:"total"`
	Page     int      `json:"page"`
	Limit    int      `json:"limit"`
	Offset   int      `json:"offset"`
	DomainID int      `db:"domain_id" json:"domain_id"`
}

var sqlInsert = "insert into post (post_raw,domain_id) values (:post_raw,:domain_id) returning post_id;"
var sqlUpdate = "update post set post_raw=:post_raw where post_id=:post_id;"

func (obj *Object) Create() error {
	db := postgres.Db
	rows, err := db.NamedQuery(sqlInsert, &obj)

	if err != nil {
		log.Println("Error creating company")
		log.Println(err)
		return err
	}

	if rows.Next() {
		rows.Scan(&obj.PostID)
	}
	defer rows.Close()
	return err
}

func (obj *Object) Update() error {
	db := postgres.Db
	_, err := db.NamedExec(sqlUpdate, &obj)
	if err != nil {
		log.Println(err)
		log.Println("Error updating post")
	}
	return err
}

func (obj *Object) Save() error {
	if obj.PostID == 0 {
		return obj.Create()
	}
	return obj.Update()
}

//Delete  data
func (obj *Object) Delete() error {
	err := crud.Delete("post", obj.PostID)
	return err
}

var sqlList = "select post_id,post_raw from post where domain_id=$1 limit $2 offset $3;"

func (ol *List) Fetch() error {
	db := postgres.Db
	ol.Offset = cpmath.Offset(ol.Page, ol.Limit)
	err := db.Select(&ol.Data, sqlList, ol.DomainID, ol.Limit, ol.Offset)
	if err != nil {
		log.Println(err)
		log.Println("Error fetching post list")
	}

	return err
}
