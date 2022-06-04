CREATE TYPE "entry_type" AS ENUM (
  'debit',
  'credit'
);

CREATE TABLE "users" (
                         "id" bigserial PRIMARY KEY,
                         "name" varchar NOT NULL,
                         "email" varchar UNIQUE NOT NULL,
                         "hashed_password" varchar NOT NULL,
                         "created_at" timestamptz NOT NULL DEFAULT (now()),
                         "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "wallets" (
                           "id" bigserial PRIMARY KEY,
                           "user_id" bigint NOT NULL,
                           "balance" bigint NOT NULL,
                           "currency" varchar NOT NULL,
                           "created_at" timestamptz NOT NULL DEFAULT (now()),
                           "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "entries" (
                           "id" bigserial PRIMARY KEY,
                           "wallet_id" bigint NOT NULL,
                           "entry_type" entry_type,
                           "amount" bigint NOT NULL,
                           "created_at" timestamptz NOT NULL DEFAULT (now()),
                           "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "product" (
                           "id" bigserial PRIMARY KEY,
                           "name" varchar NOT NULL,
                           "price" bigint NOT NULL,
                           "created_at" timestamptz NOT NULL DEFAULT (now()),
                           "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "wallets" ("user_id");

CREATE INDEX ON "entries" ("wallet_id");

COMMENT ON COLUMN "entries"."amount" IS 'can be negative or positive';

ALTER TABLE "wallets" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "entries" ADD FOREIGN KEY ("wallet_id") REFERENCES "wallets" ("id") ON DELETE CASCADE;
