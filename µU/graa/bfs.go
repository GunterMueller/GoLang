package gra

// (c) Christian Maurer   v. 171112 - license see µU.go

import (
  "sort"
  "µU/ker"
  . "µU/obj"
)

func insert (s []*vertex, v *vertex, i uint) []*vertex {
  l := uint(len (s))
  if i > l { i = l }
  s1 := make ([]*vertex, l + 1)
  copy (s1[:i], s[:i])
  s1[i] = v
  copy (s1[i+1:], s[i:])
  return s1
}

func (x *graph) preBreadth() {
  for v := x.vAnchor.nextV; v != x.vAnchor; v = v.nextV {
    v.dist = inf
    v.predecessor = nil
  }
  x.colocal.dist = 0
}

func (x *graph) ActPred (p Pred) {
  v := x.vAnchor.nextV
  if v == x.vAnchor { return }
  if ! p (x.local.Any) { return }
  x.ClrSub()
  if ! x.ConnCond (p) { return }
  x.preBreadth()
  if x.eAnchor.Any == nil {
    x.bfs (p)
  } else {
    x.searchShortestPath (p)
  }
  x.path = nil
  for v := x.local; v != nil; v = v.predecessor {
   x.path = insert (x.path, v, 0)
  }
  x.defineSubgraph (x.local)
}

func (x *graph) Act() {
  x.ActPred (AllTrue)
}

// Lit.: CLR 23.2, CLRS 22.2
// TODO spec
func (x *graph) bfs (p Pred) {
  var qu []*vertex
  qu = append (qu, x.colocal)
  for len (qu) > 0 {
    v := qu[0]
    if len (qu) == 1 {
      qu = nil
    } else {
      qu = qu [1:]
    }
    for n := v.nbPtr.nextNb; n != v.nbPtr; n = n.nextNb {
      if n.outgoing && n.to.dist == inf && p (n.to.Any) {
        if x.demo [Breadth] {
          var n1 *neighbour
          if n.to.predecessor == nil {
            // TODO then what ?
          } else {
            n1 = n.to.predecessor.nbPtr.nextNb
            for n1.from != n.to.predecessor {
              n1 = n1.nextNb
              if n1.nextNb == n1 { ker.Oops() }
            }
            x.writeE (n1.edgePtr.Any, false)
            x.writeV (n1.from.Any, true)
            wait()
          }
        }
        n.to.dist = v.dist + 1
        n.to.predecessor = v
        qu = append (qu, n.to)
      }
    }
  }
}

type vSeq []*vertex

func (vs vSeq) Less (i, j int) bool {
  if vs[i].dist == vs[j].dist {
    if vs[i] == vs[j] { return false }
    return i < j
  }
  return vs[i].dist < vs[j].dist
}

// Algorithm of Dijkstra, Lit.: CLR 25.1-2, CLRS 24.2-3
// Pre: dist == inf, predecessor == nil for all vertices.
// TODO spec
func (x *graph) searchShortestPath (p Pred) {
  v := x.vAnchor.nextV
  vs := make (vSeq, x.nVertices)
  for i, v := 0, x.vAnchor.nextV; v != x.vAnchor; i, v = i + 1, v.nextV {
    vs[i] = v
  }
  sort.Slice (vs, vs.Less)
  for len (vs) > 0 {
    v = vs[0]
    if len (vs) == 1 {
      vs = nil
    } else {
      vs = vs[1:]
    }
    var d uint32
    for n := v.nbPtr.nextNb; n != v.nbPtr; n = n.nextNb {
      if n.outgoing && n.to != v.predecessor && p (n.to.Any) {
        if v.dist == inf {
          d = inf
        } else {
          d = v.dist + uint32(Val(n.edgePtr.Any))
        }
        if d < n.to.dist {
          if x.demo [Breadth] {
            if n.to.predecessor != nil {
              n1 := n.to.predecessor.nbPtr.nextNb
              for n1.from != n.to.predecessor {
                n1 = n1.nextNb
                if n1.nextNb == n1 { ker.Oops() }
              }
              x.writeE (n1.edgePtr.Any, false)
              x.writeV (n.to.Any, false)
            }
            x.writeE (n.edgePtr.Any, true)
            x.writeV (n.to.Any, true)
            wait()
          }
          n.to.dist, n.to.predecessor = d, v
// put the changed nb.to into the right position in vs:
          sort.Slice (vs, vs.Less)
        }
      }
    }
  }
}
