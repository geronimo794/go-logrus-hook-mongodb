# Mongolog: Simple Logrus Hook for MongoDB
Inspired by https://github.com/weekface/mgorus
## Usage

```go
package main

import (
	"github.com/sirupsen/logrus"
	"github.com/geronimo794/go-mongolog"
)

func main() {
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
```

