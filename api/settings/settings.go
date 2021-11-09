package settings

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
)

type DBConnectionSettings struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Database string `json:"database"`
}

type AppSettings struct {
	DbConnection DBConnectionSettings `json:"dbConnection"`
	Port         string               `json:"port"`
}

func GetSettings() AppSettings {
	env := GetEnvironment()
	file := GetSettingsFile(env)
	return ParseSettings(file)
}

func ParseSettings(file string) AppSettings {
	settingsFile, err := os.Open(file)
	if err != nil {
		panic(err.Error())
	}
	rawData, err := ioutil.ReadAll(settingsFile)
	if err != nil {
		panic(err.Error())
	}
	var settings AppSettings
	err = json.Unmarshal(rawData, &settings)
	if err != nil {
		panic(err.Error())
	}
	return settings
}

func GetSettingsFile(env Environment) string {
	switch env {
	case Dev:
		return "settings.dev.json"
	case Test:
		return "settings.test.json"
	case Prod:
		return "settings.prod.json"
	default:
		return "settings.dev.json"
	}
}

// TODO: Move all this env stuff to its own file
type Environment int64

const (
	Dev  Environment = 0
	Test Environment = 1
	Prod Environment = 2
)

func GetEnvironment() Environment {
	const buildEnv string = "CUSTODIRE_ENV"
	switch env := strings.ToLower(os.Getenv(buildEnv)); env {
	case "dev":
		return Dev
	case "test":
		return Test
	case "prod":
		return Prod
	default:
		return Dev
	}
}
