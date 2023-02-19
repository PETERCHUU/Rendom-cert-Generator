package postgre

import (
	"database/sql"
	"fmt"
	"linuxWebUI/model/config"
	"os"

	_ "github.com/lib/pq"
)

type PostgreDB struct {
	*sql.DB
}

// after make this ssl safe
func Login(db *config.DB) *sql.DB {
	connection := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", db.Ac, db.Pw, db.Ip, db.Port, db.Dataname)
	sql, err := sql.Open("postgres", connection)
	if err != nil {
		print(err.Error())
		os.Exit(1)
	}
	defer sql.Close()
	return sql
}

func (m *PostgreDB) Login() {
}
