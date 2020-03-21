CREATE TABLE common.health_constant (
	id uuid NOT NULL DEFAULT uuid_generate_v4(),
	date_time timestamptz NOT NULL DEFAULT now(),
	temperature float4 NOT NULL,
	is_tired boolean NOT NULL,
	has_dry_cough boolean NOT NULL,
	has_shortness_of_breath boolean NOT NULL,
	has_headache boolean NOT NULL,
	has_runny_nose boolean NOT NULL,
	has_nasal_congestion boolean NOT NULL,
	has_sore_throat boolean NOT NULL,
	has_muscle_pain boolean NOT NULL,
	patient_id uuid NOT NULL,
	has_diarrhea boolean NOT NULL,
	CONSTRAINT health_constant_pk PRIMARY KEY (id),
	CONSTRAINT health_constant_fk FOREIGN KEY (patient_id) REFERENCES common.patient(patient_id)
);
CREATE INDEX health_constant_date_time_idx ON common.health_constant (date_time);
CREATE INDEX health_constant_temperature_idx ON common.health_constant (temperature,is_tired,has_dry_cough,has_shortness_of_breath);
CREATE INDEX health_constant_has_diarrhea_idx ON common.health_constant (has_diarrhea,has_muscle_pain,has_sore_throat,has_nasal_congestion,has_ranny_nose,has_headache);
