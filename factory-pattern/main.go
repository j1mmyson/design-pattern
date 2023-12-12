package main

import (
	"factory-pattern/factory"
	"fmt"
)

func main() {

	championFactory := factory.NewChampionFactory()

	ash := championFactory.CreateChampion("ash")

	ash.Move()
	ash.Attack()
	fmt.Println("Champion's name: " + ash.GetName())
}
