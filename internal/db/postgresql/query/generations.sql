-- name: CreateGeneration :one
INSERT INTO generations (year)
	VALUES (@year::varchar(4))
RETURNING
	*, FALSE AS active;

-- name: GetGeneration :one
SELECT
	g.*,
	CASE WHEN g.id = (
		SELECT
			s.related_id
		FROM
			settings s
		WHERE
			s.setting_type = 'active_generation'
		LIMIT 1) THEN
		TRUE
	ELSE
		FALSE
	END AS active
FROM
	generations g
WHERE
	id = $1
LIMIT 1;

-- name: ListGeneration :many
SELECT
	g.id,
	g.year,
	g.created_at,
	CASE WHEN a.id = g.id THEN
		TRUE
	ELSE
		FALSE
	END AS active
FROM
	generations g
	LEFT JOIN GET_ACTIVE_GENERATION() a ON g.id = a.id
LIMIT $1 offset $2;

-- name: UpdateGeneration :one
UPDATE
	generations
SET
	year = @year::varchar(4)
WHERE
	id = @generationId
RETURNING
	*,
	(
		CASE WHEN (
			SELECT
				id
			FROM
				GET_ACTIVE_GENERATION()) = id THEN
			TRUE
		ELSE
			FALSE
		END) AS ACTIVE;

-- name: DeleteGeneration :execrows
DELETE FROM generations
WHERE id = @id;

-- name: GetActiveGeneration :one
SELECT
	*,
	TRUE AS active
FROM
	GET_ACTIVE_GENERATION();

-- name: SetActiveGeneration :one
WITH activeGen AS (
	UPDATE
		settings s
	SET
		related_id = @generationId
	WHERE
		s.setting_type = 'active_generation'
	RETURNING
		*
)
SELECT
	g.*,
	TRUE AS active
FROM
	generations g,
	activeGen
WHERE
	g.id = activeGen.related_id;
