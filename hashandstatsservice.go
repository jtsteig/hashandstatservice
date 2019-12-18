package hashandstatsservice

import (
	encodedhash "github.com/jtsteig/encodedhash"
	hashmodel "github.com/jtsteig/hashmodels"
)

// HashStatsService is the orchestration entry point for storing entries and stats.
type HashStatsService struct {
	HashRepository *hashmodel.HashRepository
}

// CreateEmptyHashEntry creates an empty row for later updating with a hash value and time.
func (service *HashStatsService) CreateEmptyHashEntry() (int64, error) {
	countID, err := service.HashRepository.CreateEmptyHashEntry()
	if err != nil {
		return -1, err
	}
	return countID, nil
}

// StoreValue takes a plain string value, calculates the hash and stores it and the elapsed time to persistence.
func (service *HashStatsService) StoreValue(countID int64, value string) error {
	hash, duration := encodedhash.CalculateHash(value)

	err := service.HashRepository.UpdateHashWithValues(countID, hash, duration.Microseconds())
	if err != nil {
		return err
	}

	return nil
}

// GetHash returns the HashStats for a countID or an error if not found.
func (service *HashStatsService) GetHash(countID int64) (hashmodel.HashStat, error) {
	hashStat, err := service.HashRepository.GetHashStat(countID)
	if err != nil {
		return hashmodel.HashStat{}, err
	}
	return hashStat, err
}

// GetTotalStats returns the total stats for all runs and returns an error if anything goes amiss.
func (service *HashStatsService) GetTotalStats() (hashmodel.TotalStats, error) {
	totalStats, err := service.HashRepository.GetTotalStats()
	if err != nil {
		return hashmodel.TotalStats{}, err
	}
	return totalStats, err
}

// Close cleans up everything. This MUST be called to close connections.
func (service *HashStatsService) Close() error {
	if err := service.Close(); err != nil {
		return err
	}
	return nil
}
