-- +goose Up
CREATE TABLE IF NOT EXISTS closs_company(
    ked_codigo TEXT NOT NULL PRIMARY KEY,
    ked_nombre TEXT NOT NULL DEFAULT '',
    ked_status INTEGER NOT NULL DEFAULT 0,
    ked_enlace TEXT NOT NULL DEFAULT '',
    ked_agen TEXT NOT NULL DEFAULT 'mcbo',
    created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TEXT DEFAULT NULL
);
CREATE TABLE IF NOT EXISTS closs_config(
    cnfg_idconfig TEXT NOT NULL,
    cnfg_clase TEXT NOT NULL DEFAULT '',
    cnfg_tipo TEXT NOT NULL,
    cnfg_valnum REAL NOT NULL DEFAULT 0.00,
    cnfg_valsino INTEGER NOT NULL DEFAULT 0,
    cnfg_valtxt TEXT NOT NULL,
    cnfg_lentxt INTEGER NOT NULL DEFAULT 0,
    cnfg_valfch TEXt NOT NULL DEFAULT '1000-01-01',
    cnfg_activa INTEGER NOT NULL DEFAULT 0,
    cnfg_etiq TEXT NOT NULL,
    cnfg_ttip TEXT NOT NULL,
    username TEXT NOT NULL DEFAULT '',
    created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TEXT DEFAULT NULL,
    PRIMARY KEY (cnfg_idconfig, username)
);


-- +goose Down


