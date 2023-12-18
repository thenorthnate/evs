package evs

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetStack(t *testing.T) {
	stack := GetStack(0)
	stackStr := stack.String()
	require.Contains(t, stackStr, "TestGetStack")
}
