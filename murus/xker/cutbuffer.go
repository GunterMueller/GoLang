package xker

// (c) murus.org  v. 150124 - license see murus.go

// #include <stdio.h>
// #include <stdlib.h>
// #include <X11/X.h>
// #include <X11/Xlib.h>
// #include <X11/Xlibint.h>
/*
cut (Display *d, char* s, int n) {
  int r0, r1, r2;
  r0 = XStoreBytes (d, s, n);
  r1 = XStoreBuffer (d, s, n, 1);
  r2 = XStoreBuffer (d, s, n, 2);
  printf ("%d, %d, %d\n", r0, r1, r2);
}
char *global;
char *paste (Display *d, int* n) {
  global = XFetchBytes (d, n);
  return global;
}
void xfree () {
  Xfree (global);
}
*/
import
  "C"
import (
  "unsafe"
)

func (x *window) Cut (s string) {
  cs, n:= C.CString (s), C.int(len (s))
  defer C.free (unsafe.Pointer (cs))
  C.cut (dpy, cs, n)
}

func (x *window) Paste() string {
  var (cs *C.char; n C.int)
  defer C.free (unsafe.Pointer (cs))
  cs = C.paste (dpy, &n)
  s:= C.GoStringN (cs, n)
  C.xfree()
  x.Flush()
  return s
}
