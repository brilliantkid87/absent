CREATE OR REPLACE FUNCTION master_position.update_master_position(params jsonb)
 RETURNS boolean
 LANGUAGE plpgsql
AS $function$
DECLARE
    updated_position RECORD;
BEGIN
    
    UPDATE master_position.master_position
    SET
        department_id = COALESCE((params->>'department_id')::INTEGER, department_id),
        position_name = COALESCE(params->>'position_name', position_name),
        updated_at = CURRENT_TIMESTAMP,
        updated_by = COALESCE(params->>'updated_by', updated_by)
    WHERE position_id = (params->>'position_id')::INTEGER
    RETURNING
        position_id,
        department_id,
        position_name,
        created_at,
        created_by,
        updated_at,
        updated_by,
        deleted_at
    INTO updated_position;

    
    IF NOT FOUND THEN
        RETURN false; 
    END IF;

    RETURN true;  
END;
$function$
;
