package todo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

// Add appends a ToDo item to the List.
func (l *List) Add(task string) {
	t := item{
		Task:      task,
		CreatedAt: time.Now(),
	}
	*l = append(*l, t)
}

// Complete marks a ToDo item as complete at index i.
func (l *List) Complete(i int) error {
	if i <= 0 || i > len(*l) {
		return fmt.Errorf("item %d does not exist, "+
			"it was not marked as complete", i)
	}
	(*l)[i-1].Complete = true
	(*l)[i-1].CompletedAt = time.Now()
	return nil
}

// Delete removes a ToDo item from the list at index i.
func (l *List) Delete(i int) error {
	if i <= 0 || i > len(*l) {
		return fmt.Errorf("item %d does not exist, "+
			"delete failed", i)
	}
	lst := *l
	*l = append(lst[:i-1], lst[i:]...)
	return nil
}

// Save encodes the List as JSON using the provided file name.
func (l *List) Save(filename string) error {
	lst, err := json.MarshalIndent(*l, "", " ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, lst, 0644)
}

// Get decodes a JSON file into a List.
func (l *List) Get(filename string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // no errors for non-existant lists
		}
		return err
	}
	if len(file) == 0 {
		return nil // no errors for empty lists
	}
	return json.Unmarshal(file, l)
}

// List represents a list of ToDo items.
type List []item

// item represents a ToDo item.
type item struct {
	Task        string
	Complete    bool
	CreatedAt   time.Time
	CompletedAt time.Time
}
