package repository

import "c19/connector/pgsql"


type SecurityRepository interface {

}

type repository struct {
	db *pgsql.DB
}

func NewSecurityRepository(db *pgsql.DB) SecurityRepository {
	return repository{db: db}
}