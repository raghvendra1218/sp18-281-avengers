package main

import (
	//"encoding/json"
	"encoding/json"
	"fmt"
	//"log"

	"github.com/go-redis/redis"
)

func main() {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	employee1 := &Employee{Firstname: "John", Lastname: "Smith", Gender: "Male", Age: "27", Salary: 60000}
	data, err := json.Marshal(&employee1)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
	err = client.Set("key", data, 0).Err()
	if err != nil {
		panic(err)
	}

}