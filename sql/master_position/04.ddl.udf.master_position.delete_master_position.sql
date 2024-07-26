CREATE OR REPLACE FUNCTION master_position.delete_master_position(params jsonb)
 RETURNS boolean
 LANGUAGE plpgsql
AS $function$
DECLARE
    pst_id INTEGER;
BEGIN
    
    pst_id := (params->>'position_id')::INTEGER;

    
    UPDATE master_position.master_position
    SET deleted_at = CURRENT_TIMESTAMP
    WHERE position_id = pst_id;

    
    IF FOUND THEN
        RETURN true;  
    ELSE
        RETURN false; 
    END IF;
END;
$function$
;
