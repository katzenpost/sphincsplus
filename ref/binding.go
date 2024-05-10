//go:build (cgo && ((linux && amd64) || (darwin && amd64) || (darwin && arm64) || (windows && amd64))) || (sphincs_haraka_128f || sphincs_haraka_128s || sphincs_haraka_192f || sphincs_haraka_192s || sphincs_haraka_256f || sphincs_haraka_256s || sphincs_sha2_128f || sphincs_sha2_128s || sphincs_sha2_192f || sphincs_sha2_192s || sphincs_sha2_256f || sphincs_sha2_256s || sphincs_shake_128f || sphincs_shake_128s || sphincs_shake_192f || sphincs_shake_192s || sphincs_shake_256f || sphincs_shake_256s) || (!sphincs_haraka_128f && !sphincs_haraka_128s && !sphincs_haraka_192f && !sphincs_haraka_192s && !sphincs_haraka_256f && !sphincs_haraka_256s && !sphincs_sha2_128f && !sphincs_sha2_128s && !sphincs_sha2_192f && !sphincs_sha2_192s && !sphincs_sha2_256f && !sphincs_sha2_256s && !sphincs_shake_128f && !sphincs_shake_128s && !sphincs_shake_192f && !sphincs_shake_192s && !sphincs_shake_256f && !sphincs_shake_256s)

package sphincsplus

