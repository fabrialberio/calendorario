-- name: GetSubject :one
SELECT *
FROM "subject"
WHERE "id" = $1
LIMIT 1;
-- name: ListSubjects :many
SELECT *
FROM "subject";
-- name: CreateSubject :one
INSERT INTO "subject" ("name", "color_hex_value")
VALUES ($1, $2)
RETURNING *;
-- name: UpdateSubject :one
UPDATE "subject"
SET "name" = $2,
    "color_hex_value" = $3
WHERE "id" = $1
RETURNING *;
-- name: DeleteSubject :exec
DELETE FROM "subject"
WHERE "id" = $1;
