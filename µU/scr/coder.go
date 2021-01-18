package scr

// (c) Christian Maurer   v. 210104 - license see µU.go

func (X *screen) Codelen (w, h uint) uint {
  if underX {
    return X.XWindow.Codelen (w, h)
  }
  return X.Console.Codelen (w, h)
}

func (X *screen) Encode (x, y, w, h uint) []byte {
  if underX {
    return X.XWindow.Encode (x, y, w, h)
  }
  return X.Console.Encode (x, y, w, h)
}

func (X *screen) Decode (s []byte) {
  if underX {
    X.XWindow.Decode (s)
  } else {
    X.Console.Decode (s)
  }
}
