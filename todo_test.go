package todo

import (
	"io/ioutil"
	"os"
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
		list           List
		complete       int
		expectedErrStr string
	}{
		{list: l, complete: 5, expectedErrStr: "item 5 does not exist, " +
			"it was not marked as complete"},
		{list: l, complete: 1},
		{list: l, complete: 2},
		{list: l, complete: 1}, // already completed
	}

	for _, want := range entries {
		err := want.list.Complete(want.complete)
		if want.expectedErrStr == "" {
			require.NoError(t, err)
			require.Equal(t, want.list[want.complete-1].Complete, true)
		} else {
			require.EqualError(t, err, want.expectedErrStr)

		}
	}
}

func TestDelete(t *testing.T) {
	l := []item{
		item{Task: "task 1"},
		item{Task: "task 2"},
		item{Task: "task 3"},
	}

	entries := []struct {
		list           List
		delete         int
		expectedErrStr string
	}{
		{list: l, delete: 4, expectedErrStr: "item 4 does not exist, " +
			"delete failed"},
		{list: l, delete: -4, expectedErrStr: "item -4 does not exist, " +
			"delete failed"},
		{list: l, delete: 1},
		{list: l, delete: 2},
		{list: l, delete: 1},
		{list: l, delete: 1},
	}

	for _, want := range entries {
		err := want.list.Delete(want.delete)
		if want.expectedErrStr == "" {
			require.NoError(t, err)
		} else {
			require.EqualError(t, err, want.expectedErrStr)
		}
	}
}

func TestSaveGet(t *testing.T) {
	tf, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatalf("could not create temp file: %s", err)
	}
	defer os.Remove(tf.Name())

	// test Save
	l1 := List([]item{
		item{Task: "task 1", Complete: true},
		item{Task: "task 2", Complete: false},
	})
	err = l1.Save(tf.Name())
	require.NoError(t, err)

	// test Get
	l2 := List{}
	err = l2.Get(tf.Name())
	require.NoError(t, err)
	require.Equal(t, l2[0], l1[0])
	require.Equal(t, l2[1], l1[1])
}
