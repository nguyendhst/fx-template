-- Create "user" table
CREATE TABLE "public"."user" ("id" serial NOT NULL, "name" text NOT NULL, "email" text NOT NULL, "password" text NOT NULL, "created_at" timestamptz NOT NULL DEFAULT now(), "updated_at" timestamptz NOT NULL DEFAULT now(), PRIMARY KEY ("id"));
-- Create index "user_email_key" to table: "user"
CREATE UNIQUE INDEX "user_email_key" ON "public"."user" ("email");
