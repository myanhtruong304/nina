CREATE TABLE "users" (
  "username" varchar PRIMARY KEY NOT NULL,
  "hashed_pwd" varchar NOT NULL,
  "user_email" varchar UNIQUE NOT NULL,
  "pwd_changed_at" timestamp DEFAULT '0001-01-01 00:00:00Z' NOT NULL,
  "created_at" timestamp DEFAULT (now()) NOT NULL
);

CREATE TABLE "projects" (
  "project_name" varchar PRIMARY KEY NOT NULL,
  "symbol" varchar,
  "contract_address" varchar,
  "explorer" varchar,
  "website" varchar,
  "twitter" varchar,
  "facebook" varchar,
  "linkedin" varchar,
  "telegram" varchar,
  "medium" varchar,
  "whitepaper" varchar,
  "email" varchar,
  "owner" varchar NOT NULL,
  "created_at" timestamp DEFAULT (now()) NOT NULL
);

CREATE TABLE "contents" (
  "id" varchar UNIQUE NOT NULL,
  "content" varchar NOT NULL,
  "word_count" varchar NOT NULL,
  "schedule_time" varchar,
  "facebook_check" varchar,
  "twitter_check" varchar,
  "linkedin_check" varchar,
  "content_type" varchar,
  "project_name" varchar NOT NULL,
  "image_text" varchar,
  "image_link" varchar,
  "created_at" timestamp DEFAULT (now())  NOT NULL,
  "last_updated_at" timestamp DEFAULT (now())  NOT NULL
);

CREATE INDEX ON "contents" ("id");

CREATE INDEX ON "projects" ("owner");

ALTER TABLE "contents" ADD CONSTRAINT "project_content_key" UNIQUE ("content", "project_name");

ALTER TABLE "contents" ADD FOREIGN KEY ("project_name") REFERENCES "projects" ("project_name");

ALTER TABLE "projects" ADD CONSTRAINT "project_owner_key" UNIQUE ("owner", "project_name");

ALTER TABLE "projects" ADD FOREIGN KEY ("owner") REFERENCES "users" ("username");
