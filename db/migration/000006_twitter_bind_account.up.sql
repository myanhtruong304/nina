CREATE TABLE "twitter_bind_account" (
  "id" int PRIMARY KEY,
  "project_name" varchar NOT NULL,
  "access_token" varchar UNIQUE NOT NULL,
  "access_token_secret" varchar UNIQUE NOT NULL
);