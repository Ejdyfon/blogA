-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: ListFeeds :many
select * from feeds;

-- name: GetFeedByUrl :one
select * from feeds where url = $1 LIMIT 1;

-- name: GetFeedFollowsForUser :many
select    feed_follows.*,
    feeds.name AS feed_name,
    users.name AS user_name
from feed_follows
INNER JOIN feeds on feed_follows.feed_id = feeds.id
INNER JOIN users on  feed_follows.user_id = users.id
where users.name = $1;

-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
)
RETURNING *)
SELECT
    inserted_feed_follow.*,
    feeds.name AS feed_name,
    users.name AS user_name
FROM inserted_feed_follow 
INNER JOIN feeds on inserted_feed_follow.feed_id = feeds.id
INNER JOIN users on  inserted_feed_follow.user_id = users.id;

-- name: DeleteFeedFollowByUserAndFeed :exec
delete from feed_follows where user_id=$1 and feed_id=$2;