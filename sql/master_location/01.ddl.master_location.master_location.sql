-- master_location.master_location definition

-- Drop table

-- DROP TABLE master_location.master_location;

CREATE TABLE master_location.master_location (
	location_id serial4 NOT NULL,
	location_name varchar(255) NOT NULL,
	created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	created_by varchar(255) NOT NULL,
	updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	updated_by varchar(255) NULL,
	deleted_at timestamp NULL,
	CONSTRAINT master_location_pkey PRIMARY KEY (location_id)
);