package repository

const (
	insertNewCountry = `INSERT INTO common.country(id, name, iso_code)
					VALUES(:id, :name, :isocode)`
	insertNewTown = `INSERT INTO common.town(name, country_id, longitude, latitude)
					VALUES(:name, :countryid, :longitude, :latitude) 
					RETURNING id`
	insertNewDistrict = `INSERT INTO common.district(name, town_id)
					VALUES(:name, :townid) 
					RETURNING id`

	getCountries = `SELECT * FROM common.country`

	getLocalizations = `SELECT d.id, c.id as "code", CONCAT(c.name, ' ', t.name, ' ', d.name) as position
						from common.district d
						inner join common.town t on t.id = d.town_id 
						inner join common.country c on c.id = t.country_id `
)
