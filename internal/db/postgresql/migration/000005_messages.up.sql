CREATE TABLE "messages" (
  "id" uuid PRIMARY KEY DEFAULT uuid_generate_v4 (),
  "full_name" varchar(255) NOT NULL,
  "phone_number" varchar(18) NOT NULL,
  "question" TEXT NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "deleted_at" timestamp DEFAULT (now())
);