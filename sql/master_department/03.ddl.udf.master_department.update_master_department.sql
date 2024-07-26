CREATE OR REPLACE FUNCTION master_department.update_master_department(params jsonb)
 RETURNS boolean
 LANGUAGE plpgsql
AS $function$
DECLARE
    updated_count INTEGER;
BEGIN

    UPDATE master_department.master_department
    SET
        department_name = COALESCE(params->>'department_name', department_name),
        updated_at = CURRENT_TIMESTAMP,
        updated_by = COALESCE(params->>'updated_by', updated_by)
    WHERE department_id = (params->>'department_id')::INTEGER;

    GET DIAGNOSTICS updated_count = ROW_COUNT;


    IF updated_count = 0 THEN
        RETURN false;
    END IF;

    RETURN true;
END;
$function$
;
