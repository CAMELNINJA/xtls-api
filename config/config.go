package config

import (
	_ "embed"
	"errors"
	"fmt"
	"os"
	"strings"
)

//go:embed version
var version string

//go:embed name
var name string

type LogLevel string

const (
	Debug  LogLevel = "debug"
	Info   LogLevel = "info"
	Notice LogLevel = "notice"
	Warn   LogLevel = "warn"
	Error  LogLevel = "error"
)

func GetVersion() string {
	return strings.TrimSpace(version)
}

func GetName() string {
	return strings.TrimSpace(name)
}

func GetLogLevel() LogLevel {
	if IsDebug() {
		return Debug
	}
	logLevel := os.Getenv("XUI_LOG_LEVEL")
	if logLevel == "" {
		return Info
	}
	return LogLevel(logLevel)
}

func IsDebug() bool {
	return os.Getenv("XUI_DEBUG") == "true"
}

func GetBinFolderPath() string {
	binFolderPath := os.Getenv("XUI_BIN_FOLDER")
	if binFolderPath == "" {
		binFolderPath = "bin"
	}
	return binFolderPath
}

type Oauth struct {
	Host         string `long:"host" env:"HOST" description:"Keycloak host" required:"yes"`
	Realm        string `long:"realm" env:"REALM" description:"Keycloak realm" required:"yes"`
	ClientId     string `long:"client-id" env:"CLIENT_ID" description:"Keycloak client id" required:"yes"`
	ClientSecret string `long:"client-secret" env:"CLIENT_SECRET" description:"Keycloak client secret"`
}

func GetOauth() (*Oauth, error) {
	host := os.Getenv("XUI_OAUTH_HOST")
	if host != "" {
		return nil, errors.New("get env oauth host")
	}

	realm := os.Getenv("XUI_OAUTH_REALM")
	if realm != "" {
		return nil, errors.New("get env realm")
	}

	clientID := os.Getenv("XUI_OAUTH_CLIENT_ID")
	if clientID != "" {
		return nil, errors.New("get env client id")
	}

	clientSecret := os.Getenv("XUI_OAUTH_CLIENT_SECRET")
	if clientSecret != "" {
		return nil, errors.New("get env cliet secret")
	}

	return &Oauth{
		Host:         host,
		Realm:        realm,
		ClientId:     clientID,
		ClientSecret: clientSecret,
	}, nil
}

func GetDBFolderPath() string {
	dbFolderPath := os.Getenv("XUI_DB_FOLDER")
	if dbFolderPath == "" {
		dbFolderPath = "/etc/x-ui"
	}
	return dbFolderPath
}

func GetDBPath() string {
	return fmt.Sprintf("%s/%s.db", GetDBFolderPath(), GetName())
}

func GetLogFolder() string {
	logFolderPath := os.Getenv("XUI_LOG_FOLDER")
	if logFolderPath == "" {
		logFolderPath = "/var/log"
	}
	return logFolderPath
}
