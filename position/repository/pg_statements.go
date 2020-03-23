package repository

const (
	insertNewCountry = `INSERT INTO common.country(id, name)
					VALUES(:id, :name)`
	insertNewTown = `INSERT INTO common.town(name, country_id)
					VALUES(:name, :countryid) 
					RETURNING id`
	insertNewDistrict = `INSERT INTO common.district(name, town_id)
					VALUES(:name, :townid) 
					RETURNING id`
)
