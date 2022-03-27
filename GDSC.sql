CREATE TYPE "setting_type" AS ENUM (
	'active_generation'
);

CREATE TYPE "division_type" AS ENUM (
	'LEAD',
	'COLEAD',
	'DIVISION'
);

CREATE TABLE "roles" (
	"id" uuid PRIMARY KEY DEFAULT (UUID_GENERATE_V4()),
	"name" varchar(255) NOT NULL,
	"created_at" timestamp NOT NULL DEFAULT (NOW())
);

CREATE TABLE "divisions" (
	"id" uuid PRIMARY KEY DEFAULT (UUID_GENERATE_V4()),
	"name" varchar(255) NOT NULL,
	"generation_id" uuid NOT NULL,
	"type" division_type,
	"created_at" timestamp NOT NULL DEFAULT (NOW())
);

CREATE TABLE "settings" (
	"value" varchar(255),
	"related_id" uuid,
	"setting_type" setting_type UNIQUE NOT NULL
);

CREATE TABLE "generations" (
	"id" uuid PRIMARY KEY DEFAULT (UUID_GENERATE_V4()),
	"year" varchar(4) UNIQUE NOT NULL,
	"created_at" timestamp NOT NULL DEFAULT (NOW())
);

CREATE TABLE "core_teams" (
	"division_id" uuid PRIMARY KEY NOT NULL,
	"member_id" uuid NOT NULL
);

CREATE TABLE "members" (
	"id" uuid PRIMARY KEY DEFAULT (UUID_GENERATE_V4()),
	"full_name" varchar(255) NOT NULL,
	"university" varchar(255) NOT NULL DEFAULT 'Sriwijaya University',
	"role_id" uuid NOT NULL,
	"division_id" uuid NOT NULL,
	"picture_path" varchar(255),
	"created_at" timestamp NOT NULL DEFAULT (NOW()),
	"deleted_at" timestamp DEFAULT (NOW())
);

CREATE TABLE "messages" (
	"id" uuid PRIMARY KEY DEFAULT (UUID_GENERATE_V4()),
	"full_name" varchar(255),
	"phone_number" varchar(18),
	"question" text,
	"created_at" timestamp NOT NULL DEFAULT (NOW()),
	"deleted_at" timestamp DEFAULT (NOW())
);

ALTER TABLE "members"
	ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("id") ON DELETE
	RESTRICT ON UPDATE RESTRICT;

ALTER TABLE "members"
	ADD FOREIGN KEY ("division_id") REFERENCES "divisions" ("id") ON DELETE
	RESTRICT ON UPDATE RESTRICT;

ALTER TABLE "divisions"
	ADD FOREIGN KEY ("generation_id") REFERENCES "generations" ("id") ON
	DELETE RESTRICT ON UPDATE RESTRICT;

ALTER TABLE "core_teams"
	ADD FOREIGN KEY ("division_id") REFERENCES "divisions" ("id") ON DELETE
	RESTRICT ON UPDATE RESTRICT;

ALTER TABLE "core_teams"
	ADD FOREIGN KEY ("member_id") REFERENCES "members" ("id") ON DELETE
	RESTRICT ON UPDATE RESTRICT;

CREATE INDEX ON "settings" ("setting_type");

CREATE INDEX ON "generations" ("year");

CREATE INDEX ON "core_teams" ("division_id");

CREATE INDEX ON "core_teams" ("member_id");

CREATE INDEX ON "members" ("full_name");

CREATE INDEX ON "messages" ("full_name");
