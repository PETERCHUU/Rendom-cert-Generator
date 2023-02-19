package database

import (
	"linuxWebUI/model/Database/maria"
	"linuxWebUI/model/Database/mssql"
	"linuxWebUI/model/Database/mysqldb"
	"linuxWebUI/model/Database/postgre"
	"linuxWebUI/model/config"
)

type Sqldo interface {
	Login()
}

func Newmysql(db *config.DB) Sqldo {
	switch db.Type {
	case config.Postgresql:
		return &postgre.PostgreDB{
			DB: postgre.Login(db),
		}
	case config.Mysql:
		return &mysqldb.Mysqldb{
			DB: mysqldb.Login(db),
		}
	case config.Mariadb:
		return &maria.Mariadb{
			DB: maria.Login(db),
		}
	case config.Mssql:
		return &mysqldb.Mysqldb{
			DB: mssql.Login(db),
		}
	}
	return nil
}
