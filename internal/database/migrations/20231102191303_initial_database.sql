-- Create "user" table
CREATE TABLE "user" ("id" uuid NOT NULL, "email" character varying NOT NULL, "password_hash" character varying NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, PRIMARY KEY ("id"));
-- Create index "user_email_key" to table: "user"
CREATE UNIQUE INDEX "user_email_key" ON "user" ("email");
-- Create "dish" table
CREATE TABLE "dish" ("id" uuid NOT NULL, "name_de" character varying NOT NULL, "name_en" character varying NULL, PRIMARY KEY ("id"));
-- Create index "dish_name_de_key" to table: "dish"
CREATE UNIQUE INDEX "dish_name_de_key" ON "dish" ("name_de");
-- Create "dish_alias" table
CREATE TABLE "dish_alias" ("alias_name" character varying NOT NULL, "normalized_alias_name" character varying NOT NULL, "dish" uuid NOT NULL, PRIMARY KEY ("alias_name"), CONSTRAINT "dish_alias_dish_aliases" FOREIGN KEY ("dish") REFERENCES "dish" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
-- Create "location" table
CREATE TABLE "location" ("id" uuid NOT NULL, "external_id" bigint NOT NULL, "name" character varying NOT NULL, "visible" boolean NOT NULL DEFAULT false, PRIMARY KEY ("id"));
-- Create index "location_external_id_key" to table: "location"
CREATE UNIQUE INDEX "location_external_id_key" ON "location" ("external_id");
-- Create index "location_name_key" to table: "location"
CREATE UNIQUE INDEX "location_name_key" ON "location" ("name");
-- Create "occurrence" table
CREATE TABLE "occurrence" ("id" uuid NOT NULL, "date" timestamptz NOT NULL, "kj" bigint NULL, "kcal" bigint NULL, "fat" bigint NULL, "saturated_fat" bigint NULL, "carbohydrates" bigint NULL, "sugar" bigint NULL, "fiber" bigint NULL, "protein" bigint NULL, "salt" bigint NULL, "price_student" bigint NULL, "price_staff" bigint NULL, "price_guest" bigint NULL, "not_available_after" timestamptz NULL, "dish" uuid NOT NULL, "location" uuid NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "occurrence_dish_dish_occurrences" FOREIGN KEY ("dish") REFERENCES "dish" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, CONSTRAINT "occurrence_location_occurrences" FOREIGN KEY ("location") REFERENCES "location" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
-- Create "review" table
CREATE TABLE "review" ("id" uuid NOT NULL, "display_name" character varying NULL, "stars" bigint NOT NULL, "text" character varying NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "accepted_at" timestamptz NULL, "occurrence" uuid NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "review_occurrence_reviews" FOREIGN KEY ("occurrence") REFERENCES "occurrence" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
-- Create "image" table
CREATE TABLE "image" ("id" uuid NOT NULL, "image_hash" character varying NOT NULL, "review" uuid NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "image_review_images" FOREIGN KEY ("review") REFERENCES "review" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);
-- Create "occurrence_side_dishes" table
CREATE TABLE "occurrence_side_dishes" ("occurrence" uuid NOT NULL, "dish" uuid NOT NULL, PRIMARY KEY ("occurrence", "dish"), CONSTRAINT "occurrence_side_dishes_dish" FOREIGN KEY ("dish") REFERENCES "dish" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT "occurrence_side_dishes_occurrence" FOREIGN KEY ("occurrence") REFERENCES "occurrence" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);
-- Create "tag" table
CREATE TABLE "tag" ("key" character varying NOT NULL, "name" character varying NOT NULL, "description" character varying NOT NULL, "short_name" character varying NULL, "priority" character varying NOT NULL DEFAULT 'HIDE', "is_allergy" boolean NOT NULL DEFAULT false, PRIMARY KEY ("key"));
-- Create "occurrence_tags" table
CREATE TABLE "occurrence_tags" ("occurrence" uuid NOT NULL, "tag" character varying NOT NULL, PRIMARY KEY ("occurrence", "tag"), CONSTRAINT "occurrence_tags_occurrence" FOREIGN KEY ("occurrence") REFERENCES "occurrence" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT "occurrence_tags_tag" FOREIGN KEY ("tag") REFERENCES "tag" ("key") ON UPDATE NO ACTION ON DELETE CASCADE);
