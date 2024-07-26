CREATE OR REPLACE FUNCTION master_location.getall_master_location()
 RETURNS jsonb
 LANGUAGE plpgsql
AS $function$
DECLARE
    all_deleted BOOLEAN;
BEGIN
    
	SELECT NOT EXISTS (
        SELECT 1
        FROM master_location.master_location
        WHERE deleted_at IS NULL
    ) INTO all_deleted;
   
   	IF all_deleted THEN
        RETURN '[]'::jsonb;
    END IF;
   
	RETURN (
        SELECT jsonb_agg(row_to_json(ml))
        FROM master_location.master_location ml
        WHERE ml.deleted_at IS NULL
    );
END;
$function$
;
