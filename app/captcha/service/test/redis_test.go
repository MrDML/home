package test

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"testing"
)

func TestExampleClient(t *testing.T) {

	var ctx = context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "leon", // no password set
		DB:       0,  // use default DB
	})

	err := rdb.Set(ctx, "name", "leon", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("name", val)

	val2, err := rdb.Get(ctx, "age").Result()
	if err == redis.Nil {
		fmt.Println("age does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("age", val2)
	}
	// Output: key value
	// key2 does not exist


}