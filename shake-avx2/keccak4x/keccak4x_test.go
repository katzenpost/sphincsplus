//go:build cgo && ((linux && amd64) || (darwin && amd64) || (darwin && arm64) || (windows && amd64)) && (sphincs_shake_128f || sphincs_shake_128s || sphincs_shake_192f || sphincs_shake_192s || sphincs_shake_256f || sphincs_shake_256s)

package keccak4x

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestString(t *testing.T) {
	require.True(t, Keccak4x_enabled())
}
