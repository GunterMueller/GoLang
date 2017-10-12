package fig2

// (c) Christian Maurer   v. 170917 - license see µu.go

import (
//  "µu/font"
  . "µu/obj"
  "µu/col"
  "µu/psp"
)
type
  Type byte; const ( // 2-dimensional elementary figures on the screen
  Points = iota // Sequence of at most 512 points
  Segments // line segment[s]
  Polygon
  Curve // Bezier curve
  InfLine // given by two different points
  Rectangle // borders parallel to the screen borders
  Circle
  Ellipse
  Text // of almost 40 characters
  Image // in the ppm-format
  NTypes
)
type
  Figure2 interface {

  Object
  Marker

// x is of kind k.
  SetType (k Type)

// x is of the kind, that was selected interactively by the user.
  Select ()

// Returns the position of the
// - first point (of the first line), if the kind of x is <= Line,
// - top left corner of x, if x is of kind Rectangle or Image,
// - middle point of x, if x is of kind Circle or Ellipse,
// - bottom left corner of first characer, if is of kind Text.
  Pos () (int, int)

// Returns true, iff the point at (a, b) has a distance
// of at most t pixels from x.
  On (a, b int, t uint) bool

// x is moved by (a, b).
  Move (a, b int)

// Returns true, iff the the mouse cursor is in the interior of x
// or has a distance of not more than t pixels from its boundary.
  UnderMouse (t uint) bool

// x has the colour c.
  SetColour (c col.Colour)

// Returns the colour of x.
  Colour () col.Colour

// If x is a text, then it has the font size f.
// Otherwise, nothing has happened.
//  SetFont (f font.Size)

// x is drawn at its position to the screen
// in the backgroundcolour of the screen.
  Erase ()

// x is drawn at its position in its colour to the screen.
  Write ()

// All pixels of x on the screen are inverted in their colour.
  Invert ()

// Pre: x has a defined kind.
// If x was empty before, then x now is the figure interactively generated by the user,
// otherwise, x now has the position interactively chosen by the user.
  Edit ()

// TODO Spec
  Print (psp.PostscriptPage)
}

// Returns a new empty figure of undefined Type.
func New() Figure2 { return new_() }
