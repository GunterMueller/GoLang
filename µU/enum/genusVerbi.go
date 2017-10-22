package enum

// (c) Christian Maurer   v. 140522 - license see µU.go

const (
  UndefGenusVerbi = uint8(iota)
  Akt
  Pass
  NGeneraVerbi
)
var
  lGenusVerbi, sGenusVerbi []string =
  []string { "", "Aktiv", "Passiv" },
  []string { "", "Akt.", "Pass." }

func init() {
  l[GenusVerbi], s[GenusVerbi] = lGenusVerbi, sGenusVerbi
}
