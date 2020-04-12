-- common.test_result definition

-- Drop table

-- DROP TABLE common.test_result;

CREATE TABLE common.test_result (
	id uuid NOT NULL DEFAULT uuid_generate_v4(),
	patient_id uuid NOT NULL,
	test_code varchar NOT NULL,
	date_time timestamptz NOT NULL DEFAULT now(),
	is_infected bool NOT NULL,
	is_reinfection bool NOT NULL,
	health_status varchar NOT NULL,
	CONSTRAINT test_resul_pk PRIMARY KEY (id)
);


-- common.test_result foreign keys

ALTER TABLE common.test_result ADD CONSTRAINT test_result_fk FOREIGN KEY (patient_id) REFERENCES common.patient(id);

-- Permissions

ALTER TABLE common.test_result OWNER TO "C19";
GRANT ALL ON TABLE common.test_result TO "C19";