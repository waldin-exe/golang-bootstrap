CREATE TABLE pegawais (
    id SERIAL PRIMARY KEY,
    nama VARCHAR(100) NOT NULL,
    tgl_lahir DATE,
    alamat TEXT,
    no_telepon VARCHAR(15),
    jenis_pegawai VARCHAR(30),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(100),
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_by VARCHAR(100),
    deleted_at TIMESTAMP DEFAULT NULL,
    deleted_by VARCHAR(100) DEFAULT NULL
);

INSERT INTO pegawais (
    nama, tgl_lahir, alamat, no_telepon, jenis_pegawai, created_by, updated_by
) VALUES (
    'Owner', '1999-01-01', 'Depok', '0895902123', 'owner', '1', '1'
);