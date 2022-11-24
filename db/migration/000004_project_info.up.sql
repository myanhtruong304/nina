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