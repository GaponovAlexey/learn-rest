CREATE TABLE
  users (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    email VARCHAR NOT NULL UNIQUE,
    encrypted_password VARCHAR NOT NULL
  )

  