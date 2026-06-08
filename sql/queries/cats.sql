-- name: GetCats :many
SELECT * FROM cats;

-- name: GetCat :one
SELECT * FROM cats WHERE id = ? LIMIT 1;

-- name: CreateCat :one
INSERT INTO	cats (name, age) VALUES (?, ?) RETURNING *; 

-- name: UpdateCat :exec
UPDATE cats SET name = ?, age= ? 	WHERE id = ?;

-- name: DeleteCat :exec
DELETE FROM cats WHERE id = ?;
