package nchan

// (c) Christian Maurer   v. 190402 - license see nU.go

import ("strconv"; "time"; "net"; . "nU/obj")

func (x *netChannel) Chan() (chan Any, chan Any) {
  return x.in, x.out
}

func (x *netChannel) serve (c net.Conn) {
  var r int
  for {
    r, x.error = c.Read (x.Stream)
    if r == 0 {
      break
    }
    if x.Any == nil {
      x.uint = uint(Decode (uint(0), x.Stream[:c0]).(uint))
      x.in <- x.Stream[c0:c0+x.uint]
      a := <-x.out
      _, x.error = c.Write(append(Encode(Codelen(a)), Encode(a)...))
    } else {
      x.in <- Decode (Clone (x.Any), x.Stream[:r])
      _, x.error = c.Write (Encode(<-x.out))
    }
  }
  c.Close()
}

func newn (a Any, h string, p uint16, s bool) NetChannel {
  x := new(netChannel)
  x.Any = Clone(a)
  x.uint = Codelen(a)
  if a == nil {
    x.uint = maxWidth
  }
  x.Stream = make(Stream, x.uint)
  x.in, x.out = make(chan Any), make(chan Any)
  x.isServer = s
  ps := ":" + strconv.Itoa(int(p))
  if x.isServer {
    x.Listener, x.error = net.Listen (network, ps)
    go func() {
      for {
        if c, e := x.Listener.Accept(); e == nil { // NOT x.Conn, x.error !
          go x.serve (c) // see above remark
        }
      }
    }()
  } else { // client
    for {
      if x.Conn, x.error = net.Dial (network, h + ps); x.error == nil {
        break
      }
      time.Sleep(500 * 1e6)
    }
  }
  return x
}
