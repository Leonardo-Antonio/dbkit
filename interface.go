package dbkit

import (
	"database/sql"
	"log"
)

type driver string

const MYSQL driver = "MYSQL"

type IDBKit interface {
	NewItem(table string, data any) (sql.Result, error)
	SelectItems(data QueryInfo) ResponseQuery
}

func New(driver driver, db *sql.DB) IDBKit {
	switch driver {
	case MYSQL:
		return newKit(db)
	}

	log.Println("❌ error: driver " + driver + " not configured")
	return nil
}
