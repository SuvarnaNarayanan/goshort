package main

import (
	"log"
	"net"
	"time"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"goshort/m/v2/api"
	"goshort/m/v2/api/impl"
	"goshort/m/v2/config"
	"goshort/m/v2/internal/db"
	"goshort/m/v2/internal/redis"

)

func main() {
	cfg := config.LoadConfig()

	dbConn, err := db.NewDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	redisClient, err := redis.NewClient(cfg.RedisURL)
	if err != nil {
		log.Fatalf("Failed to initialize Redis: %v", err)
	}

	server := impl.NewGoShortApiServer(dbConn , redisClient)
	e := echo.New()
	e.Use(middleware.Logger() , 
	      middleware.Recover() , 
		  middleware.TimeoutWithConfig(middleware.TimeoutConfig{
        Timeout: 10 * time.Second,
	}))
	serverHandler := api.NewStrictHandler(server, nil)
	api.RegisterHandlers(e, serverHandler)

	e.Logger.Fatal(e.Start(net.JoinHostPort("localhost", cfg.Port)))
}
