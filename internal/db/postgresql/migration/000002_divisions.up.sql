CREATE TABLE "divisions" (
	"id" uuid PRIMARY KEY DEFAULT uuid_generate_v4 (),
	"name" varchar(255) NOT NULL,
	"created_at" timestamp NOT NULL DEFAULT (NOW())
);
