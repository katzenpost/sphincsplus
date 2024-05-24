//go:build cgo && shake_a64 && (!haraka_aesni && !sha2_avx2 && !shake_avx2 && !ref)
package params

var ImplementationName = "shake-a64"