//#cgo darwin/amd64 CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=haraka-128 -D THASH=robust -D PARAMS=sphincs-haraka-128f
//#cgo darwin/amd64 test CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=haraka-128 -D THASH=robust -D PARAMS=sphincs-haraka-128f
//#cgo darwin/arm64 CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=haraka-128 -D THASH=robust -D PARAMS=sphincs-haraka-128f
//#cgo darwin/arm64 test CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=haraka-128 -D THASH=robust -D PARAMS=sphincs-haraka-128f
//#cgo haraka_aesni CFLAGS: -O3 -std=c99 -D CGO=1 -march=native -fomit-frame-pointer -flto -D HASH=haraka-128 -D THASH=robust -D PARAMS=sphincs-haraka-128f
//#cgo linux/amd64 CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=haraka-128 -D THASH=robust -D PARAMS=sphincs-haraka-128f
//#cgo linux/amd64 test CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=haraka-128 -D THASH=robust -D PARAMS=sphincs-haraka-128f
//#cgo linux/arm64 CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=haraka-128 -D THASH=robust -D PARAMS=sphincs-haraka-128f
//#cgo linux/arm64 test CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=haraka-128 -D THASH=robust -D PARAMS=sphincs-haraka-128f
//#cgo ref CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=haraka-128 -D THASH=robust -D PARAMS=sphincs-haraka-128f
//#cgo sha2_avx2 CFLAGS: -O3 -std=c99 -D CGO=1 -march=native -fomit-frame-pointer -flto -D HASH=sha2-128 -D THASH=robust -D PARAMS=sphincs-sha2-128f
//#cgo shake_a64 CFLAGS: -O3 -std=c99 -D CGO=1 -march=native -fomit-frame-pointer -flto -D HASH=shake-128 -D THASH=robust -D PARAMS=sphincs-shake-128f
//#cgo shake_avx2 CFLAGS: -O3 -std=c99 -D CGO=1 -march=native -fomit-frame-pointer -flto -D HASH=shake-128 -D THASH=robust -D PARAMS=sphincs-shake-128f
//#cgo sphincs_haraka_128f haraka robust CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=haraka -D THASH=robust -D PARAMS=sphincs-haraka-128f
//#cgo sphincs_haraka_128f haraka simple CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=haraka -D THASH=simple -D PARAMS=sphincs-haraka-128f
//#cgo sphincs_haraka_128f sha2 robust CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=sha2 -D THASH=robust -D PARAMS=sphincs-haraka-128f
//#cgo sphincs_haraka_128f sha2 simple CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=sha2 -D THASH=simple -D PARAMS=sphincs-haraka-128f
//#cgo sphincs_haraka_128f shake robust CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=shake -D THASH=robust -D PARAMS=sphincs-haraka-128f
//#cgo sphincs_haraka_128f shake simple CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=shake -D THASH=simple -D PARAMS=sphincs-haraka-128f
//#cgo sphincs_haraka_128s haraka robust CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=haraka -D THASH=robust -D PARAMS=sphincs-haraka-128s
//#cgo sphincs_haraka_128s haraka simple CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=haraka -D THASH=simple -D PARAMS=sphincs-haraka-128s
//#cgo sphincs_haraka_128s sha2 robust CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=sha2 -D THASH=robust -D PARAMS=sphincs-haraka-128s
//#cgo sphincs_haraka_128s sha2 simple CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=sha2 -D THASH=simple -D PARAMS=sphincs-haraka-128s
//#cgo sphincs_haraka_128s shake robust CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=shake -D THASH=robust -D PARAMS=sphincs-haraka-128s
//#cgo sphincs_haraka_128s shake simple CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=shake -D THASH=simple -D PARAMS=sphincs-haraka-128s
//#cgo sphincs_haraka_192f haraka robust CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=haraka -D THASH=robust -D PARAMS=sphincs-haraka-192f
//#cgo sphincs_haraka_192f haraka simple CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=haraka -D THASH=simple -D PARAMS=sphincs-haraka-192f
//#cgo sphincs_haraka_192f sha2 robust CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=sha2 -D THASH=robust -D PARAMS=sphincs-haraka-192f
//#cgo sphincs_haraka_192f sha2 simple CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=sha2 -D THASH=simple -D PARAMS=sphincs-haraka-192f
//#cgo sphincs_haraka_192f shake robust CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=shake -D THASH=robust -D PARAMS=sphincs-haraka-192f
//#cgo sphincs_haraka_192f shake simple CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=shake -D THASH=simple -D PARAMS=sphincs-haraka-192f
//#cgo sphincs_haraka_192s haraka robust CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=haraka -D THASH=robust -D PARAMS=sphincs-haraka-192s
//#cgo sphincs_haraka_192s haraka simple CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=haraka -D THASH=simple -D PARAMS=sphincs-haraka-192s
//#cgo sphincs_haraka_192s sha2 robust CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=sha2 -D THASH=robust -D PARAMS=sphincs-haraka-192s
//#cgo sphincs_haraka_192s sha2 simple CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=sha2 -D THASH=simple -D PARAMS=sphincs-haraka-192s
//#cgo sphincs_haraka_192s shake robust CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=shake -D THASH=robust -D PARAMS=sphincs-haraka-192s
//#cgo sphincs_haraka_192s shake simple CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=shake -D THASH=simple -D PARAMS=sphincs-haraka-192s
//#cgo sphincs_haraka_256f haraka robust CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=haraka -D THASH=robust -D PARAMS=sphincs-haraka-256f
//#cgo sphincs_haraka_256f haraka simple CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=haraka -D THASH=simple -D PARAMS=sphincs-haraka-256f
//#cgo sphincs_haraka_256f sha2 robust CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=sha2 -D THASH=robust -D PARAMS=sphincs-haraka-256f
//#cgo sphincs_haraka_256f sha2 simple CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=sha2 -D THASH=simple -D PARAMS=sphincs-haraka-256f
//#cgo sphincs_haraka_256f shake robust CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=shake -D THASH=robust -D PARAMS=sphincs-haraka-256f
//#cgo sphincs_haraka_256f shake simple CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=shake -D THASH=simple -D PARAMS=sphincs-haraka-256f
//#cgo sphincs_haraka_256s haraka robust CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=haraka -D THASH=robust -D PARAMS=sphincs-haraka-256s
//#cgo sphincs_haraka_256s haraka simple CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=haraka -D THASH=simple -D PARAMS=sphincs-haraka-256s
//#cgo sphincs_haraka_256s sha2 robust CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=sha2 -D THASH=robust -D PARAMS=sphincs-haraka-256s
//#cgo sphincs_haraka_256s sha2 simple CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=sha2 -D THASH=simple -D PARAMS=sphincs-haraka-256s
//#cgo sphincs_haraka_256s shake robust CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=shake -D THASH=robust -D PARAMS=sphincs-haraka-256s
//#cgo sphincs_haraka_256s shake simple CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=shake -D THASH=simple -D PARAMS=sphincs-haraka-256s
//#cgo sphincs_sha2_128f haraka robust CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=haraka -D THASH=robust -D PARAMS=sphincs-sha2-128f
//#cgo sphincs_sha2_128f haraka simple CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=haraka -D THASH=simple -D PARAMS=sphincs-sha2-128f
//#cgo sphincs_sha2_128f sha2 robust CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=sha2 -D THASH=robust -D PARAMS=sphincs-sha2-128f
//#cgo sphincs_sha2_128f sha2 simple CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=sha2 -D THASH=simple -D PARAMS=sphincs-sha2-128f
//#cgo sphincs_sha2_128f shake robust CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=shake -D THASH=robust -D PARAMS=sphincs-sha2-128f
//#cgo sphincs_sha2_128f shake simple CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=shake -D THASH=simple -D PARAMS=sphincs-sha2-128f
//#cgo sphincs_sha2_128s haraka robust CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=haraka -D THASH=robust -D PARAMS=sphincs-sha2-128s
//#cgo sphincs_sha2_128s haraka simple CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=haraka -D THASH=simple -D PARAMS=sphincs-sha2-128s
//#cgo sphincs_sha2_128s sha2 robust CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=sha2 -D THASH=robust -D PARAMS=sphincs-sha2-128s
//#cgo sphincs_sha2_128s sha2 simple CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=sha2 -D THASH=simple -D PARAMS=sphincs-sha2-128s
//#cgo sphincs_sha2_128s shake robust CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=shake -D THASH=robust -D PARAMS=sphincs-sha2-128s
//#cgo sphincs_sha2_128s shake simple CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=shake -D THASH=simple -D PARAMS=sphincs-sha2-128s
//#cgo sphincs_sha2_192f haraka robust CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=haraka -D THASH=robust -D PARAMS=sphincs-sha2-192f
//#cgo sphincs_sha2_192f haraka simple CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=haraka -D THASH=simple -D PARAMS=sphincs-sha2-192f
//#cgo sphincs_sha2_192f sha2 robust CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=sha2 -D THASH=robust -D PARAMS=sphincs-sha2-192f
//#cgo sphincs_sha2_192f sha2 simple CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=sha2 -D THASH=simple -D PARAMS=sphincs-sha2-192f
//#cgo sphincs_sha2_192f shake robust CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=shake -D THASH=robust -D PARAMS=sphincs-sha2-192f
//#cgo sphincs_sha2_192f shake simple CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=shake -D THASH=simple -D PARAMS=sphincs-sha2-192f
//#cgo sphincs_sha2_192s haraka robust CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=haraka -D THASH=robust -D PARAMS=sphincs-sha2-192s
//#cgo sphincs_sha2_192s haraka simple CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=haraka -D THASH=simple -D PARAMS=sphincs-sha2-192s
//#cgo sphincs_sha2_192s sha2 robust CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=sha2 -D THASH=robust -D PARAMS=sphincs-sha2-192s
//#cgo sphincs_sha2_192s sha2 simple CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=sha2 -D THASH=simple -D PARAMS=sphincs-sha2-192s
//#cgo sphincs_sha2_192s shake robust CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=shake -D THASH=robust -D PARAMS=sphincs-sha2-192s
//#cgo sphincs_sha2_192s shake simple CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=shake -D THASH=simple -D PARAMS=sphincs-sha2-192s
//#cgo sphincs_sha2_256f haraka robust CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=haraka -D THASH=robust -D PARAMS=sphincs-sha2-256f
//#cgo sphincs_sha2_256f haraka simple CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=haraka -D THASH=simple -D PARAMS=sphincs-sha2-256f
//#cgo sphincs_sha2_256f sha2 robust CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=sha2 -D THASH=robust -D PARAMS=sphincs-sha2-256f
//#cgo sphincs_sha2_256f sha2 simple CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=sha2 -D THASH=simple -D PARAMS=sphincs-sha2-256f
//#cgo sphincs_sha2_256f shake robust CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=shake -D THASH=robust -D PARAMS=sphincs-sha2-256f
//#cgo sphincs_sha2_256f shake simple CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=shake -D THASH=simple -D PARAMS=sphincs-sha2-256f
//#cgo sphincs_sha2_256s haraka robust CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=haraka -D THASH=robust -D PARAMS=sphincs-sha2-256s
//#cgo sphincs_sha2_256s haraka simple CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=haraka -D THASH=simple -D PARAMS=sphincs-sha2-256s
//#cgo sphincs_sha2_256s sha2 robust CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=sha2 -D THASH=robust -D PARAMS=sphincs-sha2-256s
//#cgo sphincs_sha2_256s sha2 simple CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=sha2 -D THASH=simple -D PARAMS=sphincs-sha2-256s
//#cgo sphincs_sha2_256s shake robust CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=shake -D THASH=robust -D PARAMS=sphincs-sha2-256s
//#cgo sphincs_sha2_256s shake simple CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=shake -D THASH=simple -D PARAMS=sphincs-sha2-256s
//#cgo sphincs_shake_128f haraka robust CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=haraka -D THASH=robust -D PARAMS=sphincs-shake-128f
//#cgo sphincs_shake_128f haraka simple CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=haraka -D THASH=simple -D PARAMS=sphincs-shake-128f
//#cgo sphincs_shake_128f sha2 robust CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=sha2 -D THASH=robust -D PARAMS=sphincs-shake-128f
//#cgo sphincs_shake_128f sha2 simple CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=sha2 -D THASH=simple -D PARAMS=sphincs-shake-128f
//#cgo sphincs_shake_128f shake robust CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=shake -D THASH=robust -D PARAMS=sphincs-shake-128f
//#cgo sphincs_shake_128f shake simple CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=shake -D THASH=simple -D PARAMS=sphincs-shake-128f
//#cgo sphincs_shake_128s haraka robust CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=haraka -D THASH=robust -D PARAMS=sphincs-shake-128s
//#cgo sphincs_shake_128s haraka simple CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=haraka -D THASH=simple -D PARAMS=sphincs-shake-128s
//#cgo sphincs_shake_128s sha2 robust CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=sha2 -D THASH=robust -D PARAMS=sphincs-shake-128s
//#cgo sphincs_shake_128s sha2 simple CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=sha2 -D THASH=simple -D PARAMS=sphincs-shake-128s
//#cgo sphincs_shake_128s shake robust CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=shake -D THASH=robust -D PARAMS=sphincs-shake-128s
//#cgo sphincs_shake_128s shake simple CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=shake -D THASH=simple -D PARAMS=sphincs-shake-128s
//#cgo sphincs_shake_192f haraka robust CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=haraka -D THASH=robust -D PARAMS=sphincs-shake-192f
//#cgo sphincs_shake_192f haraka simple CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=haraka -D THASH=simple -D PARAMS=sphincs-shake-192f
//#cgo sphincs_shake_192f sha2 robust CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=sha2 -D THASH=robust -D PARAMS=sphincs-shake-192f
//#cgo sphincs_shake_192f sha2 simple CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=sha2 -D THASH=simple -D PARAMS=sphincs-shake-192f
//#cgo sphincs_shake_192f shake robust CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=shake -D THASH=robust -D PARAMS=sphincs-shake-192f
//#cgo sphincs_shake_192f shake simple CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=shake -D THASH=simple -D PARAMS=sphincs-shake-192f
//#cgo sphincs_shake_192s haraka robust CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=haraka -D THASH=robust -D PARAMS=sphincs-shake-192s
//#cgo sphincs_shake_192s haraka simple CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=haraka -D THASH=simple -D PARAMS=sphincs-shake-192s
//#cgo sphincs_shake_192s sha2 robust CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=sha2 -D THASH=robust -D PARAMS=sphincs-shake-192s
//#cgo sphincs_shake_192s sha2 simple CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=sha2 -D THASH=simple -D PARAMS=sphincs-shake-192s
//#cgo sphincs_shake_192s shake robust CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=shake -D THASH=robust -D PARAMS=sphincs-shake-192s
//#cgo sphincs_shake_192s shake simple CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=shake -D THASH=simple -D PARAMS=sphincs-shake-192s
//#cgo sphincs_shake_256f haraka robust CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=haraka -D THASH=robust -D PARAMS=sphincs-shake-256f
//#cgo sphincs_shake_256f haraka simple CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=haraka -D THASH=simple -D PARAMS=sphincs-shake-256f
//#cgo sphincs_shake_256f sha2 robust CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=sha2 -D THASH=robust -D PARAMS=sphincs-shake-256f
//#cgo sphincs_shake_256f sha2 simple CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=sha2 -D THASH=simple -D PARAMS=sphincs-shake-256f
//#cgo sphincs_shake_256f shake robust CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=shake -D THASH=robust -D PARAMS=sphincs-shake-256f
//#cgo sphincs_shake_256f shake simple CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=shake -D THASH=simple -D PARAMS=sphincs-shake-256f
//#cgo sphincs_shake_256s haraka robust CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=haraka -D THASH=robust -D PARAMS=sphincs-shake-256s
//#cgo sphincs_shake_256s haraka simple CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=haraka -D THASH=simple -D PARAMS=sphincs-shake-256s
//#cgo sphincs_shake_256s sha2 robust CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=sha2 -D THASH=robust -D PARAMS=sphincs-shake-256s
//#cgo sphincs_shake_256s sha2 simple CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=sha2 -D THASH=simple -D PARAMS=sphincs-shake-256s
//#cgo sphincs_shake_256s shake robust CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=shake -D THASH=robust -D PARAMS=sphincs-shake-256s
//#cgo sphincs_shake_256s shake simple CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=shake -D THASH=simple -D PARAMS=sphincs-shake-256s
//#cgo windows/amd64 CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=haraka-128 -D THASH=robust -D PARAMS=sphincs-haraka-128f
//#cgo windows/amd64 test CFLAGS: -O3 -std=c99 -D CGO=1 -D HASH=haraka-128 -D THASH=robust -D PARAMS=sphincs-haraka-128f
//#include "api.h"
import "C"
import (
	"fmt"
	"unsafe"

	"github.com/ioerror/sphincsplus/ref/params"
)

