package mongolog_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/geronimo794/go-mongolog"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
func TestNewHook_Success(t *testing.T) {
	var mongoHost = os.Getenv("MONGO_DB_HOST_NATIVE")
	var mongoUsername = os.Getenv("MONGO_DB_USERNAME")
	var mongoPassword = os.Getenv("MONGO_DB_PASSWORD")
	var mongoPort = os.Getenv("MONGO_DB_PORT")

	log := logrus.New()
	hook, err := mongolog.NewHook(mongoHost, mongoPort, mongoUsername, mongoPassword, db, coll)
	if err == nil {
		log.Hooks.Add(hook)
	} else {
		fmt.Print(err)
	}

	log.WithFields(logrus.Fields{
		"name": "Ach Rozikin",
	}).Error("TestNewHook_Success")

	log.Warn("Warning log: TestNewHook_Success")
	log.Trace("Trace log: TestNewHook_Success")
	log.Error("Error log: TestNewHook_Success")
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
	}).Error("TestNewHookConnectionString_Success")

	log.Warn("Warning log")
	log.Trace("Trace log")
	log.Error("Error log")
}
func TestNewHookClient_Success(t *testing.T) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connectionString))
	if err != nil {
		fmt.Print(err)
	}

	log := logrus.New()
	hook, err := mongolog.NewHookClient(client, db, coll)
	if err == nil {
		log.Hooks.Add(hook)
	} else {
		fmt.Print(err)
	}

	log.WithFields(logrus.Fields{
		"name": "Ach Rozikin",
	}).Error("TestNewHookClient_Success")

	log.Warn("Warning log: TestNewHookClient_Success")
	log.Trace("Trace log: TestNewHookClient_Success")
	log.Error("Error log: TestNewHookClient_Success")
}
func TestNewHookDatabase_Success(t *testing.T) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connectionString))
	if err != nil {
		fmt.Print(err)
	}

	log := logrus.New()
	hook, err := mongolog.NewHookDatabase(client.Database(db), coll)
	if err == nil {
		log.Hooks.Add(hook)
	} else {
		fmt.Print(err)
	}

	log.WithFields(logrus.Fields{
		"name": "Ach Rozikin",
	}).Error("TestNewHookDatabase_Success")

	log.Warn("Warning log: TestNewHookDatabase_Success")
	log.Trace("Trace log: TestNewHookDatabase_Success")
	log.Error("Error log: TestNewHookDatabase_Success")
}
func TestNewHookCollection_Success(t *testing.T) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connectionString))
	if err != nil {
		fmt.Print(err)
	}

	log := logrus.New()
	hook, err := mongolog.NewHookCollection(client.Database(db).Collection(coll))
	if err == nil {
		log.Hooks.Add(hook)
	} else {
		fmt.Print(err)
	}

	log.WithFields(logrus.Fields{
		"name": "Ach Rozikin",
	}).Error("TestNewHookCollection_Success")

	log.Warn("Warning log: TestNewHookCollection_Success")
	log.Trace("Trace log: TestNewHookCollection_Success")
	log.Error("Error log: TestNewHookCollection_Success")
}
