package factory

import (
	"abstract-factory-pattern/champion"
	"abstract-factory-pattern/region"
)

// IChampionFactory is an interface for the abstract factory pattern
type IChampionFactory interface {
	CreateADCarry() champion.IADCarry
	CreateSupport() champion.ISupport
	CreateTanker() champion.ITanker
}

// GetChampionFactory returns a new champion factory
func GetChampionFactory(rg region.Region) IChampionFactory {
	switch rg {
	case region.Demacia:
		return NewDemaciaChampionFactory()
	case region.Freljord:
		return NewFreljordChampionFactory()
	default:
		return nil
	}
}
