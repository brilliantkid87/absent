CREATE OR REPLACE FUNCTION master_position.getall_master_position()
 RETURNS jsonb
 LANGUAGE plpgsql
AS $function$
DECLARE
    all_deleted BOOLEAN;
BEGIN
	
	SELECT NOT EXISTS (
        SELECT 1
        FROM master_position.master_position
        WHERE deleted_at IS NULL
    ) INTO all_deleted;
   
   IF all_deleted THEN
        RETURN '[]'::jsonb;
    END IF;
   
	RETURN (
		SELECT jsonb_agg(row_to_json(mp))
		FROM master_position.master_position mp
       	WHERE mp.deleted_at IS NULL
    );
END;
$function$
;
