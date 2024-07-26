CREATE OR REPLACE FUNCTION employee.update_employee(params jsonb)
 RETURNS boolean
 LANGUAGE plpgsql
AS $function$
DECLARE
    updated_employee RECORD;
BEGIN

    UPDATE employee.employee
    SET
        employee_code = COALESCE(NULLIF(params->>'employee_code', ''), employee_code),
        employee_name = COALESCE(NULLIF(params->>'employee_name', ''), employee_name),
        "password" = COALESCE(NULLIF(params->>'password', ''), "password"),
        department_id = COALESCE(NULLIF((params->>'department_id')::INT, department_id), department_id),
        position_id = COALESCE(NULLIF((params->>'position_id')::INT, position_id), position_id),
        superior = COALESCE(NULLIF((params->>'superior')::INT, superior), superior),
        created_by = COALESCE(NULLIF(params->>'created_by', ''), created_by),
        updated_by = COALESCE(NULLIF(params->>'updated_by', ''), updated_by),
        deleted_at = COALESCE(NULLIF(params->>'deleted_at', '')::TIMESTAMP, deleted_at),
        updated_at = CURRENT_TIMESTAMP
    WHERE employee_id = (params->>'employee_id')::INTEGER
    RETURNING
        employee_id,
        employee_code,
        employee_name,
        "password",
        department_id,
        position_id,
        superior,
        created_at,
        created_by,
        updated_at,
        updated_by,
        deleted_at
    INTO updated_employee;


    IF NOT FOUND THEN
        RETURN false; 
    END IF;

    RETURN true;  
END;
$function$
;
