package weapon

import "fmt"

// Fireball is a struct that implements the IStrategy interface
type Fireball struct {
	Damage int
	Speed  int
}

// Fire fires the weapon
func (f *Fireball) Fire() {
	fmt.Printf("Fireball fired with %d damage and %d speed\n", f.Damage, f.Speed)
	// fmt.Println("Fireball fired")
	return
}

// GetDamage returns the damage of the weapon
func (f *Fireball) GetDamage() int {
	return f.Damage
}

// GetAttackSpeed returns the attack speed of the weapon
func (f *Fireball) GetAttackSpeed() int {
	return f.Speed
}
