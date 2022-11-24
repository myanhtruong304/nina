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