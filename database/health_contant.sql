CREATE TABLE common.health_constant (
	id uuid NOT NULL DEFAULT common.uuid_generate_v4(),
	patient_id uuid NOT NULL,
	date_time timestamptz NOT NULL DEFAULT now(),
	temperature float4 NOT NULL,
	is_tired bool NOT NULL,
	has_dry_cough bool NOT NULL,
	has_been_in_contact_with_infected_person bool NOT NULL,
	has_shortness_of_breath bool NOT NULL,
	has_headache bool NOT NULL,
	has_runny_nose bool NOT NULL,
	has_nasal_congestion bool NOT NULL,
	has_sore_throat bool NOT NULL,
	has_muscle_pain bool NOT NULL,
	has_diarrhea bool NOT NULL,
	CONSTRAINT health_constant_pk PRIMARY KEY (id),
	CONSTRAINT health_constant_fk FOREIGN KEY (patient_id) REFERENCES common.patient(id)
);
CREATE INDEX health_constant_date_time_idx ON common.health_constant USING btree (date_time);
CREATE INDEX health_constant_has_diarrhea_idx ON common.health_constant USING btree (has_diarrhea, has_muscle_pain, has_sore_throat, has_nasal_congestion, has_runny_nose, has_headache);
CREATE INDEX health_constant_temperature_idx ON common.health_constant USING btree (temperature, is_tired, has_dry_cough, has_shortness_of_breath);
