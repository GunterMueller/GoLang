package obj

// (c) Christian Maurer   v. 200908 - license see nU.go

import ("runtime"; "math"; "strconv")

type Coder interface {

// Returns the number of bytes, that are needed
// to serialize x uniquely revertibly.
  Codelen() uint

// x.Eq (x.Decode (x.Encode())
  Encode() Stream

// Pre: b is a result of y.Encode() for some y
//      of the same type as x.
// x.Eq(y); x.Encode() == b, i.e. those slices coincide.
  Decode (Stream)
}

// Returns Codelen(int(0)) (== Codelen(uint(0))).
func C0() uint { return c0 }

// Pre: a is atomic or implements Object.
// Returns the codelength of a.
func Codelen (a Any) uint { return codelen(a) }

// Pre: a is atomic or implements Object.
// Returns a as encoded byte sequence.
func Encode (a Any) Stream { return encode(a) }

// Pre: a is atomic or implements Object.
//      bs is a encoded byte sequence.
// Returns the object, that was encoded.
func Decode (a Any, bs Stream) Any { return decode(a,bs) }

// Returns a byte sequence of length 16,
// that encodes a, b, c, d.
// func Encode4 (a, b, c, d uint32) Stream { return encode4 (a,b,c,d) }

// Pre: len(bs) == 16;
//      bs encodes 4 numbers of type uint32.
// Returns those 4 numbers.
// func Decode4 (bs Stream) (uint32, uint32, uint32, uint32) { return decode4(bs) }

// Returns true, iff a implements Coder.
func IsCoder (a Any) bool { return isCoder(a) }

// Returns true, iff a is atomic or implements Coder.
func AtomicOrCoder (a Any) bool { return Atomic(a) || isCoder(a) }

func isCoder (a Any) bool {
  _, c := a.(Coder)
  return c
}

func fail (a Any) {
  panic ("nU only [en|de]codes atomic types and objects of type string, {[Bool|Uint|Any]}Stream or Coder !")
}

var c0 uint

func init() {
  switch runtime.GOARCH {
  case "amd64":
    c0 = 8
  case "386":
    c0 = 4
  default:
    panic ("$GOARCH not treated") // Wer hilft mir mit MAC-OS ?
  }
}

func codelen (a Any) uint {
  if a == nil { return 0 }
  switch a.(type) {
  case bool, int8, uint8:
    return 1
  case int16, uint16:
    return 2
  case int32, uint32, float32:
    return 4
  case int, uint:
    return c0
  case int64, uint64, float64, complex64:
    return 8
  case complex128:
    return 16
  case string:
    return uint(len(a.(string)))
  case BoolStream:
    return uint(len(a.(BoolStream)))
  case Stream:
    return uint(len(a.(Stream)))
  case IntStream:
    return c0 + c0 * uint(len(a.(IntStream)))
  case UintStream:
    return c0 + c0 * uint(len(a.(UintStream)))
  case AnyStream:
    y := c0
    for _, b := range a.(AnyStream) {
      y += uint(codelen(b))
    }
    return y
  case Object:
    return (a.(Object)).Codelen()
  }
  fail (a)
  panic("") // shut up, compiler
}

