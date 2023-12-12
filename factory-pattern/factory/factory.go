package factory

import "factory-pattern/champion"

type championFactory struct {
}

// NewChampionFactory returns a new champion factory
func NewChampionFactory() *championFactory {
	return &championFactory{}
}

func (cf *championFactory) CreateChampion(championType string) champion.Champion {
	switch championType {
	case "ash":
		return &champion.Ash{Name: "Ash"}

	case "janna":
		return &champion.Janna{Name: "Janna"}

	default:
		return nil
	}
}
