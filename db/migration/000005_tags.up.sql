CREATE TABLE "tags" (
  "tag" varchar UNIQUE NOT NULL,
  "project_name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);