func encode (a Any) Stream {
  if a == nil {
    return nil
  }
  var bs Stream
  switch a.(type) {
  case bool:
    bs = make(Stream, 1)
    if a.(bool) { bs[0]++ }
  case int8:
    bs = make(Stream, 1)
    bs[0] = a.(byte)
  case int16:
    n, x := 2, a.(int16)
    bs = make(Stream, n)
    for i := 0; i < n; i++ {
      bs[i] = byte(x)
      x >>= 8
    }
  case int32:
    n, x := 4, a.(int32)
    bs = make(Stream, n)
    for i := 0; i < n; i++ {
      bs[i] = byte(x)
      x >>= 8
    }
  case int:
    n, x := codelen(int(0)), a.(int)
    bs = make(Stream, n)
    for i := uint(0); i < n; i++ {
      bs[i] = byte(x)
      x >>= 8
    }
  case int64:
    n, x := 8, a.(int64)
    bs = make(Stream, n)
    for i := 0; i < n; i++ {
      bs[i] = byte(x)
      x >>= 8
    }
  case uint8:
    bs = make(Stream, 1)
    bs[0] = a.(uint8)
  case uint16:
    n, x := 2, a.(uint16)
    bs = make(Stream, n)
    for i := 0; i < n; i++ {
      bs[i] = byte(x)
      x >>= 8
    }
  case uint32:
    n, x := 4, a.(uint32)
    bs = make(Stream, n)
    for i := 0; i < n; i++ {
      bs[i] = byte(x)
      x >>= 8
    }
  case uint:
    x := a.(uint)
    bs = make(Stream, c0)
    for i := uint(0); i < c0; i++ {
      bs[i] = byte(x)
      x >>= 8
    }
  case uint64:
    n, x := 8, a.(uint64)
    bs = make(Stream, n)
    for i := 0; i < n; i++ {
      bs[i] = byte(x)
      x >>= 8
    }
  case float32:
    n, x := 4, math.Float32bits (a.(float32))
    bs = make(Stream, n)
    for i := 0; i < n; i++ {
      bs[i] = byte(x)
      x >>= 8
    }
  case float64:
    n, x := 8, math.Float64bits (a.(float64))
    bs = make(Stream, n)
    for i := 0; i < n; i++ {
      bs[i] = byte(x)
      x >>= 8
    }
  case complex64:
    n, c := 8, a.(complex64)
    x, y := math.Float32bits (real(c)), math.Float32bits (imag(c))
    bs = make(Stream, n)
    for i := 0; i < n / 2; i++ {
      bs[i] = byte(x)
      bs[i + n/2] = byte(y)
      x >>= 8; y >>= 8
    }
  case complex128:
    n, c := 16, a.(complex128)
    x, y := math.Float64bits (real(c)), math.Float64bits (imag(c))
    bs = make(Stream, n)
    for i := 0; i < n / 2; i++ {
      bs[i] = byte(x)
      bs[i + n/2] = byte(y)
      x >>= 8; y >>= 8
    }
  case string:
    return (Stream)(a.(string))
  case BoolStream:
    n := len (a.(BoolStream))
    ys := make(Stream, n)
    for i := 0; i < n; i++ {
      ys[i] = 0; if a.(BoolStream)[i] { ys[i] = 1 }
    }
    copy (bs, ys)
  case Stream:
    return a.(Stream)
  case IntStream:
    us := a.(IntStream)
    n := len(us)
    c := int(c0)
    bs = make(Stream, c * (n + 1))
    copy (bs[:c], encode(n))
    i := c
    for j := 0; j < n; j++ {
      copy (bs[i:i+c], encode(us[j]))
      i += c
    }
  case UintStream:
    us := a.(UintStream)
    n := uint(len(us))
    c := c0
    bs = make(Stream, c * (n + 1))
    copy (bs[:c], encode(n))
    i := c
    for j := uint(0); j < n; j++ {
      copy (bs[i:i+c], encode(us[j]))
      i += c
    }
  case AnyStream:
    as := a.(AnyStream)
    n := uint(len(as))
    c := c0
    for j := uint(0); j < n; j++ {
      c += c0 + codelen(as[j])
    }
    bs = make(Stream, c)
    copy (bs[:c0], encode(n))
    i := c0
    for j := uint(0); j < n; j++ {
      k := codelen(as[j])
      copy (bs[i:i+c0], encode(k))
      i += c0
      copy (bs[i:i+k], encode(as[j]))
      i += k
    }
  case Object:
    return a.(Object).Encode()
  default:
    fail (a)
  }
  return bs
}

func chk (b Stream, n int) {
  if len(b) < n {
    panic ("nU/obj/coder.go chk: len(b) == " + strconv.Itoa(len(b)) + " < n == " + strconv.Itoa(n))
  }
}

