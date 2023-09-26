-- name: CreateUser :one
INSERT INTO "user" ("name", "email", "password") VALUES ($1, $2, $3) RETURNING *;

-- name: Fetch :many
SELECT * FROM "user";

-- name: GetByEmail :one
SELECT * FROM "user" WHERE "email" = $1;

-- name: GetByID :one
SELECT * FROM "user" WHERE "id" = $1;