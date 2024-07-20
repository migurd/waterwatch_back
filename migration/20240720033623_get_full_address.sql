-- +goose Up
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION get_full_address(address_id BIGINT)
RETURNS VARCHAR AS $$
DECLARE
  full_address VARCHAR;
BEGIN
  SELECT
    CONCAT(
      'Calle ', ca.street, ', ',
      'Número ', ca.house_number, ', ',
      'Colonia ', ca.neighborhood, ', ',
      ca.city, ', ',
      ca.state, ', ',
      'Código Postal ', ca.postal_code
    )::VARCHAR INTO full_address
  FROM
    public.client_address ca
  WHERE
    ca.id = get_full_address.address_id;

  RETURN full_address;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP FUNCTION IF EXISTS get_full_address(BIGINT);
-- +goose StatementEnd
