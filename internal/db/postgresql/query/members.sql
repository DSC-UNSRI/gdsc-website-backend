-- name: CreateMember :one
INSERT INTO MEMBERS (full_name, university, role_id, division_id, picture_path)
	VALUES ($1, $2, $3, $4, $5)
RETURNING
	*;

-- name: ListMembers :many
SELECT
	*
FROM
	members
LIMIT $1 OFFSET $2;

-- name: GetMember :one
SELECT
	*
FROM
	members
WHERE
	id = @memberId;

-- name: UpdateMember :one
UPDATE
	members
SET
	full_name = @name,
	university = @university,
	role_id = @roleId,
	division_id = @divisionId,
	picture_path = @picturePath
RETURNING
	*;

-- name: DeleteMember :exec
DELETE FROM members
WHERE id = @memberId;
