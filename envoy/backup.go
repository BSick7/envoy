package envoy

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"github.com/hashicorp/consul/api"
	"io"
)

func Backup(config Config, out io.Writer) error {
	gw := gzip.NewWriter(out)
	defer gw.Close()
	tw := tar.NewWriter(gw)
	defer tw.Close()

	consul, err := config.NewConsulClient()
	if err != nil {
		return fmt.Errorf("error creating consul client: %s", err)
	}

	kvs, _, err := consul.KV().List("/", nil)
	if err != nil {
		return fmt.Errorf("error retrieving consul data: %s", err)
	}

	for _, kv := range kvs {
		if err := writeKVToTar(tw, kv); err != nil {
			return fmt.Errorf("error writing %s: %s", kv.Key, err)
		}
	}
	return nil
}

func writeKVToTar(tw *tar.Writer, kv *api.KVPair) error {
	header := &tar.Header{}
	header.Name = kv.Key
	header.Size = int64(len(kv.Value))
	header.Mode = int64(0666)

	if err := tw.WriteHeader(header); err != nil {
		return fmt.Errorf("error writing k/v header: %s", err)
	}
	if _, err := tw.Write(kv.Value); err != nil {
		return fmt.Errorf("error writing k/v: %s", err)
	}
	return nil
}
