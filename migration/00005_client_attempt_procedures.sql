-- +goose Up
-- +goose StatementBegin

-- when user fails the password for user
CREATE OR REPLACE PROCEDURE increase_attempts_client(username_input VARCHAR)
LANGUAGE plpgsql
AS $$
DECLARE
  current_attempts INT;
  last_attempt_time TIMESTAMP;
BEGIN
  SELECT account_security.attempts, account_security.last_attempt
  INTO current_attempts, last_attempt_time
  FROM public.account_security
  JOIN public.account
    ON account_security.account_client_id = account.client_id
  WHERE account.username = username_input;

  IF NOT FOUND THEN
    RAISE EXCEPTION 'Username not found';
  END IF;

  IF current_attempts >= 5 AND NOW() - last_attempt_time <= INTERVAL '1 minute' THEN
    RAISE EXCEPTION 'Please wait for a minute before trying again';
  END IF;

  IF current_attempts >= 5 THEN
    UPDATE public.account_security
    SET 
      attempts = 1,
      last_attempt = NOW()
    FROM public.account
    WHERE account.username = username_input
      AND account_security.account_client_id = account.client_id;
  ELSE
    UPDATE public.account_security
    SET 
      attempts = attempts + 1,
      last_attempt = NOW()
    FROM public.account
    WHERE account.username = username_input
      AND account_security.account_client_id = account.client_id;
  END IF;
END;
$$;

-- when user logs in successfully
CREATE OR REPLACE PROCEDURE reset_attempts_client(username_input VARCHAR)
LANGUAGE plpgsql
AS $$
BEGIN
  UPDATE public.account_security
  SET 
    attempts = 0,
    last_attempt = NOW()
  FROM public.account
  WHERE account.username = username_input
    AND account_security.account_client_id = account.client_id;

  IF NOT FOUND THEN
    RAISE EXCEPTION 'Username not found';
  END IF;
END;
$$;

-- get attempts
CREATE OR REPLACE PROCEDURE get_attempts_client(username_input VARCHAR, OUT attempts_count INT)
LANGUAGE plpgsql
AS $$
BEGIN
  SELECT account_security.attempts
  INTO attempts_count
  FROM public.account_security
  JOIN public.account
    ON account_security.account_client_id = account.client_id
  WHERE account.username = username_input;

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
DROP PROCEDURE IF EXISTS get_attempts_client(VARCHAR);
-- +goose StatementEnd
