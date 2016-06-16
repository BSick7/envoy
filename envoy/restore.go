package envoy

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"github.com/hashicorp/consul/api"
	"io"
	"io/ioutil"
)

func Restore(config Config, r io.Reader) error {
	gr, err := gzip.NewReader(r)
	if err != nil {
		return fmt.Errorf("error opening backup file: %s", err)
	}
	defer gr.Close()
	tr := tar.NewReader(gr)

	consul, err := config.NewConsulClient()
	if err != nil {
		return fmt.Errorf("error creating consul client: %s", err)
	}

	for {
		header, err := tr.Next()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return fmt.Errorf("error reading backup file: %s", err)
		}

		switch header.Typeflag {
		case tar.TypeDir:
			continue
		case tar.TypeReg:
			fallthrough
		case tar.TypeRegA:
			b, err := ioutil.ReadAll(r)
			if err != nil {
				return fmt.Errorf("error reading backup value: %s", err)
			}
			_, err = consul.KV().Put(&api.KVPair{
				Key:   header.Name,
				Value: b,
			}, nil)
			if err != nil {
				return fmt.Errorf("error writing k/v to consul: %s", err)
			}
		}
	}
}
