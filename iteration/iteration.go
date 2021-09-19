package iteration

// Repeat returns 'char' repeated 'repeatCount' times in a single string e.g
// char = 'a' -> repeated = 'aaaaa'
func Repeat(char string, repeatCount int) (repeated string) {
	for i := 0; i < repeatCount; i++ {
		repeated += char
	}
	return repeated
}
