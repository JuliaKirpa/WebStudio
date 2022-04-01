package migrations

import (
	"database/sql"
	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upNewTables, downNewTables)
}

func upNewTables(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	return nil
}

func downNewTables(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
