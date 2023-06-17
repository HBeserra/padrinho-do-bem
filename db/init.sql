
DROP DATABASE padrinho_db;

create database if not exists padrinho_db;

USE padrinho_db;

CREATE TABLE `user` (
    id INT PRIMARY KEY AUTO_INCREMENT,
    full_name VARCHAR(255) NOT NULL,
    display_name VARCHAR(255),
    primary_email VARCHAR(255) UNIQUE NOT NULL,
    document_type ENUM('CPF', 'CNPJ') NOT NULL,
    document_number INT UNSIGNED NOT NULL,
    UNIQUE (document_type,document_number)
);

CREATE TABLE `user_password` (
    user_id INT,
    password_hash BINARY(60),
    update_date datetime NOT NULL,
    FOREIGN KEY fk_user (user_id)
        REFERENCES user(id)
        ON DELETE CASCADE
);





