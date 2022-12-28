package mongolog_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/geronimo794/go-mongolog"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var connectionString string
var db string
var coll string

func init() {
	// Load .env data put to os environment
	if err := godotenv.Load(); err != nil {
		panic("No .env file found")
	}

	/**
	* Connection variable definition
	**/
	// Connection string
	connectionString = os.Getenv("MONGO_DB_CONNECTION_STRING")

	// Per variable definition
	var mongoHost = os.Getenv("MONGO_DB_HOST_NATIVE")
	var mongoUsername = os.Getenv("MONGO_DB_USERNAME")
	var mongoPassword = os.Getenv("MONGO_DB_PASSWORD")
	var mongoPort = os.Getenv("MONGO_DB_PORT")

	// If mongo db connection string is empty then create the connection from variabel
	if len(connectionString) == 0 {
		connectionString = "mongodb://" + mongoUsername + ":" + mongoPassword + "@" + mongoHost + ":" + mongoPort + "/?retryWrites=true&w=majority"
	}

	// Set test db and collection
	db = "db-test"
	coll = "collection-test"
}

func TestNewHookConnectionString_Success(t *testing.T) {
	log := logrus.New()
	hook, err := mongolog.NewHookConnectionString(connectionString, db, coll)
	if err == nil {
		log.Hooks.Add(hook)
	} else {
		fmt.Print(err)
	}

	log.WithFields(logrus.Fields{
		"name": "Ach Rozikin",
	})

	log.Warn("Warning log")
	log.Trace("Trace log")
	log.Error("Error log")

}
