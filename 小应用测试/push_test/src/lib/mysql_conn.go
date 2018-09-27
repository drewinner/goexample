package lib

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

func ConnMysql() (db *sql.DB) {

	db, err := sql.Open("mysql", "root:123456@/gopush")
	if err != nil {
		log.Println("connection mysql error = ", err)
		return nil
	}
	return db

}
