package redis

import (
	"context"
	"encoding/json"
	"fmt"
	goredis "github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"

	"testing"
)

type Address struct {
	City  string
	State string
}

type User struct {
	Name string
	Age  int
	Addr Address
}

var conn = Connection{
	Address:  "localhost:6379",
	Password: "",
	Database: 0,
}

var redisInstance = NewRedis(conn)

func TestSetGet(t *testing.T) {
	user := User{
		Name: "John",
		Age:  30,
		Addr: Address{
			City:  "New York",
			State: "NY",
		},
	}

	if err := redisInstance.Set("key", user); err != nil {
		fmt.Println(err)
	}

	var value User
	if err := redisInstance.Get("key", &value); err != nil {
		fmt.Println(err)
	}

	assert.EqualValuesf(t, user.Name, value.Name, "Name should be equal")
	assert.EqualValuesf(t, user.Age, value.Age, "Age should be equal")
	assert.EqualValuesf(t, user.Addr.City, value.Addr.City, "City should be equal")
	assert.EqualValuesf(t, user.Addr.State, value.Addr.State, "State should be equal")
}

func TestPublish(_ *testing.T) {
	user := User{
		Name: "John",
		Age:  30,
		Addr: Address{
			City:  "New York",
			State: "NY",
		},
	}

	err := redisInstance.Publish("channel", user)
	if err != nil {
		return
	}
}

func TestSubscribe(_ *testing.T) {
	err := redisInstance.Subscribe(context.Background(), "channel", func(msg *goredis.Message) {
		var user User
		err := json.Unmarshal([]byte(msg.Payload), &user)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(user)
	})

	if err != nil {
		return
	}
}
