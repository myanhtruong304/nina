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

CREATE TABLE "image_content" (
  "id" int PRIMARY KEY,
  "project_name" varchar NOT NULL,
  "image_content" varchar NOT NULL,
  "content_id" int NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL,
  "link" varchar UNIQUE
);

CREATE TABLE "projects_info" (
  "project_name" varchar PRIMARY KEY NOT NULL,
  "symbol" varchar,
  "contract_address" varchar,
  "explorer" varchar,
  "twitter" varchar,
  "facebook" varchar,
  "linkedin" varchar,
  "medium" varchar,
  "telegram" varchar,
  "website" varchar,
  "git" varchar,
  "cmc" varchar,
  "coingecko" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "tags" (
  "tag" varchar UNIQUE NOT NULL,
  "project_name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "twitter_bind_account" (
  "id" int PRIMARY KEY,
  "project_name" varchar NOT NULL,
  "access_token" varchar UNIQUE NOT NULL,
  "access_token_secret" varchar UNIQUE NOT NULL
);

CREATE INDEX ON "content" ("project_name");

CREATE INDEX ON "content" ("content_type");

CREATE INDEX ON "content" ("upload");

CREATE INDEX ON "image_content" ("project_name");

CREATE INDEX ON "projects_info" ("project_name");

ALTER TABLE "content" ADD FOREIGN KEY ("image_link") REFERENCES "image_content" ("link");

ALTER TABLE "content" ADD FOREIGN KEY ("project_name") REFERENCES "projects_info" ("project_name");

ALTER TABLE "image_content" ADD FOREIGN KEY ("project_name") REFERENCES "projects_info" ("project_name");

ALTER TABLE "image_content" ADD FOREIGN KEY ("content_id") REFERENCES "content" ("id");

ALTER TABLE "twitter_bind_account" ADD FOREIGN KEY ("project_name") REFERENCES "projects_info" ("project_name");

ALTER TABLE "tags" ADD FOREIGN KEY ("project_name") REFERENCES "projects_info" ("project_name");