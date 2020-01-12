package lock

// (c) Christian Maurer   v. 190811 - license see µU.go

import (
  "µU/obj"
  . "µU/atomic"
)
type
  cas struct {
             uint "0 or 1, initially 0"
             }

func newCAS() Locker {
  return new(cas)
}

func (x *cas) Lock() {
  for ! CompareAndSwap (&x.uint, 0, 1) {
    obj.Nothing()
  }
}

func (x *cas) Unlock() {
  x.uint = 0
}
