package dsnbuilder

import (
	"errors"
	"fmt"
)

var (
	dialectPostgres = "postgres"
	dialectMysql    = "mysql"

	errDialectNotFound         = errors.New("dialect not found")
	errDBCredentialNotComplete = errors.New("db credential not complete")
)

//DsnBuilder helps to generate a DSN for supported dialects,
//now is MySQL and Postgres
type DsnBuilder struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

func New(host string, port int, username string, password string, database string) *DsnBuilder {
	return &DsnBuilder{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
		Database: database,
	}
}

//Build return string of DSN, return error if dialect given is not in supported
//or lack of credential data
func (builder *DsnBuilder) Build(dialect string) (string, error) {
	if dialect == dialectMysql {
		return builder.mysql()
	}

	if dialect == dialectPostgres {
		return builder.postgres()
	}

	return "", errDialectNotFound
}

func (builder *DsnBuilder) mysql() (string, error) {
	return "", nil
}

func (builder *DsnBuilder) postgres() (string, error) {

	//Check if minimum configuration dsn is fulfilled
	if builder.Host == "" ||
		builder.Database == "" ||
		builder.Username == "" {
		return "", errDBCredentialNotComplete
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		builder.Host, builder.Port, builder.Username, builder.Database, builder.Password)

	return dsn, nil
}
