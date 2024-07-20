-- +goose Up
-- +goose StatementBegin

-- Appointment Type
INSERT INTO appointment_type(name)
VALUES('INSTALACIÓN');
INSERT INTO appointment_type(name)
VALUES('MANTENIMIENTO');

-- IoT Device
INSERT INTO iot_device(serial_key) VALUES('AAAAA-AAAAA-AAAAA');
INSERT INTO iot_device(serial_key) VALUES('BBBBB-BBBBB-BBBBB');
INSERT INTO iot_device(serial_key) VALUES('CCCCC-CCCCC-CCCCC');
INSERT INTO iot_device(serial_key) VALUES('XXXXX-XXXXX-XXXXX');
INSERT INTO iot_device(serial_key) VALUES('YYYYY-YYYYY-YYYYY');
INSERT INTO iot_device(serial_key) VALUES('ZZZZZ-ZZZZZ-ZZZZZ');

-- Employee Type
INSERT INTO employee_type(name) VALUES('INSTALADOR');
INSERT INTO employee_type(name) VALUES('MANTENEDOR');
INSERT INTO employee_type(name) VALUES('INSTALADOR Y MANTENEDOR');

-- Employee
INSERT INTO employee(employee_type_id, first_name, last_name, curp, status)
VALUES((SELECT id FROM employee_type WHERE name = 'INSTALADOR'), 'Instalador', 'Num 1', 'AAAAAAAAAAAAAAAAAA', TRUE);
INSERT INTO employee(employee_type_id, first_name, last_name, curp, status)
VALUES((SELECT id FROM employee_type WHERE name = 'INSTALADOR'), 'Instalador', 'Num 2', 'BBBBBBBBBBBBBBBBBB', TRUE);
INSERT INTO employee(employee_type_id, first_name, last_name, curp, status)
VALUES((SELECT id FROM employee_type WHERE name = 'MANTENEDOR'), 'Mantenedor', 'Num 1', 'CCCCCCCCCCCCCCCCCC', TRUE);
INSERT INTO employee(employee_type_id, first_name, last_name, curp, status)
VALUES((SELECT id FROM employee_type WHERE name = 'MANTENEDOR'), 'Mantenedor', 'Num 2', 'DDDDDDDDDDDDDDDDDD', TRUE);
INSERT INTO employee(employee_type_id, first_name, last_name, curp, status)
VALUES((SELECT id FROM employee_type WHERE name = 'INSTALADOR Y MANTENEDOR'), 'Todólogo', 'Num 1', 'EEEEEEEEEEEEEEEEEE', TRUE);
INSERT INTO employee(employee_type_id, first_name, last_name, curp, status)
VALUES((SELECT id FROM employee_type WHERE name = 'INSTALADOR Y MANTENEDOR'), 'Todólogo', 'Num 2', 'FFFFFFFFFFFFFFFFFF', TRUE);

-- Employee email & phone number
INSERT INTO employee_email(employee_id, email) VALUES(1, 'a@gmail.com');
INSERT INTO employee_email(employee_id, email) VALUES(2, 'b@gmail.com');
INSERT INTO employee_email(employee_id, email) VALUES(3, 'c@gmail.com');
INSERT INTO employee_email(employee_id, email) VALUES(4, 'd@gmail.com');
INSERT INTO employee_email(employee_id, email) VALUES(5, 'e@gmail.com');
INSERT INTO employee_email(employee_id, email) VALUES(6, 'f@gmail.com');
INSERT INTO employee_phone_number(employee_id, country_code, phone_number) VALUES(1, '52', '1111111111');
INSERT INTO employee_phone_number(employee_id, country_code, phone_number) VALUES(2, '52', '2222222222');
INSERT INTO employee_phone_number(employee_id, country_code, phone_number) VALUES(3, '52', '3333333333');
INSERT INTO employee_phone_number(employee_id, country_code, phone_number) VALUES(4, '52', '4444444444');
INSERT INTO employee_phone_number(employee_id, country_code, phone_number) VALUES(5, '52', '5555555555');
INSERT INTO employee_phone_number(employee_id, country_code, phone_number) VALUES(6, '52', '6666666666');

-- Employee account
INSERT INTO employee_account(employee_id, username, password)
VALUES(1, 'a', '123');
INSERT INTO employee_account(employee_id, username, password)
VALUES(2, 'b', '123');
INSERT INTO employee_account(employee_id, username, password)
VALUES(3, 'c', '123');
INSERT INTO employee_account(employee_id, username, password)
VALUES(4, 'd', '123');
INSERT INTO employee_account(employee_id, username, password)
VALUES(5, 'e', '123');
INSERT INTO employee_account(employee_id, username, password)
VALUES(6, 'f', '123');

