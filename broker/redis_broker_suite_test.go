package broker_test

import (
	"strconv"

	"github.com/benchapman/redis-broker/broker"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pivotal-cf/brokerapi"

	"testing"
)

var _ = Describe(".RedisService", func() {
	var (
		redisBroker broker.RedisService
	)

	JustBeforeEach(func() {
		redisBroker = broker.New([16]string{})
	})

	Describe(".Services", func() {
		It("works", func() {
			Expect(redisBroker.Services()[0].Name).To(Equal("Shared Redis"))
		})
	})

	Describe(".Provision", func() {
		Context("if there are databases available", func() {
			It("should return a nil error", func() {
				_, err := redisBroker.Provision(
					"Pikachu",
					brokerapi.ProvisionDetails{},
					false,
				)
				Expect(err).To(BeNil())
			})
		})

		Context("if there are no databases available", func() {
			JustBeforeEach(func() {
				var fullArray [16]string
				for i := range fullArray {
					fullArray[i] = strconv.Itoa(i)
				}
				redisBroker = broker.New(fullArray)
			})

			It("should return an error if it can't provision a redis db", func() {
				_, err := redisBroker.Provision(
					"Pikachu",
					brokerapi.ProvisionDetails{},
					false,
				)
				Expect(err).ToNot(BeNil())
			})
		})
	})

	Describe(".Deprovision", func() {
		Context("when a service does not exist", func() {
			It("should return an error", func() {
				_, err := redisBroker.Deprovision(
					"Pikachu",
					brokerapi.DeprovisionDetails{},
					false,
				)
				Expect(err).ToNot(BeNil())
			})
		})

		Context("given a service has been provisioned", func() {
			JustBeforeEach(func() {
				redisBroker = broker.New([16]string{"Pikachu"})
			})

			It("should return no error", func() {
				_, err := redisBroker.Deprovision(
					"Pikachu",
					brokerapi.DeprovisionDetails{},
					false,
				)
				Expect(err).To(BeNil())
			})
		})
	})
})

var _ = Describe("Functional test", func() {
	It("should provision and then deprovision", func() {
		redisBroker := broker.New([16]string{})

		_, err := redisBroker.Provision(
			"Pikachu",
			brokerapi.ProvisionDetails{},
			false,
		)
		Expect(err).To(BeNil())

		_, err = redisBroker.Deprovision(
			"Pikachu",
			brokerapi.DeprovisionDetails{},
			false,
		)
		Expect(err).To(BeNil())

		_, err = redisBroker.Deprovision(
			"Pikachu",
			brokerapi.DeprovisionDetails{},
			false,
		)
		Expect(err).ToNot(BeNil())
	})
})

func TestRedisBroker(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RedisBroker Suite")
}
