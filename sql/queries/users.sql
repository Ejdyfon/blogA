-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name)
VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING *;

-- name: GetAllUsers :many
select * from users;

-- name: DeleteAllUsers :exec
delete from users;

-- name: GetUser :one
select * from users where name = $1 LIMIT 1;

-- name: GetUserById :one
select * from users where id = $1;