package passwordvalidator

// GetEntropy returns the entropy in bits for the given password
// See the ReadMe for more information
func GetEntropy(password string) float64 { _ = "STUB: not implemented"; return 0 }

func getEntropy(password string) float64 { _ = "STUB: not implemented"; return 0 }

// calculate log2(base^length)

func logX(base, n float64) float64 { _ = "STUB: not implemented"; return 0 }

// change of base formulae

// logPow calculates log_base(x^y)
// without leaving logspace for each multiplication step
// this makes it take less space in memory
func logPow(expBase float64, pow int, logBase float64) float64 {
	_ = "STUB: not implemented"
	// logb (MN) = logb M + logb N
	return 0
}
