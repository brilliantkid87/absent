CREATE OR REPLACE FUNCTION attendance.delete_attendance(params jsonb)
 RETURNS boolean
 LANGUAGE plpgsql
AS $function$
DECLARE
    att_id INTEGER;
BEGIN
    
    att_id := (params->>'attendance_id')::INTEGER;

    
    UPDATE attendance.attendance
    SET deleted_at = CURRENT_TIMESTAMP
    WHERE attendance_id = att_id;

    
    IF FOUND THEN
        RETURN true;  
    ELSE
        RETURN false; 
    END IF;
END;
$function$
;
