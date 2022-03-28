package postgresql

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/stretchr/testify/require"
)

var isAlreadyExecutedSettings bool

func setupSettingsTable(t *testing.T, relatedId *uuid.UUID) {
	if isAlreadyExecutedSettings {
		return
	}

	setting, err := querier.GetSetting(context.Background(), SettingTypeActiveGeneration)
	if string(setting.SettingType) != "" {
		isAlreadyExecutedSettings = true
		return
	}

	require.ErrorIs(t, pgx.ErrNoRows, err)

	var params CreateSettingParams
	params.Settingtype = SettingTypeActiveGeneration
	if relatedId != nil {
		var uuid uuid.NullUUID
		uuid.Scan(relatedId.String())
		params.Relatedid = uuid
	}

	createSettingHelper(t, true, &params)
	isAlreadyExecutedSettings = true
}

func TestCreateSetting(t *testing.T) {
	deleteSettingsHelper(t)
	createSettingHelper(t, true, nil)
}

func createSettingHelper(t *testing.T, validateSetting bool, arg *CreateSettingParams) (Setting, error) {
	var err error

	var param CreateSettingParams
	if arg != nil {
		param = *arg
	} else {
		param.Value = ""
		param.Relatedid = uuid.NullUUID{}
		param.Settingtype = SettingTypeActiveGeneration
	}

	setting, err := querier.CreateSetting(context.Background(), param)
	if validateSetting {
		require.NoError(t, err)
		if arg != nil {
			if arg.Relatedid.Valid {
				require.True(t, setting.RelatedID.Valid)
			}
			if arg.Value != "" {
				require.True(t, setting.Value.Valid)
			}
		} else {
			require.False(t, setting.Value.Valid)
			require.False(t, setting.RelatedID.Valid)
		}
	}
	return setting, err
}

func deleteSettingsHelper(t *testing.T) {
	_, err := querier.db.Exec(context.Background(), "DELETE FROM settings")
	require.NoError(t, err)
}

func TestUpdateSetting(t *testing.T) {
	deleteSettingsHelper(t)
	setting, _ := createSettingHelper(t, true, nil)
	finalYear := "1999"

	gen, _ := createGenerationHelper(t, &finalYear)
	var nullUuid uuid.NullUUID
	err := nullUuid.Scan(gen.ID.String())
	require.NoError(t, err)
	require.True(t, nullUuid.Valid)

	newSetting, err := querier.UpdateSetting(context.Background(), UpdateSettingParams{
		Settingtype: SettingTypeActiveGeneration,
		Relatedid:   nullUuid,
	})
	require.NoError(t, err)
	require.NotEmpty(t, newSetting)
	require.False(t, newSetting.Value.Valid)
	require.True(t, newSetting.RelatedID.Valid)
	require.NotEqual(t, setting.RelatedID, newSetting.RelatedID)
	require.NotEqual(t, setting.RelatedID.Valid, newSetting.RelatedID.Valid)
}

func TestGetSetting(t *testing.T) {
	deleteSettingsHelper(t)
	setting, _ := createSettingHelper(t, true, nil)

	newSetting, err := querier.GetSetting(context.Background(), SettingTypeActiveGeneration)
	require.NoError(t, err)
	require.NotEmpty(t, newSetting)
	require.Equal(t, newSetting.RelatedID, setting.RelatedID)
	require.Equal(t, newSetting.SettingType, setting.SettingType)
	require.Equal(t, newSetting.Value, setting.Value)
}
