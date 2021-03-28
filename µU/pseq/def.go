package pseq

// (c) Christian Maurer   v. 210322 - license see µU.go

import
  . "µU/obj"
type
  PersistentSequence interface {

  Seeker // hence Collector
  Persistor
}

// Pre: a is atomic or of a type implementing Equaler and Coder.
// Returns a new empty persistent sequence for objects of the type of a.
func New (a Any) PersistentSequence { return new_(a) }
