-- name: CreateDivision :one
INSERT INTO divisions (name, generation_id, type)
	VALUES (@NAME::varchar(255), (
			SELECT
				g.id
			FROM
				GET_ACTIVE_GENERATION() g
			LIMIT 1),
		@type)
RETURNING
	*;

-- name: DeleteDivision :execrows
DELETE FROM divisions
WHERE id = @division_id;

-- name: UpdateDivision :one
UPDATE
	divisions
SET
	name = @NAME::varchar(255)
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

-- name: ListAllDivisions :many
SELECT
	*
FROM
	divisions
LIMIT $1 OFFSET $2;

-- name: ListActiveDivisions :many
SELECT
	*
FROM
	divisions
WHERE
	generation_id = (
		SELECT
			g.id
		FROM
			GET_ACTIVE_GENERATION() g)
LIMIT $1 OFFSET $2;
