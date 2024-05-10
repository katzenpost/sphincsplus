//go:build ((linux && amd64) || (darwin && amd64) || (darwin && arm64) || (windows && amd64)) || (sphincs_haraka_128f || sphincs_haraka_128s || sphincs_haraka_192f || sphincs_haraka_192s || sphincs_haraka_256f || sphincs_haraka_256s || sphincs_sha2_128f || sphincs_sha2_128s || sphincs_sha2_192f || sphincs_sha2_192s || sphincs_sha2_256f || sphincs_sha2_256s || sphincs_shake_128f || sphincs_shake_128s || sphincs_shake_192f || sphincs_shake_192s || sphincs_shake_256f || sphincs_shake_256s)

package tests

import (
	sphincsplus "github.com/ioerror/sphincsplus/ref"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSignVerifyVector(t *testing.T) {
	privKey1, pubKey1 := sphincsplus.NewKeypair()
	message1 := []byte("i am a message")
	signature1 := privKey1.Sign(message1)
	require.True(t, pubKey1.Verify(signature1, message1))

	t.Logf("SPHINCS+ scheme parameters: %s", sphincsplus.Name())
	t.Logf("SPHINCS+ scheme implementation: %s", sphincsplus.Implementation())
	t.Logf("priv key %x", privKey1.Bytes())
	t.Logf("pub key %x", pubKey1.Bytes())
	t.Logf("sig %x", signature1)

	require.True(t, pubKey1.Verify(signature1, message1))

}
