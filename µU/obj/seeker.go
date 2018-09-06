package obj

// (c) Christian Maurer   v. 180902 - license see µU.go

type
  Seeker interface { // makes sense only for objects of type Collector

//  Collector

// Returns Num(), iff Offc(); returns otherwise
// the position of the actual object of x (starting at 0).
  Pos() uint

// The actual object of x is its p-th object, iff p < Num();
// otherwise Offc() == true.
  Seek (p uint)
}

func IsSeeker (a Any) bool {
  if a == nil { return false }
  _, ok := a.(Seeker)
  return ok
}

/*
type
  SeekerIterator interface {

  Iterator
  Seeker
}
*/
