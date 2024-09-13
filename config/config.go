package config

import "os"

func GetRapidApiKey() string {
	return os.Getenv("RAPID_API_KEY")
}

func GetDynamoDbTableName() string {
	return os.Getenv("TABLE_NAME")
}

func GetRedisAddr() string {
	return os.Getenv("REDIS_ADDR")
}

func GetRedisPassword() string {
	return os.Getenv("REDIS_PASSWORD")
}
