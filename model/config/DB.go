package config

const (
	Mysql uint8 = iota
	Postgresql
	Mssql
	Mariadb
)

type (
	DB struct {
		Type     uint8  `yaml:"dbtype"` //env:"DB_Type"
		Ac       string `yaml:"dbac"`   //env:"DB_Ac"
		Pw       string `yaml:"dbpw"`   // env:"DB_Pw"
		Ip       string `yaml:"dbip"`   //env:"DB_Ip"
		Port     string `yaml:"dbport"` //env:"DB_Port"
		Dataname string
	}
)
