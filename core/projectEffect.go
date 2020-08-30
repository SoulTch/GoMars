package core

var always = func(d belongsToPlayer) bool {
	return true
}

var skip = func(d belongsToPlayer) { /* do nothing */ }

var projectEffects = map[string]func([]interface{}) effect{
	"nop": func(param []interface{}) effect {
		return effect{
			condition: always,
			effect:    skip}
	},
	"increase_megacredit": func(param []interface{}) effect {
		return effect{
			condition: always,
			effect: func(d belongsToPlayer) {
				d.getPlayer().resources[int(megacredit)] += param[0].(int)
			}}
	}}
