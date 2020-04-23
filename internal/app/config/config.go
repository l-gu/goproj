package config

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/l-gu/goproject/internal/app/logger"
)

// default values
const DEFAULT_WEB_DIR = "web"
const DEFAULT_WEB_PORT = "8080"
const DEFAULT_DAO_TYPE = 1
const DEFAULT_BOLT_FILE_NAME = "bolt.db"

type Configuration struct {
	WebPort      string
	WebDir       string
	DaoType      int
	BoltFileName string
}

func NewConfiguration() Configuration {
	conf := Configuration{}
	conf.WebPort = DEFAULT_WEB_PORT
	conf.WebDir = DEFAULT_WEB_DIR
	conf.DaoType = DEFAULT_DAO_TYPE
	conf.BoltFileName = DEFAULT_BOLT_FILE_NAME
	return conf
}

const LOG_PREFIX = "config"

var configuration = NewConfiguration()

// initialize configuration from the given JSON file
func InitConfig(configFile string) error {
	logger.Log(LOG_PREFIX, "Init configuration from file :", configFile)
	file, err := os.Open(configFile)
	if err != nil {
		return errors.New("Cannot open configuration file '" + configFile + "'")
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&configuration)
	if err != nil {
		return errors.New("Cannot decode configuration file '" + configFile + "'")
	}

	logger.Log(LOG_PREFIX, "Configuration initialized :")
	logger.Log(LOG_PREFIX, " . WebPort      :", configuration.WebPort)
	logger.Log(LOG_PREFIX, " . WebDir       :", configuration.WebDir)
	logger.Log(LOG_PREFIX, " . DaoType      :", configuration.DaoType)
	logger.Log(LOG_PREFIX, " . BoltFileName :", configuration.BoltFileName)

	return nil
}

func GetWebPort() string {
	return configuration.WebPort
}
func GetWebDir() string {
	return configuration.WebDir
}
func GetDaoType() int {
	return configuration.DaoType
}
func GetBoltFileName() string {
	return configuration.BoltFileName
}
