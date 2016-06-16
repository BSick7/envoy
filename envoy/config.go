package envoy

import "github.com/hashicorp/consul/api"

type Config struct {
	HttpAddress string
	AclToken    string
}

func (c Config) NewConsulClient() (*api.Client, error) {
	config := api.DefaultConfig()
	config.Address = c.HttpAddress
	config.Token = c.AclToken
	return api.NewClient(config)
}
