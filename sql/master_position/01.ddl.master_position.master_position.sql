-- master_position.master_position definition

-- Drop table

-- DROP TABLE master_position.master_position;

CREATE TABLE master_position.master_position (
	position_id serial4 NOT NULL,
	department_id int4 NOT NULL,
	position_name varchar(255) NOT NULL,
	created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	created_by varchar(255) NOT NULL,
	updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	updated_by varchar(255) NULL,
	deleted_at timestamp NULL,
	CONSTRAINT master_position_pkey PRIMARY KEY (position_id)
);


-- master_position.master_position foreign keys

ALTER TABLE master_position.master_position ADD CONSTRAINT master_position_department_id_fkey FOREIGN KEY (department_id) REFERENCES master_department.master_department(department_id);