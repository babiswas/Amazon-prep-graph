package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Person struct {
	Name string `json:"name"`
	Id   int64  `json:"id"`
}

func main() {
	var person [5]Person
	for i := 0; i < 5; i++ {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Enter the name")
		scanner.Scan()
		name := scanner.Text()
		fmt.Println("Enter the id")
		scanner.Scan()
		id, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		person[i] = Person{name, id}
	}
	redis := make(map[string]string)
	for _, p := range person {
		data, err := json.Marshal(p)
		if err != nil {
			fmt.Println("Error occured", err)
		}
		redis[p.Name] = string(data)
	}
	fmt.Println(redis)
        for key,value:=range redis{
           fmt.Println(key,value)
        }
}
