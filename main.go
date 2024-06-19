package main

import (
	"flag"

	apiservice "myProj/api_service"
	"myProj/dao"
	digcontainer "myProj/dig_container"
	log "myProj/sdk/log"
)

func main() {

	migrate := flag.Bool("migrate", false, "Run database migrations")
	flag.Parse()
	if *migrate {
		dao.MigrateDB()
		return
	}

	digc := digcontainer.NewDigC()
	err := digc.Invoke(func(api apiservice.ApiServiceInterface) {
		api.Run(":2333")

	})
	if err != nil {
		log.Fatal(err)
	}

}
