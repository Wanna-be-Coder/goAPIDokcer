package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/teachschool/simplebank/util"
)

func createNewEntry(t *testing.T) Entry {
	account := createNewAccount(t)

	arg := CreateEntriesParams{
		Balance:   util.RandomBalance(),
		AccountID: account.ID,
	}
	entry, err := testQueries.CreateEntries(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Balance, entry.Balance)

	require.NotEmpty(t, entry.ID)
	require.NotEmpty(t, entry.CreatedAt)

	return entry

}

func TestCreateEntry(t *testing.T) {

	createNewEntry(t)

}

func TestGetEntry(t *testing.T) {
	entry := createNewEntry(t)

	rEntry, err := testQueries.GetEntries(context.Background(), entry.ID)

	require.NoError(t, err)
	require.NotEmpty(t, rEntry)

	require.Equal(t, rEntry.AccountID, entry.AccountID)
	require.Equal(t, rEntry.Balance, entry.Balance)

}

func TestUpdateEntry(t *testing.T) {
	entry := createNewEntry(t)

	arg := UpdateEntriesParams{
		ID:      entry.ID,
		Balance: util.RandomBalance(),
	}
	rEntry, err := testQueries.UpdateEntries(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, rEntry)

	require.Equal(t, entry.ID, rEntry.ID)
	require.Equal(t, arg.Balance, rEntry.Balance)

}

func TestDeleteEntry(t *testing.T) {
	entry := createNewEntry(t)

	err := testQueries.DeleteEntries(context.Background(), entry.ID)

	require.NoError(t, err)

	deleted, err := testQueries.GetEntries(context.Background(), entry.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, deleted)

}

func TestListEntries(t *testing.T) {
	for i := 0; i < 10; i++ {
		createNewEntry(t)
	}
	arg := ListEntriesParams{
		Limit:  5,
		Offset: 5,
	}
	entries, err := testQueries.ListEntries(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}

}
