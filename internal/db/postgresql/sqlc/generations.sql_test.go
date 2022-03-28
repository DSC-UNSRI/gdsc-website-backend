package postgresql

import (
	"context"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v4"
	"github.com/stretchr/testify/require"
)

var isAlreadyExecutedSetupActiveGeneration bool

func setupActiveGeneration(t *testing.T, shouldExecutedOnce bool) {
	if shouldExecutedOnce && isAlreadyExecutedSetupActiveGeneration {
		return
	}

	var uuid uuid.NullUUID
	gen, _ := createGenerationHelper(t, nil)

	uuid.Scan(gen.ID.String())

	setupSettingsTable(t, &gen.ID)
	isAlreadyExecutedSetupActiveGeneration = true
}

func createGenerationHelper(t *testing.T, year *string) (CreateGenerationRow, error) {
	var finalYear string
	if year == nil {
		finalYear = faker.YearString()
	} else {
		finalYear = *year
	}
	gen, err := querier.CreateGeneration(context.Background(), finalYear)
	if year == nil {
		require.NoError(t, err)
		require.NotEmpty(t, gen)
		require.NotZero(t, gen.Year)
		require.NotEmpty(t, gen.ID)
		require.NotZero(t, gen.CreatedAt)
		require.False(t, gen.Active)
	}

	return gen, err
}

func TestCreateGeneration(t *testing.T) {
	createGenerationHelper(t, nil)
}

func deleteGenerations(t *testing.T) {
	_, err := querier.db.Exec(context.Background(), "DELETE FROM generations")
	require.NoError(t, err)
}

func TestCreateGenerationDuplicate(t *testing.T) {
	faker.SetGenerateUniqueValues(true)
	year := faker.YearString()
	gen, err := createGenerationHelper(t, &year)
	require.NoError(t, err)
	require.NotEmpty(t, gen)

	gen, err = createGenerationHelper(t, &year)
	require.Error(t, err)
	require.Empty(t, gen)
	require.IsType(t, &pgconn.PgError{}, err)
	var pgErr *pgconn.PgError
	pgErr, ok := err.(*pgconn.PgError)
	require.True(t, ok)
	require.Equal(t, pgerrcode.UniqueViolation, pgErr.Code)
}

func TestGetGeneration(t *testing.T) {
	gen, _ := createGenerationHelper(t, nil)
	newGen, err := querier.GetGeneration(context.Background(), gen.ID)
	require.NoError(t, err)

	require.NotEmpty(t, newGen)
	require.Equal(t, gen.ID, newGen.ID)
	require.Equal(t, gen.CreatedAt, newGen.CreatedAt)
	require.Equal(t, gen.Year, newGen.Year)
}

func TestGetAndUpdateActiveGeneration(t *testing.T) {
	faker.SetGenerateUniqueValues(true)
	setupSettingsTable(t, nil)

	gen, _ := createGenerationHelper(t, nil)

	var id uuid.NullUUID
	err := id.Scan(gen.ID.String())
	require.NoError(t, err)

	oldGen, err := querier.SetActiveGeneration(context.Background(), id)
	require.NoError(t, err)
	require.NotEmpty(t, oldGen)

	genActive, err := querier.GetActiveGeneration(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, genActive)

	require.Equal(t, genActive.ID.String(), gen.ID.String())
	require.Equal(t, gen.ID.String(), genActive.ID.String())
	require.Equal(t, genActive.CreatedAt, gen.CreatedAt)
	require.Equal(t, genActive.Year, gen.Year)
	require.True(t, genActive.Active)

	require.Equal(t, oldGen.ID.String(), genActive.ID.String())
	require.Equal(t, oldGen.CreatedAt, genActive.CreatedAt)
	require.Equal(t, oldGen.Year, genActive.Year)
	require.Equal(t, oldGen.Active, genActive.Active)
	require.True(t, oldGen.Active)
}

func TestListGenerations(t *testing.T) {
	setupActiveGeneration(t, true)
	faker.SetGenerateUniqueValues(true)
	var offset int
	row := querier.db.QueryRow(context.Background(), "SELECT COUNT(*) FROM generations")
	row.Scan(&offset)

	count := 3
	for i := 0; i < count; i++ {
		createGenerationHelper(t, nil)
	}

	listGen, err := querier.ListGeneration(context.Background(), ListGenerationParams{
		Limit:  int32(count),
		Offset: int32(offset),
	})
	require.NoError(t, err)
	require.Len(t, listGen, count)

	for _, v := range listGen {
		require.NotEmpty(t, v)
	}
}

func TestUpdateGeneration(t *testing.T) {
	faker.SetGenerateUniqueValues(true)
	gen, _ := createGenerationHelper(t, nil)

	year := faker.YearString()

	newGen, err := querier.UpdateGeneration(context.Background(), UpdateGenerationParams{
		Generationid: gen.ID,
		Year:         year,
	})
	require.NoError(t, err)
	require.NotEmpty(t, newGen)
	require.Equal(t, year, newGen.Year)
	require.False(t, newGen.Active)
	require.NotEqual(t, gen.Year, newGen.Year)
}

func TestDeleteGeneration(t *testing.T) {
	gen, _ := createGenerationHelper(t, nil)

	rows, err := querier.DeleteGeneration(context.Background(), gen.ID)
	require.NoError(t, err)
	require.NotZero(t, rows)

	_, err = querier.GetGeneration(context.Background(), gen.ID)
	require.ErrorIs(t, pgx.ErrNoRows, err)
}
