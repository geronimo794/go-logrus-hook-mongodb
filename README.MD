# Mongolog
Simple [Logrus](https://github.com/sirupsen/logrus) <img src="http://i.imgur.com/hTeVwmJ.png" width="20" height="20" alt=":walrus:" class="emoji" title=":walrus:"/> hook for MongoDB.

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



_This repo inspired by [mgorus](https://github.com/weekface/mgorus)_


