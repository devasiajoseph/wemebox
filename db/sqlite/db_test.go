package sqlite

import (
	"log"
	"testing"
)

func TestConnect(t *testing.T) {
	InitDB()
	schema := `CREATE TABLE place (
		country text,
		city text NULL,
		telcode integer);`

	// execute a query on the server
	result, err := DB.Exec(schema)

	if err != nil {
		log.Println(err)
		t.Errorf("Error creating schema")
	}

	log.Println(result)

}
