-- name: GetUser :one
select * from users where name = $1 LIMIT 1;

-- name: GetUserById :one
select * from users where id = $1;