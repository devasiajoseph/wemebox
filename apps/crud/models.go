package crud

import (
	"log"

	"github.com/devasiajoseph/wemebox/db/postgres"
)

//Delete  data
func Delete(obj string, objID int) error {
	db := postgres.Db

	_, err := db.Exec("delete from "+obj+" where post_id=$1;", objID)
	if err != nil {
		log.Println("models:Error deleting " + obj)
		log.Println(err)
	}

	return err
}
