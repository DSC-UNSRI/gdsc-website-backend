package postgresql_test

import (
	"context"
	"log"
	"testing"
	"time"

	postgresql "github.com/DSC-UNSRI/gdsc-website-backend/internal/db/postgresql/sqlc"
	"github.com/bxcodec/faker/v3"
	"github.com/jackc/pgx/v4"
	"github.com/stretchr/testify/require"
)

func createDivision(t *testing.T) postgresql.Division {
	division, err := querier.CreateDivision(context.Background(), faker.Name())
	require.NoError(t, err)
	require.NotEmpty(t, division)
	require.NotZero(t, division.ID)
	require.IsType(t, division.CreatedAt, time.Now())
	require.NotEmpty(t, division.Name)

	return division
}

func createDivisionLongString(t *testing.T) postgresql.Division {
	longString := "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz"
	division, err := querier.CreateDivision(context.Background(), longString)
	if err != nil {
		log.Fatalf("something went wrong %v", err)
	}

	require.Equal(t, longString[:255], division.Name)

	return division
}

func TestCreateDivision(t *testing.T) {
	createDivision(t)
	createDivisionLongString(t)
}

func TestUpdateDivision(t *testing.T) {
	faker.SetGenerateUniqueValues(true)
	division := createDivision(t)
	randomName := faker.Name()

	updatedDivision, err := querier.UpdateDivision(context.Background(), postgresql.UpdateDivisionParams{
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
	division := createDivision(t)

	newDivision, err := querier.GetDivision(context.Background(), division.ID)
	require.NoError(t, err)
	require.NotEmpty(t, newDivision)
	require.Equal(t, newDivision.ID, division.ID)
	require.Equal(t, division.Name, newDivision.Name)
	require.Equal(t, division.CreatedAt, newDivision.CreatedAt)
}

func TestListDivision(t *testing.T) {
	var count int = 5
	for i := 0; i < count; i++ {
		createDivision(t)
	}

	divisions, err := querier.ListDivisions(context.Background(), postgresql.ListDivisionsParams{
		Limit:  5,
		Offset: 0,
	})
	require.NoError(t, err)
	require.Len(t, divisions, count)
	for _, v := range divisions {
		require.NotEmpty(t, v)
	}
}

func TestDeleteDivision(t *testing.T) {
	division := createDivision(t)
	rows, err := querier.DeleteDivision(context.Background(), division.ID)
	require.NoError(t, err)
	require.NotZero(t, rows)
	_, err = querier.GetDivision(context.Background(), division.ID)
	require.ErrorIs(t, pgx.ErrNoRows, err)
}
