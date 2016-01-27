package broker

import (
	"fmt"

	"github.com/pivotal-cf/brokerapi"
)

//RedisDatabase ...
type RedisDatabase struct {
	index              int
	assignedInstanceID string
}

//RedisService ...
type RedisService struct {
	databases [16]RedisDatabase
}

//New ...
func New(names [16]string) (broker RedisService) {
	for i := range broker.databases {
		broker.databases[i] = RedisDatabase{
			i,
			names[i],
		}
	}
	return
}

//Services ...
func (*RedisService) Services() (service []brokerapi.Service) {
	return []brokerapi.Service{
		brokerapi.Service{
			ID:            "29e7fc30-5941-44c9-8062-a8ec169f0202",
			Name:          "Shared Redis",
			Description:   "",
			Bindable:      true,
			Tags:          []string{"pivotal", "cf"},
			PlanUpdatable: false,
			Plans: []brokerapi.ServicePlan{
				brokerapi.ServicePlan{
					ID:          "be6701cc-43da-401d-8bad-06220108e4d9",
					Name:        "Basic",
					Description: "",
				},
			},
		},
	}
}

//Provision ...
func (broker *RedisService) Provision(
	instanceID string,
	details brokerapi.ProvisionDetails,
	asyncAllowed bool,
) (brokerapi.ProvisionedServiceSpec, error) {
	var (
		indexOfDatabase int
		err             error
	)

	if indexOfDatabase, err = broker.indexOfDatabase(""); err != nil {
		return brokerapi.ProvisionedServiceSpec{}, err
	}

	database := &broker.databases[indexOfDatabase]
	database.assignedInstanceID = instanceID
	return brokerapi.ProvisionedServiceSpec{}, nil
}

//LastOperation ...
func (*RedisService) LastOperation(instanceID string) (service brokerapi.LastOperation, err error) {
	return
}

//Deprovision ...
func (broker *RedisService) Deprovision(
	instanceID string,
	details brokerapi.DeprovisionDetails,
	asyncAllowed bool,
) (brokerapi.IsAsync, error) {
	var (
		indexOfDatabase int
		err             error
	)

	asyncFalse := brokerapi.IsAsync(false)

	if indexOfDatabase, err = broker.indexOfDatabase(instanceID); err != nil {
		return asyncFalse, err
	}

	database := &broker.databases[indexOfDatabase]
	database.assignedInstanceID = ""
	return asyncFalse, err
}

//Bind ...
func (*RedisService) Bind(instanceID, bindingID string, details brokerapi.BindDetails) (service brokerapi.Binding, err error) {
	return
}

//Unbind ...
func (*RedisService) Unbind(instanceID, bindingID string, details brokerapi.UnbindDetails) (err error) {
	return
}

//Update ...
func (*RedisService) Update(instanceID string, details brokerapi.UpdateDetails, asyncAllowed bool) (service brokerapi.IsAsync, err error) {
	return
}

func (broker *RedisService) indexOfDatabase(assignedInstanceID string) (int, error) {
	for i, database := range broker.databases {
		if database.assignedInstanceID == assignedInstanceID {
			return i, nil
		}
	}
	return 0, fmt.Errorf("Database instance ID %s not found.", assignedInstanceID)
}
