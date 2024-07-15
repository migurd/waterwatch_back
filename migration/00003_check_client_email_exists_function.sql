-- +goose Up
-- +goose StatementBegin
BEGIN;

CREATE OR REPLACE FUNCTION check_client_email_exists(email_to_check VARCHAR)
RETURNS BOOLEAN AS $$
DECLARE
  email_exists BOOLEAN;
BEGIN
  SELECT EXISTS (
    SELECT 1
    FROM public.client_email
    WHERE email = email_to_check
  ) INTO email_exists;
  
  RETURN email_exists;
END;
$$ LANGUAGE plpgsql;

END;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
BEGIN;

DROP FUNCTION IF EXISTS check_client_email_exists(VARCHAR);

END;
-- +goose StatementEnd
