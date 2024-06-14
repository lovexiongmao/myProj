package main

import (
	digcontainer "myProj/dig_container"

	apiservice "myProj/api_service"
	log "myProj/sdk/log"
)


func main() {
	digc := digcontainer.NewDigC()
	err := digc.Invoke(func(api apiservice.ApiService) {
		api.Run(":2333")

	})
	if err != nil {
		log.Fatal(err)
	}

}
