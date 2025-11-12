CREATE TYPE role AS ENUM (
    'administrator',
    'teacher',
    'secretary'
);
CREATE TABLE "class" (
    "id" BIGSERIAL PRIMARY KEY,
    "name" TEXT NOT NULL,
    "term_id" BIGINT NOT NULL,
    "program_id" BIGINT NOT NULL
);
CREATE TABLE "lesson" (
    "id" BIGSERIAL PRIMARY KEY,
    "class_id" BIGINT NOT NULL,
    "teacher_id" BIGINT NOT NULL,
    "date" DATE NULL,
    "start_time" TIME(0) WITHOUT TIME ZONE NOT NULL,
    "end_time" TIME(0) WITHOUT TIME ZONE NOT NULL
);
CREATE TABLE "program" (
    "id" BIGSERIAL PRIMARY KEY,
    "name" TEXT NOT NULL
);
CREATE TABLE "subject" (
    "id" BIGSERIAL PRIMARY KEY,
    "name" TEXT NOT NULL,
    "color_id" BIGINT NOT NULL,
    "teacher_id" BIGINT NOT NULL
);
CREATE TABLE "subject_color" (
    "id" BIGSERIAL PRIMARY KEY,
    "hex_value" bytea NOT NULL
);
CREATE TABLE "teacher" (
    "id" BIGSERIAL PRIMARY KEY,
    "preferred_program_id" BIGINT NOT NULL,
    "user_id" BIGINT NOT NULL,
    "subject_id" BIGINT NOT NULL
);
CREATE TABLE "term" (
    "id" BIGSERIAL PRIMARY KEY,
    "start_date" DATE NOT NULL,
    "end_date" DATE NOT NULL
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
ALTER TABLE "teacher"
ADD CONSTRAINT "teacher_preferred_program_id_foreign" FOREIGN KEY("preferred_program_id") REFERENCES "program"("id");
ALTER TABLE "class"
ADD CONSTRAINT "class_term_id_foreign" FOREIGN KEY("term_id") REFERENCES "term"("id");
ALTER TABLE "class"
ADD CONSTRAINT "class_program_id_foreign" FOREIGN KEY("program_id") REFERENCES "program"("id");
ALTER TABLE "lesson"
ADD CONSTRAINT "lesson_class_id_foreign" FOREIGN KEY("class_id") REFERENCES "class"("id");
ALTER TABLE "teacher"
ADD CONSTRAINT "teacher_user_id_foreign" FOREIGN KEY("user_id") REFERENCES "user"("id");
ALTER TABLE "teacher"
ADD CONSTRAINT "teacher_subject_id_foreign" FOREIGN KEY("subject_id") REFERENCES "subject"("id");
ALTER TABLE "lesson"
ADD CONSTRAINT "lesson_teacher_id_foreign" FOREIGN KEY("teacher_id") REFERENCES "teacher"("id");
