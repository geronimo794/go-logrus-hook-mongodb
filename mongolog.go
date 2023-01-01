package mongolog

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewHook(h string, p string, u string, pass string, db string, coll string) (*hook, error) {
	cs := "mongodb://" + u + ":" + pass + "@" + h + ":" + p + "/?retryWrites=true&w=majority"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(cs))
	if err != nil {
		return nil, err
	}

	return newHookStruct(client.Database(db).Collection(coll)), nil
}
func NewHookConnectionString(cs string, db string, coll string) (*hook, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(cs))
	if err != nil {
		return nil, err
	}

	return newHookStruct(client.Database(db).Collection(coll)), nil
}
func NewHookClient(client *mongo.Client, db string, coll string) (*hook, error) {
	return newHookStruct(client.Database(db).Collection(coll)), nil

}
func NewHookDatabase(database *mongo.Database, coll string) (*hook, error) {
	return newHookStruct(database.Collection(coll)), nil
}
func NewHookCollection(collection *mongo.Collection) (*hook, error) {
	return newHookStruct(collection), nil
}

/**
* Hook struct for Logrus hook interface
**/
type hook struct {
	c            *mongo.Collection
	isAsync      bool
	writeTimeout time.Duration
	ctx          context.Context
	failoverFile *os.File
}

// Function to create struct with default value
func newHookStruct(C *mongo.Collection) *hook {
	return &hook{
		c:            C,
		isAsync:      false,
		writeTimeout: 0,
		ctx:          context.Background(),
		failoverFile: nil,
	}
}

func (h *hook) Fire(entry *logrus.Entry) error {
	if h.isAsync {
		go h.fireProcess(entry)
		return nil
	} else {
		return h.fireProcess(entry)
	}

}
func (h *hook) Levels() []logrus.Level {
	return logrus.AllLevels
}
func (h *hook) SetIsAsync(IsAsync bool) {
	h.isAsync = IsAsync
}
func (h *hook) SetWriteTimeout(Dur time.Duration) {
	h.writeTimeout = Dur
}
func (h *hook) SetContext(Ctx context.Context) {
	h.ctx = Ctx
}
func (h *hook) SetFailoverFilePath(F string) error {
	file, err := os.OpenFile(F, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return fmt.Errorf("failed to set failover file: %v", err)
	}
	h.failoverFile = file
	return nil
}

// Private function for internal process
func (h *hook) fireProcess(entry *logrus.Entry) error {
	ctx := h.ctx

	// If write timeout greater than 0
	if h.writeTimeout > 0 {
		var ctxCancelFunc context.CancelFunc
		ctx, ctxCancelFunc = context.WithTimeout(ctx, h.writeTimeout)
		defer ctxCancelFunc()
	}

	data := make(logrus.Fields)
	data["level"] = entry.Level.String()
	data["time"] = entry.Time
	data["message"] = entry.Message
	failoverText := entry.Time.String() + "\t" + entry.Level.String() + "\t" + entry.Message

	for k, v := range entry.Data {
		if errData, isError := v.(error); logrus.ErrorKey == k && v != nil && isError {
			data[k] = errData.Error()
			failoverText += "\t" + errData.Error()
		} else {
			data[k] = v
			failoverText += "\t" + v.(string)
		}
	}
	_, err := h.c.InsertOne(ctx, data)

	if err != nil {
		// If failoverFile is setted then append to text
		if h.failoverFile != nil {
			// Write log to file
			if _, err := h.failoverFile.WriteString(failoverText + "\n"); err != nil {
				return fmt.Errorf("failed to save log to failoverfile: %v", err)
			}

			// Write mongodb error to file
			if _, err := h.failoverFile.WriteString(err.Error() + "\n"); err != nil {
				return fmt.Errorf("failed to save mongodb error to failoverfile: %v", err)
			}
		}

		return fmt.Errorf("failed to save log to mongodb: %v", err)
	}

	return nil
}
