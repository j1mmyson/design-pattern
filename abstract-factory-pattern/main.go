package main

import (
	"abstract-factory-pattern/factory"
	"abstract-factory-pattern/region"
)

func main() {

	freljordFactory := factory.GetChampionFactory(region.Freljord)
	demaciaFactory := factory.GetChampionFactory(region.Demacia)

	ash := freljordFactory.CreateADCarry()
	ash.Move()
	ash.GetName()
	ash.Fire()

	braum := freljordFactory.CreateSupport()
	braum.Heal()

	sejuani := freljordFactory.CreateTanker()
	sejuani.Shield()

	vayne := demaciaFactory.CreateADCarry()
	vayne.Fire()

	lux := demaciaFactory.CreateSupport()
	lux.Heal()

	garen := demaciaFactory.CreateTanker()
	garen.Shield()

}
