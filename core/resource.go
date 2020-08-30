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

type payMethod int

const (
	withIron = payMethod(iota)
	withTitanium
	withHeat
)
