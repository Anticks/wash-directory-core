-- Create "users" table
CREATE TABLE "users" ("id" uuid NOT NULL, "first_name" character varying NOT NULL, "last_name" character varying NOT NULL, "email" character varying NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, PRIMARY KEY ("id"));
-- Create index "users_email_key" to table: "users"
CREATE UNIQUE INDEX "users_email_key" ON "users" ("email");
-- Create "washer_profiles" table
CREATE TABLE "washer_profiles" ("id" uuid NOT NULL, "service_details" character varying NULL, "availability" character varying NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "user_washer_profile" uuid NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "washer_profiles_users_washer_profile" FOREIGN KEY ("user_washer_profile") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
-- Create index "washer_profiles_user_washer_profile_key" to table: "washer_profiles"
CREATE UNIQUE INDEX "washer_profiles_user_washer_profile_key" ON "washer_profiles" ("user_washer_profile");
