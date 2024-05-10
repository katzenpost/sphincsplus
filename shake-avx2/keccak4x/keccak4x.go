package keccak4x

//#cgo shake_avx2 CFLAGS: -O3 -std=c99 -D CGO=1 -march=native -fomit-frame-pointer -flto -D HASH=shake-128 -D THASH=robust -D PARAMS=sphincs-shake-128f
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
import (
	"C"
)

func Keccak4x_enabled() bool {
	return true
}