-- Employee account security
INSERT INTO employee_account_security(employee_account_employee_id)
VALUES(1);
INSERT INTO employee_account_security(employee_account_employee_id)
VALUES(2);
INSERT INTO employee_account_security(employee_account_employee_id)
VALUES(3);
INSERT INTO employee_account_security(employee_account_employee_id)
VALUES(4);
INSERT INTO employee_account_security(employee_account_employee_id)
VALUES(5);
INSERT INTO employee_account_security(employee_account_employee_id)
VALUES(6);

-- client
INSERT INTO client(first_name, last_name)
VALUES('Angel', 'Qui');

INSERT INTO client_address(client_id, "state", city, street, house_number, neighborhood, postal_code)
VALUES(1, 'Sinaloa', 'Mazatlán', 'Calle Pepito', '2222', 'Venadillo', '60060');

INSERT INTO client_phone_number(client_id, country_code, phone_number)
VALUES(1, '52', '9999999999');

INSERT INTO client_email(client_id, email)
VALUES(1, 'angel@gmail.com');

INSERT INTO account(client_id, username, "password", "status")
VALUES(1, 'angelqui', '123', TRUE);

INSERT INTO account_security(account_client_id, is_password_encrypted)
VALUES(1, FALSE);

-- enable first SAA for testing
INSERT INTO appointment(appointment_type_id, address_id, client_id, employee_id, details, requested_date)
VALUES((SELECT id FROM appointment_type WHERE "name" = 'INSTALACIÓN'), 1, 1, 1, 'huwwwyyyy up! my tinaco is all caca', '2024-10-10');

INSERT INTO saa_type("name", "description", capacity, diameter, height)
VALUES('TINACO TRICAPA 1100 L EQUIPADO', 'Tinaco Sistema Mejor Agua (SMA) 1100 L Equipado Rotoplas fabricado con polietileno lineal de baja densidad, color beige por fuera y blanco por dentro.', 1100, 110, 137);

INSERT INTO saa(appointment_id, saa_type_id, iot_device_id)
VALUES(1, 1, (SELECT id FROM iot_device WHERE serial_key = 'AAAAA-AAAAA-AAAAA'));

INSERT INTO saa_description(saa_id, "name", "description")
VALUES(1, 'Tinaco 1', 'Tinaco que está en el techo de mi abuelita Pancha');

UPDATE appointment SET done_date = '2024-10-10';

UPDATE iot_device SET "status" = TRUE WHERE serial_key = 'AAAAA-AAAAA-AAAAA';

-- create contacts
INSERT INTO contact("name", photo_url)
VALUES('WaterWatch', 'https://pbs.twimg.com/media/Fkhdzf9VQAE4fbn.jpg');
INSERT INTO contact("name", photo_url)
VALUES('SHITSU', 'https://wallpapercave.com/wp/wp8856774.jpg');
INSERT INTO contact("name", photo_url)
VALUES('Ayuda', 'https://static.zerochan.net/Rudeus.Greyrat.full.4003848.jpg');

INSERT INTO contact_email(contact_id, email)
VALUES(1, 'ayuda1@waterwatch.com');
INSERT INTO contact_email(contact_id, email)
VALUES(1, 'ayuda2@waterwatch.com');

INSERT INTO contact_email(contact_id, email)
VALUES(2, 'ayuda1@shitsu.com');
INSERT INTO contact_email(contact_id, email)
VALUES(2, 'ayuda2@shitsu.com');

INSERT INTO contact_email(contact_id, email)
VALUES(3, 'superayuda1@waterwatch.com');
INSERT INTO contact_email(contact_id, email)
VALUES(3, 'superayuda2@waterwatch.com');
INSERT INTO contact_email(contact_id, email)
VALUES(3, 'superayuda3@waterwatch.com');

INSERT INTO contact_phone_number(contact_id, country_code, phone_number)
VALUES(1, '52', '1111111111');
INSERT INTO contact_phone_number(contact_id, country_code, phone_number)
VALUES(2, '52', '2222222222');
INSERT INTO contact_phone_number(contact_id, country_code, phone_number)
VALUES(3, '52', '3333333333');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
TRUNCATE TABLE appointment_type CASCADE;
TRUNCATE TABLE iot_device CASCADE;
TRUNCATE TABLE employee_type CASCADE;
-- +goose StatementEnd
