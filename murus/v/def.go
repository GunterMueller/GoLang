package v

// (c) murus.org  v. 170809 - license see murus.go

import
  "murus/col"

func Colours() (col.Colour, col.Colour, col.Colour) { return colours() }
func String() string { return v.String() }
func Want (y, m, d uint) { want (y,m,d) }
