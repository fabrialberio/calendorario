-- name: GetTerm :one
SELECT *
FROM "term"
WHERE "id" = $1
LIMIT 1;
-- name: ListTerms :many
SELECT *
FROM "term";
-- name: CreateTerm :one
INSERT INTO "term" ("name", "start_date", "end_date")
VALUES ($1, $2, $3)
RETURNING *;
-- name: UpdateTerm :one
UPDATE "term"
SET "name" = $2,
    "start_date" = $3,
    "end_date" = $4
WHERE "id" = $1
RETURNING *;
-- name: DeleteTerm :exec
DELETE FROM "term"
WHERE "id" = $1;
