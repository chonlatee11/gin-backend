CREATE EXTENSION IF NOT EXISTS "uuid-ossp" cascade;
CREATE EXTENSION IF NOT EXISTS "pgcrypto" cascade;
CREATE EXTENSION IF NOT EXISTS tablefunc cascade;

CREATE TABLE "users" (
  "id" uuid PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "email" varchar NOT NULL UNIQUE,
  "password" varchar NOT NULL,
  "create_at" timestamp NOT NULL DEFAULT (CURRENT_TIMESTAMP),
  "update_at" timestamp DEFAULT (CURRENT_TIMESTAMP)
);
