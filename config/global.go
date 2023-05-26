package config

const (
	QueryProductKey = "QueryProduct"
	QueryOrderKey   = "QueryOrder"
)

// variável global (singleton) sem garantia que o valor não será alterado
// contras: valor pode ser alterado em qualquer lugar do código
var QueryMap = map[string]string{
	QueryProductKey: "SELECT * FROM product",
	QueryOrderKey:   "SELECT * FROM order",
}
