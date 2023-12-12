package champion

import "fmt"

// Janna is a champion
type Janna struct {
	Name string
}

// Move moves the champion
func (j *Janna) Move() {
	fmt.Println("Ash moved")
	return
}

// Attack attacks the champion
func (j *Janna) Attack() {
	fmt.Println("Ash attacked")
	return
}

// GetName returns the name of the champion
func (j *Janna) GetName() string {
	return j.Name
}
