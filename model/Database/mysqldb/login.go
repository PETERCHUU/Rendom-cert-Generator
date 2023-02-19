package mysqldb

import (
	"database/sql"
	"fmt"
	"linuxWebUI/model/config"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Mysqldb struct {
	*sql.DB
}

func Login(dbinfo *config.DB) *sql.DB {
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbinfo.Ac, dbinfo.Pw, dbinfo.Ip, dbinfo.Port, dbinfo.Dataname)
	sql, err := sql.Open("mysql", connection)
	if err != nil {
		print(err.Error())
		os.Exit(1)
	}
	defer sql.Close()
	return sql
}

func (m *Mysqldb) Login() {}
