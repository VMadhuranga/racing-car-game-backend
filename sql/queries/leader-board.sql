-- name: AddUserToLeaderBoard :exec
INSERT INTO leader_board (id, best_time, user_id)
VALUES ($1, $2, $3);

-- name: UpdateUserBestTimeByUserId :exec
UPDATE leader_board SET best_time = $1 WHERE user_id = $2;
