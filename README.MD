# Mongolog
Simple [Logrus](https://github.com/sirupsen/logrus) <img src="http://i.imgur.com/hTeVwmJ.png" width="20" height="20" alt=":walrus:" class="emoji" title=":walrus:"/> hook for MongoDB.<br />
Main feature: **Async Feature**, **Write Timeout**, **Context Setting**, and **Failover File Logging**
## Download and install
```bash
$ go get -u github.com/geronimo794/go-mongolog
```
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
## Feature
1. **Async Feature**(_default=false_): Create the process of writing to mongodb asyncronushly. To activate the feature, call **SetIsAsync** with parameter **true**. If you use this feaature, make sure don't close the application before the process is done _(You can use time.Sleep(XXX * time.Second) to prevent application from close it's self)_
	```go
	log := logrus.New()
	hook, err := mongolog.NewHook(mongoHost, mongoPort, mongoUsername, mongoPassword, db, coll)
	if err == nil {
		hook.SetIsAsync(true)
		log.Hooks.Add(hook)
	} else {
		fmt.Print(err)
	}
	```
1. **Write Timeout**(_default=0_): Add write timeout feature to set max timeout when writing data to the mongodb instance. You can set it via **SetWriteTimeout** method.
	```go
	log := logrus.New()
	hook, err := mongolog.NewHook(mongoHost, mongoPort, mongoUsername, mongoPassword, db, coll)
	if err == nil {
		hook.SetWriteTimeout(1 * time.Second)
		log.Hooks.Add(hook)
	} else {
		fmt.Print(err)
	}
	```
1. **Context Setting**(_default=context.Background()_): You can set custom context to pass it to the hook. You can set it via **SetContext** method.
	```go
	log := logrus.New()
	hook, err := mongolog.NewHook(mongoHost, mongoPort, mongoUsername, mongoPassword, db, coll)
	if err == nil {
		// Create context with timeout
		ctx, ctxCancelFunc := context.WithTimeout(context.TODO(), 1*time.Second)
		hook.SetContext(ctx)
		log.Hooks.Add(hook)

		defer ctxCancelFunc()
	} else {
		fmt.Print(err)
	}
	```
1. **Failover File Logging**(_default=nil_): Set the failover file to save the log when the application failed to save the log on mongodb. You can set it via **SetFailoverFilePath** method.
	```go
	log := logrus.New()
	hook, err := mongolog.NewHook(mongoHost, mongoPort, mongoUsername, mongoPassword, db, coll)
	if err == nil {
		err = hook.SetFailoverFilePath("mongolog-failover.log")
		if err == nil {
			log.Hooks.Add(hook)
		} else {
			fmt.Print(err)
		}
	} else {
		fmt.Print(err)
	}
	```
*You can use **SetIsAsync** and **SetWriteTimeout** combination to prevent your application for goroutine leak. <br />
*If you are using **SetWriteTimeout** and **Context Setting** at once. **Context Setting** value will be ignored by the hook process.<br />
*Be aware when using the **Failover File Logging**, keep this log is clean and clear (It's mean mongodb always success to save the log)

## Planned Update
1. **Failover File Logging**: Create failover mechanisme when failed to write on mongodb database. **[DONE]**

## License
_MIT_
