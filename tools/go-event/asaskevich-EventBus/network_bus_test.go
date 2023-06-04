package EventBus

import (
	"testing"

	"github.com/asaskevich/EventBus"
)

func TestNewServer(t *testing.T) {
	serverBus := EventBus.NewServer(":2010", "/_server_bus_", EventBus.New())
	_ = serverBus.Start()
	if serverBus == nil {
		t.Log("New server EventBus not created!")
		t.Fail()
	}
	select {}
	//serverBus.Stop()
}

func TestNewClient(t *testing.T) {
	clientBus := EventBus.NewClient(":2015", "/_client_bus_", EventBus.New())
	_ = clientBus.Start()
	if clientBus == nil {
		t.Log("New client EventBus not created!")
		t.Fail()
	}
	clientBus.Stop()
}

func TestServerPublish(t *testing.T) {
	serverBus := EventBus.NewServer(":2020", "/_server_bus_b", EventBus.New())
	serverBus.Start()

	fn := func(a int) {
		if a != 10 {
			t.Fail()
		}
	}

	clientBus := EventBus.NewClient(":2025", "/_client_bus_b", EventBus.New())
	clientBus.Start()

	clientBus.Subscribe("topic", fn, ":2020", "/_server_bus_b")

	serverBus.EventBus().Publish("topic", 10)

	clientBus.Stop()
	serverBus.Stop()
}

func TestNetworkBus(t *testing.T) {
	networkBusA := EventBus.NewNetworkBus(":2035", "/_net_bus_A")
	networkBusA.Start()

	networkBusB := EventBus.NewNetworkBus(":2030", "/_net_bus_B")
	networkBusB.Start()

	fnA := func(a int) {
		if a != 10 {
			t.Fail()
		}
	}
	networkBusA.Subscribe("topic-A", fnA, ":2030", "/_net_bus_B")
	networkBusB.EventBus().Publish("topic-A", 10)

	fnB := func(a int) {
		if a != 20 {
			t.Fail()
		}
	}
	networkBusB.Subscribe("topic-B", fnB, ":2035", "/_net_bus_A")
	networkBusA.EventBus().Publish("topic-B", 20)

	networkBusA.Stop()
	networkBusB.Stop()
}
