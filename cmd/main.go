package main

import (
	"context"
	"github.com/fmo/tm-players/config"
	"github.com/fmo/tm-players/internal/adapters/cache/redis"
	"github.com/fmo/tm-players/internal/adapters/cli"
	"github.com/fmo/tm-players/internal/adapters/database/dynamodb"
	"github.com/fmo/tm-players/internal/adapters/player-data/transfermarkt"
	"github.com/fmo/tm-players/internal/application/core/api"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	environment := os.Getenv("ENVIRONMENT")
	if environment != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file")
		}
	}

	ctx := context.Background()

	cacheAdapter, err := redis.NewAdapter(config.GetRedisAddr(), config.GetRedisPassword())
	if err != nil {
		log.Fatalf("Failed to connect to redis. Error: %v", err)
	}

	playerAdapter, err := transfermarkt.NewAdapter(config.GetRapidApiKey(), cacheAdapter)
	if err != nil {
		log.Fatalf("Failed to connect to transfermarkt. Error: %v", err)
	}

	dbAdapter, err := dynamodb.NewAdapter(config.GetDynamoDbTableName())
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}

	application := api.NewApplication(playerAdapter, dbAdapter)
	cliAdapter := cli.NewAdapter(application)
	cliAdapter.Run(ctx)
}
