package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/Wanna-be-Coder/goAPIDokcer/util"
	"github.com/stretchr/testify/require"
)

func createNewTransfer(t *testing.T) Transfer {
	account1 := createNewAccount(t)
	account2 := createNewAccount(t)
	arg := CreateTransfersParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomBalance(),
	}
	tranfer, err := testQueries.CreateTransfers(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, tranfer)

	require.Equal(t, arg.FromAccountID, tranfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, tranfer.ToAccountID)
	require.Equal(t, arg.Amount, tranfer.Amount)

	require.NotEmpty(t, tranfer.ID)
	require.NotEmpty(t, tranfer.CreatedAt)

	return tranfer

}

func TestCreateTransfer(t *testing.T) {

	createNewTransfer(t)

}

func TestGetTransfer(t *testing.T) {
	transfer := createNewTransfer(t)

	rTransfer, err := testQueries.GetTransfers(context.Background(), transfer.ID)

	require.NoError(t, err)
	require.NotEmpty(t, rTransfer)

	require.Equal(t, transfer.FromAccountID, rTransfer.FromAccountID)
	require.Equal(t, transfer.ToAccountID, rTransfer.ToAccountID)
	require.Equal(t, transfer.Amount, rTransfer.Amount)

	require.NotEmpty(t, rTransfer.ID)
	require.NotEmpty(t, rTransfer.CreatedAt)

}

func TestUpdateTransfer(t *testing.T) {
	transfer := createNewTransfer(t)

	arg := UpdatetransfersParams{
		ID:     transfer.ID,
		Amount: util.RandomBalance(),
	}
	rTransfer, err := testQueries.Updatetransfers(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, rTransfer)

	require.Equal(t, transfer.FromAccountID, rTransfer.FromAccountID)
	require.Equal(t, transfer.ToAccountID, rTransfer.ToAccountID)
	require.Equal(t, arg.Amount, rTransfer.Amount)

	require.NotEmpty(t, rTransfer.ID)
	require.NotEmpty(t, rTransfer.CreatedAt)

}

func TestDeleteTransfer(t *testing.T) {
	transfer := createNewTransfer(t)

	err := testQueries.DeleteTransfers(context.Background(), transfer.ID)

	require.NoError(t, err)

	deleted, err := testQueries.GetTransfers(context.Background(), transfer.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, deleted)

}

func TestListTransfer(t *testing.T) {
	for i := 0; i < 10; i++ {
		createNewTransfer(t)
	}
	arg := ListTransfersParams{
		Limit:  5,
		Offset: 5,
	}
	tranfers, err := testQueries.ListTransfers(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, tranfers, 5)

	for _, transfer := range tranfers {
		require.NotEmpty(t, transfer)
	}

}
