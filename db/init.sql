DROP DATABASE padrinho_db;

create database if not exists padrinho_db;

USE padrinho_db;

CREATE TABLE `address`
(
    id           INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    address_name varchar(60),
    street       VARCHAR(255) NOT NULL,
    number       INT UNSIGNED NOT NULL,
    city         VARCHAR(255) NOT NULL,
    state        CHAR(2),
    postal_code  INT UNSIGNED NOT NULL,
    complement   varchar(255)
);

CREATE TABLE `user`
(
    id              INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    type            ENUM ('administrator', 'user') DEFAULT 'user' NOT NULL,
    full_name       VARCHAR(255)                                  NOT NULL,
    display_name    VARCHAR(255),
    primary_email   VARCHAR(255) UNIQUE                           NOT NULL,
    document_type   ENUM ('CPF', 'CNPJ')                          NOT NULL,
    document_number INT UNSIGNED                                  NOT NULL,
    address_id      INT UNSIGNED,
    UNIQUE (document_type, document_number),
    FOREIGN KEY fk_user_address (address_id)
        REFERENCES address (id)
        ON DELETE set null
);

CREATE TABLE `user_password`
(
    user_id       INT UNSIGNED,
    password_hash BINARY(60),
    update_date   datetime NOT NULL,
    FOREIGN KEY fk_user (user_id)
        REFERENCES user (id)
        ON DELETE CASCADE
);

CREATE TABLE guardian
(
    id                 INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    type               enum ('parents', 'family_member', 'legal_guardian', 'orphanage') NOT NULL,
    document_number    INT UNSIGNED,
    document_type      ENUM ('cpf', 'cnpj'),
    legal_name         VARCHAR(255),
    public_description TINYTEXT,
#     bank_account
    verified           TINYINT(1) UNSIGNED default false,
    verified_by        INT UNSIGNED,
    verification_date  DATETIME,
    created_date       datetime            default now()                                NOT NULL,
    update_date        datetime                                                         NOT NULL,
    FOREIGN KEY fk_verify (verified_by)
        REFERENCES user (id)
        ON DELETE CASCADE
);

CREATE TABLE user_guardian_access
(
    guardian_id INT UNSIGNED NOT NULL,
    user_id     INT UNSIGNED NOT NULL,
    UNIQUE (guardian_id, user_id),
    FOREIGN KEY fk_guardian_access_user (user_id)
        REFERENCES user (id)
        ON DELETE CASCADE,
    FOREIGN KEY fk_guardian_access (guardian_id)
        REFERENCES guardian (id)
        ON DELETE CASCADE
);

CREATE TRIGGER update_guardian_date
    BEFORE UPDATE
    ON guardian
    FOR EACH ROW
BEGIN
    SET NEW.update_date = NOW();
    SET NEW.created_date = OLD.created_date;
end;

CREATE TABLE `children`
(
    id                 INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    cpf                INT UNSIGNED unique,
    name               varchar(255) NOT NULL,
    brith_date         date         NOT NULL,
    gender             enum ('male', 'female'),
    address_city       VARCHAR(60),
    address_state      VARCHAR(60),
    current_situation  VARCHAR(255),
    background_history TEXT,
    legal_guardian     INT UNSIGNED,
    created_date       DATETIME     NOT NULL DEFAULT NOW(),
    update_date        DATETIME     NOT NULL,
    FOREIGN KEY fk_children_legal_guardian (legal_guardian)
        REFERENCES guardian (id)
        ON DELETE CASCADE
);

CREATE TRIGGER update_child_date
    BEFORE UPDATE
    ON children
    FOR EACH ROW
BEGIN
    SET NEW.update_date = NOW();
    SET NEW.created_date = OLD.created_date;
end;

CREATE TABLE interest
(
    id       INT UNSIGNED PRIMARY KEY,
    interest VARCHAR(255) UNIQUE
);


CREATE TABLE children_interest
(
    interest_id INT UNSIGNED NOT NULL,
    children_id INT UNSIGNED NOT NULL,
    UNIQUE (interest_id, children_id),
    FOREIGN KEY fk_children_id (children_id)
        REFERENCES children (id)
        ON DELETE CASCADE,
    FOREIGN KEY fk_interest_id (interest_id)
        REFERENCES interest (id)
        ON DELETE CASCADE
)

