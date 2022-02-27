package repository

import (
	"database/sql"
	"testing"

	_ "github.com/proullon/ramsql/driver"
)

func TestInsert(t *testing.T) {

	db, err := sql.Open("ramsql", "TestInsert")
	if err != nil {
		t.Errorf("unable to open database")
	}
	defer db.Close()

}
