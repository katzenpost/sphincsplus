//go:build ((linux && amd64) || (darwin && amd64) || (darwin && arm64) || (windows && amd64)) && (sphincs_haraka_128f || sphincs_haraka_128s || sphincs_haraka_192f || sphincs_haraka_192s || sphincs_haraka_256f || sphincs_haraka_256s || sphincs_sha2_128f || sphincs_sha2_128s || sphincs_sha2_192f || sphincs_sha2_192s || sphincs_sha2_256f || sphincs_sha2_256s || sphincs_shake_128f || sphincs_shake_128s || sphincs_shake_192f || sphincs_shake_192s || sphincs_shake_256f || sphincs_shake_256s)

package sphincsplus

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSchemeName(t *testing.T) {
	t.Logf("SPHINCS+ scheme parameters: %s", SchemeName())
	require.Equal(t, SignatureName, SchemeName())
}