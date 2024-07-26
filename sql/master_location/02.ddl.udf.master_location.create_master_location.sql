CREATE OR REPLACE FUNCTION master_location.create_master_location(params jsonb)
 RETURNS integer
 LANGUAGE plpgsql
AS $function$
DECLARE
    ret_id INTEGER;
BEGIN
   
    INSERT INTO master_location.master_location (
        location_name,
        created_at,
        created_by,
        updated_at,
        updated_by,
        deleted_at
    )
    VALUES (
        params ->> 'location_name',
        CURRENT_TIMESTAMP,
        params ->> 'created_by',
        CURRENT_TIMESTAMP,
        params ->> 'created_by',
        NULL
    )
    RETURNING location_id INTO ret_id;

    RETURN ret_id;
END;
$function$
;

