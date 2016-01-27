package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/benchapman/redis-broker/broker"
	"github.com/pivotal-cf/brokerapi"
	"github.com/pivotal-golang/lager"
)

var ipAddress string

func init() {
	flag.StringVar(&ipAddress, "ip", "", "Exposed IP address of the local machine")
	flag.Parse()
}

func main() {
	if ipAddress == "" {
		fmt.Println("IP address must be set")
		os.Exit(1)
	}

	serviceBroker := broker.New([16]string{}, ipAddress)
	logger := lager.NewLogger("redis-service-broker")
	credentials := brokerapi.BrokerCredentials{
		Username: "admin",
		Password: "admin",
	}

	brokerAPI := brokerapi.New(&serviceBroker, logger, credentials)
	fmt.Println("Listening on port 3000")
	http.Handle("/", brokerAPI)
	http.ListenAndServe(":3000", nil)
}
