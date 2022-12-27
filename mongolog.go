package mongolog

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewHookConnectionString(cs string, db string, coll string) (*hook, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(cs))
	if err != nil {
		return nil, err
	}

	return &hook{c: client.Database(db).Collection(coll)}, nil
}

/**
* Hook struct for Logrus hook interface
**/
type hook struct {
	c *mongo.Collection
}

func (h *hook) Fire(entry *logrus.Entry) error {
	data := make(logrus.Fields)
	data["level"] = entry.Level.String()
	data["time"] = entry.Time
	data["message"] = entry.Message

	for k, v := range entry.Data {
		if errData, isError := v.(error); logrus.ErrorKey == k && v != nil && isError {
			data[k] = errData.Error()
		} else {
			data[k] = v
		}
	}
	_, err := h.c.InsertOne(context.TODO(), data)

	if err != nil {
		return fmt.Errorf("failed to save log: %v", err)
	}

	return nil
}

func (h *hook) Levels() []logrus.Level {
	return logrus.AllLevels
}
