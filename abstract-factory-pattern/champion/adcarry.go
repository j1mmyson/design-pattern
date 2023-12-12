package champion

import "fmt"

type IADCarry interface {
	Fire()
	IChampion
}

type Vayne struct {
}

func (v *Vayne) Fire() {
	fmt.Println("Vayne is firing")
}

func (v *Vayne) Move() {
	fmt.Println("Vayne moved")
	return
}

func (v *Vayne) GetName() string {
	return "Vayne"
}

type Ashe struct {
}

func (a *Ashe) Fire() {
	fmt.Println("Ashe is firing")
}

func (a *Ashe) Move() {
	fmt.Println("Ashe moved")
	return
}

func (a *Ashe) GetName() string {
	return "Ashe"
}
