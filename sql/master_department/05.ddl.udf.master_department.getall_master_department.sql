CREATE OR REPLACE FUNCTION master_department.getall_master_department()
 RETURNS jsonb
 LANGUAGE plpgsql
AS $function$
DECLARE
    all_deleted BOOLEAN;
BEGIN
    
    SELECT NOT EXISTS (
        SELECT 1
        FROM master_department.master_department
        WHERE deleted_at IS NULL
    ) INTO all_deleted;

    
    IF all_deleted THEN
        RETURN '[]'::jsonb;
    END IF;

    
    RETURN (
        SELECT jsonb_agg(row_to_json(md))
        FROM master_department.master_department md
        WHERE md.deleted_at IS NULL
    );
END;
$function$
;
