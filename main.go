package main

import (
	"fmt"
	"net/http"

	"github.com/benchapman/redis-broker/broker"
	"github.com/pivotal-cf/brokerapi"
	"github.com/pivotal-golang/lager"
)

func main() {
	serviceBroker := new(broker.RedisService)
	logger := lager.NewLogger("redis-service-broker")
	credentials := brokerapi.BrokerCredentials{
		Username: "admin",
		Password: "admin",
	}

	brokerAPI := brokerapi.New(serviceBroker, logger, credentials)
	fmt.Println("Listening on port 3000")
	http.Handle("/", brokerAPI)
	http.ListenAndServe(":3000", nil)
}
