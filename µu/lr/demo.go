package lr

// (c) Christian Maurer   v. 170411 - license see µu.go

import
  . "µu/obj"
var
  writeL, writeR = Ignore, Ignore

func InstallDemo (l, r Op) {
  writeL, writeR = l, r
}
