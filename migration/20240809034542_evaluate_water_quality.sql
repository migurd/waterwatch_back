-- +goose Up
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION evaluate_water_quality(
  water_level double precision,
  water_level2 double precision,
  ph_level double precision
) 
RETURNS TABLE (
  status TEXT,
  description TEXT
) 
LANGUAGE plpgsql
AS $$
BEGIN
  RETURN QUERY
  SELECT 
    CASE 
      WHEN ph_level < 6.5 OR ph_level > 8.5 OR water_level < 20 or water_level < 20 THEN 'Malo'
      WHEN water_level >= 20 AND water_level <= 60 THEN 'Bueno'
      WHEN water_level >= 20 AND water_level <= 60 THEN 'Bueno'
      ELSE 'Excelente'
    END AS status,
    CONCAT(
      CASE WHEN ph_level < 6.5 THEN 'El pH es demasiado bajo (ácido). ' ELSE '' END,
      CASE WHEN ph_level > 8.5 THEN 'El pH es demasiado alto (alcalino). ' ELSE '' END,
      CASE WHEN water_level < 20 THEN 'Primer contenedor: El nivel de agua es demasiado bajo. ' ELSE '' END,
      CASE WHEN water_level >= 20 AND water_level <= 60 THEN 'Primer contenedor: El nivel de agua es aceptable. ' ELSE '' END,
      CASE WHEN water_level > 60 THEN 'Primer contenedor: El agua está en buen estado. ' ELSE '' END,
      CASE WHEN water_level2 < 20 THEN 'Segundo contenedor: El nivel de agua es demasiado bajo. ' ELSE '' END,
      CASE WHEN water_level2 >= 20 AND water_level2 <= 60 THEN 'Segundo contenedor: El nivel de agua es aceptable. ' ELSE '' END,
      CASE WHEN water_level2 > 60 THEN 'Segundo contenedor: El agua está en buen estado.' ELSE '' END
    ) AS description;
END;
$$;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP FUNCTION IF EXISTS evaluate_water_quality(double precision, double precision, double precision);
-- +goose StatementEnd
