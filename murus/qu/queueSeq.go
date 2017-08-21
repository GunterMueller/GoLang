package qu

// (c) murus.org  v. 170620 - license see murus.go

import (
  . "murus/obj"
  "murus/seq"
)
type
  squeue struct {
                seq.Sequence
                }

func news(a Any) Queue {
  return &squeue { seq.New(a) }
}

func (x *squeue) Ins (a Any) {
  x.Seek(x.Num ())
  x.Sequence.Ins(a)
}

func (x *squeue) Get() Any {
  if x.Empty() { return nil }
  x.Seek(0)
  defer x.Sequence.Del()
  return x.Sequence.Get()
}
