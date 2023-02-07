CREATE TABLE
    IF NOT EXISTS users (
        id char(36) NOT NULL,
        name varchar(255) NOT NULL,
        email varchar(255) NOT NULL,
        password varchar(255) DEFAULT NULL,
        connected_account_id bigint (20) unsigned DEFAULT NULL,
        created_at timestamp NULL DEFAULT NULL,
        updated_at timestamp NULL DEFAULT NULL,
        email_verified_at timestamp NULL DEFAULT NULL,
        PRIMARY KEY (id),
        UNIQUE KEY users_email_unique (email)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;