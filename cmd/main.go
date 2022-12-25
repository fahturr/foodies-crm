package main

import (
	"fmt"
	"foodies-crm/apps"
	"foodies-crm/config"
	"foodies-crm/pkg/database"
)

func main() {
	err := config.LoadConfig("./cmd/.env")
	if err != nil {
		panic(err)
	}

	db := database.NewDatabase()

	port := config.GetString(config.APP_PORT)
	factory := apps.NewRouterFactory(db)
	router, err := factory.Create(apps.ROUTER_GIN, fmt.Sprintf(":%s", port))
	if err != nil {
		panic(err)
	}

	executor := apps.NewRouterExecutor()
	err = executor.Execute(router)
	if err != nil {
		panic(err)
	}
}
