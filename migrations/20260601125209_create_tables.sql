-- +goose Up
CREATE TABLE IF NOT EXISTS "department"(
"id" BIGSERIAL PRIMARY KEY,
"name" TEXT NOT NULL,
"parent_id" BIGINT,
"created_at" TIMESTAMP NOT NULL DEFAULT now(),
FOREIGN KEY ("parent_id") REFERENCES "department"("id") ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS "employee"(
"id" BIGSERIAL PRIMARY KEY,
"department_id" BIGINT NOT NULL,
"full_name" TEXT NOT NULL,
"position" TEXT NOT NULL,
"hired_at" DATE,
"created_at" TIMESTAMP NOT NULL DEFAULT now(),
FOREIGN KEY ("department_id") REFERENCES "department"("id") ON DELETE CASCADE
);
-- +goose Down
DROP TABLE employee;
DROP TABLE department;
 