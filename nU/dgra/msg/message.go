package msg

// (c) Christian Maurer   v. 190402 - license see µU.go

import (
  . "nU/obj"
  "nU/dgra/status"
)
type
  message struct {
                 byte "message type"
               s status.Status
                 }

func new_() Message {
  x := new(message)
  x.byte = NTypes
  x.s = status.New()
  return x
}

func (x *message) imp (Y Any) *message {
  y, _ := Y.(*message)
  return y
}

func (x *message) Set (t byte, s status.Status) {
  x.byte, x.s = t, s
}

func (x *message) Type() byte {
  return x.byte
}

func (x *message) String() string {
  switch x.byte {
  case 0:
    return "Ask   "
  case 1:
    return "Accept"
  case 2:
    return "Update"
  case 3:
    return "YourC."
  case 4:
    return "Leader"
  }
  return "NTypes"
}

func (x *message) Eq (Y Any) bool {
  y := x.imp(Y)
  return x.byte == y.byte && x.s.Eq (y.s)
}

func (x *message) Copy (Y Any) {
  y := x.imp(Y)
  x.byte = y.byte
  x.s.Copy (y.s)
}

func (x *message) Clone() Any {
  y := new_()
  y.Set (x.byte, x.s)
  return y
}

func (x *message) Status() status.Status {
  return x.s
}

func (x *message) Codelen() uint {
  return 1 + x.s.Codelen()
}

func (x *message) Encode() Stream {
  return append(Encode(x.byte), x.s.Encode()...)
}

func (x *message) Decode (s Stream) {
  x.byte = s[0]
  x.s.Decode (s[1:])
}
