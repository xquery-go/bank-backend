package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/v4n1lla-1ce/mini-bank/util"
)

func TestCreateTransfer(t *testing.T) {
	// create new transfer
	args := CreateTransferParams{
		FromAccountID: createRandomAccount(t).ID,
		ToAccountID:   createRandomAccount(t).ID,
		Amount:        util.RandomTransactionAmount(),
	}

	// create transfer with a random amount
	transfer, err := testQueries.CreateTransfer(context.Background(), args)

	// no errors and not empty
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	// transfer created must equal args given
	require.Equal(t, args.FromAccountID, transfer.FromAccountID)
	require.Equal(t, args.ToAccountID, transfer.ToAccountID)
	require.Equal(t, args.Amount, transfer.Amount)

	// make sure auto fields aren't 0
	require.NotZero(t, transfer.CreatedAt)
	require.NotZero(t, transfer.ID)
}

func TestGetTransfer(t *testing.T) {
	// create new transfer
	args := CreateTransferParams{
		FromAccountID: createRandomAccount(t).ID,
		ToAccountID:   createRandomAccount(t).ID,
		Amount:        util.RandomTransactionAmount(),
	}

	transfer, _ := testQueries.CreateTransfer(context.Background(), args)

	// get transfer by id
	fetchedTransfer, err := testQueries.GetTransfer(context.Background(), transfer.ID)

	// no error and not emppty
	require.NoError(t, err)
	require.NotEmpty(t, fetchedTransfer)

	// comparison
	require.Equal(t, transfer.ID, fetchedTransfer.ID)
	require.Equal(t, transfer.FromAccountID, fetchedTransfer.FromAccountID)
	require.Equal(t, transfer.ToAccountID, fetchedTransfer.ToAccountID)
	require.Equal(t, transfer.Amount, fetchedTransfer.Amount)
	require.WithinDuration(t, transfer.CreatedAt, fetchedTransfer.CreatedAt, time.Second)
}

func TestListTransfers(t *testing.T) {
	// create transfer between 2 accounts
	acc1 := createRandomAccount(t)
	acc2 := createRandomAccount(t)

	// simulate 20 transactions between 2 accounts
	for i := 0; i < 20; i++ {
		randTransferParams := CreateTransferParams{
			FromAccountID: acc1.ID,
			ToAccountID:   acc2.ID,
			Amount:        util.RandomTransactionAmount(),
		}

		testQueries.CreateTransfer(context.Background(), randTransferParams)
	}

	arg := ListTransfersParams{
		FromAccountID: acc1.ID,
		ToAccountID:   acc2.ID,
		Limit:         100,
		Offset:        0,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)

	// no errors and 20 length
	require.NoError(t, err)
	require.Len(t, transfers, 20)

	// loop through all transfers and make sure it isnt empty and is
	// from the correct account to the correct account
	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
		require.Equal(t, transfer.FromAccountID, arg.FromAccountID)
		require.Equal(t, transfer.ToAccountID, arg.ToAccountID)
	}
}
