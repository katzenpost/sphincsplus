//go:build cgo && haraka_aesni && (!sha2_avx2 && !shake_a64 && !shake_avx2 && !ref)
package params

var ImplementationName = "haraka-aesni"
