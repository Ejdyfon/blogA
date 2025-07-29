-- name: GetUser :one
select * from users where name = $1 LIMIT 1;