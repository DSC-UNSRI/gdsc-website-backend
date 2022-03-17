-- name: CreateDivision :one
INSERT INTO divisions (name)
	VALUES (@name::varchar(255))
RETURNING
	*;

-- name: DeleteDivision :execrows
DELETE FROM divisions
WHERE id = @division_id;

-- name: UpdateDivision :one
UPDATE
	divisions
SET
	name = @name::varchar(255)
WHERE
	id = @divisionId
RETURNING
	*;

-- name: GetDivision :one
SELECT
	*
FROM
	divisions
WHERE
	id = @divisionId
LIMIT 1;

-- name: ListDivisions :many
SELECT
	*
FROM
	divisions
LIMIT $1 OFFSET $2;
