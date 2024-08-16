-- +goose Up
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION get_employee_details(emp_id bigint)
RETURNS TABLE(full_name text, phone_number text, email text) AS $$
BEGIN
  RETURN QUERY (
    SELECT 
      e.first_name || ' ' || e.last_name AS full_name, 
      '+' || ep.country_code || ' ' || ep.phone_number AS phone_number,
      ee.email::text  -- Cast email to text
    FROM public.employee e
    LEFT JOIN public.employee_phone_number ep ON e.id = ep.employee_id
    LEFT JOIN public.employee_email ee ON e.id = ee.employee_id
    WHERE e.id = emp_id
  );
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd

-- +goose StatementBegin
CREATE OR REPLACE FUNCTION get_address_state_city(address_id bigint)
RETURNS text AS $$
BEGIN
  RETURN (
    SELECT state || ', ' || city
    FROM public.client_address
    WHERE id = address_id
  );
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP FUNCTION IF EXISTS get_employee_details(bigint);
-- +goose StatementEnd

-- +goose StatementBegin
DROP FUNCTION IF EXISTS get_address_state_city(bigint);
-- +goose StatementEnd