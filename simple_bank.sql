CREATE TYPE "Currency" AS ENUM (
  'USD',
  'EUR',
  'MYR',
  'SGD'
);

CREATE TABLE "user" (
  "id" bigserial PRIMARY KEY,
  "first_name" varchar,
  "last_name" varchar,
  "email" varchar UNIQUE NOT NULL,
  "created_at" timestampz DEFAULT (now()),
  "deleted_at" timestampz
);

CREATE TABLE "accounts" (
  "id" bigserial PRIMARY KEY,
  "owner_id" bigint NOT NULL,
  "balance" bigint NOT NULL,
  "currency" Currency NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  "deleted_at" timestampz
);

CREATE TABLE "entries" (
  "id" bigserial PRIMARY KEY,
  "account_id" bigint,
  "amount" bigint NOT NULL,
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "transfer" (
  "id" bigserial PRIMARY KEY,
  "from_account_id" bigint NOT NULL,
  "to_account_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestampz DEFAULT (now())
);

CREATE INDEX ON "entries" ("account_id");

CREATE INDEX ON "transfer" ("from_account_id");

CREATE INDEX ON "transfer" ("to_account_id");

CREATE INDEX ON "transfer" ("from_account_id", "to_account_id");

COMMENT ON COLUMN "entries"."amount" IS 'can be positive or negative';

COMMENT ON COLUMN "transfer"."amount" IS 'must be positive';

ALTER TABLE "accounts" ADD FOREIGN KEY ("owner_id") REFERENCES "user" ("id");

ALTER TABLE "entries" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfer" ADD FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfer" ADD FOREIGN KEY ("to_account_id") REFERENCES "accounts" ("id");
