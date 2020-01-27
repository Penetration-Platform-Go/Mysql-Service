package model

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Result Define
type Result struct {
	IsValid bool
	Value   *sql.Rows
}

var db *sql.DB

func init() {
	db, _ = sql.Open("mysql", "root:123456@/Platform")
}
