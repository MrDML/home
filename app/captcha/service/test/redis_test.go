package test

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"math/rand"
	"strings"
	"testing"
	"time"
)

func TestExampleClient(t *testing.T) {

	s := GetRandCode(6)
	fmt.Println("code=",s)

}

func redisTest() {
	var ctx = context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "leon", // no password set
		DB:       0,      // use default DB
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




func GetRandCode(size int) string{
	if size > 9 || size < 0 {
		size = 6
	}
	rand.Seed(time.Now().Unix())
	var collection  = []string{"1","2","3","4","5","6","7","8","9"}
	res := make([]string,size)
	for i := 0; i < size; i++ {
		index := rand.Intn(size)
		res[i] = collection[index]
	}
	return strings.Join(res,"")
}
