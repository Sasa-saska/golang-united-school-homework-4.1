package reverse_string

func ReverseString(input string) (output string) {
	var k int
	var STR = []rune(input)

	for range STR {
		k++
	}
	for i := k; i > 0; i-- {
		I := i - 1
		output += string(STR[I:i])
	}
	return output
}
