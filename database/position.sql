CREATE TABLE common.country (
	id int NOT NULL,
	"name" varchar NOT NULL,
	CONSTRAINT country_pk PRIMARY KEY (id)
);

CREATE TABLE common.town (
	id uuid NOT NULL DEFAULT common.uuid_generate_v4(),
	"name" varchar NOT NULL,
	country_id int NULL,
	CONSTRAINT town_pk PRIMARY KEY (id),
	CONSTRAINT town_fk FOREIGN KEY (country_id) REFERENCES common.country(id)
);


CREATE TABLE common.district (
	id uuid NOT NULL DEFAULT common.uuid_generate_v4(),
	"name" varchar NOT NULL,
	town_id uuid NOT NULL,
	CONSTRAINT district_pk PRIMARY KEY (id),
	CONSTRAINT district_fk FOREIGN KEY (town_id) REFERENCES common.town(id)
);

ALTER TABLE common.patient ADD district_id uuid NOT NULL;
ALTER TABLE common.patient ADD CONSTRAINT patient_fk FOREIGN KEY (district_id) REFERENCES common.district(id);
