CREATE TABLE "roles" (
  "id" uuid PRIMARY KEY DEFAULT uuid_generate_v4 (),
  "name" varchar(255) NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "divisions" (
  "id" uuid PRIMARY KEY DEFAULT uuid_generate_v4 (),
  "name" varchar(255) NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "core_teams" (
  "division_id" uuid UNIQUE NOT NULL,
  "member_id" uuid NOT NULL,
);

CREATE TABLE "members" (
  "id" uuid PRIMARY KEY DEFAULT uuid_generate_v4 (),
  "full_name" varchar(255) NOT NULL,
  "university" varchar(255) NOT NULL DEFAULT 'Sriwijaya University',
  "role_id" uuid NOT NULL,
  "division_id" uuid,
  "picture_path" varchar(255),
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "deleted_at" timestamp DEFAULT (now())
);

CREATE TABLE "messages" (
  "id" uuid PRIMARY KEY DEFAULT uuid_generate_v4 (),
  "full_name" varchar(255),
  "phone_number" varchar(18),
  "question" TEXT,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "deleted_at" timestamp DEFAULT (now())
);

ALTER TABLE "divisions" ADD FOREIGN KEY ("id") REFERENCES "core_teams" ("division_id");

ALTER TABLE "members" ADD FOREIGN KEY ("id") REFERENCES "core_teams" ("member_id");

ALTER TABLE "members" ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("id");

ALTER TABLE "members" ADD FOREIGN KEY ("division_id") REFERENCES "divisions" ("id");

CREATE INDEX ON "core_teams" ("division_id");

CREATE INDEX ON "core_teams" ("member_id");

CREATE INDEX ON "members" ("full_name");

CREATE INDEX ON "messages" ("full_name");
