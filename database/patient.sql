//add this before create table statement
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE common.patient (
	patient_id uuid NOT NULL DEFAULT uuid_generate_v4(),
	phone_number varchar NOT NULL,
	age int NULL,
	is_diabetic boolean NULL DEFAULT false,
	is_hypertensive boolean NULL DEFAULT false,
	is_asthmastic boolean NULL DEFAULT false,
	is_cardio_ischemic boolean NULL DEFAULT false,
	has_lung_disease boolean NULL DEFAULT false,
	has_kidney_disease boolean NULL DEFAULT false,
	is_smoker boolean NULL DEFAULT false,
	is_return_from_travel boolean NULL DEFAULT false,
	longitude double precision NOT NULL,
	latitude double precision NOT NULL,
	created_at timestamptz NOT NULL DEFAULT now(),
	CONSTRAINT patient_pk PRIMARY KEY (patient_id),
	CONSTRAINT patient_un UNIQUE (phone_number)
);

