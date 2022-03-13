CREATE TABLE
  "members" (
    "id" uuid PRIMARY KEY DEFAULT uuid_generate_v4 (),
    "full_name" VARCHAR(255) NOT NULL,
    "university" VARCHAR(255) NOT NULL DEFAULT 'Sriwijaya University',
    "role_id" uuid NOT NULL,
    "division_id" uuid,
    "picture_path" VARCHAR(255),
    "created_at" TIMESTAMP NOT NULL DEFAULT (now()),
    "deleted_at" TIMESTAMP DEFAULT (now())
  );