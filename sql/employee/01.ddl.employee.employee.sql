-- employee.employee definition

-- Drop table

-- DROP TABLE employee.employee;

CREATE TABLE employee.employee (
	employee_id serial4 NOT NULL,
	employee_code varchar(255) NOT NULL,
	employee_name varchar(255) NOT NULL,
	"password" varchar(255) NOT NULL,
	department_id int4 NOT NULL,
	position_id int4 NOT NULL,
	superior int4 NULL,
	created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	created_by varchar(255) NOT NULL,
	updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	updated_by varchar(255) NULL,
	deleted_at timestamp NULL,
	CONSTRAINT employee_employee_code_key UNIQUE (employee_code),
	CONSTRAINT employee_pkey PRIMARY KEY (employee_id)
);


-- employee.employee foreign keys

ALTER TABLE employee.employee ADD CONSTRAINT employee_department_id_fkey FOREIGN KEY (department_id) REFERENCES master_department.master_department(department_id);
ALTER TABLE employee.employee ADD CONSTRAINT employee_position_id_fkey FOREIGN KEY (position_id) REFERENCES master_position.master_position(position_id);