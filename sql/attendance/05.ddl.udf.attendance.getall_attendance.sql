CREATE OR REPLACE FUNCTION attendance.getall_attendance()
 RETURNS jsonb
 LANGUAGE plpgsql
AS $function$
DECLARE
	all_deleted BOOLEAN;
BEGIN
    SELECT NOT EXISTS (
        SELECT 1
        FROM attendance.attendance
        WHERE deleted_at IS NULL
    ) INTO all_deleted;
   	
   	IF all_deleted THEN
        RETURN '[]'::jsonb;
    END IF;
   
	RETURN (
        SELECT jsonb_agg(row_to_json(a))
        FROM attendance.attendance a
        where a.deleted_at IS NULL
    );
END;
$function$
;
