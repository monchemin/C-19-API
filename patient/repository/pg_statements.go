package repository

const (
	insertNewPatient = `INSERT INTO common.patient(
			phone_number, 
			age, 
			weight, 
			is_diabetic, 
			is_hypertensive, 
			is_asthmatic, 
			is_cardio_ischemic, 
			has_lung_disease, 
			has_kidney_disease, 
			is_smoker, 
			is_return_from_travel,
			longitude, 
			latitude,
			district_id,
			gender)
	VALUES(
			:phonenumber, 
			:age,
			:weight,
			:isdiabetic, 
			:ishypertensive, 
			:isasthmatic, 
			:iscardioischemic, 
			:haslungdisease, 
			:haskidneydisease, 
			:issmoker, 
			:isreturnfromtravel, 
			:longitude,
			:latitude,
			:districtid,
			:gender) 
	RETURNING id`

	insertNewConstant = `INSERT INTO common.health_constant(
							patient_id,
							temperature,
							is_tired ,
							has_dry_cough,
							has_been_in_contact_with_infected_person,
							has_shortness_of_breath,
							has_headache ,
							has_runny_nose ,
							has_nasal_congestion ,
							has_sore_throat,
							has_muscle_pain,
							has_diarrhea
							)
					VALUES(	:patientid,
							:temperature,
							:istired,
							:hasdrycough,
							:hasbeenincontactwithinfectedperson,
							:hasshortnessofbreath,
							:hasheadache,
							:hasrunnynose,
							:hasnasalcongestion,
							:hassorethroat,
							:hasmusclepain,
							:hasdiarrhea)
					RETURNING id, date_time`

	getPatient = `SELECT p.*, d.name as "district_name", d.town_id, 
	            t.name as "town_name", t.longitude as "town_longitude", t.latitude as "town_latitude",
	            c.id as "country_code", c.name as "country_name", c.iso_code as "iso_code"
				FROM common.patient p
				inner join common.district d on d.id = p.district_id
				inner join common.town t on t.id = d.town_id 
				inner join common.country c on c.id = t.country_id
 				WHERE p.phone_number = $1 OR p.id::TEXT = $1`

	getPatientById = `SELECT p.* FROM common.patient p WHERE p.id = $1::uuid`

	getPatientHealthConstants = `SELECT hc.* FROM common.health_constant hc WHERE hc.patient_id = $1 ORDER BY hc.date_time DESC`
)

type Search int

const (
	ID Search = iota
	PHONE
)
