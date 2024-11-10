package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/v4n1lla-1ce/mini-bank/util"
)

func TestCreateTransaction(t *testing.T) {
	// test invalid transaction with account that doesn't exist
	invalidArg := CreateTransactionParams{
		AccountID: util.RandomInvalidAccountId(),
		Amount:    util.RandomTransactionAmount(),
	}

	invalidTransaction, err := testQueries.CreateTransaction(context.Background(), invalidArg)

	require.Error(t, err)
	require.Empty(t, invalidTransaction)

	validArg := CreateTransactionParams{
		AccountID: createRandomAccount(t).ID,
		Amount:    util.RandomTransactionAmount(),
	}

	validTransaction, err := testQueries.CreateTransaction(context.Background(), validArg)
	require.NoError(t, err)
	require.NotEmpty(t, validTransaction)

	require.Equal(t, validArg.AccountID, validTransaction.AccountID)
	require.Equal(t, validArg.Amount, validTransaction.Amount)

	require.NotZero(t, validTransaction.CreatedAt)
	require.NotZero(t, validTransaction.ID)
}
