
ALTER TABLE "users" 
ADD COLUMN IF NOT EXISTS "user_name" VARCHAR UNIQUE NOT NULL,
ADD COLUMN IF NOT EXISTS "hashed_password" VARCHAR NOT NULL, 
ADD COLUMN IF NOT EXISTS "password_changed_at" timestamptz NOT NULL DEFAULT (now());

ALTER TABLE "accounts" ADD CONSTRAINT "user_currency_key" UNIQUE ("id", "currency");
