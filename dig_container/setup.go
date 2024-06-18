package digcontainer

import (
	apiservice "myProj/api_service"
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
	container.Provide(services.NewUserService)
	return container
}
