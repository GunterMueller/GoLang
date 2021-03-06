package internal

// (c) Christian Maurer   v. 170316 - license see murus.go

import (
  "os"
//  . "murus/ker"
  "murus/str"
//  "murus/env"
  "murus/errh"
)
const
  rights = 0644
/*
type
  clients byte; const (user = iota; group; world)
type
  accesses byte; const (readable = iota; writable; executable)
*/
type
  file struct {
         file *os.File
       exists bool
       offset,
    endoffset uint64
       isOpen bool
              error
              }
/*
var (
  statuscode [3][3]uint // accesses, clients
  caller, callerGroup uint // uid, gid of the calling process
)
*/
/*
func Owner (N string) uint {
  fi, err := Stat (N)
  if err != nil {
    return MaxCard
  }
  return uint(fi.Gid)
}
*/
/*
func FileGroup (N string) uint {
  fi, err := Stat (N)
  if err == nil {
    return uint(fi.Uid)
  }
  return MaxCard
}
*/
/*
func accessible (N string, a accesses) bool {
  fi, err := Stat (N)
  if err == nil {
    return caller == Owner (N)       && statuscode [a, user]  IN fi.Mode ||
           callerGroup == Gruppe (N) && statuscode [a, group] IN fi.Mode
  }
  return false
}
*/

func directLength (N string) uint64 {
  fi, err := os.Stat (N)
  if err == nil {
    return uint64(fi.Size())
  }
  return 0
}

func erase (N string) {
  _, err := os.Stat (N)
  if err != nil {
    os.Remove (N)
  }
}

func new_() File {
  f := new (file)
  f.file = nil
  f.exists = false
  f.offset = 0
  f.endoffset = 0
  f.isOpen = false
  return f
//  return &f // prüfen, ob file = nil
}

func (f *file) Fin() {
  f.flush()
}

/*
func (f *file) report (a, b string, n uint) {
  if f.error == nil {
//    errh.Error2 (a + b, n, "ok", 0)
  } else {
    errh.Error2 (a + " Error " + b, n, f.error.String(), 0)
  }
}
*/

func (f *file) Name (N string) {
  if str.Empty (N) { return }
  f.flush()
  var fi os.FileInfo
  fi, f.error = os.Stat (N)
  f.exists = f.error == nil
  if f.exists { // is there a file with name N (?)
//    if ! fi.IsRegular() { errh.Error0(N + " is no regular file"); Fin(); Exit (1) } // nothing goes
//    if fi.Permission() != rights { errh.Error0(N + " has no rights"); Fin(); Exit (1) } // nothing goes
    f.file, f.error = os.OpenFile (N, os.O_RDWR, rights) // ; f.report ("define", "OpenFile", 0)
    if f.error == nil {
      f.endoffset = uint64(fi.Size())
      _ = f.file.Close() // ; f.report ("define", "Close", 0)
    } else {
      f.file = nil
      println (&os.PathError { "define", N, f.error })
    }
  } else { // there is no file with name N (?)
    if os.IsPermission (f.error) { println ("no permission ") }
    f.file, f.error = os.Create (N) // ; f.report ("define", "Create", 0)
    if f.error == nil {
      f.endoffset = 0
      _ = f.file.Close() // ; f.report ("define", "Close", 0)
      f.exists = true
    } else {
      f.file = nil
    }
  }
  f.offset = 0
  f.isOpen = false
}

func (f *file) Rename (s string) {
  f.flush()
  if f.isOpen {
    _ = f.file.Sync() // ; f.report ("redefine", "Sync", 0)
    _ = f.file.Close() // ; f.report ("redefine", "Close", 0)
    f.isOpen = false
  }
  if f.file.Name() == s { return }
  _ = os.Rename (f.file.Name(), s) // ; f.report ("redefine", "Rename", 0)
  f.exists = true // = Stat (&name, f.status) == 0
  if ! f.exists { /* f.report ("redefine", "Stat", 0) */ }
  f.offset = 0
//  f.endoffset = status.Byteanzahl // !!!
  f.isOpen = false
}

