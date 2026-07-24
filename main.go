package main

import (
	"fmt"

	"github.com/dewialvi/digital-channel-monitoring/config"
)

func main() {
	cfg := config.LoadConfig()

	fmt.Println("Digital Channel Monitoring System")
	fmt.Printf("Environment: %s\n", cfg.AppEnv)
	fmt.Printf("Server akan berjalan di port: %s\n", cfg.AppPort)
}