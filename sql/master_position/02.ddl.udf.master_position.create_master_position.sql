CREATE OR REPLACE FUNCTION master_position.create_master_position(params jsonb)
 RETURNS integer
 LANGUAGE plpgsql
AS $function$
DECLARE
    ret_id INTEGER;
BEGIN
    
    INSERT INTO master_position.master_position (
        department_id,
        position_name,
        created_at,
        created_by,
        deleted_at
    )
    VALUES (
        (params->>'department_id')::INTEGER,
        params->>'position_name',
        CURRENT_TIMESTAMP,
        params->>'created_by',
        NULL
    )
    RETURNING position_id INTO ret_id;

    RETURN ret_id;
END;
$function$
;
