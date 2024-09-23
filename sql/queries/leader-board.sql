-- name: AddUserToLeaderBoard :exec
INSERT INTO leader_board (id, best_time, user_id)
VALUES ($1, $2, $3);
