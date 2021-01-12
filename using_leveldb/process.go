package main

import (
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"strings"
)

func ProcessInput(inp string, db *leveldb.DB) bool {
	inp = strings.Replace(inp, "  ", " ", -1)
	split := strings.Split(inp, " ")
	if split[0] == "SET" {
		ProcessSet(split, db)
		return true
	} else if split[0] == "GET" {
		ProcessGet(split, db)
		return true
	} else if split[0] == "DELETE" {
		ProcessDelete(split, db)
		return true
	} else {
		return false
	}
}

func ProcessSet(split []string, db *leveldb.DB) bool {
	if db == nil {
		fmt.Println("no db open")
		return false
	}
	if len(split) != 3 {
		fmt.Println("invalid SET command")
		return false
	}
	err := db.Put([]byte(split[1]), []byte(split[2]), nil)
	if err != nil {
		panic("error setting value")
	}
	fmt.Println("set")
	return true
}

func ProcessGet(split []string, db *leveldb.DB) bool {
	if db == nil {
		fmt.Println("no db open")
		return false
	}
	if len(split) != 2 {
		fmt.Println("invalid GET command")
		return false
	}
	data, err := db.Get([]byte(split[1]), nil)
	if err != nil {
		fmt.Println("error getting value")
		return false
	}
	fmt.Println(string(data))
	return true
}

func ProcessDelete(split []string, db *leveldb.DB) bool {
	if db == nil {
		fmt.Println("no db open")
		return false
	}
	if len(split) != 2 {
		fmt.Println("invalid DELETE command")
		return false
	}
	err := db.Delete([]byte(split[1]), nil)
	if err != nil {
		fmt.Println("error deleting value")
		return false
	}
	fmt.Println("delete")
	return true
}
