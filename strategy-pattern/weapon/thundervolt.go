package weapon

import "fmt"

// ThunderVolt is a struct that implements the IStrategy interface
type ThunderVolt struct {
	Damage int
	Speed  int
}

// Fire fires the weapon
func (t *ThunderVolt) Fire() {
	fmt.Printf("ThunderVolt fired with %d damage and %d speed\n", t.Damage, t.Speed)
	// fmt.Println("ThunderVolt fired")
	return
}

// GetDamage returns the damage of the weapon
func (t *ThunderVolt) GetDamage() int {
	return t.Damage
}

// GetAttackSpeed returns the attack speed of the weapon
func (t *ThunderVolt) GetAttackSpeed() int {
	return t.Speed
}
