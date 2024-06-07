package main

import (
	"fmt"
	"log/slog"

	"github.com/saladin2098/month3/lesson11/voting_gateway/api"
	"github.com/saladin2098/month3/lesson11/voting_gateway/api/handlers"
	"github.com/saladin2098/month3/lesson11/voting_gateway/config"
)

func main() {
	cfg := config.Load()
	handler := handlers.NewHandler()
	r := api.NewGin(handler)
	fmt.Printf("Server started on port %s\n", cfg.HTTPPort)
	err := r.Run(cfg.HTTPPort)
	if err!= nil {
        slog.Error("error",err)
    }
}
