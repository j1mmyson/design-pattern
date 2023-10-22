package weapon

import "fmt"

// IceBall is a struct that implements the IStrategy interface
type IceBall struct {
	Damage int
	Speed  int
}

// Fire fires the weapon
func (i *IceBall) Fire() {
	fmt.Printf("IceBall fired with %d damage and %d speed\n", i.Damage, i.Speed)
	// fmt.Println("IceBall fired")
	return
}

// GetDamage returns the damage of the weapon
func (i *IceBall) GetDamage() int {
	return i.Damage
}

// GetAttackSpeed returns the attack speed of the weapon
func (i *IceBall) GetAttackSpeed() int {
	return i.Speed
}
