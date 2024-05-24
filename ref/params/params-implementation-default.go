//go:build cgo && (!haraka_aesni && !ref && !sha2_avx2 && !shake_a64 && !shake_avx2) 
package params

var ImplementationName = "ref"
