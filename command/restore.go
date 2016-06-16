package command

import (
	"github.com/BSick7/envoy/envoy"
	"os"
	"strings"
)

type RestoreCommand struct {
	Meta
}

func (c *RestoreCommand) Run(args []string) int {
	flags := c.Meta.FlagSet("restore")
	var httpAddr string
	flags.StringVar(&httpAddr, "http-address", "", "Consul HTTP Address")
	var token string
	flags.StringVar(&token, "token", "", "Consul ACL Token")
	if err := flags.Parse(args); err != nil {
		return 1
	}

	conf := envoy.Config{
		HttpAddress: httpAddr,
		AclToken:    token,
	}

	var filename string
	extra := flags.Args()
	if len(extra) > 0 {
		filename = extra[0]
	}

	if filename == "" {
		if err := envoy.Restore(conf, os.Stdin); err != nil {
			c.Ui.Error(err.Error())
			return 1
		}
	} else {
		file, err := os.Open(filename)
		if err != nil {
			c.Ui.Error(err.Error())
			return 1
		}
		defer file.Close()

		if err := envoy.Restore(conf, file); err != nil {
			c.Ui.Error(err.Error())
			return 1
		}
	}

	return 0
}

func (c *RestoreCommand) Synopsis() string {
	return "Restores consul k/v store"
}

func (c *RestoreCommand) Help() string {
	helpText := `
Usage: envoy restore [options] [filename]

Self Options:

  -http-address=127.0.0.1:8500  Consul HTTP Address.
                                Fallback to CONSUL_HTTP_ADDR.
                                Fallback to 127.0.0.1:8500.

  -token=acl-token              Issue consul API calls with ACL Token.
                                Fallback to CONSUL_HTTP_TOKEN.

  [filename]                    Optional tar.gz.
                                Will use stdin if missing.
`
	return strings.TrimSpace(helpText)
}
