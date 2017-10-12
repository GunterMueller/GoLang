package ker

// (c) Christian Maurer   v. 170101 - license see µu.go

const (
  T6 = 54
  tw = 38 //           0          1         2          3          4         5
)         //           012 345678901234567890123456789 0123456789 0123456789012
var       //                                                   !    !
  p6t []byte = []byte("P6\n# (c) 1986-2017   Christian Maurer\n____ ____\n255\n")

func p6txt (n, k uint) {
  for i := 0; i < 4; i++ {
    p6t[int(k) + 3 - i] = '0' + byte(n % 10)
    n = n / 10
  }
}

func p6number (b []byte) (uint, int) {
  i:= 0
  for '0' <= b[i] && b[i] <= '9' { i ++ }
  n := uint(0)
  for j := 0; j < i; j++ {
    n = 10 * n + uint(b[j] - '0')
  }
  return n, i
}

func P6Txt (w, h uint, bs []byte) int {
  p6txt (w, tw)
  p6txt (h, tw + 5)
  j := T6
  copy (bs[:j], p6t)
  return j
}

func P6dec (bs []byte) (uint, uint, uint, int) {
  w, h, fix := uint(0), uint(0), uint(0)
  p6 := string (bs[:2])
  if p6 != "P6" { return 0, 0, 0, 0 }
  i, di := 3, 0
  if bs[i] == '#' { // ignore comment
    for {
      i++
      if bs[i] < ' ' {
        i++ // ignore LF
        break
      }
    }
  }
  w, di = p6number (bs[i:])
  i += 1 + di // ignore LF or space
  h, di = p6number (bs[i:])
  i += 1 + di
  fix, di = p6number (bs[i:])
  i += 1 + di
  return w, h, fix, i
}

func P6Size (bs []byte) (uint, uint) {
  w, h, fix, _ := P6dec (bs)
  if fix != 255 { w, h = uint(0), uint(0) }
  return w, h
}
