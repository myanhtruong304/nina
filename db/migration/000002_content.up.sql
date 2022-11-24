CREATE TABLE "content" (
  "id" int PRIMARY KEY,
  "project_name" varchar NOT NULL,
  "content" varchar NOT NULL,
  "char_count" int NOT NULL,
  "image_link" varchar,
  "image_id" int,
  "platform" varchar NOT NULL,
  "content_type" varchar NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "upload" boolean NOT NULL
);