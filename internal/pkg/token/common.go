package token

const (
	Delimiter = "%"
)

type aliaser interface {
	// Alias returns an alias for the given package.
	Alias(string) string
}

// toExpr removes surrounding delimiters
func toExpr(expr string) (string, bool) {
	runes := []rune(expr)
	if len(runes) < 2 {
		return "", false
	}

	if string(runes[0]) != Delimiter || string(runes[len(runes)-1]) != Delimiter {
		return "", false
	}

	return string(runes[1 : len(runes)-1]), true
}
