package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	PublicHost 				string
	Port 	   				string
	DBUser	   				string
	DBPassword 				string
	DBAddress  				string
	DBName	   				string
	JWTExpirationInSeconds 	int64
	JWTSecret 				string
}

var Envs = initConfig()

func initConfig() Config {
	return Config {
		PublicHost: getEnv("PUBLIC_HOST", "devtype.cvqwmo6u455l.us-east-1.rds.amazonaws.com"),
		Port: getEnv("PORT", "3306"),
		DBUser: getEnv("DB_USER", "admin"),
		DBPassword: getEnv("DB_PASSWORD", "00zL}L71w,I0"),
		DBAddress: fmt.Sprintf("%s:%s", getEnv("DB_HOST", "devtype.cvqwmo6u455l.us-east-1.rds.amazonaws.com"), getEnv("DB_PORT", "3306")),
		DBName: getEnv("DB_NAME", "devtype"),
		JWTSecret: getEnv("JWT_SECRET", "secret"),
		JWTExpirationInSeconds: getEnvAsInt("JWT_EXP", 3600 * 24 * 7),
	}
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if val, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return fallback
		}
		return i
	}
	return fallback
}