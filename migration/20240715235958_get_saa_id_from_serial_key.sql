-- +goose Up
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION get_saa_id_by_serial_key(serial_key_param VARCHAR)
RETURNS BIGINT AS $$
DECLARE
  saa_id BIGINT;
BEGIN
  SELECT s.id INTO saa_id
  FROM public.saa s
  JOIN public.iot_device i ON s.iot_device_id = i.id
  WHERE i.serial_key = serial_key_param;

  RETURN saa_id;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP FUNCTION IF EXISTS get_saa_id_by_serial_key(VARCHAR)
-- +goose StatementEnd
