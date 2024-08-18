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
VALUES((SELECT id FROM employee_type WHERE name = 'INSTALADOR'), 'Arturo', 'Rivas Paredes', 'AAAAAAAAAAAAAAAAAA', TRUE);
INSERT INTO employee(employee_type_id, first_name, last_name, curp, status)
VALUES((SELECT id FROM employee_type WHERE name = 'INSTALADOR'), 'Mandujano', 'González Pérez', 'BBBBBBBBBBBBBBBBBB', TRUE);
INSERT INTO employee(employee_type_id, first_name, last_name, curp, status)
VALUES((SELECT id FROM employee_type WHERE name = 'MANTENEDOR'), 'Pablo', 'Urrea Solís', 'CCCCCCCCCCCCCCCCCC', TRUE);
INSERT INTO employee(employee_type_id, first_name, last_name, curp, status)
VALUES((SELECT id FROM employee_type WHERE name = 'MANTENEDOR'), 'Adrián', 'Osuna Morales', 'DDDDDDDDDDDDDDDDDD', TRUE);
INSERT INTO employee(employee_type_id, first_name, last_name, curp, status)
VALUES((SELECT id FROM employee_type WHERE name = 'INSTALADOR Y MANTENEDOR'), 'Juan', 'Tirado Romero', 'EEEEEEEEEEEEEEEEEE', TRUE);
INSERT INTO employee(employee_type_id, first_name, last_name, curp, status)
VALUES((SELECT id FROM employee_type WHERE name = 'INSTALADOR Y MANTENEDOR'), 'Abel', 'Sánchez López', 'FFFFFFFFFFFFFFFFFF', TRUE);

-- Employee email & phone number
INSERT INTO employee_email(employee_id, email) VALUES(1, 'arturo_instalador@waterwatch.com');
INSERT INTO employee_email(employee_id, email) VALUES(2, 'mandujano_instalador@waterwatch.com');
INSERT INTO employee_email(employee_id, email) VALUES(3, 'pablo_mantenedora@waterwatch.com');
INSERT INTO employee_email(employee_id, email) VALUES(4, 'adrian_mantenedor@gmail.com');
INSERT INTO employee_email(employee_id, email) VALUES(5, 'juanjose_todologo@gmail.com');
INSERT INTO employee_email(employee_id, email) VALUES(6, 'abelsanzdev_todologo@gmail.com');
INSERT INTO employee_phone_number(employee_id, country_code, phone_number) VALUES(1, '52', '6692583938');
INSERT INTO employee_phone_number(employee_id, country_code, phone_number) VALUES(2, '52', '6692531290');
INSERT INTO employee_phone_number(employee_id, country_code, phone_number) VALUES(3, '52', '6692843492');
INSERT INTO employee_phone_number(employee_id, country_code, phone_number) VALUES(4, '52', '6692024543');
INSERT INTO employee_phone_number(employee_id, country_code, phone_number) VALUES(5, '52', '6692268561');
INSERT INTO employee_phone_number(employee_id, country_code, phone_number) VALUES(6, '52', '6692164023');

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
VALUES((SELECT id FROM appointment_type WHERE "name" = 'INSTALACIÓN'), 1, 1, 1, 'Mi casa es un changarro ya viejito, pero sí me gustaría tener agua acá bien monitoreada como ustedes promocionan.', '2024-10-10');

INSERT INTO saa_type("name", "description", capacity, diameter, height)
VALUES('Garrafón de agua cortado', 'Se utilizó este contenedor como tinaco para poder hacer un test con un contenedor de agua.', 15, 25, 35);
INSERT INTO saa_type("name", "description", capacity, diameter, height)
VALUES('Pecera de prueba', 'Pecera utilizada para la prueba de demostración que blah blah', 38, 30, 30);

INSERT INTO saa(appointment_id, saa_type_id, saa_type_id2, iot_device_id)
VALUES(1, 1, 2, (SELECT id FROM iot_device WHERE serial_key = 'AAAAA-AAAAA-AAAAA'));

INSERT INTO saa_description(saa_id, "name", "description")
VALUES(1, 'Tinaco 1', 'Tinaco que está en el techo de mi abuelita Pancha');

INSERT INTO saa_record(saa_id, water_level, water_level2, ph_level, date)
VALUES(1, 80, 100, 7, NOW());

UPDATE appointment SET done_date = '2024-10-10' WHERE id = 1;

UPDATE iot_device SET "status" = TRUE WHERE serial_key = 'AAAAA-AAAAA-AAAAA';

-- first maintenance thingy
INSERT INTO appointment(appointment_type_id, address_id, client_id, employee_id, details, requested_date)
VALUES((SELECT id FROM appointment_type WHERE "name" = 'MANTENIMIENTO'), 1, 1, 3, '¡Hay agua de coco por todas partes, apresúrense!', '2024-08-12');

UPDATE appointment SET done_date = '2024-08-12' WHERE id = 2;

-- create contacts
INSERT INTO contact("name", photo_url)
VALUES('WaterWatch', 'https://imgur.com/hbmnZqH.png');
INSERT INTO contact("name", photo_url)
VALUES('SHITSU', 'https://imgur.com/sSm3TLY.png');
INSERT INTO contact("name", photo_url)
VALUES('Jumapam', 'https://imgur.com/emtEJot.png');
INSERT INTO contact("name", photo_url)
VALUES('Rotoplas', 'https://imgur.com/NG8R6Vv.png');

INSERT INTO contact_email(contact_id, email)
VALUES(1, 'peraza_rh@waterwatch.com');
INSERT INTO contact_email(contact_id, email)
VALUES(1, 'kevin_rh@waterwatch.com');

INSERT INTO contact_email(contact_id, email)
VALUES(2, 'angelrodev@shitsu.com');
INSERT INTO contact_email(contact_id, email)
VALUES(2, 'abigailnatalia@shitsu.com');

INSERT INTO contact_email(contact_id, email)
VALUES(3, 'apoyojumapam@jumapam.com');
INSERT INTO contact_email(contact_id, email)
VALUES(3, 'soportejumapam@jumapam.com');
INSERT INTO contact_email(contact_id, email)
VALUES(3, 'ayudajumapam@jumapam.com');

INSERT INTO contact_email(contact_id, email)
VALUES(4, 'carlosrojas@rotoplas.com');
INSERT INTO contact_email(contact_id, email)
VALUES(4, 'motavelasco@rotoplas.com');

INSERT INTO contact_phone_number(contact_id, country_code, phone_number)
VALUES(1, '52', '6692667793');
INSERT INTO contact_phone_number(contact_id, country_code, phone_number)
VALUES(2, '52', '6693362323');
INSERT INTO contact_phone_number(contact_id, country_code, phone_number)
VALUES(3, '52', '6692617697');
INSERT INTO contact_phone_number(contact_id, country_code, phone_number)
VALUES(4, '52', '6699946809');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
TRUNCATE TABLE appointment_type CASCADE;
TRUNCATE TABLE iot_device CASCADE;
TRUNCATE TABLE employee_type CASCADE;
-- +goose StatementEnd
