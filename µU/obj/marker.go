package obj

// (c) Christian Maurer   v. 180902 - license see µU.go

type
  Marker interface {

// x is marked, iff m.
  Mark (m bool)

// Returns true, iff x is marked.
  Marked () bool
}

func IsMarker (a Any) bool {
  if a == nil { return false }
  _, ok := a.(Marker)
  return ok
}
