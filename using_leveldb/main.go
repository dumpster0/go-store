package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/syndtr/goleveldb/leveldb"
)

func prompt() {
	fmt.Print("dummy-db>>")
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	path := os.Args[1]

	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		panic("error opening db")
	}
	defer db.Close()

	for {
		prompt()
		inp, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("error reading input")
		}

		inp = inp[:len(inp)-1]
		if inp == "EXIT" {
			break
		}
		if len(inp) > 0 {
			ProcessInput(inp, db)
		}

	}

}
