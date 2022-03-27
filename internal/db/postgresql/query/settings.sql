-- name: CreateSetting :one
INSERT INTO settings (value, related_id, setting_type)
	VALUES (STRING_NULL_OR_TRUNCATE(@value), @relatedId, @settingType)
RETURNING
	*;

-- name: GetSetting :one
SELECT
	*
FROM
	settings
WHERE
	setting_type = @settingType;

-- name: UpdateSetting :one
UPDATE
	settings
SET
	value = STRING_NULL_OR_TRUNCATE(@value),
	related_id = @relatedId,
	setting_type = @settingType
RETURNING
	*;
