-- +goose Up
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION evaluate_water_quality(
  water_level double precision,
  ph_level double precision,
  is_contaminated boolean
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
      WHEN is_contaminated OR ph_level < 6.5 OR ph_level > 8.5 OR water_level < 20 THEN 'MALO'
      WHEN water_level >= 20 AND water_level <= 60 THEN 'MEDIANO'
      ELSE 'BUENO'
    END AS status,
    CONCAT(
      CASE WHEN is_contaminated THEN 'El agua está contaminada. ' ELSE '' END,
      CASE WHEN ph_level < 6.5 THEN 'El pH es demasiado bajo (ácido). ' ELSE '' END,
      CASE WHEN ph_level > 8.5 THEN 'El pH es demasiado alto (alcalino). ' ELSE '' END,
      CASE WHEN water_level < 20 THEN 'El nivel de agua es demasiado bajo. ' ELSE '' END,
      CASE WHEN water_level >= 20 AND water_level <= 60 THEN 'El nivel de agua es aceptable. ' ELSE '' END,
      CASE WHEN water_level > 60 THEN 'El agua está en buen estado.' ELSE '' END
    ) AS description;
END;
$$;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP FUNCTION IF EXISTS evaluate_water_quality(double precision, double precision, boolean);
-- +goose StatementEnd
