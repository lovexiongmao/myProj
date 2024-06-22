package digcontainer

import (
	apiservice "myProj/apiservice"
	"myProj/clients"
	"myProj/conf"
	"myProj/dao"
	"myProj/services"

	"go.uber.org/dig"
)

func NewDigC() *dig.Container {
	container := dig.New()
	container.Provide(apiservice.NewApiService)
	container.Provide(dao.NewDB)
	container.Provide(conf.LoadConfig)
	container.Provide(clients.NewUserHandler)
	container.Provide(clients.NewAccountHandler)
	container.Provide(services.NewUserService)
	container.Provide(services.NewAccountService)
	return container
}
