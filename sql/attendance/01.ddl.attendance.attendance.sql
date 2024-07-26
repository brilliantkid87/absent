-- attendance.attendance definition

-- Drop table

-- DROP TABLE attendance.attendance;

CREATE TABLE attendance.attendance (
	attendance_id serial4 NOT NULL,
	employee_id int4 NOT NULL,
	location_id int4 NOT NULL,
	absent_in timestamp NULL,
	absent_out timestamp NULL,
	created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	created_by varchar(255) NOT NULL,
	updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	updated_by varchar(255) NULL,
	deleted_at timestamp NULL,
	CONSTRAINT attendance_pkey PRIMARY KEY (attendance_id)
);


-- attendance.attendance foreign keys

ALTER TABLE attendance.attendance ADD CONSTRAINT attendance_employee_id_fkey FOREIGN KEY (employee_id) REFERENCES employee.employee(employee_id);
ALTER TABLE attendance.attendance ADD CONSTRAINT attendance_location_id_fkey FOREIGN KEY (location_id) REFERENCES master_location.master_location(location_id);