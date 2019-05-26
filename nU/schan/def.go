package schan

// (c) Christian Maurer   v. 171104 - license see nU.go

import . "nU/obj"

type SynchronousChannel interface {

  Send (a Any)

  Recv() Any
}

func New (a Any) SynchronousChannel { return new_(a) }
