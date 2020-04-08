package memdb

import (
	"testing"

	"github.com/Terry2415/De-KV/leveldb/testutil"
)

func TestMemDB(t *testing.T) {
	testutil.RunSuite(t, "MemDB Suite")
}