var (

	// PublicKeySize is the size in bytes of the public key.
	PublicKeySize int = C.CRYPTO_PUBLICKEYBYTES

	// PrivateKeySize is the size in bytes of the private key.
	PrivateKeySize int = C.CRYPTO_SECRETKEYBYTES

	// SignatureSize is the size in bytes of the signature.
	SignatureSize int = C.CRYPTO_BYTES

	// ErrPublicKeySize indicates the raw data is not the correct size for a public key.
	ErrPublicKeySize error = fmt.Errorf("%s: raw public key data size is wrong", Name())

	// ErrPrivateKeySize indicates the raw data is not the correct size for a private key.
	ErrPrivateKeySize error = fmt.Errorf("%s: raw private key data size is wrong", Name())
)

func Name() string {
	return params.Name()
}

func SchemeName() string {
	return params.Name()
}

func Implementation() string {
	return params.Implementation()
}

// NewKeypair generates a new Sphincs+ keypair.
func NewKeypair() (*PrivateKey, *PublicKey) {
	privKey := &PrivateKey{
		privateKey: make([]byte, C.CRYPTO_SECRETKEYBYTES),
	}
	pubKey := &PublicKey{
		publicKey: make([]byte, C.CRYPTO_PUBLICKEYBYTES),
	}
	C.crypto_sign_keypair((*C.uchar)(unsafe.Pointer(&pubKey.publicKey[0])),
		(*C.uchar)(unsafe.Pointer(&privKey.privateKey[0])))
	return privKey, pubKey
}

