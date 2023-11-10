-- Modify "user" table
ALTER TABLE "user"
    ADD COLUMN "username" character varying NOT NULL DEFAULT SUBSTRING(gen_random_uuid()::text, 1, 16),
    ADD COLUMN "role" character varying NULL;
-- Create index "user_username_key" to table: "user"
CREATE UNIQUE INDEX "user_username_key" ON "user" ("username");
