package repository

import _ "c19/connector/pgsql"

type PatientRepository interface {

}

type repository struct {
	db *pgsql.DB
}
