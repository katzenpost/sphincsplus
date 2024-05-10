//go:build cgo && shake_avx2 && (!haraka_aesni && !sha2_avx2 && !shake_a64 && !ref)
package params

var ImplementationName = "shake-avx2"
