package weapon

// IStrategy is an interface for the strategy pattern
type IStrategy interface {
	Fire()
	GetDamage() int
	GetAttackSpeed() int
}
