-- name: GetClass :one
SELECT *
FROM "class"
WHERE "id" = $1
LIMIT 1;
-- name: ListClasses :many
SELECT *
FROM "class";
-- name: CreateClass :one
INSERT INTO "class" ("grade", "section", "term_id", "program_id")
VALUES ($1, $2, $3, $4)
RETURNING *;
-- name: UpdateClass :one
UPDATE "class"
SET "grade" = $2,
    "section" = $3,
    "term_id" = $4,
    "program_id" = $5
WHERE "id" = $1
RETURNING *;
-- name: DeleteClass :exec
DELETE FROM "class"
WHERE "id" = $1;
