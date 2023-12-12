package champion

import "fmt"

type ITanker interface {
	Shield()
	IChampion
}

type Garen struct {
}

func (g *Garen) Shield() {
	fmt.Println("Garen is shielding")
	return
}

func (g *Garen) Move() {
	fmt.Println("Garen moved")
	return
}

func (g *Garen) GetName() string {
	return "Garen"
}

type Sejuani struct {
}

func (s *Sejuani) Shield() {
	fmt.Println("Sejuani is shielding")
	return
}

func (s *Sejuani) Move() {
	fmt.Println("Sejuani moved")
	return
}

func (s *Sejuani) GetName() string {
	return "Sejuani"
}
