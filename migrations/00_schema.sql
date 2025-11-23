CREATE TYPE role AS ENUM (
    'administrator',
    'teacher',
    'secretary'
);
CREATE TABLE "term" (
    "id" BIGSERIAL PRIMARY KEY,
    "name" TEXT NOT NULL,
    "start_date" DATE NOT NULL,
    "end_date" DATE NOT NULL
);
CREATE TABLE "program" (
    "id" BIGSERIAL PRIMARY KEY,
    "name" TEXT NOT NULL
);
CREATE TABLE "user" (
    "id" BIGSERIAL PRIMARY KEY,
    "username" text NOT NULL,
    "name" text NOT NULL,
    "surname" text NOT NULL,
    "password_hash" bytea,
    "role" role NOT NULL,
    "created_at" TIMESTAMP NOT NULL,
    "deleted" BOOLEAN NOT NULL
);
CREATE TABLE "vacation" (
    "id" BIGSERIAL PRIMARY KEY,
    "name" TEXT NOT NULL,
    "start_date" DATE NOT NULL,
    "end_date" DATE NOT NULL,
    "term_id" BIGINT NOT NULL REFERENCES "term"
);
CREATE TABLE "class" (
    "id" BIGSERIAL PRIMARY KEY,
    "grade" INT NOT NULL,
    "section" TEXT NOT NULL,
    "term_id" BIGINT NOT NULL REFERENCES "term",
    "program_id" BIGINT NOT NULL REFERENCES "program"
);
CREATE TABLE "subject" (
    "id" BIGSERIAL PRIMARY KEY,
    "name" TEXT NOT NULL,
    "color_hex_value" bytea NOT NULL
);
CREATE TABLE "teacher" (
    "user_id" BIGINT PRIMARY KEY REFERENCES "user",
    "contract_start_date" DATE NOT NULL,
    "contract_end_date" DATE NOT NULL
);
CREATE TABLE "lesson" (
    "id" BIGSERIAL PRIMARY KEY,
    "class_id" BIGINT NOT NULL REFERENCES "class",
    "teacher_id" BIGINT NOT NULL REFERENCES "teacher",
    "start_time" TIME(0) WITHOUT TIME ZONE NOT NULL,
    "end_time" TIME(0) WITHOUT TIME ZONE NOT NULL,
    "first_date" DATE NOT NULL,
    "last_date" DATE NOT NULL
);
CREATE TABLE "teacher_subject_assignment" (
    "term_id" BIGINT NOT NULL REFERENCES "term",
    "teacher_id" BIGINT NOT NULL REFERENCES "teacher",
    "subject_id" BIGINT NOT NULL REFERENCES "subject",
    PRIMARY KEY ("term_id", "teacher_id", "subject_id")
)
