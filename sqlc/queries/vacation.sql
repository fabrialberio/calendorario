-- name: GetVacation :one
SELECT *
FROM "vacation"
WHERE "id" = $1
LIMIT 1;
-- name: ListVacations :many
SELECT *
FROM "vacation";
-- name: ListVacationWithTermID :many
SELECT *
FROM "vacation"
WHERE "term_id" = $1;
-- name: CreateVacation :one
INSERT INTO "vacation" ("name", "start_date", "end_date", "term_id")
VALUES ($1, $2, $3, $4)
RETURNING *;
-- name: UpdateVacation :one
UPDATE "vacation"
SET "name" = $2,
    "start_date" = $3,
    "end_date" = $4,
    "term_id" = $5
WHERE "id" = $1
RETURNING *;
-- name: DeleteVacation :exec
DELETE FROM "vacation"
WHERE "id" = $1;
