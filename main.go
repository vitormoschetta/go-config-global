package main

import (
	"fmt"
	"test/config"
)

func main() {
	fmt.Println("Constantes: ")
	fmt.Println(config.QueryProduct)
	fmt.Println(config.QueryOrder)
	fmt.Println()

	fmt.Println("Variavel Global: ")
	fmt.Println(config.QueryMap[config.QueryProductKey])
	fmt.Println(config.QueryMap[config.QueryOrderKey])
	fmt.Println()

	fmt.Println("Singleton: ")
	configuration := config.NewConfiguration("SELECT * FROM product", "SELECT * FROM order")
	queryMap := configuration.GetQueryMap()
	fmt.Println(queryMap[config.QueryProductKey])
	fmt.Println(queryMap[config.QueryOrderKey])
}
