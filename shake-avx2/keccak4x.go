package sphincsplus

import (
  "github.com/ioerror/sphincsplus/shake-avx2/keccak4x"
)

func Keccak4xEnabled() bool {
  return keccak4x.Keccak4x_enabled()
}
