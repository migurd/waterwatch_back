-- +goose Up
-- +goose StatementBegin
-- This script was generated by the ERD tool in pgAdmin 4.
-- Please log an issue at https://github.com/pgadmin-org/pgadmin4/issues/new/choose if you find any bugs, including reproduction steps.
BEGIN;


CREATE TABLE IF NOT EXISTS public.saa
(
    id bigserial NOT NULL,
    client_id bigint NOT NULL,
    saa_type_id bigint NOT NULL,
    iot_device_id bigint NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.saa_maintenance
(
    id bigserial NOT NULL,
    saa_id bigint NOT NULL,
    details character varying(1023) COLLATE pg_catalog."default" NOT NULL,
    requested_date timestamp with time zone NOT NULL,
    done_date timestamp with time zone,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.saa_record
(
    id bigserial NOT NULL,
    saa_id bigint NOT NULL,
    water_level double precision NOT NULL,
    ph_level double precision NOT NULL,
    is_contaminated boolean NOT NULL DEFAULT false,
    date timestamp with time zone NOT NULL,
    PRIMARY KEY (id)
);

COMMENT ON TABLE public.saa_record
    IS 'Water level is in percentage';

CREATE TABLE IF NOT EXISTS public.saa_type
(
    id bigserial NOT NULL,
    name character varying COLLATE pg_catalog."default" NOT NULL,
    description character varying COLLATE pg_catalog."default" NOT NULL,
    capacity integer NOT NULL,
    PRIMARY KEY (id)
);

COMMENT ON TABLE public.saa_type
    IS 'capacity is in mL';

CREATE TABLE IF NOT EXISTS public.account
(
    client_id bigint NOT NULL,
    username character varying(255) COLLATE pg_catalog."default" NOT NULL,
    password character varying(255) COLLATE pg_catalog."default" NOT NULL,
    PRIMARY KEY (client_id),
    CONSTRAINT account_username_key UNIQUE (username)
);

CREATE TABLE IF NOT EXISTS public.account_security
(
    user_id bigint NOT NULL,
    attempts integer NOT NULL,
    max_attempts integer DEFAULT 5,
    last_attempt timestamp with time zone NOT NULL,
    last_time_password_changed timestamp with time zone,
    is_password_encrypted boolean DEFAULT FALSE,
    CONSTRAINT account_security_pkey PRIMARY KEY (user_id)
);

CREATE TABLE IF NOT EXISTS public.client_address
(
    client_id bigint NOT NULL,
    state character varying(255) COLLATE pg_catalog."default" NOT NULL,
    city character varying(255) COLLATE pg_catalog."default" NOT NULL,
    street character varying(255) COLLATE pg_catalog."default" NOT NULL,
    house_number character varying(255) COLLATE pg_catalog."default" NOT NULL,
    neighborhood character varying(255) COLLATE pg_catalog."default" NOT NULL,
    postal_code character varying(255) COLLATE pg_catalog."default" NOT NULL,
    PRIMARY KEY (client_id)
);

CREATE TABLE IF NOT EXISTS public.appointment
(
    id bigserial NOT NULL,
    client_id bigint NOT NULL,
    employee_id bigint,
    requested_date timestamp with time zone NOT NULL,
    done_date timestamp with time zone,
    CONSTRAINT appointment_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.employee
(
    id bigserial NOT NULL,
    first_name character varying(255) COLLATE pg_catalog."default" NOT NULL,
    last_name character varying(255) COLLATE pg_catalog."default" NOT NULL,
    curp character varying(18) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT employee_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.goose_db_version
(
    id serial NOT NULL,
    version_id bigint NOT NULL,
    is_applied boolean NOT NULL,
    tstamp timestamp without time zone DEFAULT now(),
    CONSTRAINT goose_db_version_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.iot_device
(
    id bigserial NOT NULL,
    serial_key character varying(23) COLLATE pg_catalog."default" NOT NULL,
    status boolean NOT NULL DEFAULT false,
    PRIMARY KEY (id),
    UNIQUE (serial_key)
);

CREATE TABLE IF NOT EXISTS public.client
(
    id bigserial NOT NULL,
    address_id bigint DEFAULT 1,
    first_name character varying(255) COLLATE pg_catalog."default" NOT NULL,
    last_name character varying(255) COLLATE pg_catalog."default" NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.client_email
(
    client_id bigint NOT NULL,
    email character varying(255) COLLATE pg_catalog."default" NOT NULL,
    PRIMARY KEY (client_id),
    UNIQUE (email)
);

CREATE TABLE IF NOT EXISTS public.client_phone_number
(
    client_id bigint NOT NULL,
    country_code character varying(3) NOT NULL,
    phone_number character varying(10) COLLATE pg_catalog."default" NOT NULL,
    PRIMARY KEY (client_id),
    UNIQUE (phone_number)
);

CREATE TABLE IF NOT EXISTS public.employee_email
(
    employee_id bigint NOT NULL,
    email character varying(100) NOT NULL,
    PRIMARY KEY (employee_id),
    UNIQUE (email)
);

CREATE TABLE IF NOT EXISTS public.employee_phone_number
(
    employee_id bigint NOT NULL,
    country_code character varying(3) NOT NULL,
    phone_number character varying(10) NOT NULL,
    PRIMARY KEY (employee_id),
    UNIQUE (phone_number)
);

CREATE TABLE IF NOT EXISTS public.saa_specific_address
(
    saa_id bigint NOT NULL,
    name character varying(50) NOT NULL,
    description character varying(255) NOT NULL,
    PRIMARY KEY (saa_id)
);

ALTER TABLE IF EXISTS public.saa
    ADD FOREIGN KEY (client_id)
    REFERENCES public.client (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


ALTER TABLE IF EXISTS public.saa
    ADD FOREIGN KEY (saa_type_id)
    REFERENCES public.saa_type (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


ALTER TABLE IF EXISTS public.saa
    ADD FOREIGN KEY (iot_device_id)
    REFERENCES public.iot_device (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


ALTER TABLE IF EXISTS public.saa_maintenance
    ADD FOREIGN KEY (saa_id)
    REFERENCES public.saa (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


ALTER TABLE IF EXISTS public.saa_record
    ADD FOREIGN KEY (saa_id)
    REFERENCES public.saa (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


ALTER TABLE IF EXISTS public.account
    ADD FOREIGN KEY (client_id)
    REFERENCES public.client (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


ALTER TABLE IF EXISTS public.account_security
    ADD CONSTRAINT account_security_user_id_fkey FOREIGN KEY (user_id)
    REFERENCES public.account (client_id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;
CREATE INDEX IF NOT EXISTS account_security_pkey
    ON public.account_security(user_id);


ALTER TABLE IF EXISTS public.client_address
    ADD FOREIGN KEY (client_id)
    REFERENCES public.client (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


ALTER TABLE IF EXISTS public.appointment
    ADD CONSTRAINT appointment_employee_id_fkey FOREIGN KEY (employee_id)
    REFERENCES public.employee (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


ALTER TABLE IF EXISTS public.appointment
    ADD FOREIGN KEY (client_id)
    REFERENCES public.client (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


ALTER TABLE IF EXISTS public.client_email
    ADD FOREIGN KEY (client_id)
    REFERENCES public.client (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


ALTER TABLE IF EXISTS public.client_phone_number
    ADD FOREIGN KEY (client_id)
    REFERENCES public.client (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


ALTER TABLE IF EXISTS public.employee_email
    ADD FOREIGN KEY (employee_id)
    REFERENCES public.employee (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


ALTER TABLE IF EXISTS public.employee_phone_number
    ADD FOREIGN KEY (employee_id)
    REFERENCES public.employee (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


ALTER TABLE IF EXISTS public.saa_specific_address
    ADD FOREIGN KEY (saa_id)
    REFERENCES public.saa (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;

END;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
BEGIN;


DROP TABLE IF EXISTS public.saa CASCADE;
DROP TABLE IF EXISTS public.saa_record CASCADE;
DROP TABLE IF EXISTS public.saa_type CASCADE;
DROP TABLE IF EXISTS public.saa_specific_address CASCADE;
DROP TABLE IF EXISTS public.saa_maintenance CASCADE;
DROP TABLE IF EXISTS public.account CASCADE;
DROP TABLE IF EXISTS public.account_security CASCADE;
DROP TABLE IF EXISTS public.client CASCADE;
DROP TABLE IF EXISTS public.client_address CASCADE;
DROP TABLE IF EXISTS public.client_email CASCADE;
DROP TABLE IF EXISTS public.client_phone_number CASCADE;
DROP TABLE IF EXISTS public.iot_device CASCADE;
DROP TABLE IF EXISTS public.employee CASCADE;
DROP TABLE IF EXISTS public.employee_email CASCADE;
DROP TABLE IF EXISTS public.employee_phone_number CASCADE;
DROP TABLE IF EXISTS public.appointment CASCADE;

END;
-- +goose StatementEnd
