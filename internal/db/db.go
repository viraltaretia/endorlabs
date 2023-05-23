package db

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"myapp/internal/object"
	"myapp/pkg/errors"
	"myapp/pkg/logging"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

var logger logging.Logger

type DB struct {
	client *redis.Client
}

// Gets environment variable if available else default value
func getEnv(envVar string, defaultVal string) string {
	ret := os.Getenv(envVar)
	if ret == "" {
		return defaultVal
	}
	return ret
}

func NewDB() (*DB, error) {
	redisURL := getEnv("REDIS_ADDR", "redis:6379")
	redisPass := getEnv("REDIS_PASSWORD", "pass@123")

	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisURL,
		Password: redisPass,
		DB:       0,
	})

	// Test the connection to Redis
	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %s", err)
	}

	return &DB{
		client: redisClient,
	}, nil
}

func (d *DB) Close() error {
	return d.client.Close()
}

func (d *DB) Store(ctx context.Context, obj object.Object) error {
	id := uuid.New().String()
	key := strings.Split(obj.GetKind(), ".")[1] + ":" + obj.GetName() + ":" + id
	obj.SetID(key)

	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		return fmt.Errorf("failed to serialize object: %s", err)
	}

	if err := d.client.Set(ctx, key, jsonBytes, 0).Err(); err != nil {
		return fmt.Errorf("failed to save object to Redis: %s", err)
	}

	return nil
}

// GetObjectByID will retrieve the object with the provided ID.
func (r *DB) GetObjectByID(ctx context.Context, id string) (object.Object, error) {
	var obj object.Object

	bytes, err := r.client.Get(ctx, id).Result()
	if err == redis.Nil {
		return obj, errors.ErrObjectNotFound
	} else if err != nil {
		return obj, err
	}

	key := strings.Split(id, ":")
	obj, err = object.CreateObj(key[0])
	if err != nil {
		logger.Errorf("error while creating object of type: %s", key[0])
		return obj, err
	}

	if err := json.Unmarshal([]byte(bytes), &obj); err != nil {
		return obj, err
	}

	return obj, nil

}

// GetObjectByName will retrieve the object with the given name.
func (r *DB) GetObjectByName(ctx context.Context, name string) ([]object.Object, error) {
	matchString := fmt.Sprintf("*:%s:*", name)
	keys, err := r.client.Keys(ctx, matchString).Result()
	if err != nil {
		logger.Errorf("error while getting keys")
	}
	values, _ := r.client.MGet(ctx, keys...).Result()
	var objects []object.Object
	for index, val := range values {
		var obj object.Object

		kind := strings.Split(keys[index], ":")[0]
		obj, err = object.CreateObj(kind)
		if err != nil {
			return nil, err
		}
		err := json.Unmarshal([]byte(val.(string)), &obj)
		if err != nil {
			logger.Errorf("error while unmarshalling data ", err)
			return nil, err
		}
		objects = append(objects, obj)
	}
	return objects, nil
}

// ListObjects will return a list of all objects of the given kind.
func (r *DB) ListObjects(ctx context.Context, kind string) ([]object.Object, error) {
	var objects []object.Object
	iter := r.client.Scan(ctx, 0, fmt.Sprintf("%s:*", kind), 0).Iterator()
	for iter.Next(ctx) {
		bytes, err := r.client.Get(ctx, iter.Val()).Bytes()
		if err != nil {
			return nil, err
		}

		obj, err := object.CreateObj(kind)
		if err != nil {
			return nil, err
		}
		if err := json.Unmarshal(bytes, &obj); err != nil {
			return nil, err
		}

		objects = append(objects, obj)
	}

	if err := iter.Err(); err != nil {
		return nil, err
	}
	return objects, nil
}

// DeleteObject will delete the object with the given ID.
func (r *DB) DeleteObject(ctx context.Context, id string) error {
	if err := r.client.Del(ctx, id).Err(); err != nil {
		return err
	}

	return nil
}
