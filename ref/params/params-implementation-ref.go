//go:build cgo && ref && (!haraka_aesni && !sha2_avx2 && !shake_a64 && !shake_avx2)
package params

var ImplementationName = "ref"
