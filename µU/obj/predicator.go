package obj

// (c) Christian Maurer   v. 210314 - license see µU.go

type
  Predicator interface {

  Collector

// Returns the number of those elements in x, for which p returns true.
  NumPred (p Pred) uint

// Returns NumPred(p) == Num(), i.e. returns true, iff p returns true
// on all elements in x (particularly if x has no elements).
  All (p Pred) bool

// Returns true, iff there is an element in x, for which p returns true.
// In that case the actual element of x is for b/!b the last/first such
// element, otherwise the actual element of x is the same as before.
  ExPred (p Pred, b bool) bool

// Returns true, iff there is an element in x in direction f
// from the actual element of x, for which p returns true.
// In that case the actual element of x is for f/!f the
// next/previous such element, otherwise the actual element of x
// is the same as before.
  StepPred (p Pred, f bool) bool

// Pre: y is a collector of elements of the same type as those in x.
// y consists exactly of those elements in x before
// (in their order in x), for which p returns true.
// The actual element of x is undefined; x is unchanged.
  Filter (y Collector, p Pred)

// Pre: See Filter.
// y contains exactly those elements in x (in their order in x),
// for which p returns true, and exactly those elements are
// removed from x. The actual elements of x and y are undefined.
  Cut (y Collector, p Pred)

// In x all elements, for which p returns true, are removed.
// If the actual element of x was one of them, now it is undefined.
  ClrPred (p Pred)
}

func IsPredicator (a Any) bool {
  if a == nil { return false }
  _, ok := a.(Predicator)
  return ok
}
