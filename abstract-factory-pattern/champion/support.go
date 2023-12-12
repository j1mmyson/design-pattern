package champion

type ISupport interface {
	Heal()
	IChampion
}

type Lux struct {
}

func (l *Lux) Move() {
	println("Lux moved")
	return
}

func (l *Lux) GetName() string {
	return "Lux"
}

func (l *Lux) Heal() {
	println("Lux is healing")
	return
}

type Braum struct {
}

func (b *Braum) Move() {
	println("Braum moved")
	return
}

func (b *Braum) GetName() string {
	return "Braum"
}

func (b *Braum) Heal() {
	println("Braum is healing")
	return
}
