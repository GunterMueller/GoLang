package obj

// (c) Christian Maurer   v. 171227 - license see nU.go

type (
// statements
  Stmt func()
  StmtSpectrum func (uint)
// conditions
  Cond func() bool
  CondSpectrum func (uint) bool
// Nvalued functions
  NFunc func() uint
  NFuncSpectrum func (uint) uint
// operations
  Op func (Any)
  OpSpectrum func (Any, uint)
// functions
  Func func (Any) Any
  FuncSpectrum func (Any, uint) Any
// predicates
  Pred func (Any) bool
  PredSpectrum func (Any, uint) bool
// conditioned operations
  CondOp func (Any, bool)
  CondOp2 func (Any, Any, bool)
  CondOpSpectrum func (Any, bool, uint)
// relations
  Rel func (Any, Any) bool
  RelSpectrum func (Any, Any, uint) bool
)

// Stmt[Spectrum]
func Nothing() { }
func NothingSp (i uint) { }

// Op[Spectrum]
func Ignore (a Any) { }
func IgnoreSp (a Any, i uint) { }

// NFunc[Spectrum]
func Null() uint { return 0 }
func NullSp(i uint) uint { return 0 }

// Cond[Spectrum]
func True() bool { return true }
func TrueSp (i uint) bool { return true }

// Func[Spectrum]
func Id (a Any) Any { return a }
func IdSp (a Any, i uint) Any { return a }
func Nil (a Any) Any { return nil }
func NilSp (a Any, i uint) Any { return nil }

// Pred[Spectrum]
func AllTrue (a Any) bool { return true }
func AllTrueSp (a Any, i uint) bool { return true }

// CondOp[Spectrum]
func CondIgnore (a Any, b bool) { }
func CondIgnore2 (a, a1 Any, b bool) { }
func CondIgnoreSp (a Any, b bool, i uint) { }

// we get rid of TravPred by:
func PredOp2Op (p Pred, o Op) Op {
  return func (a Any) { if p(a) { o(a) } }
}

// we get rid of TravCond by:
func PredCondOp2Op (p Pred, o CondOp) Op {
  return func (a Any) { if p(a) { o(a, true) } else { o(a, false) } }
}
