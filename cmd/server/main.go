package main

import (
	"fmt"
	"os"

	"stack-bm/internal/config"
	"stack-bm/internal/database"
	"stack-bm/internal/router"
)

func main() {
	config.LoadConfig()

	database.InitDB()

	r := router.SetupRouter()

	addr := fmt.Sprintf(":%s", config.AppConfig.Server.Port)
	fmt.Printf("Server starting on %s (mode: %s)\n", addr, config.AppConfig.Server.Mode)

	if err := r.Run(addr); err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
		os.Exit(1)
	}
}
