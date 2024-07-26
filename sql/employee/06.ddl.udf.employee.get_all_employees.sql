CREATE OR REPLACE FUNCTION employee.get_all_employees(params jsonb)
 RETURNS jsonb
 LANGUAGE plpgsql
AS $function$
DECLARE
    result jsonb;
BEGIN
    SELECT jsonb_agg(jsonb_build_object(
        'employee_id', e.employee_id,
        'employee_code', e.employee_code,
        'employee_name', e.employee_name,
        'password', e.password,
        'department_id', e.department_id,
        'position_id', e.position_id,
        'superior', e.superior,
        'created_at', e.created_at,
        'created_by', e.created_by,
        'updated_at', e.updated_at,
        'updated_by', e.updated_by
    ))
    INTO result
    FROM employee.employee e
    WHERE (params ->> 'employee_name' IS NULL OR e.employee_name = params ->> 'employee_name');

    RETURN COALESCE(result, '[]'::jsonb);
END;
$function$
;
