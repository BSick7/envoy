package command

import (
	"github.com/BSick7/envoy/envoy"
	"os"
	"strings"
)

type BackupCommand struct {
	Meta
}

func (c *BackupCommand) Run(args []string) int {
	flags := c.Meta.FlagSet("backup")
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
		if err := envoy.Backup(conf, os.Stdout); err != nil {
			c.Ui.Error(err.Error())
			return 1
		}
	} else {
		file, err := os.Create(filename)
		if err != nil {
			c.Ui.Error(err.Error())
			return 1
		}
		defer file.Close()

		if err := envoy.Backup(conf, file); err != nil {
			c.Ui.Error(err.Error())
			return 1
		}
	}

	return 0
}

func (c *BackupCommand) Synopsis() string {
	return "Backup consul k/v store"
}

func (c *BackupCommand) Help() string {
	helpText := `
Usage: envoy backup [options] [filename]

Self Options:

  -http-address=127.0.0.1:8500  Consul HTTP Address.
                                Fallback to CONSUL_HTTP_ADDR.
                                Fallback to 127.0.0.1:8500.

  -token=acl-token              Issue consul API calls with ACL Token.
                                Fallback to CONSUL_HTTP_TOKEN.

  [filename]                    Optional filename to emit the tar.gz.
                                Will emit to stdout if missing.
`
	return strings.TrimSpace(helpText)
}
