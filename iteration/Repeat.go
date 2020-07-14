package iteraton

func Repeat(character string) string {
	repeated := ""
	for i := 0; i < 5; i++ {
		repeated += character
	}
	return repeated
}
