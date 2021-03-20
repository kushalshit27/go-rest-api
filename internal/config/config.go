package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config ...
type Config struct {
	Port              string
	OauthRedirectURL  string
	DBURL             string
	DBName            string
	DBUser            string
	DBPassword        string
	DBPort            int
	APIPrefix         string
	DeviceSaveRoute   string
	DeviceTagRoute    string
	DeviceToken       string
	DeviceSport       string
	IsProduction      bool
	GithubID          string
	GithubSecret      string
	GithubState       string
	GithubCallbackURL string
}


// Load config from .env file
func Load(path string) *Config {
	err := godotenv.Load(path)
	if err != nil {
		log.Fatal("error with config file", err)
	}

	c := Config{}
	c.Port = os.Getenv("RS_PORT")
	if c.Port == "" {
		c.Port = "4200"
	}

	c.OauthRedirectURL = os.Getenv("RS_REDIRECT_URI")
	if c.OauthRedirectURL == "" {
		c.OauthRedirectURL = "http://localhost:8080"
	}

	c.DBURL = os.Getenv("RS_DB_HOST")
	if c.DBURL == "" {
		c.DBURL = "localhost"
	}

	mode := os.Getenv("RS_MODE")
	if mode != "dev" {
		c.IsProduction = true
	}

	dbPortStr := os.Getenv("RS_DB_PORT")
	c.DBPort, err = strconv.Atoi(dbPortStr)
	if err != nil {
		c.DBPort = 5432
	}

	c.DBName = os.Getenv("RS_DB_NAME")
	c.DBUser = os.Getenv("RS_DB_USER")
	c.DBPassword = os.Getenv("RS_DB_PASSWORD")
	c.APIPrefix = os.Getenv("RS_API_PREFIX")
	c.DeviceTagRoute = os.Getenv("RS_DEVICE_TAG_ROUTE")
	c.DeviceSaveRoute = os.Getenv("RS_DEVICE_SAVE_ROUTE")
	c.APIPrefix = os.Getenv("RS_API_PREFIX")
	c.DeviceToken = os.Getenv("RS_DEVICE_TOKEN")
	c.DeviceSport = os.Getenv("RS_DEVICE_SPORT")
	c.GithubID = os.Getenv("RS_GITHUB_CLIENT_ID")
	c.GithubSecret = os.Getenv("RS_GITHUB_SECRET")
	c.GithubState = os.Getenv("RS_GITHUB_STATE")
	c.GithubCallbackURL = os.Getenv("RS_GITHUB_CALLBACK_URL")

	return &c
}