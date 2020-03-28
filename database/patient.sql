//add this before create table statement
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE common.patient (
	id uuid NOT NULL DEFAULT common.uuid_generate_v4(),
	phone_number varchar NOT NULL,
	age int4 NULL,
	weight float8 NULL,
	is_diabetic bool NULL DEFAULT false,
	is_hypertensive bool NULL DEFAULT false,
	is_asthmatic bool NULL DEFAULT false,
	is_cardio_ischemic bool NULL DEFAULT false,
	has_lung_disease bool NULL DEFAULT false,
	has_kidney_disease bool NULL DEFAULT false,
	is_smoker bool NULL DEFAULT false,
	is_return_from_travel bool NULL DEFAULT false,
	created_at timestamptz NOT NULL DEFAULT now(),
	longitude float8 NOT NULL,
	latitude float8 NOT NULL,
	district_id uuid NOT NULL,
	gender varchar(1) NULL,
	CONSTRAINT patient_pk PRIMARY KEY (id),
	CONSTRAINT patient_un UNIQUE (phone_number)
);

