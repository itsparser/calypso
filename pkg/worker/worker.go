package worker

import (
	"go.uber.org/dig"

	"github.com/workfoxes/calypso/pkg/config"
	"github.com/workfoxes/calypso/pkg/log"
)

type Worker struct {
	Name      string `json:"name"`
	container *dig.Container
	config    *config.Config
}

// New : This will create new Worker object
func New(name string) *Worker {
	worker := &Worker{Name: name}
	DefaultProviders(worker)
	return worker
}

// DefaultProviders : will provide all the default provider in the server start
func DefaultProviders(worker *Worker) {
	worker.AddProvider(config.GetConfig)
	worker.AddProvider(log.Init)
	//worker.AddProvider(_redis.New)
	//worker.Invoker(func(l _logger.Logger) {
	//	log.L = l
	//})
	worker.Invoker(func(_config *config.Config) {
		config.C = _config
	})
	//worker.Invoker(func(r *redis.Client) {
	//	_redis.R = r
	//})
}

// AddProvider : This will add new provider to the server container
func (worker *Worker) AddProvider(constructor interface{}, opts ...dig.ProvideOption) {
	err := worker.container.Provide(constructor, opts...)
	if err != nil {
		panic(err)
	}
}

// Invoker : This will add new provider to the server container
func (worker *Worker) Invoker(function interface{}, opts ...dig.ProvideOption) {
	err := worker.container.Invoke(function)
	if err != nil {
		panic(err)
	}
}
