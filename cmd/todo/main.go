package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/kellydanma/todo"
)

var (
	flagTask     *string
	flagList     *bool
	flagComplete *int

	// default file name
	todoFileName = ".todo.json"
)

func main() {
	// parse flags
	flagTask = flag.String("task", "", "task to be added to the todo list")
	flagList = flag.Bool("list", false, "list all tasks")
	flagComplete = flag.Int("complete", 0, "item marked as completed")
	flag.Parse()

	l := &todo.List{}
	if f := os.Getenv("TODO_FILENAME"); f != "" {
		todoFileName = f
	}
	if err := l.Get(todoFileName); err != nil {
		log.Fatalf(err.Error())
	}

	switch {
	case *flagList:
		// list todo items
		for i, item := range *l {
			complete := "[ ]"
			if item.Complete {
				complete = "[✓]"
			}
			fmt.Printf("%2d. %-30s %s\n", i+1, item.Task, complete)
		}
	case *flagComplete > 0:
		// mark given item as complete
		if err := l.Complete(*flagComplete); err != nil {
			log.Fatalf(err.Error())
		}
		if err := l.Save(todoFileName); err != nil {
			log.Fatalf(err.Error())
		}
	case *flagTask != "":
		// add item to todo list
		l.Add(*flagTask)
		if err := l.Save(todoFileName); err != nil {
			log.Fatalf(err.Error())
		}
	default:
		// invalid flag
		fmt.Fprintln(os.Stderr, "invalid option")
		os.Exit(1)
	}
}
