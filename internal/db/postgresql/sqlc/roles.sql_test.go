package postgresql

import (
	"context"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/jackc/pgx/v4"
	"github.com/stretchr/testify/require"
)

func createRoleHelper(t *testing.T, name *string) Role {
	var roleName string
	if name != nil {
		roleName = *name
	} else {
		roleName = faker.Name()
	}
	role, err := querier.CreateRole(context.Background(), roleName)
	require.NoError(t, err)

	require.NotEmpty(t, role)
	require.NotEmpty(t, role.Name)
	require.NotEmpty(t, role.CreatedAt)
	require.NotEmpty(t, role.ID)
	require.NotEmpty(t, role.ID.String())

	return role
}

func TestCreateRole(t *testing.T) {
	createRoleHelper(t, nil)
}

func TestUpdateRole(t *testing.T) {
	faker.SetGenerateUniqueValues(true)
	role := createRoleHelper(t, nil)

	updatedName := faker.Name()
	updatedRole, err := querier.UpdateRole(context.Background(), UpdateRoleParams{
		Name:   updatedName,
		Roleid: role.ID,
	})
	require.NoError(t, err)
	require.NotEmpty(t, updatedRole)
	require.Equal(t, updatedName, updatedRole.Name)
	require.Equal(t, updatedRole.ID, role.ID)
}

func TestDeleteRole(t *testing.T) {
	role := createRoleHelper(t, nil)

	rows, err := querier.DeleteRole(context.Background(), role.ID)
	require.NoError(t, err)
	require.Equal(t, int64(1), rows)

	_, err = querier.GetRole(context.Background(), role.ID)
	require.ErrorIs(t, pgx.ErrNoRows, err)
}

func TestGetRole(t *testing.T) {
	role := createRoleHelper(t, nil)

	newRole, err := querier.GetRole(context.Background(), role.ID)
	require.NoError(t, err)

	require.NotEmpty(t, newRole)
	require.Equal(t, role.ID, newRole.ID)
	require.Equal(t, role.CreatedAt, newRole.CreatedAt)
	require.Equal(t, role.Name, newRole.Name)
}

func TestListRole(t *testing.T) {
	count := 5
	for i := 0; i < count; i++ {
		createRoleHelper(t, nil)
	}

	roles, err := querier.ListRole(context.Background(), ListRoleParams{
		Limit:  int32(count),
		Offset: 0,
	})
	require.NoError(t, err)
	require.Len(t, roles, count)

	for _, role := range roles {
		require.NotEmpty(t, role)
	}
}
