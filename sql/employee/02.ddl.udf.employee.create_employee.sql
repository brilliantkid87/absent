CREATE OR REPLACE FUNCTION employee.create_employee(params jsonb)
 RETURNS integer
 LANGUAGE plpgsql
AS $function$
DECLARE
    new_employee_id INTEGER;
BEGIN
    INSERT INTO employee.employee (
        employee_name,
        password,
        employee_code,
        department_id,
        position_id,
        superior,
        created_at,
        created_by,
        updated_at,
        updated_by,
        deleted_at
    )
    VALUES (
        params ->> 'employee_name',
        params ->> 'password',
        params ->> 'employee_code',
        (params ->> 'department_id')::INTEGER,
        (params ->> 'position_id')::INTEGER,
        (params ->> 'superior')::INTEGER,
        CURRENT_TIMESTAMP,
        params ->> 'created_by',
        CURRENT_TIMESTAMP,
        params ->> 'created_by',
        NULL
    )
    RETURNING employee_id INTO new_employee_id;

    RETURN new_employee_id;
END;
$function$
;
