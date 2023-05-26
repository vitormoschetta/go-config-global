package config

// objeto singleton com garantia que o valor não será alterado
type configuration struct {
	queryMap map[string]string
}

var instance *configuration

func NewConfiguration(queryProduct, queryOrder string) *configuration {
	if instance == nil {
		instance = &configuration{
			queryMap: map[string]string{
				QueryProductKey: queryProduct,
				QueryOrderKey:   queryOrder,
			},
		}
	}
	return instance
}

func (c *configuration) GetQueryMap() map[string]string {
	return c.queryMap
}
