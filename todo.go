package todo

import (
	"fmt"
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
		return fmt.Errorf("item %d does not exist,"+
			"it was not marked as complete", i)
	}
	(*l)[i-1].Complete = true
	(*l)[i-1].CompletedAt = time.Now()
	return nil
}

// Delete removes a ToDo item from the list at index i.
func (l *List) Delete(i int) error {
	if i <= 0 || i > len(*l) {
		return fmt.Errorf("item %d does not exit,"+
			"delete failed", i)
	}
	lst := *l
	*l = append(lst[:i-1], lst[i:]...)
	return nil
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
