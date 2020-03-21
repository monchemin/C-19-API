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
			latitude)
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
			:latitude) 
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
					RETURNING id`

	getPatient = `SELECT p.* FROM common.patient p WHERE p.phone_number = $1`

	getPatientHealthConstants = `SELECT hc.* FROM common.health_constant hc WHERE hc.patient_id = $1 ORDER BY hc.date_time DESC`
)
