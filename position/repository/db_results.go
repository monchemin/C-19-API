package repository

import "github.com/google/uuid"

type CountryResult struct {
	ID      string    `db:"id"`
	Name    string    `db:"name"`
	IsoCode string    `db:"iso_code"`
	UID     uuid.UUID `db:"uid"`
}
