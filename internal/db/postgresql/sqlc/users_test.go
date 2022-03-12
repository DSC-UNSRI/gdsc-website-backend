package postgresql

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	user, err := queries.CreateUser(context.Background(), CreateUserParams{
		Name:      "Arya",
		Birthdate: time.Date(2002, time.April, 10, 10, 10, 10, 10, time.Local),
	})

	t.Logf("ini user %v", user)
	assert.NoError(t, err)
	assert.NotEmpty(t, user)
	assert.NotZero(t, user.Uuid.ID())
}

func TestListUsers(t *testing.T) {
	users, err := queries.ListUsers(context.Background(), ListUsersParams{
		Limit:  10,
		Offset: 0,
	})

	assert.NoError(t, err)
	assert.NotEmpty(t, users)
	for _, v := range users {
		fmt.Printf("nama : %v\n", v.Name)
	}
}
