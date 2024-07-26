CREATE OR REPLACE FUNCTION master_location.update_master_location(params jsonb)
 RETURNS boolean
 LANGUAGE plpgsql
AS $function$
DECLARE
    updated_location RECORD;
BEGIN
    -- Perform the update
    UPDATE master_location.master_location
    SET
        location_name = COALESCE(params->>'location_name', location_name),
        updated_at = CURRENT_TIMESTAMP,
        updated_by = COALESCE(params->>'updated_by', updated_by)
    WHERE location_id = (params->>'location_id')::INTEGER
    RETURNING
        location_id,
        location_name,
        created_at,
        created_by,
        updated_at,
        updated_by,
        deleted_at
    INTO updated_location;

    -- Check if any rows were updated
    IF NOT FOUND THEN
        RETURN false; -- Update failed or location not found
    END IF;

    RETURN true;  -- Successfully updated
END;
$function$
;
