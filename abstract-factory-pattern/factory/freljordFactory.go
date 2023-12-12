package factory

import "abstract-factory-pattern/champion"

// FreljordChampionFactory is a factory for Freljord champions
type FreljordChampionFactory struct {
}

// NewFreljordChampionFactory creates a new FreljordChampionFactory
func NewFreljordChampionFactory() *FreljordChampionFactory {
	return &FreljordChampionFactory{}
}

// CreateADCarry creates a new ADCarry champion
func (f *FreljordChampionFactory) CreateADCarry() champion.IADCarry {
	return &champion.Ashe{}
}

// CreateSupport creates a new Support champion
func (f *FreljordChampionFactory) CreateSupport() champion.ISupport {
	return &champion.Braum{}
}

// CreateTanker creates a new Tanker champion
func (f *FreljordChampionFactory) CreateTanker() champion.ITanker {
	return &champion.Sejuani{}
}
