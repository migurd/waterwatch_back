-- +goose Up
-- +goose StatementBegin
BEGIN;


CREATE TABLE IF NOT EXISTS public.account
(
    client_id bigint NOT NULL,
    username character varying(255) COLLATE pg_catalog."default" NOT NULL,
    password character varying(255) COLLATE pg_catalog."default" NOT NULL,
    status boolean DEFAULT true,
    CONSTRAINT account_pkey PRIMARY KEY (client_id),
    CONSTRAINT account_username_key UNIQUE (username)
);

CREATE TABLE IF NOT EXISTS public.account_security
(
    account_client_id bigint NOT NULL,
    attempts integer DEFAULT 0,
    is_password_encrypted boolean DEFAULT false,
    last_attempt timestamp with time zone,
    last_time_password_changed timestamp with time zone,
    CONSTRAINT account_security_pkey PRIMARY KEY (account_client_id)
);

CREATE TABLE IF NOT EXISTS public.appointment
(
    id bigserial NOT NULL,
    appointment_type_id bigint NOT NULL,
    address_id bigint NOT NULL,
    employee_id bigint,
    details character varying(255) COLLATE pg_catalog."default",
    requested_date timestamp with time zone NOT NULL,
    done_date timestamp with time zone,
    CONSTRAINT appointment_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.appointment_type
(
    id bigserial NOT NULL,
    name character varying(50) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT appointment_type_pkey PRIMARY KEY (id)
);

COMMENT ON TABLE public.appointment_type
    IS '1. Installation; 2. Maintenance';

CREATE TABLE IF NOT EXISTS public.client
(
    id bigserial NOT NULL,
    first_name character varying(255) COLLATE pg_catalog."default" NOT NULL,
    last_name character varying(255) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT client_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.client_address
(
    id bigserial NOT NULL,
    client_id bigint NOT NULL,
    state character varying(255) COLLATE pg_catalog."default" NOT NULL,
    city character varying(255) COLLATE pg_catalog."default" NOT NULL,
    street character varying(255) COLLATE pg_catalog."default" NOT NULL,
    house_number character varying(255) COLLATE pg_catalog."default" NOT NULL,
    neighborhood character varying(255) COLLATE pg_catalog."default" NOT NULL,
    postal_code character varying(255) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT client_address_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.client_email
(
    client_id bigint NOT NULL,
    email character varying(255) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT client_email_pkey PRIMARY KEY (client_id),
    CONSTRAINT client_email_email_key UNIQUE (email)
);

CREATE TABLE IF NOT EXISTS public.client_phone_number
(
    client_id bigint NOT NULL,
    country_code character varying(3) COLLATE pg_catalog."default" NOT NULL,
    phone_number character varying(10) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT client_phone_number_pkey PRIMARY KEY (client_id),
    CONSTRAINT client_phone_number_phone_number_key UNIQUE (phone_number)
);

CREATE TABLE IF NOT EXISTS public.contact
(
    id bigserial NOT NULL,
    name character varying(50) COLLATE pg_catalog."default" NOT NULL,
    photo_url character varying(255) COLLATE pg_catalog."default",
    CONSTRAINT contact_pkey PRIMARY KEY (id)
);

COMMENT ON TABLE public.contact
    IS 'There are two types of contacts:
1. Celular
2. Correo';

CREATE TABLE IF NOT EXISTS public.contact_email
(
    contact_id bigint NOT NULL,
    contact_email character varying(50) COLLATE pg_catalog."default" NOT NULL
);

CREATE TABLE IF NOT EXISTS public.contact_phone_number
(
    contact_id bigint NOT NULL,
    country_code character varying(3) COLLATE pg_catalog."default" NOT NULL,
    phone_number character varying(10) COLLATE pg_catalog."default" NOT NULL
);

CREATE TABLE IF NOT EXISTS public.employee
(
    id bigserial NOT NULL,
    employee_type_id bigint NOT NULL,
    first_name character varying(255) COLLATE pg_catalog."default" NOT NULL,
    last_name character varying(255) COLLATE pg_catalog."default" NOT NULL,
    curp character varying(18) COLLATE pg_catalog."default" NOT NULL,
    status boolean NOT NULL DEFAULT true,
    CONSTRAINT employee_pkey PRIMARY KEY (id),
    CONSTRAINT employee_curp_key UNIQUE (curp)
);

CREATE TABLE IF NOT EXISTS public.employee_account
(
    employee_id bigint NOT NULL,
    username character varying(50) COLLATE pg_catalog."default" NOT NULL,
    password character varying(255) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT employee_account_pkey PRIMARY KEY (employee_id)
);

CREATE TABLE IF NOT EXISTS public.employee_account_security
(
    employee_account_employee_id bigint NOT NULL,
    attemts integer DEFAULT 0,
    is_password_encrypted boolean DEFAULT false,
    last_attempt time with time zone,
    last_time_password_changed time with time zone,
    CONSTRAINT employee_account_security_pkey PRIMARY KEY (employee_account_employee_id)
);

CREATE TABLE IF NOT EXISTS public.employee_email
(
    employee_id bigint NOT NULL,
    email character varying(100) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT employee_email_pkey PRIMARY KEY (employee_id),
    CONSTRAINT employee_email_email_key UNIQUE (email)
);

CREATE TABLE IF NOT EXISTS public.employee_phone_number
(
    employee_id bigint NOT NULL,
    country_code character varying(3) COLLATE pg_catalog."default" NOT NULL,
    phone_number character varying(10) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT employee_phone_number_pkey PRIMARY KEY (employee_id),
    CONSTRAINT employee_phone_number_phone_number_employee_id_key UNIQUE (phone_number)
        INCLUDE(employee_id)
);

CREATE TABLE IF NOT EXISTS public.employee_type
(
    id bigserial NOT NULL,
    name character varying(50) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT employee_type_pkey PRIMARY KEY (id)
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
    CONSTRAINT iot_device_pkey PRIMARY KEY (id),
    CONSTRAINT iot_device_serial_key_key UNIQUE (serial_key)
);

CREATE TABLE IF NOT EXISTS public.saa
(
    id bigserial NOT NULL,
    address_id bigint NOT NULL,
    saa_type_id bigint NOT NULL,
    iot_device_id bigint NOT NULL,
    CONSTRAINT saa_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.saa_description
(
    saa_id bigint NOT NULL,
    name character varying(50) COLLATE pg_catalog."default" NOT NULL,
    description character varying(255) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT saa_specific_address_pkey PRIMARY KEY (saa_id)
);

CREATE TABLE IF NOT EXISTS public.saa_maintenance
(
    id bigserial NOT NULL,
    appointment_id bigint NOT NULL,
    saa_id bigint NOT NULL,
    CONSTRAINT saa_maintenance_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.saa_record
(
    id bigserial NOT NULL,
    saa_id bigint NOT NULL,
    water_level double precision NOT NULL,
    ph_level double precision NOT NULL,
    is_contaminated boolean NOT NULL DEFAULT false,
    date timestamp with time zone NOT NULL,
    CONSTRAINT saa_record_pkey PRIMARY KEY (id)
);

COMMENT ON TABLE public.saa_record
    IS 'Water level is in percentage';

CREATE TABLE IF NOT EXISTS public.saa_type
(
    id bigserial NOT NULL,
    name character varying COLLATE pg_catalog."default" NOT NULL,
    description character varying COLLATE pg_catalog."default" NOT NULL,
    capacity integer NOT NULL,
    diameter double precision NOT NULL,
    height double precision NOT NULL,
    CONSTRAINT saa_type_pkey PRIMARY KEY (id)
);

COMMENT ON TABLE public.saa_type
    IS 'capacity must be in L; diameter and height must be in cm';

ALTER TABLE IF EXISTS public.account
    ADD CONSTRAINT account_client_id_fkey FOREIGN KEY (client_id)
    REFERENCES public.client (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;
CREATE INDEX IF NOT EXISTS account_pkey
    ON public.account(client_id);


ALTER TABLE IF EXISTS public.account_security
    ADD CONSTRAINT account_security_account_client_id_fkey FOREIGN KEY (account_client_id)
    REFERENCES public.account (client_id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;
CREATE INDEX IF NOT EXISTS account_security_pkey
    ON public.account_security(account_client_id);


ALTER TABLE IF EXISTS public.appointment
    ADD CONSTRAINT appointment_address_id_fkey FOREIGN KEY (address_id)
    REFERENCES public.client_address (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


ALTER TABLE IF EXISTS public.appointment
    ADD CONSTRAINT appointment_appointment_type_id_fkey FOREIGN KEY (appointment_type_id)
    REFERENCES public.appointment_type (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


ALTER TABLE IF EXISTS public.appointment
    ADD CONSTRAINT appointment_employee_id_fkey FOREIGN KEY (employee_id)
    REFERENCES public.employee (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


ALTER TABLE IF EXISTS public.client_address
    ADD CONSTRAINT client_address_client_id_fkey FOREIGN KEY (client_id)
    REFERENCES public.client (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


ALTER TABLE IF EXISTS public.client_email
    ADD CONSTRAINT client_email_client_id_fkey FOREIGN KEY (client_id)
    REFERENCES public.client (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;
CREATE INDEX IF NOT EXISTS client_email_pkey
    ON public.client_email(client_id);


ALTER TABLE IF EXISTS public.client_phone_number
    ADD CONSTRAINT client_phone_number_client_id_fkey FOREIGN KEY (client_id)
    REFERENCES public.client (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;
CREATE INDEX IF NOT EXISTS client_phone_number_pkey
    ON public.client_phone_number(client_id);


ALTER TABLE IF EXISTS public.contact_email
    ADD CONSTRAINT contact_email_contact_id_fkey FOREIGN KEY (contact_id)
    REFERENCES public.contact (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


ALTER TABLE IF EXISTS public.contact_phone_number
    ADD CONSTRAINT contact_phone_number_contact_id_fkey FOREIGN KEY (contact_id)
    REFERENCES public.contact (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


ALTER TABLE IF EXISTS public.employee
    ADD CONSTRAINT employee_employee_type_id_fkey FOREIGN KEY (employee_type_id)
    REFERENCES public.employee_type (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


ALTER TABLE IF EXISTS public.employee_account
    ADD CONSTRAINT employee_account_employee_id_fkey FOREIGN KEY (employee_id)
    REFERENCES public.employee (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;
CREATE INDEX IF NOT EXISTS employee_account_pkey
    ON public.employee_account(employee_id);


ALTER TABLE IF EXISTS public.employee_account_security
    ADD CONSTRAINT employee_account_security_employee_account_employee_id_fkey FOREIGN KEY (employee_account_employee_id)
    REFERENCES public.employee_account (employee_id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;
CREATE INDEX IF NOT EXISTS employee_account_security_pkey
    ON public.employee_account_security(employee_account_employee_id);


ALTER TABLE IF EXISTS public.employee_email
    ADD CONSTRAINT employee_email_employee_id_fkey FOREIGN KEY (employee_id)
    REFERENCES public.employee (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;
CREATE INDEX IF NOT EXISTS employee_email_pkey
    ON public.employee_email(employee_id);


ALTER TABLE IF EXISTS public.employee_phone_number
    ADD CONSTRAINT employee_phone_number_employee_id_fkey FOREIGN KEY (employee_id)
    REFERENCES public.employee (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;
CREATE INDEX IF NOT EXISTS employee_phone_number_pkey
    ON public.employee_phone_number(employee_id);


ALTER TABLE IF EXISTS public.saa
    ADD CONSTRAINT saa_address_id_fkey FOREIGN KEY (address_id)
    REFERENCES public.client_address (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


ALTER TABLE IF EXISTS public.saa
    ADD CONSTRAINT saa_iot_device_id_fkey FOREIGN KEY (iot_device_id)
    REFERENCES public.iot_device (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


ALTER TABLE IF EXISTS public.saa
    ADD CONSTRAINT saa_saa_type_id_fkey FOREIGN KEY (saa_type_id)
    REFERENCES public.saa_type (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


ALTER TABLE IF EXISTS public.saa_description
    ADD CONSTRAINT saa_specific_address_saa_id_fkey FOREIGN KEY (saa_id)
    REFERENCES public.saa (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;
CREATE INDEX IF NOT EXISTS saa_specific_address_pkey
    ON public.saa_description(saa_id);


ALTER TABLE IF EXISTS public.saa_maintenance
    ADD CONSTRAINT saa_maintenance_appointment_id_fkey FOREIGN KEY (appointment_id)
    REFERENCES public.appointment (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


ALTER TABLE IF EXISTS public.saa_maintenance
    ADD CONSTRAINT saa_maintenance_saa_id_fkey FOREIGN KEY (saa_id)
    REFERENCES public.saa (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


ALTER TABLE IF EXISTS public.saa_record
    ADD CONSTRAINT saa_record_saa_id_fkey FOREIGN KEY (saa_id)
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
DROP TABLE IF EXISTS public.saa_description CASCADE;
DROP TABLE IF EXISTS public.account CASCADE;
DROP TABLE IF EXISTS public.account_security CASCADE;
DROP TABLE IF EXISTS public.employee_account CASCADE;
DROP TABLE IF EXISTS public.employee_account_security CASCADE;
DROP TABLE IF EXISTS public.client CASCADE;
DROP TABLE IF EXISTS public.client_address CASCADE;
DROP TABLE IF EXISTS public.client_email CASCADE;
DROP TABLE IF EXISTS public.client_phone_number CASCADE;
DROP TABLE IF EXISTS public.iot_device CASCADE;
DROP TABLE IF EXISTS public.employee_type CASCADE;
DROP TABLE IF EXISTS public.employee CASCADE;
DROP TABLE IF EXISTS public.employee_email CASCADE;
DROP TABLE IF EXISTS public.employee_phone_number CASCADE;
DROP TABLE IF EXISTS public.appointment_type CASCADE;
DROP TABLE IF EXISTS public.appointment CASCADE;
DROP TABLE IF EXISTS public.contact CASCADE;
DROP TABLE IF EXISTS public.contact_email CASCADE;
DROP TABLE IF EXISTS public.contact_phone_number CASCADE;

END;
-- +goose StatementEnd
