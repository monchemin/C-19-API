-- DROP TABLE common.constant_indexation;

CREATE TABLE common.indexed_constants (
	date_time timestamptz NOT NULL DEFAULT now(),
	completed bool NOT NULL DEFAULT false,
	error varchar NULL,
	CONSTRAINT constant_indexation_pk PRIMARY KEY (date_time)
);

-- Permissions

ALTER TABLE common.constant_indexation OWNER TO "C19";
GRANT ALL ON TABLE common.constant_indexation TO "C19";
