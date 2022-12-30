# Mongolog
Simple [Logrus](https://github.com/sirupsen/logrus) <img src="http://i.imgur.com/hTeVwmJ.png" width="20" height="20" alt=":walrus:" class="emoji" title=":walrus:"/> hook for MongoDB.

## Usage
1. Use with existing connection string (*You can use connection string from MongoDB atlas*)
	```go
	package main

	import (
		"github.com/geronimo794/go-mongolog"
		"github.com/sirupsen/logrus"
	)

	func main() {
		// Create mongolog hook: Atlas mongodb service
		var connectionString = "mongodb+srv://<username>:<password>@<host>/?retryWrites=true&w=majority"
		var db = "db"
		var coll = "coll"
		hook, err := mongolog.NewHookConnectionString(connectionString, db, coll)

		// New logrus instance
		log := logrus.New()
		if err == nil {
			log.Hooks.Add(hook)
		} else {
			log.Panic(err)
		}

		// Create log example
		log.WithFields(logrus.Fields{
			"name": "Ach Rozikin",
		}).Error("Error log")
		log.Warn("Warning log")
		log.Error("Error log")
	}
	```
1. Use with mongoDB configuration host, port, username, password.
	```go
	package main

	import (
		"github.com/geronimo794/go-mongolog"
		"github.com/sirupsen/logrus"
	)

	func main() {
		// Create mongolog hook
		var mongoHost = "localhost"
		var mongoUsername = ""
		var mongoPassword = ""
		var mongoPort = "27017"
		var db = "db"
		var coll = "coll"
		hook, err := mongolog.NewHook(mongoHost, mongoPort, mongoUsername, mongoPassword, db, coll)

		// Create logrus instance
		log := logrus.New()
		if err == nil {
			log.Hooks.Add(hook)
		} else {
			log.Panic(err)
		}

		// Create log example
		log.WithFields(logrus.Fields{
			"name": "Ach Rozikin",
		}).Error("Error log")
		log.Warn("Warning log")
		log.Error("Error log")

	}	
	```
1. Use with pre existiong _*mongo.Client_
	```go
	package main

	import (
		"context"
		"log"
		"time"

		"github.com/geronimo794/go-mongolog"
		"github.com/sirupsen/logrus"
		"go.mongodb.org/mongo-driver/mongo"
		"go.mongodb.org/mongo-driver/mongo/options"
	)

	func main() {
		// Create mongoDB client with connection string
		// All this option is optional
		serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
		clientOptions := options.Client().
			ApplyURI("mongodb+srv://<username>:<password>@<host>/?retryWrites=true&w=majority").
			SetServerAPIOptions(serverAPIOptions)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		client, err := mongo.Connect(ctx, clientOptions)
		if err != nil {
			log.Fatal(err)
		}

		// Create mongolog with client variable
		var db = "db"
		var coll = "coll-test"
		hook, err := mongolog.NewHookClient(client, db, coll)

		// New logrus instance and hook
		log := logrus.New()
		if err == nil {
			log.Hooks.Add(hook)
		} else {
			log.Panic(err)
		}

		// Create log example
		log.WithFields(logrus.Fields{
			"name": "Ach Rozikin",
		}).Error("Error log")
		log.Warn("Warning log")
		log.Error("Error log")
	}	
	```
1. Use with pre existiong _*mongo.Database_
	```go
	package main

	import (
		"context"
		"log"
		"time"

		"github.com/geronimo794/go-mongolog"
		"github.com/sirupsen/logrus"
		"go.mongodb.org/mongo-driver/mongo"
		"go.mongodb.org/mongo-driver/mongo/options"
	)

	func main() {
		// Create mongoDB client with connection string
		// All this option is optional
		serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
		clientOptions := options.Client().
			ApplyURI("mongodb+srv://<username>:<password>@<host>/?retryWrites=true&w=majority").
			SetServerAPIOptions(serverAPIOptions)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		client, err := mongo.Connect(ctx, clientOptions)
		if err != nil {
			log.Fatal(err)
		}
		clientDb := client.Database("db")

		// Create mongolog with client database variable
		var coll = "coll-test"
		hook, err := mongolog.NewHookDatabase(clientDb, coll)

		// New logrus instance and hook
		log := logrus.New()
		if err == nil {
			log.Hooks.Add(hook)
		} else {
			log.Panic(err)
		}

		// Create log example
		log.WithFields(logrus.Fields{
			"name": "Ach Rozikin",
		}).Error("Error log")
		log.Warn("Warning log")
		log.Error("Error log")

	}	
	```
1. Use with pre existiong _*mongo.Collection_
	```go
	package main

	import (
		"context"
		"log"
		"time"

		"github.com/geronimo794/go-mongolog"
		"github.com/sirupsen/logrus"
		"go.mongodb.org/mongo-driver/mongo"
		"go.mongodb.org/mongo-driver/mongo/options"
	)

	func main() {

		// Create mongoDB client with connection string
		// All this option is optional
		serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
		clientOptions := options.Client().
			ApplyURI("mongodb+srv://<username>:<password>@<host>/?retryWrites=true&w=majority").
			SetServerAPIOptions(serverAPIOptions)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		client, err := mongo.Connect(ctx, clientOptions)
		if err != nil {
			log.Fatal(err)
		}
		clientDbCollection := client.Database("db").Collection("coll-test")

		// Create mongolog with client database collection variable
		hook, err := mongolog.NewHookCollection(clientDbCollection)

		// New logrus instance and hook
		log := logrus.New()
		if err == nil {
			log.Hooks.Add(hook)
		} else {
			log.Panic(err)
		}

		// Create log example
		log.WithFields(logrus.Fields{
			"name": "Ach Rozikin",
		}).Error("Error log")
		log.Warn("Warning log")
		log.Error("Error log")
	}
	```

## Planned Update
1. **NewHook(...)**: Create hook with mongoDB configuration like host, port, username, and password. **[DONE]**
1. **NewHookClient(...)**: Create hook with _*mongo.Client_ input **[DONE]**
1. **NewHookDatabase(...)**: Create hook with _*mongo.Database_ input **[DONE]**
1. **NewHookCollection(...)**: Create hook with _*mongo.Collection_ input **[DONE]**

## * Bonus
Create docker container of *mongoDB* and *mongo-express*
1. Create .env file and change value inside parenties
	```env
	# If it's docker container, name with the service name
	# If on host : localhost
	# If on docker : mongo-db
	MONGO_DB_HOST=mongo-db
	MONGO_DB_HOST_NATIVE=localhost #If running go application running directly without dokcer

	MONGO_DB_NAME={DB-NAME}
	MONGO_DB_COLLECTION_NAME={COLL-NAME}


	MONGO_DB_USERNAME={DB-USERNAME}
	MONGO_DB_PASSWORD={DB-PASSWORD}
	MONGO_DB_PORT={DB-PORT}
	```
1. Create docker-compose.yml
	```yml
	version: '3.1'
	services:
	# Creating database service with mongo-db
	mongo-db:
	  image: mongo
	  environment:
	    MONGO_INITDB_ROOT_USERNAME: ${MONGO_DB_USERNAME}
	    MONGO_INITDB_ROOT_PASSWORD: ${MONGO_DB_PASSWORD}
	  ports: 
	    - ${MONGO_DB_PORT}:27017
	  volumes:
	    - mongo-db-data:/data/db
	    - mongo-db-config:/data/configdb
	    - mongo-db-mongo:/etc/mongo

	# Creating database explorer with mongo-express
	mongo-db-explorer:
	  depends_on:
	    - mongo-db
	  restart: unless-stopped
	  image: mongo-express
	  ports:
	    - 8081:8081
	  environment:
	    ME_CONFIG_MONGODB_ADMINUSERNAME: ${MONGO_DB_USERNAME}
	    ME_CONFIG_MONGODB_ADMINPASSWORD: ${MONGO_DB_PASSWORD}
	    ME_CONFIG_MONGODB_URL: mongodb://${MONGO_DB_USERNAME}:${MONGO_DB_PASSWORD}@${MONGO_DB_HOST}:${MONGO_DB_PORT}/


	volumes:
	  mongo-db-data:
	  mongo-db-config:
	  mongo-db-mongo:

	```
1. Run docker compose
	```sh
	docker compose up -d --build
	```
1. You can view your *MongoDB* data with *mongo-express* on localhost:8081
1. And you can use it with this *mongolog* hook with shared *.env* configuration file.

.

_This repo inspired by [mgorus](https://github.com/weekface/mgorus)_


