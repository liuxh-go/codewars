package _kyu

/*

 */

func Interpreter(code string) string {
	result := make([]byte, 0)

	var index byte = 0
	for _, c := range []rune(code) {
		if c == '+' {
			if index == byte(255) {
				index = 0
			} else {
				index++
			}
		} else if c == '.' {
			result = append(result, index)
		}
	}

	return string(result)
}
