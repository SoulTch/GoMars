package core

// Enchant interface.
type Enchant interface {
	enable()
	disable()
}

// Enchantable .
type Enchantable interface {
	getEnchant() map[string]int
}
