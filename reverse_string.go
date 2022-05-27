package reverse_string

func ReverseString(input string) (output string) {
	var k int
	for range input {
		k++
	}
	for i := k; i > 0; i-- {
		I := i - 1
		output += input[I:i]
	}
	return output
}
