package env

import (
	"encoding/json"
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type EnvConfig struct {
	AppPort        string `split_words:"true"`
	DatabasesHost  string `split_words:"true"`
	DatabasePort   string `split_words:"true"`
	DatabaseSchema string `split_words:"true"`
	DatabaseName   string `split_words:"true"`
	DatabaseUser   string `split_words:"true"`
	DatabasePass   string `split_words:"true"`
	DataDir        string `split_words:"true"`
	WorkerCount    int    `split_words:"true"`
	QueueSize      int    `split_words:"true"`
	MigrationUrl   string `split_words:"true"`
	DatabaseSource string `split_words:"true"`
}

var EnvCfg EnvConfig

func ProcssEnv() error {
	if err := envconfig.Process("", &EnvCfg); err != nil {
		return err
	}
	return nil
}

func InitEnv(appenv *string) {
	var environment = string(*appenv)
	var configFile string
	if environment == "docker" {
		configFile = "pkg/env/docker.env"
	} else {
		configFile = "pkg/env/.env"
	}
	if err := godotenv.Load(configFile); err != nil {
		log.Fatal("unable to load .env")
	}
	if err := ProcssEnv(); err != nil {
		log.Fatal("unable to load .env")
	}
	cfgStr, _ := json.Marshal(EnvCfg)
	log.Println(string(cfgStr))
}
