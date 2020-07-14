package iteraton

func Repeat(character string, repeatCount int) string {
	repeated := ""
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}
