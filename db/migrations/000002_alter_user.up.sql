
ALTER TABLE "users" 
ADD COLUMN IF NOT EXISTS "hashed_password" VARCHAR NOT NULL, 
ADD COLUMN IF NOT EXISTS "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z';

ALTER TABLE "accounts" ADD CONSTRAINT "user_currency_key" UNIQUE ("id", "currency");
