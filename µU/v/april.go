package v

// (c) Christian Maurer   v. 201014 - license see µU.go

import (
  "strconv"
  "µU/time"
  "µU/str"
  "µU/col"
  "µU/scr"
  "µU/box"
  "µU/pseq"
)
const (
  f = "/tmp/april"
  t = ".tmp"
)
var (
  x0, y0 uint
  ok bool
  doofcol = col.Pink()
)

func rot (n int) { var b byte
  for i := 0; ! ok; i++ {
    switch i % 4 {
    case 0:
      b = '|'
    case 1:
      b = '/'
    case 2:
      b = '-'
    case 3:
      b = '\\'
    }
    scr.Write1 (b, y0 + 1, x0 + 64); time.Msleep (50)
  }
}

func doof (y, x uint) { scr.ColourF (doofcol); scr.Write ("d(o,o)f", y, x0 + x) }

func doodle (c col.Colour, n int) {
  scr.ScrColourB (c); scr.Cls(); time.Msleep(50)
  scr.ScrColourB (col.Black()); scr.Cls(); scr.ScrColourF (col.LightGreen())
  scr.Write ("The µU-tool to remove \"       \" is going to be executed, i.e.", y0, x0)
  scr.Write ("your disk will be completely reformatted, one moment please ... ", y0 + 1, x0); doof (y0, 26)
  const m = 1<<16
  x := str.New (m)
  ok = false
  for i := 0; i < 10 * n; i++ { if i == 0 { go rot (n) }
    file := pseq.New (x)
    file.Name (f + strconv.Itoa(i) + t)
    file.Clr()
    file.Ins (x)
    file.Fin()
  }
  ok = true
  scr.ColourF (col.LightWhite())
}

func April1st() {
  x0 = scr.NColumns() / 2 - 35
  y0 = scr.NLines() / 2 - 1
  scr.MousePointer (false)
  scr.Write ("Found evil software: \"       \"", y0, x0); doof (y0, 22)
  scr.ColourF (col.White());
  scr.Write ("Remove (yes/no) ?", y0 + 2, x0)
  b:= box.New(); b.Wd (3); t0:= "yes"; b.Edit (&t0, y0 + 2, x0 + 18)
  scr.WarpMouseGr (2 * int(scr.Wd()), 2 * int(scr.Ht()))
  b.Colours (col.LightOrange(), col.Black()); b.Write ("yes", y0, x0 + 16)
  doodle (col.LightWhite(), 2); doodle (col.LightWhite(), 2); doodle (col.LightYellow(), 5)
  doodle (col.Yellow(), 3); doodle (col.LightOrange(), 5); doodle (col.Orange(), 8); doodle (col.LightRed(), 3)
  t0 = str.New (70); scr.Write (t0, y0, x0); scr.Write (t0, y0 + 1, x0)
  scr.ScrColours (col.LightGreen(), col.Blue()); scr.Cls()
  scr.Colours (col.LightGreen(), col.Blue())
  scr.Write ("The µU-tool has removed \"       \" - your disk is reformatted :-)", y0, x0); doof (y0, 28)
  scr.Colours (col.LightWhite(), col.Blue())
  scr.Write ("Please install Linux, TeX, git, Go and µU completely new !", y0 + 2, x0)
  time.Sleep(30)
}
