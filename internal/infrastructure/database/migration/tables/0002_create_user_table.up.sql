CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL,
    level VARCHAR(20),
    pegawai_id INTEGER,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(100),
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_by VARCHAR(100),
    deleted_at TIMESTAMP DEFAULT NULL,
    deleted_by VARCHAR(100) DEFAULT NULL,
    CONSTRAINT fk_pegawai FOREIGN KEY (pegawai_id) REFERENCES pegawais(id)
);