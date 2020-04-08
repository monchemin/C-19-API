-- DROP TABLE common."user";

CREATE TABLE common."user" (
	id uuid NOT NULL DEFAULT uuid_generate_v4(),
	email varchar NOT NULL,
	phone_number varchar NOT NULL,
	"password" varchar NOT NULL,
	first_name varchar NOT NULL,
	last_name varchar NOT NULL,
	active bool NULL DEFAULT true,
	CONSTRAINT user_pk PRIMARY KEY (id),
	CONSTRAINT user_un UNIQUE (email)
);

-- Permissions

ALTER TABLE common."user" OWNER TO "C19";
GRANT ALL ON TABLE common."user" TO "C19";


-- DROP TABLE common."role";

CREATE TABLE common."role" (
	id int4 NOT NULL,
	code bpchar(3) NULL,
	"name" varchar NOT NULL,
	CONSTRAINT role_pk PRIMARY KEY (id)
);

-- Permissions

ALTER TABLE common."role" OWNER TO "C19";
GRANT ALL ON TABLE common."role" TO "C19";

-- DROP TABLE common.resource_type;

CREATE TABLE common.resource_type (
	id int4 NOT NULL,
	code bpchar(3) NULL,
	"name" varchar NOT NULL,
	"options" varchar NULL,
	CONSTRAINT resource_type_pk PRIMARY KEY (id),
	CONSTRAINT resource_type_un UNIQUE (code)
);

-- Permissions

ALTER TABLE common.resource_type OWNER TO "C19";
GRANT ALL ON TABLE common.resource_type TO "C19";


-- DROP TABLE common.privilege;

CREATE TABLE common.privilege (
	id uuid NOT NULL DEFAULT uuid_generate_v4(),
	user_id uuid NULL,
	resource_id uuid NOT NULL,
	resource_type_id int4 NOT NULL,
	role_id int4 NOT NULL,
	start_date date NOT NULL DEFAULT now(),
	end_date date NULL,
	CONSTRAINT privilege_pk PRIMARY KEY (id),
	CONSTRAINT privilege_un UNIQUE (user_id, resource_id, role_id),
	CONSTRAINT privilege_fk FOREIGN KEY (user_id) REFERENCES common."user"(id)
);

-- Permissions

ALTER TABLE common.privilege OWNER TO "C19";
GRANT ALL ON TABLE common.privilege TO "C19";
