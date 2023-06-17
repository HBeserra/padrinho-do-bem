
DROP DATABASE padrinho_db;

create database if not exists padrinho_db;

USE padrinho_db;

CREATE TABLE `user` (
    id INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    full_name VARCHAR(255) NOT NULL,
    display_name VARCHAR(255),
    primary_email VARCHAR(255) UNIQUE NOT NULL,
    document_type ENUM('CPF', 'CNPJ') NOT NULL,
    document_number INT UNSIGNED NOT NULL,
    UNIQUE (document_type,document_number)
);

CREATE TABLE `user_password` (
    user_id INT UNSIGNED,
    password_hash BINARY(60),
    update_date datetime NOT NULL,
    FOREIGN KEY fk_user (user_id)
        REFERENCES user(id)
        ON DELETE CASCADE
);


CREATE TABLE `children` (
    id INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    cpf INT UNSIGNED unique,
    name varchar(255) NOT NULL,
    brith_date date NOT NULL,
    gender enum('male', 'female'),
    address_city VARCHAR(60),
    address_state VARCHAR(60),
    current_situation VARCHAR(255),
    background_history TEXT,
    legal_guardian INT UNSIGNED,
    created_data DATETIME NOT NULL DEFAULT NOW(),
    update_date DATETIME NOT NULL
);

CREATE TRIGGER update_child_date
    AFTER UPDATE ON children
    FOR EACH ROW
    BEGIN
        SET NEW.update_date = NOW();
    end;




