CREATE TABLE "users" (
  "username" varchar PRIMARY KEY,
  "hashed_pwd" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "pwd_changed_at" timestamp NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "project" (
  "project_name" varchar PRIMARY KEY NOT NULL,
  "symbol" varchar NOT NULL,
  "contract_address" varchar NOT NULL,
  "owner" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()) 
);

CREATE INDEX ON "project" ("owner");

-- CREATE UNIQUE INDEX ON "project" ("owner", "project_name");

ALTER TABLE "project" ADD CONSTRAINT "project_owner_key" UNIQUE ("owner", "project_name");

ALTER TABLE "project" ADD FOREIGN KEY ("owner") REFERENCES "users" ("username");
