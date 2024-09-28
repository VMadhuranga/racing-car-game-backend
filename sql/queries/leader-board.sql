-- name: AddUserToLeaderBoard :exec
INSERT INTO leader_board (id, best_time, user_id)
VALUES ($1, $2, $3);

-- name: UpdateUserBestTimeByUserId :exec
UPDATE leader_board SET best_time = $1 WHERE user_id = $2;

-- name: GetLeaderBoard :many
SELECT l.id, l.best_time, u.username
FROM leader_board l
JOIN users u ON l.user_id = u.id
ORDER BY l.best_time DESC;
