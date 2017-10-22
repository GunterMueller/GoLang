package rw

// (c) Christian Maurer   v. 170731 - license see µU.go

// >>> 1st readers/writers problem

type
  channel struct {
  rI, rO, wI, wO,
            done chan int
                 }

func newCh() ReaderWriter {
  x := new(channel)
  x.rI, x.rO = make(chan int), make(chan int)
  x.wI, x.wO = make(chan int), make(chan int)
  x.done = make(chan int)
  go func() {
    var nR, nW uint // number of active readers/writers
    loop:
    for {
      if _, ok := <-x.done; ok { break loop }
      if nW == 0 {
        if nR == 0 {
          select {
          case <-x.rI:
            nR++
          case <-x.wI:
            nW = 1
          }
        } else { // nR > 0
          select {
          case <-x.rI:
            nR++
          case <-x.rO:
            nR--
          }
        }
      } else { // nW == 1
        select {
        case <-x.wO:
          nW = 0
        }
      }
    }
  }()
  return x
}

func (x *channel) ReaderIn() {
  x.rI <- 0
}

func (x *channel) ReaderOut() {
  x.rO <- 0
}

func (x *channel) WriterIn() {
  x.wI <- 0
}

func (x *channel) WriterOut() {
  x.wO <- 0
}

func (x *channel) Fin() {
  x.done <- 0
}
