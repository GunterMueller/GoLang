package macc

// (c) Christian Maurer   v. 171020 - license see µU.go

import (
  "µU/euro"
  "µU/host"
)
type
  MAccount interface { // A multitasking capable account.
                       // The exported functions cannot be interrupted
                       // by calls of these functions of other goroutines.

// The balance of x is incremented by e.
// Returns the new balance of x.
  Deposit (e euro.Euro) euro.Euro

// The balance of x is decremented by e.
// Returns the new balance of x.
// The calling process was blocked, until the balance of x was greater or Equal than e.
  Draw (e euro.Euro) euro.Euro
}

// All constructors return new accounts with balance 0 Euro.

// Implementation with sync Cond's.
func New() MAccount { return new_() }

// Implementation with a universal monitor.
func NewM() MAccount { return newM() }

// Implementation with a far monitor.
func NewFMon (h host.Host, p uint16, s bool) MAccount { return newf(h,p,s) }
