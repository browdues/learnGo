package nestedbrace

// NestIsValid takes a string and ensures the nesting therein is valid
func NestIsValid(s string) (result bool) {
	openerToCloser := map[rune]rune{
		rune('{'): rune('}'),
		rune('['): rune(']'),
		rune('('): rune(')'),
	}

	openers := make([]rune, 0)

	for _, c := range s {
		r := rune(c)

		if isOpener(r) {
			openers = append(openers, r)
		}

		if isCloser(r) {
			sizeOpeners := len(openers)
			if sizeOpeners < 1 {
				result = false
				break
			} else if openerToCloser[openers[sizeOpeners-1]] == r {
				result = true
				openers = openers[:sizeOpeners-1]
			} else {
				result = false
				return
			}
		}
	}

	if len(openers) != 0 {
		result = false
	}

	return
}

func isOpener(r rune) (result bool) {
	openBrace := rune('{')
	openBracket := rune('[')
	openParen := rune('(')

	result = false

	switch r {
	case openBrace:
		result = true
	case openBracket:
		result = true
	case openParen:
		result = true
	}

	return
}

func isCloser(r rune) (result bool) {
	closeBrace := rune('}')
	closeBracket := rune(']')
	closeParen := rune(')')

	result = false

	switch r {
	case closeBrace:
		result = true
	case closeBracket:
		result = true
	case closeParen:
		result = true
	}

	return
}
