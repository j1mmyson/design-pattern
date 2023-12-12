package factory

import "abstract-factory-pattern/champion"

// DemaciaChampionFactory is a factory for Demacia champions
type DemaciaChampionFactory struct {
}

// NewDemaciaChampionFactory creates a new DemaciaChampionFactory
func NewDemaciaChampionFactory() *DemaciaChampionFactory {
	return &DemaciaChampionFactory{}
}

// CreateADCarry creates a new ADCarry champion
func (d *DemaciaChampionFactory) CreateADCarry() champion.IADCarry {
	return &champion.Vayne{}
}

// CreateSupport creates a new Support champion
func (d *DemaciaChampionFactory) CreateSupport() champion.ISupport {
	return &champion.Lux{}
}

// CreateTanker creates a new Tanker champion
func (d *DemaciaChampionFactory) CreateTanker() champion.ITanker {
	return &champion.Garen{}
}
