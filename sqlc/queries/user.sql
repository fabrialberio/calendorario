-- name: GetUser :one
SELECT *
FROM "user"
WHERE "id" = $1
LIMIT 1;
-- name: GetUserWithUsername :one
SELECT *
FROM "user"
WHERE "username" = $1
LIMIT 1;
-- name: ListUsers :many
SELECT *
FROM "user";
-- name: CreateUser :one
INSERT INTO "user" (
        "username",
        "name",
        "surname",
        "password_hash",
        "role",
        "created_at",
        "deleted"
    )
VALUES (
        $1,
        $2,
        $3,
        $4,
        $5,
        CURRENT_TIMESTAMP,
        false
    )
RETURNING *;
-- name: UpdateUser :one
UPDATE "user"
SET "username" = $2,
    "name" = $3,
    "surname" = $4,
    "password_hash" = $5,
    "role" = $6,
    "deleted" = $7
WHERE "id" = $1
RETURNING *;
-- name: DeleteAuthor :exec
DELETE FROM "user"
WHERE "id" = $1;
