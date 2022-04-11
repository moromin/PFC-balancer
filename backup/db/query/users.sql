-- name: ListUsers :many
SELECT * FROM users
ORDER BY id;

-- name: Store :one
INSERT INTO users (
    nick_name,
    email,
    hashed_password
) VALUES (
    $1, $2, $3
) RETURNING *;
