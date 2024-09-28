package iteration

func Repeat(character string, numberOfIterations int) (repeated string) {
	for i := 0; i < numberOfIterations; i++ {
		repeated += character
	}
	return
}
