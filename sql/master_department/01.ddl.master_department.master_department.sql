-- master_department.master_department definition

-- Drop table

-- DROP TABLE master_department.master_department;

CREATE TABLE master_department.master_department (
	department_id int4 NOT NULL DEFAULT nextval('master_department.master_depatment_department_id_seq'::regclass),
	department_name varchar(255) NOT NULL,
	created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	created_by varchar(255) NOT NULL,
	updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	updated_by varchar(255) NULL,
	deleted_at timestamp NULL,
	CONSTRAINT master_depatment_pkey PRIMARY KEY (department_id)
);