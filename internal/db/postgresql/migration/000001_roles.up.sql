CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "roles" (
	"id" uuid PRIMARY KEY DEFAULT UUID_GENERATE_V4(),
	"name" varchar(255) NOT NULL,
	"created_at" timestamp NOT NULL DEFAULT (NOW())
);
