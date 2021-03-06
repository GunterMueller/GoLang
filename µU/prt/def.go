package prt

// (c) Christian Maurer   v. 210315 - license see µU.go

// >>> Pre: TeX is installed.

import
  "µU/font"
var
  PrintCommand = "lp"

// The actual font is f.
func SetFont (f font.Font) { setFont(f) }

// Returns the actual font.
func ActualFont() font.Font { return actualFont }

// The actual fontsize is f.
func SetFontsize (s font.Size) { setFontsize(s) }

// Returns the actual fontsize.
func ActualSize() font.Size { return actualSize }

// Returns the number of lines per page.
func NLines() uint { return nL[actualSize] }

// Returns the number of columns per line.
func NColumns() uint { return nC[actualSize] }

// Spec: See TeX.
func Voffset (mm uint) { voffset(mm) }

// Spec: See TeX.
func Footline (s string) { footline(s) }

// Pre: l < maxL; c + 1 < maxC.
// b is n line l, column c in the actual font and fontsize in the printer buffer.
func Print1 (b byte, l, c uint) { print1(b, l, c) }

// Pre: l < maxL, c + len(s) < maxC.
// s is in line l from column c in the actual font and fontsize in the printer buffer.
func Print (s string, l, c uint) { print(s, l, c) }

// All lines of the printer buffer are printed; the printer buffer is not empty.
func GoPrint() { goPrint() }
