package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/kellydanma/todo"
)

var (
	flagAdd      *bool
	flagList     *bool
	flagComplete *int

	// default file name
	todoFileName = ".todo.json"
)

func main() {
	// parse flags
	flagAdd = flag.Bool("add", false, "add task to todo list")
	flagList = flag.Bool("list", false, "list all tasks")
	flagComplete = flag.Int("complete", 0, "mark task as complete")
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
	case *flagAdd:
		// add item to todo list
		task, err := getTask(os.Stdin, flag.Args()...)
		if err != nil {
			log.Fatalf(err.Error())
		}
		l.Add(task)
		if err := l.Save(todoFileName); err != nil {
			log.Fatalf(err.Error())
		}
	default:
		// invalid flag
		fmt.Fprintln(os.Stderr, "invalid option")
		os.Exit(1)
	}
}

// getTask obtains new tasks from arguments or STDIN.
func getTask(r io.Reader, args ...string) (string, error) {
	// if args were provided, concatenate them & return
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	// otherwise, scan for single input line through reader
	s := bufio.NewScanner(r)
	s.Scan()
	if err := s.Err(); err != nil {
		return "", err
	}
	if len(s.Text()) == 0 {
		return "", fmt.Errorf("task cannot be blank")
	}
	return s.Text(), nil
}
