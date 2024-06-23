-- +goose Up
-- +goose StatementBegin

-- Appointment Type
INSERT INTO appointment_type (name)
VALUES ('CLIENTE');
INSERT INTO appointment_type (name)
VALUES ('MANTENIMIENTO');

-- IoT Device
INSERT INTO iot_device (serial_key) VALUES ('AAAAA-AAAAA-AAAAA');
INSERT INTO iot_device (serial_key) VALUES ('BBBBB-BBBBB-BBBBB');
INSERT INTO iot_device (serial_key) VALUES ('CCCCC-CCCCC-CCCCC');
INSERT INTO iot_device (serial_key) VALUES ('XXXXX-XXXXX-XXXXX');
INSERT INTO iot_device (serial_key) VALUES ('YYYYY-YYYYY-YYYYY');
INSERT INTO iot_device (serial_key) VALUES ('ZZZZZ-ZZZZZ-ZZZZZ');

-- Employee Type
INSERT INTO employee_type (id, name) VALUES (1, 'INSTALADOR');
INSERT INTO employee_type (id, name) VALUES (2, 'MANTENEDOR');

-- Employee
INSERT INTO employee (id, employee_type_id, first_name, last_name, curp, status)
VALUES (1, (SELECT id FROM employee_type WHERE name = 'INSTALADOR'), 'Instalador', 'Num 1', 'AAAAAAAAAAAAAAAAAA', TRUE);
INSERT INTO employee (id, employee_type_id, first_name, last_name, curp, status)
VALUES (2, (SELECT id FROM employee_type WHERE name = 'INSTALADOR'), 'Instalador', 'Num 2', 'BBBBBBBBBBBBBBBBBB', TRUE);
INSERT INTO employee (id, employee_type_id, first_name, last_name, curp, status)
VALUES (3, (SELECT id FROM employee_type WHERE name = 'MANTENEDOR'), 'Mantenedor', 'Num 1', 'CCCCCCCCCCCCCCCCCC', TRUE);
INSERT INTO employee (id, employee_type_id, first_name, last_name, curp, status)
VALUES (4, (SELECT id FROM employee_type WHERE name = 'MANTENEDOR'), 'Mantenedor', 'Num 2', 'DDDDDDDDDDDDDDDDDD', TRUE);

-- Employee email & phone number
INSERT INTO employee_email (employee_id, email) VALUES (1, 'a@gmail.com');
INSERT INTO employee_email (employee_id, email) VALUES (2, 'b@gmail.com');
INSERT INTO employee_email (employee_id, email) VALUES (3, 'c@gmail.com');
INSERT INTO employee_email (employee_id, email) VALUES (4, 'd@gmail.com');
INSERT INTO employee_phone_number (employee_id, country_code, phone_number) VALUES (1, '52', '1111111111');
INSERT INTO employee_phone_number (employee_id, country_code, phone_number) VALUES (2, '52', '2222222222');
INSERT INTO employee_phone_number (employee_id, country_code, phone_number) VALUES (3, '52', '3333333333');
INSERT INTO employee_phone_number (employee_id, country_code, phone_number) VALUES (4, '52', '4444444444');

-- Employee account details
INSERT INTO employee_account (employee_id, username, password)
VALUES (1, 'a', '123');
INSERT INTO employee_account_security (employee_account_employee_id)
VALUES (1);
INSERT INTO employee_account (employee_id, username, password)
VALUES (2, 'b', '123');
INSERT INTO employee_account_security (employee_account_employee_id)
VALUES (2);
INSERT INTO employee_account (employee_id, username, password)
VALUES (3, 'c', '123');
INSERT INTO employee_account_security (employee_account_employee_id)
VALUES (3);
INSERT INTO employee_account (employee_id, username, password)
VALUES (4, 'd', '123');
INSERT INTO employee_account_security (employee_account_employee_id)
VALUES (4);


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
TRUNCATE TABLE appointment_type CASCADE;
TRUNCATE TABLE iot_device CASCADE;
TRUNCATE TABLE employee_type CASCADE;
-- +goose StatementEnd
