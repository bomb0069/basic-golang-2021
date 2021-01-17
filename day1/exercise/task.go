package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type cli struct {
	users []User
}

func (c *cli) addNew(u User) {
	c.users = append(c.users, u)
}

type User struct {
	Id   int    `json:"userId`
	Name string `json:"userName"`
	Age  string `json:"userAge"`
}

func main() {
	fileName := "input.json"

	command := os.Args[1]

	switch command {
	case "add":
		com := cli{readStructFromFile(fileName)}
		name := os.Args[2]
		age := os.Args[3]
		newId := 1
		if len(com.users) > 0 {
			newId = com.users[len(com.users)-1].Id + 1
		}
		com.addNew(User{newId, name, age})
		writeJSONtoFile(fileName, com.users)
	case "list":
		fmt.Printf("%v", readJSONFromFile(fileName))

	default:
		fmt.Printf("Command Not Found")
	}
}

func writeJSONtoFile(fileName string, users []User) {
	json := structToJSON(users)

	_ = ioutil.WriteFile(fileName, json, 0644)
}

func readJSONFromFile(fileName string) string {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var stringFile string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		stringFile = stringFile + scanner.Text() + "\n"
	}
	return stringFile
}

func readStructFromFile(fileName string) []User {
	stringFile := readJSONFromFile(fileName)
	return getStructFromJSON(stringFile)
}

func getStructFromJSON(jsonString string) []User {
	var users []User
	err := json.Unmarshal([]byte(jsonString), &users)
	if err != nil {
		fmt.Println("Error ", err)
	}

	return users
}

func structToJSON(users []User) []byte {
	b, err := json.MarshalIndent(users, "", "	")
	if err != nil {
		fmt.Println("Error ", err)
	}

	return b
}