func decode (a Any, bs Stream) Any {
  if a == nil { return nil }
  switch a.(type) {
  case bool:
    chk (bs, 1)
    return bs[0] > 0
  case int8:
    chk (bs, 1)
    return int8(bs[0])
  case int16:
    n, x := 2, int16(0)
    chk (bs, n)
    for i := n - 1; i >= 0; i-- {
      x <<= 8
      x += int16(bs[i])
    }
    return x
  case int32:
    n, x := 4, int32(0)
    chk (bs, n)
    for i := n - 1; i >= 0; i-- {
      x <<= 8
      x += int32(bs[i])
    }
    return x
  case int:
    n, x := int(codelen(0)), 0
    chk (bs, int(n))
    for i := n - 1; i >= 0; i-- {
      x <<= 8
      x += int(bs[i])
    }
    return x
  case int64:
    n, x := 8, int64(0)
    chk (bs, n)
    for i := n - 1; i >= 0; i-- {
      x <<= 8
      x += int64(bs[i])
    }
    return x
  case uint8:
    chk (bs, 1)
    return bs[0]
  case uint16:
    n, x := 2, uint16(0)
    chk (bs, n)
    for i := n - 1; i >= 0; i-- {
      x <<= 8
      x += uint16(bs[i])
    }
    return x
  case uint32:
    n, x := 4, uint32(0)
    chk (bs, n)
    for i := n - 1; i >= 0; i-- {
      x <<= 8
      x += uint32(bs[i])
    }
    return x
  case uint:
    n, x := int(codelen(0)), uint(0)
    chk (bs, int(n))
    for i := n - 1; i >= 0; i-- {
      x <<= 8
      x += uint(bs[i])
    }
    return x
  case uint64:
    n, x := 8, uint64(0)
    chk (bs, n)
    for i := n - 1; i >= 0; i-- {
      x <<= 8
      x += uint64(bs[i])
    }
    return x
  case float32:
    n, x := 4, uint32(0)
    chk (bs, n)
    for i := n - 1; i >= 0; i-- {
      x <<= 8
      x += uint32(bs[i])
    }
    return math.Float32frombits (x)
  case float64:
    n, x := 8, uint64(0)
    chk (bs, n)
    for i := n - 1; i >= 0; i-- {
      x <<= 8
      x += uint64(bs[i])
    }
    return math.Float64frombits (x)
  case complex64:
    n, x, y := 8, uint32(0), uint32(0)
    chk (bs, n)
    for i := n / 2 - 1; i >= 0; i-- {
      x <<= 8; y <<= 8
      x += uint32(bs[i])
      y += uint32(bs[i + n/2])
    }
    return complex (math.Float32frombits (x), math.Float32frombits (y))
  case complex128:
    n, x, y := 16, uint64(0), uint64(0)
    chk (bs, n)
    for i := n / 2 - 1; i >= 0; i-- {
      x <<= 8; y <<= 8
      x += uint64(bs[i])
      y += uint64(bs[i + n/2])
    }
    return complex (math.Float64frombits (x), math.Float64frombits (y))
  case string:
    return string(bs)
  case BoolStream:
    n := len(bs)
    y := make(BoolStream, n)
    for i := 0; i < n; i++ {
      if bs[i] == 1 { y[i] = true }
    }
    return y
  case Stream:
    return bs
    copy (a.(Stream), bs)
  case IntStream:
    c := c0
    n := decode(0, bs[:c]).(int)
    us := make(IntStream, n)
    i := c
    for j := 0; j < n; j++ {
      us[j] = decode(0, bs[i:i+c]).(int)
      i += c
    }
    return us
  case UintStream:
    c := c0
    n := decode(uint(0), bs[:c]).(uint)
    us := make(UintStream, n)
    i := c
    for j := uint(0); j < n; j++ {
      us[j] = decode(uint(0), bs[i:i+c]).(uint)
      i += c
    }
    return us
  case AnyStream:
    n := decode(uint(0), bs[:c0]).(uint)
    as := make(AnyStream, n)
    i := c0
    for j := uint(0); j < n; j++ {
      k := decode(uint(0), bs[i:i+c0]).(uint)
      i += c0
      as[j] = bs[i:i+k] // the client has to decode this Stream himself
      i += k
    }
    return as
  case Object:
    chk (bs, int(a.(Object).Codelen()))
    a.(Object).Decode (bs)
  default:
    fail (a)
  }
  return a
}

/*/
func encode4 (a, b, c, d uint32) Stream {
  bs := make(Stream, 4 * 4)
  copy (bs[:4], encode (a))
  copy (bs[4:8], encode (b))
  copy (bs[8:12], encode (c))
  copy (bs[12:16], encode (d))
  return bs
}

func decode4 (bs Stream) (uint32, uint32, uint32, uint32) {
  a := decode (uint32(0), bs[:4]).(uint32)
  b := decode (uint32(0), bs[4:8]).(uint32)
  c := decode (uint32(0), bs[8:12]).(uint32)
  d := decode (uint32(0), bs[12:16]).(uint32)
  return a, b, c, d
}
/*/
