package iteration

// Repeat returns a character repeated n times
func Repeat(character string, nRepeat int) string {
	var repeated string

	for i := 0; i < nRepeat; i++ {
		repeated += character
	}
	return repeated
}
