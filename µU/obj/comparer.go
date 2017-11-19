package obj

// (c) Christian Maurer   v. 171112 - license see µU.go

type
  Comparer interface {

// Pre: x is of the same type as the calling object.
// Returns true, iff the calling object is smaller than x.
  Less (x Any) bool
}

func Less (a, b Any) bool { return less(a,b) }
// func Leq (a, b Any) bool { return leq(q,b) }

func less (a, b Any) bool {
  switch a.(type) {
  case bool:
    return ! a.(bool) && b.(bool)
  case byte:
    return a.(byte) < b.(byte)
  case uint16:
    return a.(uint16) < b.(uint16)
  case uint32:
    return a.(uint32) < b.(uint32)
  case uint:
    return a.(uint) < b.(uint)
  case uint64:
    return a.(uint64) < b.(uint64)
  case int8:
    return a.(int8) < b.(int8)
  case int16:
    return a.(int16) < b.(int16)
  case int32:
    return a.(int32) < b.(int32)
  case int:
    return a.(int) < b.(int)
  case int64:
    return a.(int64) < b.(int64)
  case float32:
    return a.(float32) < b.(float32)
  case float64:
    return a.(float64) < b.(float64)
  case string:
    return a.(string) < b.(string)
  case Comparer:
    switch b.(type) {
    case Comparer:
      return a.(Comparer).Less (b)
    }
  }
  return false
}

func leq (a, b Any) bool {
  switch a.(type) {
  case bool:
    return ! a.(bool) && b.(bool) ||
             a.(bool) == b.(bool)
  case byte:
    return a.(byte) <= b.(byte)
  case uint16:
    return a.(uint16) <= b.(uint16)
  case uint32:
    return a.(uint32) <= b.(uint32)
  case uint:
    return a.(uint) <= b.(uint)
  case uint64:
    return a.(uint64) <= b.(uint64)
  case int8:
    return a.(int8) <= b.(int8)
  case int16:
    return a.(int16) <= b.(int16)
  case int32:
    return a.(int32) <= b.(int32)
  case int:
    return a.(int) <= b.(int)
  case int64:
    return a.(int64) <= b.(int64)
  case float32:
    return a.(float32) <= b.(float32)
  case float64:
    return a.(float64) <= b.(float64)
  case string:
    return a.(string) <= b.(string)
  case Comparer:
    switch b.(type) {
    case Comparer:
      switch a.(type) {
      case Equaler:
        switch b.(type) {
        case Equaler:
          return a.(Comparer).Less (b) || a.(Equaler).Eq (b)
        }
      }
    }
  }
  return false
}
