package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	Port            string
	ContractAddress string
	ChainNodeUrl    string
	RabbitBrokerUrl string
	RedisUrl        string
}

var listenerEnv *Env

func InitEnv() {
	// if we're running locally, load the env from file
	_, envSpecified := os.LookupEnv("ENV_TYPE")
	if !envSpecified {
		log.Println("Env unspecified: loading from file")
		err := godotenv.Load("run.env")
		if err != nil {
			log.Fatal("Error loading env : ", err)
		}
	}

	listenerEnv = &Env{}
	listenerEnv.Port = os.Getenv("PORT")
	listenerEnv.RabbitBrokerUrl = os.Getenv("RABBIT_BROKER_URL")
	listenerEnv.RedisUrl = os.Getenv("REDIS_URL")
}

func GetEnv() *Env {
	if listenerEnv == nil {
		InitEnv()
	}
	return listenerEnv
}
