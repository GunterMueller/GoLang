package phil

// (c) Christian Maurer   v. 171125 - license see µU.go

// >>> Solution with far monitor

import (
  . "µU/obj"
  . "µU/lockn"
//  "µU/host"
  "µU/fmon"
)
type
  farMonitor struct {
                    fmon.FarMonitor
                    }

func newFM (h /* host.Host */ string, port uint16, s bool) LockerN {
  nForks := make([]uint, NPhilos)
  for i := uint(0); i < NPhilos; i++ {
    nForks[i] = 2
  }
  p := func (a Any, i uint) bool {
         if i == lock {
           p := a.(uint) // p-th philosopher
           return nForks[p] == 2
         }
         return true // unlock
       }
  f := func (a Any, i uint) Any {
         p := a.(uint) // p-th philosopher
         if i == lock {
           nForks[left(p)]--
           nForks[right(p)]--
         } else { // unlock
           nForks[left(p)]++
           nForks[right(p)]++
         }
         return p
       }
  return &farMonitor { fmon.NewN (nil, NPhilos, f, p, h, port, s) }
}

func (x *farMonitor) Lock (p uint) {
  changeStatus (p, hungry)
  x.F (p, lock)
  changeStatus (p, dining)
}

func (x *farMonitor) Unlock (p uint) {
  changeStatus (p, satisfied)
  x.F (p, unlock)
}
