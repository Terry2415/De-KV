package leveldb

import (
	"testing"

	"github.com/Terry2415/De-KV/leveldb/testutil"
)

func TestLevelDB(t *testing.T) {
	testutil.RunSuite(t, "LevelDB Suite")
}
