// +build zookeeper

package serverplugin

import (
	"testing"
	"time"

	metrics "github.com/rcrowley/go-metrics"
	"github.com/smallnest/rpcx/server"
)

func TestZookeeperRegistry(t *testing.T) {
	s := server.NewServer(nil)

	r := &ZooKeeperRegisterPlugin{
		ServiceAddress:   "tcp@127.0.0.1:8972",
		ZooKeeperServers: []string{"127.0.0.1:2181"},
		BasePath:         "/rpcx_test",
		Metrics:          metrics.NewRegistry(),
		Services:         make([]string, 1),
		UpdateInterval:   time.Minute,
	}
	err := r.Start()
	if err != nil {
		t.Fatal(err)
	}
	s.Plugins.Add(r)

	s.RegisterName("Arith", new(Arith), "")
	go s.Serve("tcp", "127.0.0.1:8972")
	defer s.Close()

	if len(r.Services) != 1 {
		t.Fatal("failed to register services in zookeeper")
	}

}