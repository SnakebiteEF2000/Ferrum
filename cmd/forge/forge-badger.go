package main

import (
	"log/slog"
	"path/filepath"

	"github.com/dgraph-io/badger/v3"
)

func OpenBadgerCon() (*badger.DB, error) {
	dir := filepath.Join(".", "data", "badger")
	opts := badger.DefaultOptions(dir).WithSyncWrites(true)

	db, err := badger.Open(opts)
	if err != nil {
		slog.Error("failed to open badgerdb", "dir", dir, "error", err)
	}
	defer db.Close()

	return db, nil
}
