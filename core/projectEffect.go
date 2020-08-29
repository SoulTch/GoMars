package core

var always = func(g *Game, p *Player, d interface{}) bool {
	return true
}

var skip = func(g *Game, p *Player, d interface{}) { /* do nothing */ }

// ProjectEffects .
var ProjectEffects = map[string]func([]int) Effect{
	"nop": func(param []int) Effect {
		return Effect{
			condition: always,
			effect:    skip}
	},
	"increase_megacredit": func(param []int) Effect {
		amount := param[0]
		return Effect{
			condition: always,
			effect: func(g *Game, p *Player, d interface{}) {
				p.getEnchant()["megacredit"] += amount
			}}
	}}
