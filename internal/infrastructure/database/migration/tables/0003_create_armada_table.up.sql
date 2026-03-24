CREATE TABLE jenis_armadas (
    id SERIAL PRIMARY KEY,
    jenis VARCHAR(50) NOT NULL,
    nama VARCHAR(50) NOT NULL
);

INSERT INTO jenis_armadas (
    jenis,
    nama
) VALUES ('elf', 'ELF'), ('hiace', 'Hiace'), ('medium-bus', 'Medium Bus'), ('bigbus', 'Bigbus');

CREATE TABLE armadas (
    id SERIAL PRIMARY KEY,
    plat_nomor VARCHAR(10) NOT NULL,
    nomor_lambung VARCHAR(10) NOT NULL,
    jumlah_seat INTEGER NOT NULL,
    merk VARCHAR(50) NOT NULL,
    tahun VARCHAR(10),
    no_kir VARCHAR(100),
    masa_berlaku_kir DATE,
    id_jenis_armada INTEGER NOT NULL,
    body VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(100),
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_by VARCHAR(100),
    deleted_at TIMESTAMP DEFAULT NULL,
    deleted_by VARCHAR(100) DEFAULT NULL,

    CONSTRAINT fk_jenis_armada FOREIGN KEY (id_jenis_armada) REFERENCES jenis_armadas(id)
);