package core

//go:generate stringer -type=resType
type resType int

const (
	megacredit resType = iota
	iron
	titanium
	plant
	energy
	heat
	resSize
)

var toResType = map[string]resType{
	"megacredit": megacredit,
	"iron":       iron,
	"titanium":   titanium,
	"plant":      plant,
	"energy":     energy,
	"heat":       heat}

//go:generate stringer -type=payWith
type payWith int

const (
	withIron payWith = iota
	withTitanium
	withHeat
	paySize
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

func (p *payMethod) Array() []string {
	ret := make([]string, 0, 5)

	for i := 0; i < int(paySize); i++ {
		if p.enabled(payWith(i)) {
			ret = append(ret, string(payWith(i)))
		}
	}

	return ret
}
