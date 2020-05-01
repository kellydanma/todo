package todo

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAdd(t *testing.T) {
	entries := []struct {
		list  List
		tasks []string
	}{
		{list: List{}, tasks: []string{""}},
		{list: List{}, tasks: []string{"clean bathroom"}},
		{list: List{}, tasks: []string{"clean bathroom", "clean kitchen"}},
	}

	for _, want := range entries {
		for _, task := range want.tasks {
			want.list.Add(task)
			require.Equal(t, want.tasks[0], want.list[0].Task)
		}
	}
}

func TestComplete(t *testing.T) {
	l := []item{
		item{Task: "task 1"},
		item{Task: "task 2"},
		item{Task: "task 3"},
	}

	entries := []struct {
		list     List
		complete int
	}{
		{list: l, complete: 1},
		{list: l, complete: 2},
		{list: l, complete: 1}, // already completed
	}

	for _, want := range entries {
		want.list.Complete(want.complete)
		require.Equal(t, want.list[want.complete-1].Complete, true)
	}
}
