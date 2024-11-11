package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/v4n1lla-1ce/mini-bank/util"
)

func TestCreateTransaction(t *testing.T) {
	arg := CreateTransactionParams{
		AccountID: createRandomAccount(t).ID,
		Amount:    util.RandomTransactionAmount(),
	}

	transaction, err := testQueries.CreateTransaction(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transaction)

	require.Equal(t, arg.AccountID, transaction.AccountID)
	require.Equal(t, arg.Amount, transaction.Amount)

	require.NotZero(t, transaction.CreatedAt)
	require.NotZero(t, transaction.ID)
}

func TestGetTransaction(t *testing.T) {
	// create account + transaction
	newAcc := createRandomAccount(t)
	arg := CreateTransactionParams{
		AccountID: newAcc.ID,
		Amount:    util.RandomTransactionAmount(),
	}

	transaction, err := testQueries.CreateTransaction(context.Background(), arg)

	fetchedTransaction, err := testQueries.GetTransaction(context.Background(), transaction.ID)
	require.NoError(t, err)
	require.NotEmpty(t, fetchedTransaction)

	require.Equal(t, transaction.ID, fetchedTransaction.ID)
	require.Equal(t, transaction.AccountID, fetchedTransaction.AccountID)
	require.Equal(t, transaction.Amount, fetchedTransaction.Amount)
	require.Equal(t, transaction.CreatedAt, fetchedTransaction.CreatedAt)
}

func TestListTransactions(t *testing.T) {
	// create new account and make many transactions
	newAcc := createRandomAccount(t)

	transactionQty := 12
	for i := 0; i < transactionQty; i++ {
		ctArgs := CreateTransactionParams{
			AccountID: newAcc.ID,
			Amount:    util.RandomTransactionAmount(),
		}
		testQueries.CreateTransaction(context.Background(), ctArgs)
	}

	// get records 2-12
	ltArgs := ListTransactionsParams{
		AccountID: newAcc.ID,
		Limit:     11,
		Offset:    1,
	}
	transactionList, err := testQueries.ListTransactions(context.Background(), ltArgs)
	require.NoError(t, err)
	require.Len(t, transactionList, 11)

	for _, transaction := range transactionList {
		require.NotEmpty(t, transaction)
	}
}
