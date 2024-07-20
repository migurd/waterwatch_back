-- +goose Up
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION get_all_contacts()
RETURNS TABLE(
  id BIGINT,
  name VARCHAR,
  photo_url VARCHAR,
  emails VARCHAR[],
  phone_numbers VARCHAR[]
) AS $$
BEGIN
  RETURN QUERY
  SELECT
    c.id,
    c.name,
    c.photo_url,
    COALESCE(ARRAY_AGG(DISTINCT ce.email), '{}') AS emails,
    COALESCE(ARRAY_AGG(DISTINCT CONCAT('+', cp.country_code, ' ', cp.phone_number)::VARCHAR), '{}') AS phone_numbers
  FROM
    public.contact c
  LEFT JOIN public.contact_email ce ON c.id = ce.contact_id
  LEFT JOIN public.contact_phone_number cp ON c.id = cp.contact_id
  GROUP BY
    c.id, c.name, c.photo_url;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP FUNCTION IF EXISTS get_all_contacts()
-- +goose StatementEnd
