/*
 * Copyright (C) Centipair Technologies Private Limited - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written By Devasia Joseph <devasiajoseph@centipair.com>, June 2019
 */

package postgres

import (
	//"database/sql"
	"fmt"

	"github.com/devasiajoseph/wemebox/core"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	//"reflect"
	"log"
)

var Db *sqlx.DB

func ConnectDB() *sqlx.DB {
	log.Println("Connecting to postgres database => ", core.Config.SQLDBName)
	psqlInfo := fmt.Sprintf("host=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		core.Config.SQLHost, core.Config.SQLUser, core.Config.SQLPassword, core.Config.SQLDBName)
	db, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	//db.SetMaxOpenConns(1000)
	return db
}

//InitDb Initialized database connection
func InitDB() {
	Db = ConnectDB()
}
