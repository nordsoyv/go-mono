package days

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsCd(t *testing.T) {
	v, str := isCd("$ cd ..")
	assert.Equal(t, true, v)
	assert.Equal(t, "..", str)
	v, str = isCd("$ cd a")
	assert.Equal(t, true, v)
	assert.Equal(t, "a", str)
	v, str = isCd("$ cd abcde")
	assert.Equal(t, true, v)
	assert.Equal(t, "abcde", str)
	v, str = isCd("$ cd \\")
	assert.Equal(t, true, v)
	assert.Equal(t, "\\", str)
	v, str = isCd("asdfasfd")
	assert.Equal(t, false, v)
	assert.Equal(t, "", str)
}

func TestIsLs(t *testing.T) {
	v := isLs("$ cd ..")
	assert.Equal(t, false, v)
	v = isLs("$ ls")
	assert.Equal(t, true, v)
}
