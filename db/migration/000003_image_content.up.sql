CREATE TABLE "image_content" (
  "id" int PRIMARY KEY,
  "project_name" varchar NOT NULL,
  "image_content" varchar NOT NULL,
  "content_id" int NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL,
  "link" varchar UNIQUE
);