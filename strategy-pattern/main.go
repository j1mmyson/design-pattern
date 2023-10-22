package main

import "strategy-pattern/weapon"

const (
	// FireballDamage is the damage of the fireball
	FireballDamage = 100
	// FireballSpeed is the speed of the fireball
	FireballSpeed = 50

	// ThunderVoltDamage is the damage of the thundervolt
	ThunderVoltDamage = 50
	// ThunderVoltSpeed is the speed of the thundervolt
	ThunderVoltSpeed = 100

	// IceBallDamage is the damage of the iceball
	IceBallDamage = 75
	// IceBallSpeed is the speed of the iceball
	IceBallSpeed = 75
)

func main() {

	player := Player{}

	player.SetWeapon(&weapon.Fireball{
		Damage: FireballDamage,
		Speed:  FireballSpeed,
	})
	player.Fire()

	player.SetWeapon(&weapon.ThunderVolt{
		Damage: ThunderVoltDamage,
		Speed:  ThunderVoltSpeed,
	})
	player.Fire()

	player.SetWeapon(&weapon.IceBall{
		Damage: IceBallDamage,
		Speed:  IceBallSpeed,
	})
	player.Fire()
}
