-- name: GetVacation :one
SELECT *
FROM "vacation"
WHERE "id" = $1
LIMIT 1;
-- name: ListVacations :many
SELECT *
FROM "vacation";
-- name: CreateVacation :one
INSERT INTO "vacation" ("name", "start_date", "end_date")
VALUES ($1, $2, $3)
RETURNING *;
-- name: UpdateVacation :one
UPDATE "vacation"
SET "name" = $2,
    "start_date" = $3,
    "end_date" = $4
WHERE "id" = $1
RETURNING *;
-- name: DeleteVacation :exec
DELETE FROM "vacation"
WHERE "id" = $1;
