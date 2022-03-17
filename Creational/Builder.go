package Creational

import "fmt"

const (
	MonolithProcess     = 1
	MicroserviceProcess = 2
)

type Director struct {
	Builder BuilderProcess
}

type BuilderProcess interface {
	SetFramework(string) BuilderProcess
	SetDatabaseType(string) BuilderProcess
	SetAsyncQueue(string) BuilderProcess
	GetApplicationDetails() string
}

type Microservice struct {
	Framework  string
	Database   string
	AsyncQueue string
}

type Monolith struct {
	Framework  string
	Database   string
	AsyncQueue string
}

func (service *Microservice) SetFramework(framework string) BuilderProcess {
	service.Framework = framework
	return service
}

func (service *Microservice) SetDatabaseType(db string) BuilderProcess {
	service.Database = db
	return service
}

func (service *Microservice) SetAsyncQueue(queue string) BuilderProcess {
	service.AsyncQueue = queue
	return service
}

func (service *Microservice) GetApplicationDetails() string {
	return fmt.Sprintf("The microservice application is written in %v framework with the support of %v database and %v async queue", service.Framework, service.Database, service.AsyncQueue)
}

func (service *Monolith) SetFramework(framework string) BuilderProcess {
	service.Framework = framework
	return service
}

func (service *Monolith) SetDatabaseType(db string) BuilderProcess {
	service.Database = db
	return service
}

func (service *Monolith) SetAsyncQueue(queue string) BuilderProcess {
	service.AsyncQueue = queue
	return service
}

func (service *Monolith) GetApplicationDetails() string {
	return fmt.Sprintf("The monolith application is written in %v framework with the support of %v database and %v async queue", service.Framework, service.Database, service.AsyncQueue)
}

func (director *Director) SetBuilder(builderCode int) {
	switch builderCode {
	case MonolithProcess:
		director.Builder = new(Monolith)
	case MicroserviceProcess:
		director.Builder = new(Microservice)
	default:
		director.Builder = nil
	}
}
