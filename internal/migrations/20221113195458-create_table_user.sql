
-- +migrate Up
CREATE TABLE "user" (
  "id" varchar PRIMARY KEY NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL,
  "deleted_at" timestamp
);
CREATE INDEX user_id_idx ON "user" USING btree (id);
CREATE INDEX user_email_idx ON "user" USING btree (email);

-- +migrate Down
DROP TABLE "user";
