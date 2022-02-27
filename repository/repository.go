package repository

import (
	"database/sql"
	"log"
)

func Schema(db *sql.DB) error {

	result, err := db.Exec(`CREATE TABLE current (
		product_id      TEXT NOT NULL,
		product_name    TEXT NOT NUll,
		product_type    TEXT NOT NULL,
		edition_id      TEXT NOT NULL
		edition_name    TEXT NOT NULL,
		burn_id         TEXT NOT NULL,
		PRIMARY KEY (product_name, edition_name)
	)`)
	if err != nil {
		return err
	}
	log.Printf("Create current : %v", result)
	return nil
}
