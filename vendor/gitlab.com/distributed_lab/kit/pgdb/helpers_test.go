package pgdb_test

import (
	"os"
	"testing"
	"time"

	"gitlab.com/distributed_lab/kit/pgdb"
)

var prefix = time.Now().UnixNano()

func getDB(t *testing.T) *pgdb.DB {
	t.Helper()
	connURL := os.Getenv("PGDB_CONN_URL")
	if connURL == "" {
		t.Skip("skipping, PGDB_CONN_URL not set")
	}
	db, err := pgdb.Open(pgdb.Opts{
		URL:                connURL,
		MaxOpenConnections: 1,
	})
	if err != nil {
		t.Fatalf("failed to open db: %v", err)
	}
	return db
}
