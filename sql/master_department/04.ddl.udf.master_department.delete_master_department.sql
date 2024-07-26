CREATE OR REPLACE FUNCTION master_department.delete_master_department(params jsonb)
 RETURNS boolean
 LANGUAGE plpgsql
AS $function$
DECLARE
    dept_id INTEGER;
BEGIN
    
    dept_id := (params->>'department_id')::INTEGER;

    
    UPDATE master_department.master_department
    SET deleted_at = CURRENT_TIMESTAMP
    WHERE department_id = dept_id;

    
    IF FOUND THEN
        RETURN true;  
    ELSE
        RETURN false; 
    END IF;
END;
$function$
;
