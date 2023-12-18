package evs

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFrameString(t *testing.T) {
	frame := Frame{
		Line:     23,
		File:     "test.go",
		Function: "TestFrameString",
	}
	require.Equal(t, "TestFrameString [test.go:23]", frame.String())
}
