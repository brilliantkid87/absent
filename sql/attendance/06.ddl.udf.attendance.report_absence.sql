CREATE OR REPLACE FUNCTION attendance.report_absence(params jsonb)
 RETURNS jsonb
 LANGUAGE plpgsql
AS $function$
BEGIN
    RETURN (
        SELECT jsonb_agg(row_to_json(r))
        FROM (
            SELECT
                a.absent_in,
                a.absent_out,
                e.employee_code,
                e.employee_name,
                d.department_name,
                p.position_name,
                l.location_name 
            FROM attendance.attendance a
            JOIN employee.employee e ON a.employee_id = e.employee_id
            JOIN master_department.master_department d ON e.department_id = d.department_id
            JOIN master_position.master_position p ON e.position_id = p.position_id
            JOIN master_location.master_location l ON a.location_id = l.location_id
            WHERE a.absent_in BETWEEN (params ->> 'start_time')::TIMESTAMP AND (params ->> 'end_time')::TIMESTAMP
            AND a.deleted_at IS NULL
        ) r
    );
END;
$function$
;
