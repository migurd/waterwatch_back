-- +goose Up
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION get_account_details(client_id BIGINT)
RETURNS TABLE(
  username VARCHAR,
  email VARCHAR,
  phone_number VARCHAR
) AS $$
BEGIN
  RETURN QUERY
  SELECT
    a.username,
    ce.email,
    CONCAT('+', cp.country_code, ' ', cp.phone_number)::VARCHAR AS phone_number
  FROM
    public.account a
  JOIN public.client_email ce ON a.client_id = ce.client_id
  JOIN public.client_phone_number cp ON a.client_id = cp.client_id
  WHERE
    a.client_id = get_account_details.client_id;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP FUNCTION IF EXISTS get_account_details(BIGINT);
-- +goose StatementEnd
