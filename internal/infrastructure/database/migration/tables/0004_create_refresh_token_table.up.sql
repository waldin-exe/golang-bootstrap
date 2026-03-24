CREATE TABLE refresh_tokens (
  id SERIAL PRIMARY KEY,
  user_email VARCHAR(255),
  token TEXT,
  expired_at TIMESTAMP,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW(),
  deleted_at TIMESTAMP DEFAULT NULL
);
