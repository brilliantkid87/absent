CREATE OR REPLACE FUNCTION attendance.update_attendance(params jsonb)
 RETURNS boolean
 LANGUAGE plpgsql
AS $function$
DECLARE
    updated_count INTEGER;
BEGIN
    
    UPDATE attendance.attendance
    SET
        employee_id = COALESCE(NULLIF((params ->> 'employee_id')::INTEGER, NULL), employee_id),
        location_id = COALESCE(NULLIF((params ->> 'location_id')::INTEGER, NULL), location_id),
        absent_in = COALESCE(NULLIF((params ->> 'absent_in')::TIMESTAMP, NULL), absent_in),
        absent_out = COALESCE(NULLIF((params ->> 'absent_out')::TIMESTAMP, NULL), absent_out),
        updated_at = CURRENT_TIMESTAMP,
        updated_by = COALESCE(NULLIF(params ->> 'updated_by', ''), updated_by)
    WHERE attendance_id = (params ->> 'attendance_id')::INTEGER;

    
    GET DIAGNOSTICS updated_count = ROW_COUNT;

    
    IF updated_count = 0 THEN
        RETURN FALSE; 
    END IF;

    RETURN TRUE;  
END;
$function$
;
