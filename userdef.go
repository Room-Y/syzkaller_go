package syzcmr

type CmrlogMutate struct {
	I                int
	Tag              int
	SquashAnyTotal   int
	SquashAnyTriage  int
	SpliceTotal      int
	SpliceTraige     int
	InsertCallTotal  int
	InsertCallTriage int
	MutateArgTotal   int
	MutateArgTriage  int
	RemoveCallTotal  int
	RemoveCallTriage int
}
