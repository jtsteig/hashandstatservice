package hashandstatsservice

import (
	"database/sql"
	"testing"

	hashmodels "github.com/jtsteig/hashmodels"
)

func TestServiceHappyPath(t *testing.T) {
	filename := "c:\\temp\\testdb.db"
	hashTable := "hashes"
	db, _ := sql.Open("sqlite3", filename)
	hashStore, initErr := hashmodels.NewHashStore(db, hashTable)
	cleanup := func() {
		hashStore.ClearStore()
		hashStore.Close()
	}
	defer cleanup()

	if initErr != nil {
		t.Errorf("Failed to init db: %q", initErr)
	}

	service := HashStatsService{hashStore}

	countID, err := service.StoreValue("OgdI800IckhuWE8rsRzxPoGfUPhP7mah14HBCJeF7Pltu6CN8Vgcs6ylTbKKQvKQCGG4qQmRfLMwrjJ9TXsG95rQ58k8kkvEsAV2kr40Z2wMvFrHYlQ3vOIl8qpImjEwpr7gZQGpCwK96iEWwtXIjGJomjCmWDgqE4dcXt4H351t7LNxR4q32VGdJ49VpREnsdbPwRxZ")
	if err != nil {
		t.Errorf("failed to store a value: %q", err)
	}
	ret, getErr := service.GetHash(countID)
	if getErr != nil {
		t.Errorf("Failed to getHash: %q", getErr)
	}
	if ret.HashValue == "" {

		t.Errorf("Didn't get appropriate hash: %q", ret.HashValue)
	}

	service.StoreValue("2")
	service.StoreValue("3")
	service.StoreValue("5555555")
	service.StoreValue("6666666666666666666666")
	service.StoreValue("77777777777777777777777777777777777")
	totalStats, totalErr := service.GetTotalStats()
	if totalErr != nil {
		t.Errorf("Failed to get totalStats: %q", totalErr)
	}
	if totalStats.Count != 6 {
		t.Errorf("Didn't get the correct number of hashes back. Expected 6, but got %d", totalStats.Count)
	}
}
