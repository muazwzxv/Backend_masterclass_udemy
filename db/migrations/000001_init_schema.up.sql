CREATE TYPE Currency AS ENUM (
  'USD',
  'EUR',
  'MYR',
  'SGD'
);

CREATE TABLE users (
  "id" bigserial PRIMARY KEY,
  "first_name" varchar,
  "last_name" varchar,
  "email" varchar UNIQUE NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  "deleted_at" timestamptz
);

CREATE TABLE accounts (
  "id" bigserial PRIMARY KEY,
  "owner_id" bigint NOT NULL,
  "balance" bigint NOT NULL,
  "currency" Currency NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  "deleted_at" timestamptz
);

CREATE TABLE entries (
  "id" bigserial PRIMARY KEY,
  "account_id" bigint,
  "amount" bigint NOT NULL,
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE transfers (
  "id" bigserial PRIMARY KEY,
  "from_account_id" bigint NOT NULL,
  "to_account_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz DEFAULT (now())
);

CREATE INDEX ON "entries" ("account_id");

CREATE INDEX ON "transfers" ("from_account_id");

CREATE INDEX ON "transfers" ("to_account_id");

CREATE INDEX ON "transfers" ("from_account_id", "to_account_id");

COMMENT ON COLUMN "entries"."amount" IS 'can be positive or negative';

COMMENT ON COLUMN "transfers"."amount" IS 'must be positive';

ALTER TABLE "accounts" ADD FOREIGN KEY ("owner_id") REFERENCES "users" ("id");

ALTER TABLE "entries" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("to_account_id") REFERENCES "accounts" ("id");

