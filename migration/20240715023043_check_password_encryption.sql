-- +goose Up
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION check_password_encryption(username_input VARCHAR)
RETURNS boolean AS $$
DECLARE
  is_encrypted boolean;
BEGIN
  SELECT "as".is_password_encrypted
  INTO is_encrypted
  FROM public.account a
  JOIN public.account_security "as"
  ON a.client_id = "as".account_client_id
  WHERE a.username = username_input;

  RETURN is_encrypted;
END;
$$ LANGUAGE plpgsql;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP FUNCTION IF EXISTS check_password_encryption(VARCHAR)
-- +goose StatementEnd
