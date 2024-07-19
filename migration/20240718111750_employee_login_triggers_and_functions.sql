-- +goose Up
-- +goose StatementBegin
BEGIN;

CREATE OR REPLACE FUNCTION get_employee_user_by_email(email_param VARCHAR)
RETURNS TABLE(username VARCHAR) AS $$
BEGIN
  RETURN QUERY
  SELECT a.username
  FROM employee_account a
  LEFT JOIN employee_email e ON a.employee_id = e.employee_id
  WHERE e.email = email_param;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION check_password_encryption_employee(username_input VARCHAR)
RETURNS boolean AS $$
DECLARE
  is_encrypted boolean;
BEGIN
  SELECT "eas".is_password_encrypted
  INTO is_encrypted
  FROM public.employee_account ea
  JOIN public.employee_account_security "eas"
  ON ea.employee_id = "eas".employee_account_employee_id
  WHERE ea.username = username_input;

  RETURN is_encrypted;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE PROCEDURE increase_attempts_employee(username_input VARCHAR)
LANGUAGE plpgsql
AS $$
DECLARE
  current_attempts INT;
  last_attempt_time TIMESTAMP;
BEGIN
  SELECT employee_account_security.attempts, employee_account_security.last_attempt
  INTO current_attempts, last_attempt_time
  FROM public.employee_account_security
  JOIN public.employee_account
    ON employee_account_security.employee_account_employee_id = employee_account.employee_id
  WHERE employee_account.username = username_input;

  IF NOT FOUND THEN
    RAISE EXCEPTION 'Username not found';
  END IF;

  IF current_attempts >= 5 AND NOW() - last_attempt_time <= INTERVAL '1 minute' THEN
    RAISE EXCEPTION 'Please wait for a minute before trying again';
  END IF;

  IF current_attempts >= 5 THEN
    UPDATE public.employee_account_security
    SET 
      attempts = 1,
      last_attempt = NOW()
    FROM public.employee_account
    WHERE employee_account.username = username_input
      AND employee_account_security.employee_account_employee_id = employee_account.employee_id;
  ELSE
    UPDATE public.employee_account_security
    SET 
      attempts = attempts + 1,
      last_attempt = NOW()
    FROM public.employee_account
    WHERE employee_account.username = username_input
      AND employee_account_security.employee_account_employee_id = employee_account.employee_id;
  END IF;
END;
$$;

CREATE OR REPLACE PROCEDURE reset_attempts_employee(username_input VARCHAR)
LANGUAGE plpgsql
AS $$
BEGIN
  UPDATE public.employee_account_security
  SET 
    attempts = 0,
    last_attempt = NOW()
  FROM public.employee_account
  WHERE employee_account.username = username_input
    AND employee_account_security.employee_account_employee_id = employee_account.employee_id;

  IF NOT FOUND THEN
    RAISE EXCEPTION 'Username not found';
  END IF;
END;
$$;

CREATE OR REPLACE PROCEDURE get_attempts_employee(username_input VARCHAR, OUT attempts_count INT)
LANGUAGE plpgsql
AS $$
BEGIN
  SELECT employee_account_security.attempts
  INTO attempts_count
  FROM public.employee_account_security
  JOIN public.employee_account
    ON employee_account_security.employee_account_employee_id = employee_account.employee_id
  WHERE employee_account.username = username_input;

  IF NOT FOUND THEN
    RAISE EXCEPTION 'Username not found';
  END IF;
END;
$$;

COMMIT;
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
BEGIN;

DROP FUNCTION IF EXISTS get_employee_user_by_email(VARCHAR);
DROP FUNCTION IF EXISTS check_password_encryption_employee(VARCHAR);
DROP PROCEDURE IF EXISTS increase_attempts_employee(VARCHAR);
DROP PROCEDURE IF EXISTS reset_attempts_employee(VARCHAR);
DROP PROCEDURE IF EXISTS get_attempts_employee(VARCHAR);

COMMIT;
-- +goose StatementEnd
