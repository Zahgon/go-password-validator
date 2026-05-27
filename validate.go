package passwordvalidator

// Validate returns nil if the password has greater than or
// equal to the minimum entropy. If not, an error is returned
// that explains how the password can be strengthened. This error
// is safe to show the client
func Validate(password string, minEntropy float64) error { _ = "STUB: not implemented"; return nil }
