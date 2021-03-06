package enum

// (c) Christian Maurer   v. 140522 - license see µu.go

const (
  UndefNumerus = uint8(iota)
  Sing
  Plur
  NNumeri
)
var
  lNum, sNum []string =
//  []string { "", "Sing.", "Plur." },
//  []string { "", "Sg.", "Pl." }
  []string { "", "Singular", "Plural" },
  []string { "", "Sing.", "Plur." }

func init() {
  l[Numerus], s[Numerus] = lNum, sNum
}
