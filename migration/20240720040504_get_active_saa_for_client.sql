-- +goose Up
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION get_all_active_saa_for_client(client_id BIGINT)
RETURNS TABLE(
  saa_id BIGINT,
  serial_key VARCHAR,
  full_address VARCHAR,
  saa_name VARCHAR,
  saa_description VARCHAR,
  water_status TEXT,
  water_description TEXT
) AS $$
BEGIN
  RETURN QUERY
  SELECT
    saa.id,
    iot.serial_key,
    get_full_address(a.address_id) AS full_address,
    sd.name as saa_name,
    sd.description as saa_description,
    COALESCE(wq.status, 'No Status') AS water_status,
    COALESCE(wq.description, 'No Description Available') AS water_description
  FROM
    public.saa saa
  JOIN public.iot_device iot ON saa.iot_device_id = iot.id
  JOIN public.appointment a ON saa.appointment_id = a.id
  JOIN public.saa_description sd ON sd.saa_id = saa.id
  LEFT JOIN LATERAL (
    SELECT status, description
    FROM evaluate_water_quality(
      COALESCE((SELECT sr.water_level FROM saa_record sr WHERE sr.saa_id = saa.id ORDER BY sr.date DESC LIMIT 1), 0),
      COALESCE((SELECT sr.ph_level FROM saa_record sr WHERE sr.saa_id = saa.id ORDER BY sr.date DESC LIMIT 1), 7.0),
      COALESCE((SELECT sr.is_contaminated FROM saa_record sr WHERE sr.saa_id = saa.id ORDER BY sr.date DESC LIMIT 1), false)
    )
  ) AS wq ON TRUE
  WHERE
    a.client_id = get_all_active_saa_for_client.client_id
    AND iot.status = TRUE;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP FUNCTION IF EXISTS get_all_active_saa_for_client(BIGINT);
-- +goose StatementEnd
