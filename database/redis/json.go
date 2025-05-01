package redis

import (
	"context"
	"encoding/json"
	"github.com/aitu-leetcode-site/core/errors"
	"reflect"
	"time"
)

type JsonClient interface {
	SetJsonObject(ctx context.Context, key string, value interface{}) error
	SetJsonObjectWithExpire(
		ctx context.Context, key string, value interface{}, expiration time.Duration,
	) error
	MSetJsonObject(ctx context.Context, keys []string, values []interface{}) error
	GetJsonObject(ctx context.Context, key string, dst interface{}) error
	// MGetJsonObjects - dst must be a slice of pointers
	MGetJsonObjects(ctx context.Context, dst interface{}, keys ...string) error
}

func (c *client) SetJsonObject(ctx context.Context, key string, value interface{}) error {
	bytes, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return c.redisCli.Set(ctx, c.genKey(key), bytes, 0).Err()
}

func (c *client) SetJsonObjectWithExpire(
	ctx context.Context, key string, value interface{}, expiration time.Duration,
) error {
	bytes, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return c.redisCli.Set(ctx, c.genKey(key), bytes, expiration).Err()
}

func (c *client) MSetJsonObject(ctx context.Context, keys []string, values []interface{}) error {
	if len(keys) != len(values) {
		return errors.New("length mismatch")
	}
	for i, k := range keys {
		err := c.redisCli.Set(ctx, c.genKey(k), values[i], 0).Err()
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *client) GetJsonObject(ctx context.Context, key string, dst interface{}) error {
	bytes, err := c.redisCli.Get(ctx, c.genKey(key)).Bytes()
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, dst)
}

func (c *client) MGetJsonObjects(ctx context.Context, dst interface{}, keys ...string) error {
	v := reflect.ValueOf(dst)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Slice {
		return errors.New("dst must be a pointer to a slice")
	}
	slice := v.Elem()
	elemType := slice.Type().Elem()
	for _, key := range keys {
		data, err := c.redisCli.Get(ctx, c.genKey(key)).Bytes()
		if err != nil {
			return err
		}
		elemPtr := reflect.New(elemType.Elem())
		err = json.Unmarshal(data, elemPtr.Interface())
		if err != nil {
			return err
		}
		slice.Set(reflect.Append(slice, elemPtr.Convert(elemType)))
	}
	return nil
}
