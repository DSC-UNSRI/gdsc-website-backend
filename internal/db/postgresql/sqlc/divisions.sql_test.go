package postgresql

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/jackc/pgx/v4"
	"github.com/stretchr/testify/require"
)

func createDivisionHelper(t *testing.T) Division {
	division, err := querier.CreateDivision(context.Background(), CreateDivisionParams{
		Name: faker.Name(),
		Type: DivisionTypeDIVISION,
	})
	require.NoError(t, err)
	require.NotEmpty(t, division)
	require.NotZero(t, division.ID)
	require.IsType(t, division.CreatedAt, time.Now())
	require.NotEmpty(t, division.Name)

	return division
}

func createDivisionLongString(t *testing.T) Division {
	longString := "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz"
	division, err := querier.CreateDivision(context.Background(), CreateDivisionParams{
		Name: longString,
		Type: DivisionTypeCOLEAD,
	})
	if err != nil {
		log.Fatalf("something went wrong %v", err)
	}

	require.Equal(t, longString[:255], division.Name)

	return division
}

func TestCreateDivision(t *testing.T) {
	setupActiveGeneration(t, true)
	createDivisionHelper(t)
	createDivisionLongString(t)
}

func TestUpdateDivision(t *testing.T) {
	setupActiveGeneration(t, true)
	faker.SetGenerateUniqueValues(true)
	division := createDivisionHelper(t)
	randomName := faker.Name()

	updatedDivision, err := querier.UpdateDivision(context.Background(), UpdateDivisionParams{
		Divisionid: division.ID,
		Name:       randomName,
	})
	require.NoError(t, err)
	require.NotEmpty(t, updatedDivision)
	require.NotZero(t, updatedDivision.ID)
	require.IsType(t, time.Now(), updatedDivision.CreatedAt)
	require.Equal(t, division.ID, updatedDivision.ID)
	require.Equal(t, randomName, updatedDivision.Name)
	require.NotEqual(t, division.Name, updatedDivision.Name)
}

func TestGetDivision(t *testing.T) {
	setupActiveGeneration(t, true)
	division := createDivisionHelper(t)

	newDivision, err := querier.GetDivision(context.Background(), division.ID)
	require.NoError(t, err)
	require.NotEmpty(t, newDivision)
	require.Equal(t, newDivision.ID, division.ID)
	require.Equal(t, division.Name, newDivision.Name)
	require.Equal(t, division.CreatedAt, newDivision.CreatedAt)
}

func TestListAllDivision(t *testing.T) {
	setupActiveGeneration(t, true)
	var count int = 5
	for i := 0; i < count; i++ {
		createDivisionHelper(t)
	}

	divisions, err := querier.ListAllDivisions(context.Background(), ListAllDivisionsParams{
		Limit:  int32(count),
		Offset: 0,
	})
	require.NoError(t, err)
	require.Len(t, divisions, count)
	for _, v := range divisions {
		require.NotEmpty(t, v)
	}
}

func deleteAllDivisions(t *testing.T) {
	_, err := querier.db.Exec(context.Background(), "DELETE FROM divisions")
	require.NoError(t, err)
}

func TestListActiveDivisions(t *testing.T) {
	setupActiveGeneration(t, true)
	deleteAllDivisions(t)
	count := 4
	for i := 0; i < count; i++ {
		createDivisionHelper(t)
	}

	divisions, err := querier.ListActiveDivisions(context.Background(), ListActiveDivisionsParams{
		Limit:  int32(count),
		Offset: 0,
	})
	require.NoError(t, err)
	require.NotEmpty(t, divisions)

	for _, v := range divisions {
		require.Equal(t, DivisionTypeDIVISION, v.Type)
	}

}

func TestDeleteDivision(t *testing.T) {
	division := createDivisionHelper(t)
	rows, err := querier.DeleteDivision(context.Background(), division.ID)
	require.NoError(t, err)
	require.NotZero(t, rows)
	_, err = querier.GetDivision(context.Background(), division.ID)
	require.ErrorIs(t, pgx.ErrNoRows, err)
}
