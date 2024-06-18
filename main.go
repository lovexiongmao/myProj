package main

import (
	digcontainer "myProj/dig_container"

	apiservice "myProj/api_service"
	log "myProj/sdk/log"
)

func main() {
	// migrate := flag.Bool("migrate", false, "Run database migrations")
	// flag.Parse()

	// if *migrate {
	// 	dao.MigrateDB()
	// 	return
	// }
	digc := digcontainer.NewDigC()
	err := digc.Invoke(func(api apiservice.ApiServiceInterface) {
		api.Run(":2333")

	})
	if err != nil {
		log.Fatal(err)
	}

}
