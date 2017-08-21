package utensil

// (c) murus.org  v. 150513 - license see murus.go

import
  "murus/col"
const (
  Papier = uint(iota)
  Tabak
  Hölzer
  NUtensils
)

func String (u uint) string { return text[u] }

func Colour (u uint) col.Colour { return colour[u] }

func Other1 (u uint) uint { return other1(u) }

func Other2 (u uint) uint { return other2(u) }

func Others (u uint) (uint, uint) { return others(u) }
