-- name: CreateRole :one
INSERT INTO roles (name)
	VALUES (@NAME::varchar(255))
RETURNING
	*;

-- name: UpdateRole :one
UPDATE
	roles
SET
	name = @name
WHERE
	id = @roleId
RETURNING
	*;

-- name: GetRole :one
SELECT
	*
FROM
	roles
WHERE
	id = @roleId;

-- name: ListRole :many
SELECT
	*
FROM
	roles
LIMIT $1 OFFSET $2;

-- name: DeleteRole :execrows
DELETE FROM roles
WHERE id = @roleId;
