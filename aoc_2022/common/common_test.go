package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddToFrontOfSlice(t *testing.T) {
	slice := []string{"A", "B"}
	slice = AddToFrontOfSlice(slice, []string{"C"})
	assert.EqualValues(t, []string{"C", "A", "B"}, slice)
}
func TestAddManyToFrontOfSlice(t *testing.T) {
	slice := []string{"A", "B"}
	slice = AddToFrontOfSlice(slice, []string{"C", "D", "E"})
	assert.EqualValues(t, []string{"C", "D", "E", "A", "B"}, slice)
}

func TestRemoveFromFrontOfSlice(t *testing.T) {
	slice := []string{"A", "B"}
	modifiedSlice, removed := RemoveFromFrontOfSlice(slice, 1)
	assert.EqualValues(t, []string{"B"}, modifiedSlice)
	assert.EqualValues(t, []string{"A"}, removed)
}

func TestRemoveManyFromFrontOfSlice(t *testing.T) {
	slice := []string{"A", "B", "C", "D"}
	modifiedSlice, removed := RemoveFromFrontOfSlice(slice, 3)
	assert.EqualValues(t, []string{"D"}, modifiedSlice)
	assert.EqualValues(t, []string{"A", "B", "C"}, removed)
}
