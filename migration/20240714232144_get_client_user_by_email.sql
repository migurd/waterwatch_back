-- +goose Up
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION get_client_user_by_email(email_param VARCHAR)
RETURNS TABLE(username VARCHAR) AS $$
BEGIN
  RETURN QUERY
  SELECT a.username
  FROM account a
  LEFT JOIN client_email e ON a.client_id = e.client_id
  WHERE e.email = email_param;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP FUNCTION IF EXISTS get_client_user_by_email(VARCHAR);
-- +goose StatementEnd
