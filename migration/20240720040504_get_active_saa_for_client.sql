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
  water_description TEXT,
  current_saa_capacity double precision,
  max_saa_capacity INT,
  saa_height INT,
  current_saa_capacity2 double precision,
  max_saa_capacity2 INT,
  saa_height2 INT,
  days_since_last_maintenance integer
  -- last_maintenance_date DATE
) AS $$
BEGIN
  RETURN QUERY
  SELECT
    saa.id,
    iot.serial_key,
    get_full_address(a.address_id) AS full_address,
    sd.name as saa_name,
    sd.description as saa_description,
    COALESCE(wq.status, 'No estado') AS water_status,
    COALESCE(wq.description, 'No hay descripción disponible') AS water_description,
    COALESCE((SELECT sr.water_level FROM saa_record sr WHERE sr.saa_id = saa.id ORDER BY sr.date DESC LIMIT 1), 0) * st.capacity / 100 AS current_saa_capacity,
    st.capacity AS saa_capacity,
    st.height AS saa_height,
    COALESCE((SELECT sr.water_level2 FROM saa_record sr WHERE sr.saa_id = saa.id ORDER BY sr.date DESC LIMIT 1), 0) * st2.capacity / 100 AS current_saa_capacity2,
    st2.capacity AS saa_capacity2,
    st2.height AS saa_height2,
    COALESCE(
      (SELECT CURRENT_DATE - MAX(ap.done_date)::DATE
       FROM appointment ap 
       WHERE ap.client_id = get_all_active_saa_for_client.client_id 
         AND ap.appointment_type_id = 2),
      -1
    ) AS days_since_last_maintenance
    -- (SELECT MAX(ap.done_date)::DATE
    --    FROM appointment ap 
    --    WHERE ap.client_id = get_all_active_saa_for_client.client_id 
    --      AND ap.appointment_type_id = 2) AS last_maintenance_date
  FROM
    public.saa saa
  JOIN public.iot_device iot ON saa.iot_device_id = iot.id
  JOIN public.appointment a ON saa.appointment_id = a.id
  JOIN public.saa_description sd ON sd.saa_id = saa.id
  JOIN public.saa_type st ON st.id = saa.saa_type_id
  JOIN public.saa_type st2 ON st2.id = saa.saa_type_id2
  LEFT JOIN LATERAL (
    SELECT status, description
    FROM evaluate_water_quality(
      COALESCE((SELECT sr.water_level FROM saa_record sr WHERE sr.saa_id = saa.id ORDER BY sr.date DESC LIMIT 1), 0),
      COALESCE((SELECT sr.water_level2 FROM saa_record sr WHERE sr.saa_id = saa.id ORDER BY sr.date DESC LIMIT 1), 0),
      COALESCE((SELECT sr.ph_level FROM saa_record sr WHERE sr.saa_id = saa.id ORDER BY sr.date DESC LIMIT 1), 7.0)
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
