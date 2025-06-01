-- Drop the verify_email table if it exists (and remove dependent objects with CASCADE)
DROP TABLE IF EXISTS "verify_emails" CASCADE;

-- Drop the column is_user_verified from the users table
ALTER TABLE "users" DROP COLUMN IF EXISTS "is_user_verified";