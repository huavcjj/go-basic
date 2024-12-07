package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

func string(ctx context.Context, client *redis.Client) {
	key := "name"
	value := "golang"
	err := client.Set(ctx, key, value, 1*time.Second).Err()
	checkError(err)

	client.Expire(ctx, key, 3*time.Second)
	time.Sleep(2 * time.Second)

	v2, err := client.Get(ctx, key).Result()
	checkError(err)
	fmt.Println(v2)

	client.Del(ctx, key)
}

func list(ctx context.Context, client *redis.Client) {
	key := "ids"
	values := []interface{}{1, 2, 3, "golang"}
	err := client.RPush(ctx, key, values...).Err()
	checkError(err)

	v2, err := client.LRange(ctx, key, 0, -1).Result()
	checkError(err)
	fmt.Println(v2)

	client.Del(ctx, key)
}

func hashTable(ctx context.Context, client *redis.Client) {
	err := client.HSet(ctx, "user1", "Name", "golang", "Age", 10, "Height", 180.5).Err()
	checkError(err)
	err = client.HSet(ctx, "user2", "Name", "python", "Age", 20, "Height", 170.0).Err()
	checkError(err)

	age, err := client.HGet(ctx, "user2", "Age").Result()
	checkError(err)
	fmt.Println(age)

	for field, value := range client.HGetAll(ctx, "user1").Val() {
		fmt.Println(field, value)
	}

	client.Del(ctx, "user1", "user2")

}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	ctx := context.TODO()

	string(ctx, client)
	list(ctx, client)
	hashTable(ctx, client)

}

func checkError(err error) {
	if err != nil {
		if err == redis.Nil {
			fmt.Println("key not found")
		} else {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
