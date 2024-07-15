-- +goose Up
-- +goose StatementBegin

-- when user fails the password for user
CREATE OR REPLACE PROCEDURE increase_attempts_client(username_input VARCHAR)
LANGUAGE plpgsql
AS $$
BEGIN
  UPDATE public.employee_account_security eas
  SET 
    attemts = attemts + 1,
    last_attempt = NOW()
  FROM public.employee_account ea
  WHERE ea.username = username_input
    AND eas.employee_account_employee_id = ea.employee_id;

  IF NOT FOUND THEN
    RAISE EXCEPTION 'Username not found';
  END IF;
END;
$$;

-- when user logs in successfully
CREATE OR REPLACE PROCEDURE reset_attempts_client(username_input VARCHAR)
LANGUAGE plpgsql
AS $$
BEGIN
  UPDATE public.employee_account_security eas
  SET 
    attemts = 0,
    last_attempt = NOW()
  FROM public.employee_account ea
  WHERE ea.username = username_input
    AND eas.employee_account_employee_id = ea.employee_id;

  IF NOT FOUND THEN
    RAISE EXCEPTION 'Username not found';
  END IF;
END;
$$;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP PROCEDURE IF EXISTS increase_attempts_client(VARCHAR);
DROP PROCEDURE IF EXISTS reset_attempts_client(VARCHAR);
-- +goose StatementEnd
