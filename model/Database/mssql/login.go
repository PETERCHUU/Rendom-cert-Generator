package mssql

import (
	"database/sql"
	"fmt"
	"linuxWebUI/model/config"
	"os"
)

type Mssqldb struct {
	*sql.DB
}

func Login(dbinfo *config.DB) *sql.DB {
	connectionurl := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", dbinfo.Ac, dbinfo.Pw, dbinfo.Ip, dbinfo.Port, dbinfo.Dataname)
	sql, err := sql.Open("sqlserver", connectionurl)
	if err != nil {
		print(err.Error())
		os.Exit(1)
	}
	defer sql.Close()
	return sql
}

func (m *Mssqldb) Login() {
}
