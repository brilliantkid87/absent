CREATE OR REPLACE FUNCTION attendance.create_attendance(params jsonb)
 RETURNS integer
 LANGUAGE plpgsql
AS $function$
DECLARE
    ret_id INTEGER;
BEGIN
    
    IF NOT EXISTS (SELECT 1 FROM employee.employee WHERE employee_id = (params ->> 'employee_id')::INTEGER) THEN
        RAISE EXCEPTION 'Employee ID % does not exist', (params ->> 'employee_id')::INTEGER;
    END IF;

    
    IF NOT EXISTS (SELECT 1 FROM master_location.master_location WHERE location_id = (params ->> 'location_id')::INTEGER) THEN
        RAISE EXCEPTION 'Location ID % does not exist', (params ->> 'location_id')::INTEGER;
    END IF;

    
    INSERT INTO attendance.attendance (
        employee_id,
        location_id,
        absent_in,
        absent_out,
        created_at,
        created_by,
        updated_at,
        deleted_at
    )
    VALUES (
        (params ->> 'employee_id')::INTEGER,
        (params ->> 'location_id')::INTEGER,
        (params ->> 'absent_in')::TIMESTAMP,
        (params ->> 'absent_out')::TIMESTAMP,
        CURRENT_TIMESTAMP,
        params ->> 'created_by',
        CURRENT_TIMESTAMP,
        NULL
    )
    RETURNING attendance_id INTO ret_id;

    RETURN ret_id;
END;
$function$
;
