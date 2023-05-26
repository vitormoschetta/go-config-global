package config

// variável global (singleton) com garantia que o valor não será alterado
// contras: valor tem que ser definido no momento da declaração
const (
	QueryProduct = "SELECT * FROM product"
	QueryOrder   = "SELECT * FROM order"
)
