package passwordvalidator

const (
	replaceChars      = `!@$&*`
	sepChars          = `_-., `
	otherSpecialChars = `"#%'()+/:;<=>?[\]^{|}~`
	lowerChars        = `abcdefghijklmnopqrstuvwxyz`
	upperChars        = `ABCDEFGHIJKLMNOPQRSTUVWXYZ`
	digitsChars       = `0123456789`
)

func getBase(password string) int { _ = "STUB: not implemented"; return 0 }
