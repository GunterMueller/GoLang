package v

// (c) Christian Maurer   v. 170809 - license see µu.go

import
  "µu/col"

func Colours() (col.Colour, col.Colour, col.Colour) { return colours() }
func String() string { return v.String() }
func Want (y, m, d uint) { want (y,m,d) }
