CREATE OR REPLACE FUNCTION master_department.create_master_department(params jsonb)
 RETURNS integer
 LANGUAGE plpgsql
AS $function$
DECLARE
    ret_id INTEGER;
BEGIN
    INSERT INTO master_department.master_department (
        department_name,
        created_by
    )
    VALUES (
        params->>'department_name',
        params->>'created_by'
    )
    RETURNING department_id INTO ret_id;
    
    RETURN ret_id;
END;
$function$
;
