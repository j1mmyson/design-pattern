package main

import "strategy-pattern/weapon"

// Player is a struct that contains a weapon
type Player struct {
	damage      int
	attackSpeed int
	weapon      weapon.IStrategy
}

// SetWeapon sets the weapon for the player
func (p *Player) SetWeapon(w weapon.IStrategy) {
	p.weapon = w
	p.damage = w.GetDamage()
	p.attackSpeed = w.GetAttackSpeed()
}

// Fire fires the weapon
func (p *Player) Fire() {
	p.weapon.Fire()
}
