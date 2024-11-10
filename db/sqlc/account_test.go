package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/v4n1lla-1ce/mini-bank/util"
)

// this function "technically" does not run as a test, but gets used by test functions
func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomBalance(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

// test account creation
func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

// test account deletion
func TestDeleteAccount(t *testing.T) {
	newAcc := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), newAcc.ID)
	require.NoError(t, err)

	deletedAcc, err := testQueries.GetAccount(context.Background(), newAcc.ID)
	require.Empty(t, deletedAcc)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
}

// test getting an account record
func TestGetAccount(t *testing.T) {
	newAcc := createRandomAccount(t)
	getNewAcc, err := testQueries.GetAccount(context.Background(), newAcc.ID)

	require.NoError(t, err)
	require.NotEmpty(t, getNewAcc)

	require.Equal(t, newAcc.ID, getNewAcc.ID)
	require.Equal(t, newAcc.Owner, getNewAcc.Owner)
	require.Equal(t, newAcc.Balance, getNewAcc.Balance)
	require.Equal(t, newAcc.Currency, getNewAcc.Currency)
	require.WithinDuration(t, newAcc.CreatedAt, getNewAcc.CreatedAt, time.Second)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	// get records 5-10 for pagination
	arg := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}

// test updating an account record
func TestUpdateAccount(t *testing.T) {
	newAcc := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID:      newAcc.ID,
		Balance: util.RandomBalance(),
	}

	updatedAcc, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedAcc)

	require.Equal(t, newAcc.ID, updatedAcc.ID)
	require.Equal(t, newAcc.Owner, updatedAcc.Owner)
	require.Equal(t, arg.Balance, updatedAcc.Balance)
	require.Equal(t, newAcc.Currency, updatedAcc.Currency)
	require.WithinDuration(t, newAcc.CreatedAt, updatedAcc.CreatedAt, time.Second)
}
