-- +goose Up
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION get_all_active_saa_for_client(client_id BIGINT)
RETURNS TABLE(
  saa_id BIGINT,
  serial_key VARCHAR,
  full_address VARCHAR,
  is_good VARCHAR
) AS $$
BEGIN
  RETURN QUERY
  SELECT
    saa.id,
    iot.serial_key,
    get_full_address(a.address_id) AS full_address,
    'BUENO'::VARCHAR AS is_good
  FROM
    public.saa saa
  JOIN public.iot_device iot ON saa.iot_device_id = iot.id
  JOIN public.appointment a ON saa.appointment_id = a.id
  WHERE
    a.client_id = get_all_active_saa_for_client.client_id
    AND iot.status = TRUE;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP FUNCTION IF EXISTS get_all_active_saa_for_client(BIGINT)
-- +goose StatementEnd
