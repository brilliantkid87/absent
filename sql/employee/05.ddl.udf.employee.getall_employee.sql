CREATE OR REPLACE FUNCTION employee.getall_employee()
 RETURNS jsonb
 LANGUAGE plpgsql
AS $function$
DECLARE
	all_deleted BOOLEAN;
BEGIN
    SELECT NOT EXISTS (
        SELECT 1
        FROM employee.employee
        WHERE deleted_at IS NULL
    ) INTO all_deleted;
   	
   	IF all_deleted THEN
        RETURN '[]'::jsonb;
    END IF;
   
	RETURN (
        SELECT jsonb_agg(row_to_json(e))
        FROM employee.employee e
        WHERE e.deleted_at IS NULL
    );
END;
$function$
;
