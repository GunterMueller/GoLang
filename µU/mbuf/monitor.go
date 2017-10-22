package mbuf

// (c) Christian Maurer   v. 170218 - license see µU.go

import (
  . "µU/obj"
  "µU/buf"
  "µU/mon"
)
type
  mbufM struct {
               mon.Monitor
               }

func newM (a Any, n uint) MBuffer {
  buffer := buf.New (a, n)
  x := new(mbufM)
  f := func (a Any, i uint) Any {
         if i == ins {
           buffer.Ins (a)
           x.Monitor.Signal (ins)
           return nil
         }
         if buffer.Num() == 0 {
           x.Monitor.Wait (ins)
         }
         return buffer.Get()
       }
  x.Monitor = mon.New (nFuncs, f)
  return x
}

func (x *mbufM) Ins (a Any) {
  x.Monitor.F (a, ins)
}

func (x *mbufM) Get() Any {
  return x.Monitor.F (nil, get)
}
