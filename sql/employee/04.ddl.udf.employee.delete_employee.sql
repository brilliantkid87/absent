CREATE OR REPLACE FUNCTION employee.delete_employee(params jsonb)
 RETURNS boolean
 LANGUAGE plpgsql
AS $function$
DECLARE
    emp_id INT;
BEGIN
    
	emp_id := (params->>'employee_id')::INTEGER;
   	UPDATE employee.employee
    SET deleted_at = CURRENT_TIMESTAMP
   	WHERE employee_id = emp_id;
    
    IF FOUND THEN
        RETURN true;
    ELSE
        RETURN false; 
    END IF;
END;
$function$
;
