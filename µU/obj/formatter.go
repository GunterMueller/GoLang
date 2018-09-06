package obj

// (c) Christian Maurer   v. 180902 - license see µU.go

type
  Format byte
type
  Formatter interface {

// Pre: f < Nformats of the objects of the type of x.
// x has the format f.
  SetFormat (f Format)

// Returns the format of x.
  GetFormat() Format
}

func IsFormatter (a Any) bool {
  if a == nil { return false }
  _, ok := a.(Formatter)
  return ok
}
