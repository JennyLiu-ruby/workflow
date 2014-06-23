package etcdutils

import (
	"testing"
	"github.com/deis/deis/tests/utils"
	"github.com/coreos/go-etcd/etcd"
	"fmt"
	"strings"
)

type EtcdHandle struct {
	Dirs []string
	Keys []string
	C    *etcd.Client
}

func getetcdClient(port string) *etcd.Client {
	IPAddress :=  utils.GetHostIpAddress()
	machines := []string{"http://"+IPAddress+":"+ port}
	c := etcd.NewClient(machines)
	return c
}

func InitetcdValues(setdir, setkeys []string, port string) *EtcdHandle {
	cli := getetcdClient(port)
	controllerHandle := new(EtcdHandle)
	controllerHandle.Dirs = setdir
	controllerHandle.Keys = setkeys
	controllerHandle.C = cli
	fmt.Println("Etcd client initialized")
	return controllerHandle
}

func SetEtcdValues(t *testing.T, keys []string, values []string, c *etcd.Client) {
	for i, key := range keys {
		_, err := c.Set(key, values[i], 0)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func Publishvalues(t *testing.T, ecli *EtcdHandle) {
	fmt.Println("Publishing ETCD Key values")
	for _, dir := range ecli.Dirs {
		_, err := ecli.C.SetDir(dir, 0)
		if err != nil {
			t.Fatal(err)
		}
	}
	for _, key := range ecli.Keys {
		switch true {
		case (strings.Contains(key, "host")):
			_, err := ecli.C.Set(key, "172.17.8.100", 0)
			if err != nil {
				t.Fatal(err)
			}
		case (strings.Contains(key, "port")):
			_, err := ecli.C.Set(key, "10881", 0)
			if err != nil {
				t.Fatal(err)
			}
		default:
			_, err := ecli.C.Set(key, "deis", 0)
			if err != nil {
				t.Fatal(err)
			}
		}
	}
}
