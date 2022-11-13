
-- +migrate Up
CREATE TABLE "tenant" (
  "id" varchar PRIMARY KEY NOT NULL,
  "owner_id" varchar,
  "business_name" varchar NOT NULL,
  "business_description" varchar,
  "business_logo" varchar,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL,
  "deleted_at" timestamp
);
ALTER TABLE "tenant" ADD FOREIGN KEY ("owner_id") REFERENCES "user" ("id");

-- +migrate Down
