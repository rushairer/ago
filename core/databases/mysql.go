package databases

import (
	_ "github.com/go-sql-driver/mysql"

	"github.com/jmoiron/sqlx"
)

func NewMysql(dsn string) (*sqlx.DB, error) {
	if mysql, err := sqlx.Open("mysql", dsn); err == nil {
		if err = mysql.Ping(); err == nil {
			return mysql, nil
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}
