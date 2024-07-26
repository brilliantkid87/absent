CREATE OR REPLACE FUNCTION master_location.delete_master_location(params jsonb)
 RETURNS boolean
 LANGUAGE plpgsql
AS $function$
DECLARE
    loc_id INTEGER;
BEGIN
    
    loc_id := (params->>'location_id')::INTEGER;

    
    UPDATE master_location.master_location
    SET deleted_at = CURRENT_TIMESTAMP
    WHERE location_id = loc_id;

    
    IF FOUND THEN
        RETURN true;  
    ELSE
        RETURN false; 
    END IF;
END;
$function$
;
