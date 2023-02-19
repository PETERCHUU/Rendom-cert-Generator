package maria

import (
	"database/sql"
	"fmt"
	"linuxWebUI/model/config"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Mariadb struct {
	*sql.DB
}

func Login(db *config.DB) *sql.DB {
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", db.Ac, db.Pw, db.Ip, db.Port, db.Dataname)
	sql, err := sql.Open("mysql", connection)
	if err != nil {
		print(err.Error())
		os.Exit(1)
	}
	defer sql.Close()
	return sql
}

func (m *Mariadb) Login() {}
