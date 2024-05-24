//go:build cgo && sha2_avx2 && (!haraka_aesni && !shake_a64 && !shake_avx2 && !ref)
package params

var ImplementationName = "sha2-avx2"
