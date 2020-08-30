package core

type resType int

const (
	megacredit = resType(iota)
	iron
	titanium
	plant
	energy
	heat
	resSize
)

type payWith int

const (
	withIron = payWith(iota)
	withTitanium
	withHeat
)

type payMethod struct {
	methods int
}

func (p *payMethod) add(x payWith) {
	p.methods ^= (1 << int(x))
}

func (p *payMethod) enabled(x payWith) bool {
	return (p.methods & (1 << int(x))) > 0
}
