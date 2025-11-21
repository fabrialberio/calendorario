-- name: GetProgram :one
SELECT *
FROM "program"
WHERE "id" = $1
LIMIT 1;
-- name: ListPrograms :many
SELECT *
FROM "program";
-- name: CreateProgram :one
INSERT INTO "program" ("name")
VALUES ($1)
RETURNING *;
-- name: UpdateProgram :one
UPDATE "program"
SET "name" = $2
WHERE "id" = $1
RETURNING *;
-- name: DeleteProgram :exec
DELETE FROM "program"
WHERE "id" = $1;
