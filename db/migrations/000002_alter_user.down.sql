
ALTER TABLE "users"
DROP COLUMN IF EXISTS "hashed_password" CASCADE,
DROP COLUMN IF EXISTS "password_changed_at" CASCADE;

ALTER TABLE IF EXISTS "accounts" 
  DROP CONSTRAINT IF EXISTS "user_currency_key";