func (f *file) Empty() bool {
  if f.exists {
    return f.endoffset == 0
  }
  return true
}

func (f *file) Clr() {
  f.open()
  f.file.Truncate (int64(0))
  f.flush()
  f.offset = 0
  f.endoffset = 0
  f.open()
}

func (f *file) Length() uint64 {
  return f.endoffset
}

func (f *file) Seek (P uint64) {
  f.offset = P
}

func (f *file) Position() uint64 {
  return f.offset
}

func (f *file) open() {
  if f.isOpen { return }
//  f, err := OpenFile (f.file.Name(), /* os.O_APPEND */ os.O_RDWR, rights) // ; f.report ("open", f.file.Name(), 0)
  f.file, f.error = os.OpenFile (f.file.Name(), /* os.O_APPEND */ os.O_RDWR, rights) // ; f.report ("open", f.file.Name(), 0)
//  if err == nil {
  if f.error == nil {
//    f.file = f
  } else {
    f.file = nil
  }
  f.isOpen = f.file != nil
}

func (f *file) read (bs []byte) uint {
  f.open()
  if ! f.isOpen { errh.Error (f.file.Name(), uint(f.offset)) }
  r := len (bs)
//  r, f.error = f.file.ReadAt (bs, int64(f.offset)) // macht Zicken
  _ /* off */, _ = f.file.Seek (int64(f.offset), 0) // ; f.report ("read", "Seek at offset", uint(off))
  r, _ = f.file.Read (bs) // ; f.report ("Read", "at offset", uint(off))
  f.offset += uint64(r)
  return uint(r)
}

func (f *file) Read (bs []byte) (int, error) {
// TODO
  f.open()
  r := len (bs)
  r, f.error = f.file.ReadAt (bs, int64 (f.offset))
  f.offset += uint64(r)
  return r, f.error
}

func (f *file) write (bs []byte) uint {
// is going to be deprecated - will be replaced by the next method
/*
  f.open()
  w := len (bs)
//  w, f.error = f.file.WriteAt (bs, int64(f.offset)) // ; f.report ("WriteAt", "", 1000000 + uint(w)
  _, _ = f.file.Seek (int64(f.offset), 0) // ; f.report ("write", "Seek at offset", uint (off))
  w, _ = f.file.Write (bs) // ; f.report ("Write", "at offset", uint (off))
  f.offset += uint64(w)
  if f.offset > f.endoffset { f.endoffset = f.offset }
  return uint(w)
*/
  w := len (bs)
  w, f.error = f.Write (bs)
  return uint (w)
}

func (f *file) Write (bs []byte) (int, error) {
  f.open()
  w := len (bs)
  w, f.error = f.file.WriteAt (bs, int64 (f.offset)) // ; f.report ("WriteAt", "", 1000000 + uint(w)
//  var off int64
//  /* off */ _, _ = f.file.Seek (int64(f.offset), 0) // ; f.report ("write", "Seek at offset", uint (off))
//  w, f.error = f.file.Write (bs) // ; f.report ("Write", "at offset", uint (off))
  f.offset += uint64(w)
  if f.offset > f.endoffset { f.endoffset = f.offset }
  return w, f.error
}

func (f *file) flush() {
  if f.isOpen {
    _ = f.file.Sync() // ; f.report ("flush", "Sync", 0)
    _ = f.file.Close() // ; f.report ("flush", "Close", 0)
    f.isOpen = false
  }
}

/*
func init() {
// nonsense - only some ideas
  for access := readable; access <= executable; access++ {
    for client := user; client <= world; client++ {
      statuscode [access, client] = 3 * (2 - client) + (2 - access)
    }
  }
  callingProgram := env.Parameter (0)
  fi, err := Stat (callingProgram)
  if err == nil {
    caller = MaxCard
    callerGroup = MaxCard
  } else {
    caller = uint(fi.Uid)
    callerGroup = uint(fi.Gid)
  }
}
*/
