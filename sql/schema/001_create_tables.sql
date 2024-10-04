-- +goose Up
CREATE TABLE IF NOT EXISTS ke_dataconex(
    ked_codigo VARCHAR(6) NOT NULL PRIMARY KEY,
    ked_nombre VARCHAR(80) NOT NULL,
    ked_status INT NOT NULL DEFAULT 0,
    ked_enlace TEXT NOT NULL,
    ked_agen VARCHAR(20) NOT NULL DEFAULT '',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);
-- +goose Down
DROP TABLE ke_dataconex;
