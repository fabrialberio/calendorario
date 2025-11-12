CREATE TYPE role AS ENUM (
    'administrator',
    'teacher',
    'secretary'
);
CREATE TABLE "term" (
    "id" BIGSERIAL PRIMARY KEY,
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
    "start_datetime" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "end_datetime" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);
CREATE TABLE "subject_color" (
    "id" BIGSERIAL PRIMARY KEY,
    "hex_value" bytea NOT NULL
);
CREATE TABLE "subject" (
    "id" BIGSERIAL PRIMARY KEY,
    "name" TEXT NOT NULL,
    "color_id" BIGINT NOT NULL REFERENCES "subject_color"
);
CREATE TABLE "class" (
    "id" BIGSERIAL PRIMARY KEY,
    "name" TEXT NOT NULL,
    "term_id" BIGINT NOT NULL REFERENCES "term",
    "program_id" BIGINT NOT NULL REFERENCES "program"
);
CREATE TABLE "teacher" (
    "id" BIGSERIAL PRIMARY KEY,
    "preferred_program_id" BIGINT NOT NULL REFERENCES "program",
    "user_id" BIGINT NOT NULL REFERENCES "user",
    "subject_id" BIGINT NOT NULL REFERENCES "subject"
);
CREATE TABLE "lesson" (
    "id" BIGSERIAL PRIMARY KEY,
    "class_id" BIGINT NOT NULL REFERENCES "class",
    "teacher_id" BIGINT NOT NULL REFERENCES "teacher",
    "date" DATE NULL,
    "start_time" TIME(0) WITHOUT TIME ZONE NOT NULL,
    "end_time" TIME(0) WITHOUT TIME ZONE NOT NULL
);
