package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/kellydanma/todo"
)

const todoFileName = ".todo.json"

func main() {
	l := &todo.List{}
	if err := l.Get(todoFileName); err != nil {
		log.Fatalf(err.Error())
	}

	switch {
	case len(os.Args) == 1:
		// list todo items
		for i, item := range *l {
			fmt.Printf("%d. %s\n", i+1, item.Task)
		}
	default:
		// add new item to list
		item := strings.Join(os.Args[1:], " ")
		l.Add(item)
		if err := l.Save(todoFileName); err != nil {
			log.Fatalf(err.Error())
		}
	}
}
