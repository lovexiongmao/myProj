package digcontainer

import (
	apiservice "myProj/api_service"

	"go.uber.org/dig"
)

func NewDigC() *dig.Container {
	DigC := dig.New()
	DigC.Provide(apiservice.NewApiService)
	return DigC
}
