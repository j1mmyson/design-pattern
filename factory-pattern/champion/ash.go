package champion

import "fmt"

// Ash is a champion
type Ash struct {
	Name string
}

// Move moves the champion
func (a *Ash) Move() {
	fmt.Println("Ash moved")
	return
}

// Attack attacks the champion
func (a *Ash) Attack() {
	fmt.Println("Ash attacked")
	return
}

// GetName returns the name of the champion
func (a *Ash) GetName() string {
	return a.Name
}
