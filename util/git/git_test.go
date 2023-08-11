package git

import (
	"testing"

	"gotest.tools/assert"
)

func TestRemoveSuffix(t *testing.T) {
	data := [][]string{
		{"hello.git", ".git", "hello"},
		{"hello", ".git", "hello"},
		{".git", ".git", ""},
	}
	for _, table := range data {
		result := removeSuffix(table[0], table[1])
		assert.Equal(t, table[2], result)
	}
}
