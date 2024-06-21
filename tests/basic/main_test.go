package basic

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddOne(t *testing.T) {
	assert.Equal(t, AddOne(2), 13, "AddOne(2) should return 3")

	assert.NotEqual(t, 2, 3)
	assert.Nil(t, nil)
}

func TestRequire(t *testing.T) {
	require.Equal(t, 2, 3)
	fmt.Println("Not executing")
}

func TestAssert(t *testing.T) {
	assert.Equal(t, 2, 3)
	fmt.Println("Executing")
}
