package repository

type CountryResult struct {
	ID                 string    `db:"id"`
	Name               string    `db:"name"`
	IsoCode            string    `db:"iso_code"`
}