// PublicKey is a public Sphincs+ key.
type PublicKey struct {
	publicKey []byte
}

// Verify checks whether the given signature is valid.
func (p *PublicKey) Verify(signature, message []byte) bool {
	ret := C.crypto_sign_verify((*C.uchar)(unsafe.Pointer(&signature[0])),
		C.size_t(len(signature)),
		(*C.uchar)(unsafe.Pointer(&message[0])),
		C.size_t(len(message)),
		(*C.uchar)(unsafe.Pointer(&p.publicKey[0])))
	if ret == 0 {
		return true
	}
	return false
}

// Bytes returns the PublicKey as a byte slice.
func (p *PublicKey) Bytes() []byte {
	return p.publicKey
}

// FromBytes loads a PublicKey from the given byte slice.
func (p *PublicKey) FromBytes(data []byte) error {
	if len(data) != PublicKeySize {
		return ErrPublicKeySize
	}

	p.publicKey = data
	return nil
}

// Verify checks whether the given signature is valid.
/*
func (p *PublicKey) Verify(signature, message []byte) bool {
	ret := C.crypto_sign_verify((*C.uchar)(unsafe.Pointer(&signature[0])),
		C.ulong(len(signature)),
		(*C.uchar)(unsafe.Pointer(&message[0])),
		C.ulong(len(message)),
		(*C.uchar)(unsafe.Pointer(&p.publicKey[0])))
	if ret == 0 {
		return true
	}
	return false
}
*/

// PrivateKey is a private Sphincs+ key.
type PrivateKey struct {
	privateKey []byte
}

// Sign signs the given message and returns the signature.
func (p *PrivateKey) Sign(message []byte) []byte {
	signature := make([]byte, C.CRYPTO_BYTES)
	sigLen := C.size_t(C.CRYPTO_BYTES)
	C.crypto_sign_signature((*C.uchar)(unsafe.Pointer(&signature[0])),
		&sigLen,
		(*C.uchar)(unsafe.Pointer(&message[0])),
		(C.size_t)(len(message)),
		(*C.uchar)(unsafe.Pointer(&p.privateKey[0])))
	return signature
}

// Bytes returns the PrivateKey as a byte slice.
func (p *PrivateKey) Bytes() []byte {
	return p.privateKey
}

// FromBytes loads a PrivateKey from the given byte slice.
func (p *PrivateKey) FromBytes(data []byte) error {
	if len(data) != PrivateKeySize {
		return ErrPrivateKeySize
	}

	p.privateKey = data
	return nil
}
